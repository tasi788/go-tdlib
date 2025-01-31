package tdlib

// StatisticalValue A value with information about its recent changes
type StatisticalValue struct {
	tdCommon
	Value                float64 `json:"value"`                  // The current value
	PreviousValue        float64 `json:"previous_value"`         // The value for the previous day
	GrowthRatePercentage float64 `json:"growth_rate_percentage"` // The growth rate of the value, as a percentage
}

// MessageType return the string telegram-type of StatisticalValue
func (statisticalValue *StatisticalValue) MessageType() string {
	return "statisticalValue"
}

// NewStatisticalValue creates a new StatisticalValue
//
// @param value The current value
// @param previousValue The value for the previous day
// @param growthRatePercentage The growth rate of the value, as a percentage
func NewStatisticalValue(value float64, previousValue float64, growthRatePercentage float64) *StatisticalValue {
	statisticalValueTemp := StatisticalValue{
		tdCommon:             tdCommon{Type: "statisticalValue"},
		Value:                value,
		PreviousValue:        previousValue,
		GrowthRatePercentage: growthRatePercentage,
	}

	return &statisticalValueTemp
}
