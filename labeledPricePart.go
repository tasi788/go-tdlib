package tdlib

// LabeledPricePart Portion of the price of a product (e.g., "delivery cost", "tax amount")
type LabeledPricePart struct {
	tdCommon
	Label  string `json:"label"`  // Label for this portion of the product price
	Amount int64  `json:"amount"` // Currency amount in the smallest units of the currency
}

// MessageType return the string telegram-type of LabeledPricePart
func (labeledPricePart *LabeledPricePart) MessageType() string {
	return "labeledPricePart"
}

// NewLabeledPricePart creates a new LabeledPricePart
//
// @param label Label for this portion of the product price
// @param amount Currency amount in the smallest units of the currency
func NewLabeledPricePart(label string, amount int64) *LabeledPricePart {
	labeledPricePartTemp := LabeledPricePart{
		tdCommon: tdCommon{Type: "labeledPricePart"},
		Label:    label,
		Amount:   amount,
	}

	return &labeledPricePartTemp
}
