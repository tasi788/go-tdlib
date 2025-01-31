package tdlib

// Venue Describes a venue
type Venue struct {
	tdCommon
	Location *Location `json:"location"` // Venue location; as defined by the sender
	Title    string    `json:"title"`    // Venue name; as defined by the sender
	Address  string    `json:"address"`  // Venue address; as defined by the sender
	Provider string    `json:"provider"` // Provider of the venue database; as defined by the sender. Currently, only "foursquare" and "gplaces" (Google Places) need to be supported
	Id       string    `json:"id"`       // Identifier of the venue in the provider database; as defined by the sender
	Type     string    `json:"type"`     // Type of the venue in the provider database; as defined by the sender
}

// MessageType return the string telegram-type of Venue
func (venue *Venue) MessageType() string {
	return "venue"
}

// NewVenue creates a new Venue
//
// @param location Venue location; as defined by the sender
// @param title Venue name; as defined by the sender
// @param address Venue address; as defined by the sender
// @param provider Provider of the venue database; as defined by the sender. Currently, only "foursquare" and "gplaces" (Google Places) need to be supported
// @param id Identifier of the venue in the provider database; as defined by the sender
// @param typeParam Type of the venue in the provider database; as defined by the sender
func NewVenue(location *Location, title string, address string, provider string, id string, typeParam string) *Venue {
	venueTemp := Venue{
		tdCommon: tdCommon{Type: "venue"},
		Location: location,
		Title:    title,
		Address:  address,
		Provider: provider,
		Id:       id,
		Type:     typeParam,
	}

	return &venueTemp
}
