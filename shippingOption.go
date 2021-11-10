package tdlib

// ShippingOption One shipping option
type ShippingOption struct {
	tdCommon
	Id         string             `json:"id"`          // Shipping option identifier
	Title      string             `json:"title"`       // Option title
	PriceParts []LabeledPricePart `json:"price_parts"` // A list of objects used to calculate the total shipping costs
}

// MessageType return the string telegram-type of ShippingOption
func (shippingOption *ShippingOption) MessageType() string {
	return "shippingOption"
}

// NewShippingOption creates a new ShippingOption
//
// @param id Shipping option identifier
// @param title Option title
// @param priceParts A list of objects used to calculate the total shipping costs
func NewShippingOption(id string, title string, priceParts []LabeledPricePart) *ShippingOption {
	shippingOptionTemp := ShippingOption{
		tdCommon:   tdCommon{Type: "shippingOption"},
		Id:         id,
		Title:      title,
		PriceParts: priceParts,
	}

	return &shippingOptionTemp
}
