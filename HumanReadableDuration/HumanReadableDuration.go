package HumanReadableDuration

import (
	"fmt"
	"strings"
)

const MinuteInSeconds = 60
const HourInSeconds = 60 * MinuteInSeconds
const DayInSeconds = 24 * HourInSeconds
const YearInSeconds = 365 * DayInSeconds

func Solution(seconds int64) string {
	if seconds == 0 {
		return "now"
	}
	dateInStrings := make([]string, 0)
	// Year
	numberOfYear := seconds / YearInSeconds
	rest := seconds % YearInSeconds
	if numberOfYear >= 1 {
		dateInStrings = append(dateInStrings, produceStringOfTimeUnit(numberOfYear, "year"))
	}
	// Day
	numberOfDay := rest / DayInSeconds
	rest = rest % DayInSeconds
	if numberOfDay >= 1 {
		dateInStrings = append(dateInStrings, produceStringOfTimeUnit(numberOfDay, "day"))
	}
	// Hour
	numberOfHour := rest / HourInSeconds
	rest = rest % HourInSeconds
	if numberOfHour >= 1 {
		dateInStrings = append(dateInStrings, produceStringOfTimeUnit(numberOfHour, "hour"))
	}
	// Minute
	numberOfMinute := rest / MinuteInSeconds
	rest = rest % MinuteInSeconds
	if numberOfMinute >= 1 {
		dateInStrings = append(dateInStrings, produceStringOfTimeUnit(numberOfMinute, "minute"))
	}
	// Second
	if rest >= 1 {
		dateInStrings = append(dateInStrings, produceStringOfTimeUnit(rest, "second"))
	}
	return produceMessage(dateInStrings)
}

func produceStringOfTimeUnit(value int64, unit string) string {
	if value > 1 {
		return fmt.Sprintf("%d %ss", value, unit)
	}
	return fmt.Sprintf("%d %s", value, unit)
}

func produceMessage(components []string) string {
	if len(components) == 1 {
		return components[0]
	}
	return strings.Join(components[0:len(components)-1], ", ") + " and " + components[len(components)-1]
}
