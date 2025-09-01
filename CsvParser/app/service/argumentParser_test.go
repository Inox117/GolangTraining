package service

import (
	"os"
	"testing"
)

func TestArgumentParser_RetrieveArguments_NoArguments(t *testing.T) {
	argumentParser := NewArgumentParser()
	os.Args = []string{"test"}
	_, err := argumentParser.RetrieveArguments()
	if err == nil {
		t.Error("Expected an error")
	}
	if err.Error() != "failed to parse arguments" {
		t.Error("Wrong error")
	}
}

func TestAppRunWithArguments(t *testing.T) {
	argumentParser := NewArgumentParser()
	os.Args = []string{
		"test", "-f=path", "-m=month",
	}
	retrieveArguments, err := argumentParser.RetrieveArguments()
	if err != nil {
		t.Error("Expected no error")
	}
	if retrieveArguments.SpecifiedMonth != "month" {
		t.Errorf("Expected month got %s", retrieveArguments.SpecifiedMonth)
	}
	if retrieveArguments.HistoryFilePath != "path" {
		t.Errorf("Expected path got %s", retrieveArguments.HistoryFilePath)
	}
}
