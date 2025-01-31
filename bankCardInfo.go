package tdlib

import (
	"encoding/json"
	"fmt"
)

// BankCardInfo Information about a bank card
type BankCardInfo struct {
	tdCommon
	Title   string                  `json:"title"`   // Title of the bank card description
	Actions []BankCardActionOpenUrl `json:"actions"` // Actions that can be done with the bank card number
}

// MessageType return the string telegram-type of BankCardInfo
func (bankCardInfo *BankCardInfo) MessageType() string {
	return "bankCardInfo"
}

// NewBankCardInfo creates a new BankCardInfo
//
// @param title Title of the bank card description
// @param actions Actions that can be done with the bank card number
func NewBankCardInfo(title string, actions []BankCardActionOpenUrl) *BankCardInfo {
	bankCardInfoTemp := BankCardInfo{
		tdCommon: tdCommon{Type: "bankCardInfo"},
		Title:    title,
		Actions:  actions,
	}

	return &bankCardInfoTemp
}

// GetBankCardInfo Returns information about a bank card
// @param bankCardNumber The bank card number
func (client *Client) GetBankCardInfo(bankCardNumber string) (*BankCardInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getBankCardInfo",
		"bank_card_number": bankCardNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var bankCardInfo BankCardInfo
	err = json.Unmarshal(result.Raw, &bankCardInfo)
	return &bankCardInfo, err
}
