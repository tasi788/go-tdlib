package tdlib

// Invoice Product invoice
type Invoice struct {
	tdCommon
	Currency                   string             `json:"currency"`                       // ISO 4217 currency code
	PriceParts                 []LabeledPricePart `json:"price_parts"`                    // A list of objects used to calculate the total price of the product
	MaxTipAmount               int64              `json:"max_tip_amount"`                 // The maximum allowed amount of tip in the smallest units of the currency
	SuggestedTipAmounts        []int64            `json:"suggested_tip_amounts"`          // Suggested amounts of tip in the smallest units of the currency
	IsTest                     bool               `json:"is_test"`                        // True, if the payment is a test payment
	NeedName                   bool               `json:"need_name"`                      // True, if the user's name is needed for payment
	NeedPhoneNumber            bool               `json:"need_phone_number"`              // True, if the user's phone number is needed for payment
	NeedEmailAddress           bool               `json:"need_email_address"`             // True, if the user's email address is needed for payment
	NeedShippingAddress        bool               `json:"need_shipping_address"`          // True, if the user's shipping address is needed for payment
	SendPhoneNumberToProvider  bool               `json:"send_phone_number_to_provider"`  // True, if the user's phone number will be sent to the provider
	SendEmailAddressToProvider bool               `json:"send_email_address_to_provider"` // True, if the user's email address will be sent to the provider
	IsFlexible                 bool               `json:"is_flexible"`                    // True, if the total price depends on the shipping method
}

// MessageType return the string telegram-type of Invoice
func (invoice *Invoice) MessageType() string {
	return "invoice"
}

// NewInvoice creates a new Invoice
//
// @param currency ISO 4217 currency code
// @param priceParts A list of objects used to calculate the total price of the product
// @param maxTipAmount The maximum allowed amount of tip in the smallest units of the currency
// @param suggestedTipAmounts Suggested amounts of tip in the smallest units of the currency
// @param isTest True, if the payment is a test payment
// @param needName True, if the user's name is needed for payment
// @param needPhoneNumber True, if the user's phone number is needed for payment
// @param needEmailAddress True, if the user's email address is needed for payment
// @param needShippingAddress True, if the user's shipping address is needed for payment
// @param sendPhoneNumberToProvider True, if the user's phone number will be sent to the provider
// @param sendEmailAddressToProvider True, if the user's email address will be sent to the provider
// @param isFlexible True, if the total price depends on the shipping method
func NewInvoice(currency string, priceParts []LabeledPricePart, maxTipAmount int64, suggestedTipAmounts []int64, isTest bool, needName bool, needPhoneNumber bool, needEmailAddress bool, needShippingAddress bool, sendPhoneNumberToProvider bool, sendEmailAddressToProvider bool, isFlexible bool) *Invoice {
	invoiceTemp := Invoice{
		tdCommon:                   tdCommon{Type: "invoice"},
		Currency:                   currency,
		PriceParts:                 priceParts,
		MaxTipAmount:               maxTipAmount,
		SuggestedTipAmounts:        suggestedTipAmounts,
		IsTest:                     isTest,
		NeedName:                   needName,
		NeedPhoneNumber:            needPhoneNumber,
		NeedEmailAddress:           needEmailAddress,
		NeedShippingAddress:        needShippingAddress,
		SendPhoneNumberToProvider:  sendPhoneNumberToProvider,
		SendEmailAddressToProvider: sendEmailAddressToProvider,
		IsFlexible:                 isFlexible,
	}

	return &invoiceTemp
}
