package paymentwall

import "fmt"

const (
	PeriodTypeDay   = "day"
	PeriodTypeWeek  = "week"
	PeriodTypeMonth = "month"
)

type period struct {
	Type   string
	Length int
}

func NewPeriod(periodType string, periodLength int) *period {
	validatePeriod(periodType, periodLength)

	return &period{
		Type:   periodType,
		Length: periodLength,
	}
}

func validatePeriod(periodType string, periodLength int) {
	if periodType != PeriodTypeDay &&
		periodType != PeriodTypeWeek &&
		periodType != PeriodTypeMonth {
		panic(fmt.Errorf("Invalid period type: '%s'", periodType))
	}
}
