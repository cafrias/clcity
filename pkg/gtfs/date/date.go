package date

import (
	"fmt"
	"time"
)

// Date represents a date with format: "YYYYMMDD"
type Date string

// FormatDate returns a Date string
func FormatDate(d time.Time) Date {
	return Date(fmt.Sprintf("%04d%02d%02d", d.Year(), d.Month(), d.Day()))
}
