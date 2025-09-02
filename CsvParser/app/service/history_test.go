package service

import (
	"encoding/csv"
	"golang_training/CsvParser/app/model"
	"reflect"
	"strings"
	"testing"
)

import (
	"errors"
)

func TestCheckHeaderValidity(t *testing.T) {
	tests := []struct {
		name          string
		headers       []string
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid headers",
			headers:     []string{"date", "amount", "content"},
			expectError: false,
		},
		{
			name:          "Invalid headers - missing date",
			headers:       []string{"amount", "content"},
			expectError:   true,
			errorContains: "invalid header, should have 3 columns",
		},
		{
			name:          "Invalid headers - missing amount",
			headers:       []string{"date", "content"},
			expectError:   true,
			errorContains: "invalid header, should have 3 columns",
		},
		{
			name:          "Invalid headers - missing content",
			headers:       []string{"date", "amount"},
			expectError:   true,
			errorContains: "invalid header, should have 3 columns",
		},
		{
			name:          "Invalid headers - empty array",
			headers:       []string{},
			expectError:   true,
			errorContains: "invalid header, should have 3 columns",
		},
		{
			name:          "Invalid headers - wrong first column",
			headers:       []string{"amount", "date", "content"},
			expectError:   true,
			errorContains: "invalid header, first column should be date",
		},
		{
			name:          "Invalid headers - wrong second column",
			headers:       []string{"date", "content", "amount"},
			expectError:   true,
			errorContains: "invalid header, second column should be amount",
		},
		{
			name:          "Invalid headers - case sensitive",
			headers:       []string{"date", "amount", "Content"},
			expectError:   true,
			errorContains: "invalid header, third column should be content",
		},
		{
			name:          "Invalid headers - extra column",
			headers:       []string{"date", "amount", "content", "extra"},
			expectError:   true,
			errorContains: "invalid header, should have 3 columns",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewHistoryService()

			err := service.checkHeaderValidity(tt.headers)

			// Vérifie si on s'attend à une erreur
			if tt.expectError && err == nil {
				t.Errorf("Expected an error for headers %v, but got nil", tt.headers)
				return
			}

			// Vérifie si on ne s'attend pas à une erreur
			if !tt.expectError && err != nil {
				t.Errorf("Did not expect an error for headers %v, but got: %v", tt.headers, err)
				return
			}

			// Si on s'attend à une erreur et qu'on en a une, vérifie si elle contient le message attendu
			if tt.expectError && err != nil && tt.errorContains != "" {
				if !errors.Is(err, errors.New(tt.errorContains)) && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Error message does not contain expected text. Got: %v, Want: %v",
						err.Error(), tt.errorContains)
				}
			}
		})
	}
}

// Test avec des en-têtes nuls
func TestCheckHeaderValidityWithNilHeaders(t *testing.T) {
	service := NewHistoryService()
	err := service.checkHeaderValidity(nil)
	if err == nil {
		t.Errorf("Expected an error for nil headers, but got nil")
	}
}

func TestCheckHeaderValidityWithEmptyHeaders(t *testing.T) {
	service := NewHistoryService()
	err := service.checkHeaderValidity([]string{})
	if err == nil {
		t.Errorf("Expected an error for empty headers, but got nil")
	}
}

