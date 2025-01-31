package tdlib

// PaymentFormTheme Theme colors for a payment form
type PaymentFormTheme struct {
	tdCommon
	BackgroundColor int32 `json:"background_color"`  // A color of the payment form background in the RGB24 format
	TextColor       int32 `json:"text_color"`        // A color of text in the RGB24 format
	HintColor       int32 `json:"hint_color"`        // A color of hints in the RGB24 format
	LinkColor       int32 `json:"link_color"`        // A color of links in the RGB24 format
	ButtonColor     int32 `json:"button_color"`      // A color of the buttons in the RGB24 format
	ButtonTextColor int32 `json:"button_text_color"` // A color of text on the buttons in the RGB24 format
}

// MessageType return the string telegram-type of PaymentFormTheme
func (paymentFormTheme *PaymentFormTheme) MessageType() string {
	return "paymentFormTheme"
}

// NewPaymentFormTheme creates a new PaymentFormTheme
//
// @param backgroundColor A color of the payment form background in the RGB24 format
// @param textColor A color of text in the RGB24 format
// @param hintColor A color of hints in the RGB24 format
// @param linkColor A color of links in the RGB24 format
// @param buttonColor A color of the buttons in the RGB24 format
// @param buttonTextColor A color of text on the buttons in the RGB24 format
func NewPaymentFormTheme(backgroundColor int32, textColor int32, hintColor int32, linkColor int32, buttonColor int32, buttonTextColor int32) *PaymentFormTheme {
	paymentFormThemeTemp := PaymentFormTheme{
		tdCommon:        tdCommon{Type: "paymentFormTheme"},
		BackgroundColor: backgroundColor,
		TextColor:       textColor,
		HintColor:       hintColor,
		LinkColor:       linkColor,
		ButtonColor:     buttonColor,
		ButtonTextColor: buttonTextColor,
	}

	return &paymentFormThemeTemp
}
