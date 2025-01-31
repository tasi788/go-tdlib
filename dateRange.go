package tdlib

// DateRange Represents a date range
type DateRange struct {
	tdCommon
	StartDate int32 `json:"start_date"` // Point in time (Unix timestamp) at which the date range begins
	EndDate   int32 `json:"end_date"`   // Point in time (Unix timestamp) at which the date range ends
}

// MessageType return the string telegram-type of DateRange
func (dateRange *DateRange) MessageType() string {
	return "dateRange"
}

// NewDateRange creates a new DateRange
//
// @param startDate Point in time (Unix timestamp) at which the date range begins
// @param endDate Point in time (Unix timestamp) at which the date range ends
func NewDateRange(startDate int32, endDate int32) *DateRange {
	dateRangeTemp := DateRange{
		tdCommon:  tdCommon{Type: "dateRange"},
		StartDate: startDate,
		EndDate:   endDate,
	}

	return &dateRangeTemp
}
