package model

// Transaction represents the structure of all operations in the given month
type Transaction struct {
	Date    string `json:"date"`
	Amount  int64  `json:"amount"`
	Content string `json:"content"`
}

// OutputData represents the processed CSV data as it should be printed
type OutputData struct {
	Period           string        `json:"period"`
	TotalIncome      int64         `json:"total_income"`
	TotalExpenditure int64         `json:"total_expenditure"`
	Transactions     []Transaction `json:"transactions"`
}
