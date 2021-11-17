package tdlib

import (
	"encoding/json"
	"fmt"
)

// ValidatedOrderInfo Contains a temporary identifier of validated order information, which is stored for one hour. Also contains the available shipping options
type ValidatedOrderInfo struct {
	tdCommon
	OrderInfoId     string           `json:"order_info_id"`    // Temporary identifier of the order information
	ShippingOptions []ShippingOption `json:"shipping_options"` // Available shipping options
}

// MessageType return the string telegram-type of ValidatedOrderInfo
func (validatedOrderInfo *ValidatedOrderInfo) MessageType() string {
	return "validatedOrderInfo"
}

// NewValidatedOrderInfo creates a new ValidatedOrderInfo
//
// @param orderInfoId Temporary identifier of the order information
// @param shippingOptions Available shipping options
func NewValidatedOrderInfo(orderInfoId string, shippingOptions []ShippingOption) *ValidatedOrderInfo {
	validatedOrderInfoTemp := ValidatedOrderInfo{
		tdCommon:        tdCommon{Type: "validatedOrderInfo"},
		OrderInfoId:     orderInfoId,
		ShippingOptions: shippingOptions,
	}

	return &validatedOrderInfoTemp
}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param orderInfo The order information, provided by the user; pass null if empty
// @param allowSave True, if the order information can be saved
func (client *Client) ValidateOrderInfo(chatId int64, messageId int64, orderInfo *OrderInfo, allowSave bool) (*ValidatedOrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "validateOrderInfo",
		"chat_id":    chatId,
		"message_id": messageId,
		"order_info": orderInfo,
		"allow_save": allowSave,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var validatedOrderInfo ValidatedOrderInfo
	err = json.Unmarshal(result.Raw, &validatedOrderInfo)
	return &validatedOrderInfo, err
}
