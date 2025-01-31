package tdlib

// Date Represents a date according to the Gregorian calendar
type Date struct {
	tdCommon
	Day   int32 `json:"day"`   // Day of the month; 1-31
	Month int32 `json:"month"` // Month; 1-12
	Year  int32 `json:"year"`  // Year; 1-9999
}

// MessageType return the string telegram-type of Date
func (date *Date) MessageType() string {
	return "date"
}

// NewDate creates a new Date
//
// @param day Day of the month; 1-31
// @param month Month; 1-12
// @param year Year; 1-9999
func NewDate(day int32, month int32, year int32) *Date {
	dateTemp := Date{
		tdCommon: tdCommon{Type: "date"},
		Day:      day,
		Month:    month,
		Year:     year,
	}

	return &dateTemp
}
