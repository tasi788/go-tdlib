package tdlib

// BankCardActionOpenUrl Describes an action associated with a bank card number
type BankCardActionOpenUrl struct {
	tdCommon
	Text string `json:"text"` // Action text
	Url  string `json:"url"`  // The URL to be opened
}

// MessageType return the string telegram-type of BankCardActionOpenUrl
func (bankCardActionOpenUrl *BankCardActionOpenUrl) MessageType() string {
	return "bankCardActionOpenUrl"
}

// NewBankCardActionOpenUrl creates a new BankCardActionOpenUrl
//
// @param text Action text
// @param url The URL to be opened
func NewBankCardActionOpenUrl(text string, url string) *BankCardActionOpenUrl {
	bankCardActionOpenUrlTemp := BankCardActionOpenUrl{
		tdCommon: tdCommon{Type: "bankCardActionOpenUrl"},
		Text:     text,
		Url:      url,
	}

	return &bankCardActionOpenUrlTemp
}
