package tdlib

import (
	"encoding/json"
	"fmt"
)

// AccountTtl Contains information about the period of inactivity after which the current user's account will automatically be deleted
type AccountTtl struct {
	tdCommon
	Days int32 `json:"days"` // Number of days of inactivity before the account will be flagged for deletion; 30-366 days
}

// MessageType return the string telegram-type of AccountTtl
func (accountTtl *AccountTtl) MessageType() string {
	return "accountTtl"
}

// NewAccountTtl creates a new AccountTtl
//
// @param days Number of days of inactivity before the account will be flagged for deletion; 30-366 days
func NewAccountTtl(days int32) *AccountTtl {
	accountTtlTemp := AccountTtl{
		tdCommon: tdCommon{Type: "accountTtl"},
		Days:     days,
	}

	return &accountTtlTemp
}

// GetAccountTtl Returns the period of inactivity after which the account of the current user will automatically be deleted
func (client *Client) GetAccountTtl() (*AccountTtl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAccountTtl",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountTtl AccountTtl
	err = json.Unmarshal(result.Raw, &accountTtl)
	return &accountTtl, err
}
