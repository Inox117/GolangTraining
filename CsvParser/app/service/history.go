package service

import (
	"golang_training/CsvParser/app/model"
	"os"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
)

// HistoryService is the service used to retrieve and process the CSV data
type HistoryService struct{}

func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

func (h *HistoryService) RetrieveHistoryFromFile(historyPath string) (*model.History, error) {
	historyFile, err := os.Open(historyPath)
	if err != nil {
		return nil, err
	}
	defer historyFile.Close()
	var historyLines []model.HistoryLine

	err = gocsv.UnmarshalFile(historyFile, &historyLines)
	if err != nil {
		return nil, err
	}

	var history model.History
	historyPerMonth := make(map[string][]model.HistoryLine)
	for _, historyLine := range historyLines {
		historyMonth := historyLine.Date[:8]
		historyMonth = strings.Replace(historyMonth, "/", "", -1)
		if historyPerMonth[historyMonth] == nil {
			historyPerMonth[historyMonth] = []model.HistoryLine{}
		}
		historyPerMonth[historyMonth] = append(historyPerMonth[historyMonth], historyLine)
	}
	history.HistoryPerMonth = historyPerMonth
	return &history, nil
}

func (h *HistoryService) ProcessHistoryForGivenMonth(history *model.History, givenMonth string) *model.OutputData {
	output := model.OutputData{
		Period: givenMonth[:4] + "/" + givenMonth[4:],
	}
	var totalIncome int64
	var totalExpenditure int64
	transactions := make([]model.Transaction, 0)
	monthlyData := history.HistoryPerMonth[givenMonth]
	for _, historyLine := range monthlyData {
		// No need to handle the "== 0" case
		// There is no point in spending or receiving 0 (Whatever currency)
		if historyLine.Amount > 0 {
			totalIncome += historyLine.Amount
		}
		if historyLine.Amount < 0 {
			totalExpenditure += historyLine.Amount
		}
		transactions = append(transactions, model.Transaction{
			Date:    historyLine.Date,
			Amount:  historyLine.Amount,
			Content: historyLine.Content,
		})
	}
	// transactions have to be sorted in descending order
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date > transactions[j].Date
	})
	output.TotalIncome = totalIncome
	output.TotalExpenditure = totalExpenditure
	output.Transactions = transactions
	return &output
}
