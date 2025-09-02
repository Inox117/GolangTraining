package service

import (
	"encoding/csv"
	"errors"
	"golang_training/CsvParser/app/model"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// HistoryService is the service used to retrieve and process the CSV data
type HistoryService struct{}

func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

func (h *HistoryService) ProcessFile(historyPath string, givenMonth string) (*model.OutputData, error) {
	file, err := os.Open(historyPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// we check the validity of the header
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}
	err = h.checkHeaderValidity(headers)
	if err != nil {
		return nil, err
	}
	// we will now process the CSV file line by line
	output, err := h.processLines(reader, givenMonth)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (h *HistoryService) checkHeaderValidity(headers []string) error {
	if len(headers) != 3 {
		return errors.New("invalid header, should have 3 columns")
	}
	if headers[0] != "date" {
		return errors.New("invalid header, first column should be date")
	}
	if headers[1] != "amount" {
		return errors.New("invalid header, second column should be amount")
	}
	if headers[2] != "content" {
		return errors.New("invalid header, third column should be content")
	}
	return nil
}

func (h *HistoryService) processLines(reader *csv.Reader, givenMonth string) (*model.OutputData, error) {
	var totalIncome int64
	var totalExpenditure int64
	transactions := make([]model.Transaction, 0)
	for {
		record, err := reader.Read()
		// Check if we have reached the end of the file
		if err == io.EOF {
			break
		}
		date := strings.Replace(record[0], "/", "", -1)
		// we check the validity of the specified month
		if date[:6] == givenMonth {
			// we convert the CSV record to a History object
			amount, err := strconv.ParseInt(record[1], 10, 64)
			if err != nil {
				return nil, err
			}
			if amount > 0 {
				totalIncome += amount
			}
			if amount < 0 {
				totalExpenditure += amount
			}
			transactions = append(transactions, model.Transaction{
				Date:    record[0],
				Amount:  amount,
				Content: record[2],
			})
		}
	}
	output := &model.OutputData{
		Period: givenMonth[:4] + "/" + givenMonth[4:],
	}
	// transactions have to be sorted in descending order
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date > transactions[j].Date
	})
	output.TotalIncome = totalIncome
	output.TotalExpenditure = totalExpenditure
	output.Transactions = transactions
	return output, nil
}
