package WeekCalculator

var dayOfTheWeek = map[string]int{
	"Mon": 1,
	"Tue": 2,
	"Wed": 3,
	"Thu": 4,
	"Fri": 5,
	"Sat": 6,
	"Sun": 7,
}

// Solution with a for loop
func Solution1(visits []string) int {
	// Visits has a length of 1 to 100 so there is at least one week
	numberOfWeeks := 1
	previousVisitDay := dayOfTheWeek[visits[0]]
	// Because we have consume the first day, let's do a good old for loop
	for i := 1; i < len(visits); i++ {
		visitDay := dayOfTheWeek[visits[i]]
		// If the visit day is lower in the order of the day than the previous then it is a new week
		if visitDay <= previousVisitDay {
			numberOfWeeks++
		}
		previousVisitDay = visitDay
	}
	return numberOfWeeks
}

// Solution with a loop using the range keyword
// I prefer Solution1 because we set previousVisitDay to the first element of the list.
// Here it looks less simple to understand the logic
func Solution2(visits []string) int {
	numberOfWeeks := 1
	previousVisitDay := 0
	for _, visit := range visits {
		visitDay := dayOfTheWeek[visit]
		if visitDay <= previousVisitDay {
			numberOfWeeks++
		}
		previousVisitDay = visitDay
	}
	return numberOfWeeks
}
