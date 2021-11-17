package tdlib

import (
	"encoding/json"
	"fmt"
)

// PaymentReceipt Contains information about a successful payment
type PaymentReceipt struct {
	tdCommon
	Title                  string          `json:"title"`                     // Product title
	Description            string          `json:"description"`               // Product description
	Photo                  *Photo          `json:"photo"`                     // Product photo; may be null
	Date                   int32           `json:"date"`                      // Point in time (Unix timestamp) when the payment was made
	SellerBotUserId        int64           `json:"seller_bot_user_id"`        // User identifier of the seller bot
	PaymentsProviderUserId int64           `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	Invoice                *Invoice        `json:"invoice"`                   // Information about the invoice
	OrderInfo              *OrderInfo      `json:"order_info"`                // Order information; may be null
	ShippingOption         *ShippingOption `json:"shipping_option"`           // Chosen shipping option; may be null
	CredentialsTitle       string          `json:"credentials_title"`         // Title of the saved credentials chosen by the buyer
	TipAmount              int64           `json:"tip_amount"`                // The amount of tip chosen by the buyer in the smallest units of the currency
}

// MessageType return the string telegram-type of PaymentReceipt
func (paymentReceipt *PaymentReceipt) MessageType() string {
	return "paymentReceipt"
}

// NewPaymentReceipt creates a new PaymentReceipt
//
// @param title Product title
// @param description Product description
// @param photo Product photo; may be null
// @param date Point in time (Unix timestamp) when the payment was made
// @param sellerBotUserId User identifier of the seller bot
// @param paymentsProviderUserId User identifier of the payment provider bot
// @param invoice Information about the invoice
// @param orderInfo Order information; may be null
// @param shippingOption Chosen shipping option; may be null
// @param credentialsTitle Title of the saved credentials chosen by the buyer
// @param tipAmount The amount of tip chosen by the buyer in the smallest units of the currency
func NewPaymentReceipt(title string, description string, photo *Photo, date int32, sellerBotUserId int64, paymentsProviderUserId int64, invoice *Invoice, orderInfo *OrderInfo, shippingOption *ShippingOption, credentialsTitle string, tipAmount int64) *PaymentReceipt {
	paymentReceiptTemp := PaymentReceipt{
		tdCommon:               tdCommon{Type: "paymentReceipt"},
		Title:                  title,
		Description:            description,
		Photo:                  photo,
		Date:                   date,
		SellerBotUserId:        sellerBotUserId,
		PaymentsProviderUserId: paymentsProviderUserId,
		Invoice:                invoice,
		OrderInfo:              orderInfo,
		ShippingOption:         shippingOption,
		CredentialsTitle:       credentialsTitle,
		TipAmount:              tipAmount,
	}

	return &paymentReceiptTemp
}

// GetPaymentReceipt Returns information about a successful payment
// @param chatId Chat identifier of the PaymentSuccessful message
// @param messageId Message identifier
func (client *Client) GetPaymentReceipt(chatId int64, messageId int64) (*PaymentReceipt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentReceipt",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentReceipt PaymentReceipt
	err = json.Unmarshal(result.Raw, &paymentReceipt)
	return &paymentReceipt, err
}
