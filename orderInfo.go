package tdlib

import (
	"encoding/json"
	"fmt"
)

// OrderInfo Order information
type OrderInfo struct {
	tdCommon
	Name            string   `json:"name"`             // Name of the user
	PhoneNumber     string   `json:"phone_number"`     // Phone number of the user
	EmailAddress    string   `json:"email_address"`    // Email address of the user
	ShippingAddress *Address `json:"shipping_address"` // Shipping address for this order; may be null
}

// MessageType return the string telegram-type of OrderInfo
func (orderInfo *OrderInfo) MessageType() string {
	return "orderInfo"
}

// NewOrderInfo creates a new OrderInfo
//
// @param name Name of the user
// @param phoneNumber Phone number of the user
// @param emailAddress Email address of the user
// @param shippingAddress Shipping address for this order; may be null
func NewOrderInfo(name string, phoneNumber string, emailAddress string, shippingAddress *Address) *OrderInfo {
	orderInfoTemp := OrderInfo{
		tdCommon:        tdCommon{Type: "orderInfo"},
		Name:            name,
		PhoneNumber:     phoneNumber,
		EmailAddress:    emailAddress,
		ShippingAddress: shippingAddress,
	}

	return &orderInfoTemp
}

// GetSavedOrderInfo Returns saved order info, if any
func (client *Client) GetSavedOrderInfo() (*OrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var orderInfo OrderInfo
	err = json.Unmarshal(result.Raw, &orderInfo)
	return &orderInfo, err
}
