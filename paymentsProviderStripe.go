package tdlib

// PaymentsProviderStripe Stripe payment provider
type PaymentsProviderStripe struct {
	tdCommon
	PublishableKey     string `json:"publishable_key"`      // Stripe API publishable key
	NeedCountry        bool   `json:"need_country"`         // True, if the user country must be provided
	NeedPostalCode     bool   `json:"need_postal_code"`     // True, if the user ZIP/postal code must be provided
	NeedCardholderName bool   `json:"need_cardholder_name"` // True, if the cardholder name must be provided
}

// MessageType return the string telegram-type of PaymentsProviderStripe
func (paymentsProviderStripe *PaymentsProviderStripe) MessageType() string {
	return "paymentsProviderStripe"
}

// NewPaymentsProviderStripe creates a new PaymentsProviderStripe
//
// @param publishableKey Stripe API publishable key
// @param needCountry True, if the user country must be provided
// @param needPostalCode True, if the user ZIP/postal code must be provided
// @param needCardholderName True, if the cardholder name must be provided
func NewPaymentsProviderStripe(publishableKey string, needCountry bool, needPostalCode bool, needCardholderName bool) *PaymentsProviderStripe {
	paymentsProviderStripeTemp := PaymentsProviderStripe{
		tdCommon:           tdCommon{Type: "paymentsProviderStripe"},
		PublishableKey:     publishableKey,
		NeedCountry:        needCountry,
		NeedPostalCode:     needPostalCode,
		NeedCardholderName: needCardholderName,
	}

	return &paymentsProviderStripeTemp
}