func TestProcessLines(t *testing.T) {
	tests := []struct {
		name          string
		csvContent    string
		givenMonth    string
		expected      *model.OutputData
		expectError   bool
		errorContains string
	}{
		{
			name: "Valid CSV content with matching month and only expenditure",
			csvContent: `date,amount,content
2022/01/25,-100000,rent
2022/01/06,-10000,debit
2022/01/05,-1000,eating out
2022/02/01,50000,salary`,
			givenMonth: "202201",
			expected: &model.OutputData{
				Period:           "2022/01",
				TotalIncome:      0,
				TotalExpenditure: -111000,
				Transactions: []model.Transaction{
					{Date: "2022/01/25", Amount: -100000, Content: "rent"},
					{Date: "2022/01/06", Amount: -10000, Content: "debit"},
					{Date: "2022/01/05", Amount: -1000, Content: "eating out"},
				},
			},
			expectError: false,
		},
		{
			name: "Valid CSV content with matching month and only income",
			csvContent: `date,amount,content
2022/01/15,200000,salary
2022/02/01,50000,salary`,
			givenMonth: "202201",
			expected: &model.OutputData{
				Period:           "2022/01",
				TotalIncome:      200000,
				TotalExpenditure: 0,
				Transactions: []model.Transaction{
					{Date: "2022/01/15", Amount: 200000, Content: "salary"},
				},
			},
			expectError: false,
		},
		{
			name: "Valid CSV content with matching month and both expenditure and income",
			csvContent: `date,amount,content
2022/01/25,-100000,rent
2022/01/15,200000,salary
2022/01/06,-10000,debit
2022/01/05,-1000,eating out
2022/02/01,50000,salary`,
			givenMonth: "202201",
			expected: &model.OutputData{
				Period:           "2022/01",
				TotalIncome:      200000,
				TotalExpenditure: -111000,
				Transactions: []model.Transaction{
					{Date: "2022/01/25", Amount: -100000, Content: "rent"},
					{Date: "2022/01/15", Amount: 200000, Content: "salary"},
					{Date: "2022/01/06", Amount: -10000, Content: "debit"},
					{Date: "2022/01/05", Amount: -1000, Content: "eating out"},
				},
			},
			expectError: false,
		},
		{
			name: "Valid CSV content with no matching month",
			csvContent: `date,amount,content
2022/02/25,-100000,rent
2022/02/06,-10000,debit
2022/02/05,-1000,eating out`,
			givenMonth: "202201",
			expected: &model.OutputData{
				Period:           "2022/01",
				TotalIncome:      0,
				TotalExpenditure: 0,
				Transactions:     []model.Transaction{},
			},
			expectError: false,
		},
		{
			name: "Invalid amount format",
			csvContent: `date,amount,content
2022/01/25,invalid,rent`,
			givenMonth:    "202201",
			expectError:   true,
			errorContains: "ParseInt",
		},
		{
			name: "Empty CSV content",
			csvContent: `date,amount,content
`,
			givenMonth: "202201",
			expected: &model.OutputData{
				Period:           "2022/01",
				TotalIncome:      0,
				TotalExpenditure: 0,
				Transactions:     []model.Transaction{},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialiser le service
			service := NewHistoryService()

			// Créer un reader CSV à partir de la chaîne de contenu
			reader := csv.NewReader(strings.NewReader(tt.csvContent))

			// Ignorer la première ligne d'en-tête qui a déjà été lue
			_, err := reader.Read()
			if err != nil && tt.csvContent != "" {
				t.Fatalf("Failed to read header: %v", err)
			}

			// Appeler la fonction à tester
			result, err := service.processLines(reader, tt.givenMonth)

			// Vérifier les erreurs
			if tt.expectError && err == nil {
				t.Errorf("Expected an error but got nil")
				return
			}

			if !tt.expectError && err != nil {
				t.Errorf("Did not expect an error but got: %v", err)
				return
			}

			if tt.expectError && err != nil {
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Error message does not contain expected text. Got: %v, Want: %v",
						err.Error(), tt.errorContains)
				}
				return
			}

			// Vérifier le résultat
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Result does not match expected. \nGot: %+v, \nWant: %+v", result, tt.expected)

				// Afficher plus de détails sur la différence
				if len(result.Transactions) != len(tt.expected.Transactions) {
					t.Errorf("Transaction length mismatch: got %d, want %d",
						len(result.Transactions), len(tt.expected.Transactions))
				} else {
					for i, got := range result.Transactions {
						want := tt.expected.Transactions[i]
						if !reflect.DeepEqual(got, want) {
							t.Errorf("Transaction at index %d differs: \ngot: %+v, \nwant: %+v", i, got, want)
						}
					}
				}

				if result.TotalIncome != tt.expected.TotalIncome {
					t.Errorf("TotalIncome mismatch: got %d, want %d",
						result.TotalIncome, tt.expected.TotalIncome)
				}

				if result.TotalExpenditure != tt.expected.TotalExpenditure {
					t.Errorf("TotalExpenditure mismatch: got %d, want %d",
						result.TotalExpenditure, tt.expected.TotalExpenditure)
				}
			}
		})
	}
}
