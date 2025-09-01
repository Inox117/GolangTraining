package model

// History is the structure used to save the data from the CSV file
type History struct {
	HistoryPerMonth map[string][]HistoryLine
}

// HistoryLine represents the structure of a line in the CSV file
type HistoryLine struct {
	Date    string `csv:"date"`
	Amount  int64  `csv:"amount"`
	Content string `csv:"content"`
}
