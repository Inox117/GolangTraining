package service

import "regexp"

var re = regexp.MustCompile(`\d{4}(?:0[1-9]|1[0-2])`)

// DateValidator is the service used to check that the specified month is valid
type DateValidator struct{}

func NewDateValidator() *DateValidator {
	return &DateValidator{}
}

func (*DateValidator) CheckSpecifiedMonthValidity(specifiedMonth string) bool {
	return re.MatchString(specifiedMonth)
}
