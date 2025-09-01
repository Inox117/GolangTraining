package model

// Arguments represents the expected arguments when calling the app
type Arguments struct {
	HistoryFilePath string `short:"f" long:"history_file" description:"file path of history file." required:"true"`
	SpecifiedMonth  string `short:"m" long:"month" description:"month for which we want the history. Format is YYYYMM." required:"true"`
}
