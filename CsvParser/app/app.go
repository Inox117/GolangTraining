package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang_training/CsvParser/app/service"
)

type CsvParserApp struct {
	argumentParser service.ArgumentParser
	dateValidator  service.DateValidator
	historyService service.HistoryService
}

func NewCsvParserApp() *CsvParserApp {
	argumentParser := service.NewArgumentParser()
	dateValidator := service.NewDateValidator()
	historyService := service.NewHistoryService()
	return &CsvParserApp{
		argumentParser: *argumentParser,
		dateValidator:  *dateValidator,
		historyService: *historyService,
	}
}

func (app *CsvParserApp) Run() error {
	// First we retrieve the arguments
	arguments, err := app.argumentParser.RetrieveArguments()
	if err != nil {
		return err
	}
	// we check the validity of the specified month
	if !app.dateValidator.CheckSpecifiedMonthValidity(arguments.SpecifiedMonth) {
		return errors.New("specified month is not valid")
	}
	// We retrieve the history from the specified file path
	historyFromFile, err := app.historyService.RetrieveHistoryFromFile(arguments.HistoryFilePath)
	if err != nil {
		return errors.New("could not retrieve history from file")
	}
	outputData := app.historyService.ProcessHistoryForGivenMonth(historyFromFile, arguments.SpecifiedMonth)
	jsonData, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		return errors.New("could not marshal json data")
	}
	fmt.Println(string(jsonData))
	return nil
}
