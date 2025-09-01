package service

import (
	"golang_training/CsvParser/app/model"
	"testing"
)

func TestHistoryService_RetrieveHistoryFromFile_FileDoesNotExist(t *testing.T) {
	historyService := NewHistoryService()
	_, err := historyService.RetrieveHistoryFromFile("data/notexist.csv")
	if err == nil {
		t.Errorf("RetrieveHistoryFromFile should have returned an error")
	}
}

func TestHistoryService_RetrieveHistoryFromFile_Success(t *testing.T) {
	historyService := NewHistoryService()
	history, err := historyService.RetrieveHistoryFromFile("./../../data/history.csv")
	if err != nil {
		t.Errorf("RetrieveHistoryFromFile should not have returned an error")
	}
	if history == nil {
		t.Errorf("RetrieveHistoryFromFile should have returned a history")
	}
	if len(history.HistoryPerMonth) != 4 {
		t.Errorf("RetrieveHistoryFromFile should have returned 4 records")
	}
	if len(history.HistoryPerMonth["202306"]) != 1 {
		t.Errorf("RetrieveHistoryFromFile should have returned 1 record for 202306")
	}
	if len(history.HistoryPerMonth["202201"]) != 3 {
		t.Errorf("RetrieveHistoryFromFile should have returned 3 records for 202201")
	}
	if len(history.HistoryPerMonth["202202"]) != 1 {
		t.Errorf("RetrieveHistoryFromFile should have returned 1 record for 202202")
	}
	if len(history.HistoryPerMonth["202303"]) != 1 {
		t.Errorf("RetrieveHistoryFromFile should have returned 1 record for 202303")
	}
}

func TestHistoryService_ProcessHistoryForGivenMonth_NoDataForGivenMonth(t *testing.T) {
	historyService := NewHistoryService()
	mockHistory := model.History{
		HistoryPerMonth: map[string][]model.HistoryLine{
			"202306": {
				{
					Date:    "2023/06/01",
					Amount:  -100,
					Content: "eating out",
				},
				{
					Date:    "2023/06/02",
					Amount:  -200,
					Content: "eating out",
				},
				{
					Date:    "2023/06/03",
					Amount:  600,
					Content: "salary",
				},
				{
					Date:    "2023/06/22",
					Amount:  400,
					Content: "bonus",
				},
			},
			"202301": {
				{
					Date:    "2023/03/01",
					Amount:  600,
					Content: "salary",
				},
			},
		},
	}
	outputData := historyService.ProcessHistoryForGivenMonth(&mockHistory, "202305")
	if outputData == nil {
		t.Errorf("ProcessHistoryForGivenMonth should have returned a output")
	}
	if outputData.Period != "2023/05" {
		t.Errorf("outputData Period should have been 2023/05 got %s", outputData.Period)
	}
	if outputData.TotalIncome != 0 {
		t.Errorf("outputData TotalIncome should have been 0 got %d", outputData.TotalIncome)
	}
	if outputData.TotalExpenditure != 0 {
		t.Errorf("outputData TotalExpenditure should have been 0 got %d", outputData.TotalExpenditure)
	}
	if len(outputData.Transactions) != 0 {
		t.Errorf("outputData should have 0 Transactions got %d", len(outputData.Transactions))
	}
}

func TestHistoryService_ProcessHistoryForGivenMonth_DataForGivenMonth(t *testing.T) {
	historyService := NewHistoryService()
	mockHistory := model.History{
		HistoryPerMonth: map[string][]model.HistoryLine{
			"202306": {
				{
					Date:    "2023/06/01",
					Amount:  -100,
					Content: "eating out",
				},
				{
					Date:    "2023/06/02",
					Amount:  -200,
					Content: "eating out",
				},
				{
					Date:    "2023/06/03",
					Amount:  600,
					Content: "salary",
				},
				{
					Date:    "2023/06/22",
					Amount:  400,
					Content: "bonus",
				},
			},
			"202301": {
				{
					Date:    "2023/03/01",
					Amount:  600,
					Content: "salary",
				},
			},
		},
	}
	outputData := historyService.ProcessHistoryForGivenMonth(&mockHistory, "202306")
	if outputData == nil {
		t.Errorf("ProcessHistoryForGivenMonth should have returned a output")
	}
	if outputData.Period != "2023/06" {
		t.Errorf("outputData Period should have been 2023/06 got %s", outputData.Period)
	}
	if outputData.TotalIncome != 1000 {
		t.Errorf("outputData TotalIncome should have been 1000 got %d", outputData.TotalIncome)
	}
	if outputData.TotalExpenditure != -300 {
		t.Errorf("outputData TotalExpenditure should have been 0 got %d", outputData.TotalExpenditure)
	}
	if len(outputData.Transactions) != 4 {
		t.Errorf("outputData should have 4 Transactions got %d", len(outputData.Transactions))
	}
	transaction := outputData.Transactions[0]
	if transaction.Date != "2023/06/01" {
		t.Errorf("transaction Date should have been 2023/06/01 got %s", transaction.Date)
	}
	if transaction.Amount != -100 {
		t.Errorf("transaction Amount should have been -100 got %d", transaction.Amount)
	}
	if transaction.Content != "eating out" {
		t.Errorf("transaction Content should have been eating out got %s", transaction.Content)
	}
	transaction = outputData.Transactions[1]
	if transaction.Date != "2023/06/02" {
		t.Errorf("transaction Date should have been 2023/06/02 got %s", transaction.Date)
	}
	if transaction.Amount != -200 {
		t.Errorf("transaction Amount should have been -200 got %d", transaction.Amount)
	}
	if transaction.Content != "eating out" {
		t.Errorf("transaction Content should have been eating out got %s", transaction.Content)
	}
	transaction = outputData.Transactions[2]
	if transaction.Date != "2023/06/03" {
		t.Errorf("transaction Date should have been 2023/06/03 got %s", transaction.Date)
	}
	if transaction.Amount != 600 {
		t.Errorf("transaction Amount should have been 600 got %d", transaction.Amount)
	}
	if transaction.Content != "salary" {
		t.Errorf("transaction Content should have been salary got %s", transaction.Content)
	}
	transaction = outputData.Transactions[3]
	if transaction.Date != "2023/06/22" {
		t.Errorf("transaction Date should have been 2023/06/22 got %s", transaction.Date)
	}
	if transaction.Amount != 400 {
		t.Errorf("transaction Amount should have been 400 got %d", transaction.Amount)
	}
	if transaction.Content != "bonus" {
		t.Errorf("transaction Content should have been bonus got %s", transaction.Content)
	}
}
