package service

import "testing"

func TestDateValidator_CheckSpecifiedMonthValidity(t *testing.T) {
	dateValidator := NewDateValidator()
	tests := []struct {
		name           string
		specifiedMonth string
		isValid        bool
	}{
		{
			name:           "january",
			specifiedMonth: "202501",
			isValid:        true,
		},
		{
			name:           "february",
			specifiedMonth: "202502",
			isValid:        true,
		},
		{
			name:           "march",
			specifiedMonth: "202503",
			isValid:        true,
		},
		{
			name:           "april",
			specifiedMonth: "202504",
			isValid:        true,
		},
		{
			name:           "may",
			specifiedMonth: "202505",
			isValid:        true,
		},
		{
			name:           "june",
			specifiedMonth: "202506",
			isValid:        true,
		},
		{
			name:           "july",
			specifiedMonth: "202507",
			isValid:        true,
		},
		{
			name:           "august",
			specifiedMonth: "202508",
			isValid:        true,
		},
		{
			name:           "september",
			specifiedMonth: "202509",
			isValid:        true,
		},
		{
			name:           "october",
			specifiedMonth: "202510",
			isValid:        true,
		},
		{
			name:           "november",
			specifiedMonth: "202511",
			isValid:        true,
		},
		{
			name:           "december",
			specifiedMonth: "202512",
			isValid:        true,
		},
		{
			name:           "Incorrect month",
			specifiedMonth: "202513",
			isValid:        false,
		},
		{
			name:           "not even a month",
			specifiedMonth: "azerty",
			isValid:        false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := dateValidator.CheckSpecifiedMonthValidity(test.specifiedMonth)
			if result != test.isValid {
				t.Error("Wrong result")
			}
		})
	}
}
