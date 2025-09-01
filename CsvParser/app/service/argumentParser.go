package service

import (
	"fmt"
	"golang_training/CsvParser/app/model"

	"github.com/jessevdk/go-flags"
)

// ArgumentParser is the service used to parse the CLI arguments
type ArgumentParser struct{}

func NewArgumentParser() *ArgumentParser {
	return &ArgumentParser{}
}

func (parser *ArgumentParser) RetrieveArguments() (*model.Arguments, error) {
	var arguments model.Arguments
	_, err := flags.Parse(&arguments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse arguments")
	}
	return &arguments, nil
}
