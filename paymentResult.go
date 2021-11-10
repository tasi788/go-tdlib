package tdlib

import (
	"encoding/json"
	"fmt"
)

// PaymentResult Contains the result of a payment request
type PaymentResult struct {
	tdCommon
	Success         bool   `json:"success"`          // True, if the payment request was successful; otherwise the verification_url will be non-empty
	VerificationUrl string `json:"verification_url"` // URL for additional payment credentials verification
}

// MessageType return the string telegram-type of PaymentResult
func (paymentResult *PaymentResult) MessageType() string {
	return "paymentResult"
}

// NewPaymentResult creates a new PaymentResult
//
// @param success True, if the payment request was successful; otherwise the verification_url will be non-empty
// @param verificationUrl URL for additional payment credentials verification
func NewPaymentResult(success bool, verificationUrl string) *PaymentResult {
	paymentResultTemp := PaymentResult{
		tdCommon:        tdCommon{Type: "paymentResult"},
		Success:         success,
		VerificationUrl: verificationUrl,
	}

	return &paymentResultTemp
}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param paymentFormId Payment form identifier returned by getPaymentForm
// @param orderInfoId Identifier returned by validateOrderInfo, or an empty string
// @param shippingOptionId Identifier of a chosen shipping option, if applicable
// @param credentials The credentials chosen by user for payment
// @param tipAmount Chosen by the user amount of tip in the smallest units of the currency
func (client *Client) SendPaymentForm(chatId int64, messageId int64, paymentFormId JSONInt64, orderInfoId string, shippingOptionId string, credentials InputCredentials, tipAmount int64) (*PaymentResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "sendPaymentForm",
		"chat_id":            chatId,
		"message_id":         messageId,
		"payment_form_id":    paymentFormId,
		"order_info_id":      orderInfoId,
		"shipping_option_id": shippingOptionId,
		"credentials":        credentials,
		"tip_amount":         tipAmount,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentResult PaymentResult
	err = json.Unmarshal(result.Raw, &paymentResult)
	return &paymentResult, err
}
