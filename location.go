package tdlib

// Location Describes a location on planet Earth
type Location struct {
	tdCommon
	Latitude           float64 `json:"latitude"`            // Latitude of the location in degrees; as defined by the sender
	Longitude          float64 `json:"longitude"`           // Longitude of the location, in degrees; as defined by the sender
	HorizontalAccuracy float64 `json:"horizontal_accuracy"` // The estimated horizontal accuracy of the location, in meters; as defined by the sender. 0 if unknown
}

// MessageType return the string telegram-type of Location
func (location *Location) MessageType() string {
	return "location"
}

// NewLocation creates a new Location
//
// @param latitude Latitude of the location in degrees; as defined by the sender
// @param longitude Longitude of the location, in degrees; as defined by the sender
// @param horizontalAccuracy The estimated horizontal accuracy of the location, in meters; as defined by the sender. 0 if unknown
func NewLocation(latitude float64, longitude float64, horizontalAccuracy float64) *Location {
	locationTemp := Location{
		tdCommon:           tdCommon{Type: "location"},
		Latitude:           latitude,
		Longitude:          longitude,
		HorizontalAccuracy: horizontalAccuracy,
	}

	return &locationTemp
}
