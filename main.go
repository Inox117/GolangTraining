package main

import (
	"golang_training/CsvParser/app"
	"log"
)

func main() {
	application := app.NewCsvParserApp()
	err := application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
