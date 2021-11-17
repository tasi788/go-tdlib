package tdlib

import (
	"encoding/json"
	"fmt"
)

// PaymentForm Contains information about an invoice payment form
type PaymentForm struct {
	tdCommon
	Id                     JSONInt64               `json:"id"`                        // The payment form identifier
	Invoice                *Invoice                `json:"invoice"`                   // Full information of the invoice
	Url                    string                  `json:"url"`                       // Payment form URL
	SellerBotUserId        int64                   `json:"seller_bot_user_id"`        // User identifier of the seller bot
	PaymentsProviderUserId int64                   `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	PaymentsProvider       *PaymentsProviderStripe `json:"payments_provider"`         // Information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
	SavedOrderInfo         *OrderInfo              `json:"saved_order_info"`          // Saved server-side order information; may be null
	SavedCredentials       *SavedCredentials       `json:"saved_credentials"`         // Information about saved card credentials; may be null
	CanSaveCredentials     bool                    `json:"can_save_credentials"`      // True, if the user can choose to save credentials
	NeedPassword           bool                    `json:"need_password"`             // True, if the user will be able to save credentials protected by a password they set up
}

// MessageType return the string telegram-type of PaymentForm
func (paymentForm *PaymentForm) MessageType() string {
	return "paymentForm"
}

// NewPaymentForm creates a new PaymentForm
//
// @param id The payment form identifier
// @param invoice Full information of the invoice
// @param url Payment form URL
// @param sellerBotUserId User identifier of the seller bot
// @param paymentsProviderUserId User identifier of the payment provider bot
// @param paymentsProvider Information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
// @param savedOrderInfo Saved server-side order information; may be null
// @param savedCredentials Information about saved card credentials; may be null
// @param canSaveCredentials True, if the user can choose to save credentials
// @param needPassword True, if the user will be able to save credentials protected by a password they set up
func NewPaymentForm(id JSONInt64, invoice *Invoice, url string, sellerBotUserId int64, paymentsProviderUserId int64, paymentsProvider *PaymentsProviderStripe, savedOrderInfo *OrderInfo, savedCredentials *SavedCredentials, canSaveCredentials bool, needPassword bool) *PaymentForm {
	paymentFormTemp := PaymentForm{
		tdCommon:               tdCommon{Type: "paymentForm"},
		Id:                     id,
		Invoice:                invoice,
		Url:                    url,
		SellerBotUserId:        sellerBotUserId,
		PaymentsProviderUserId: paymentsProviderUserId,
		PaymentsProvider:       paymentsProvider,
		SavedOrderInfo:         savedOrderInfo,
		SavedCredentials:       savedCredentials,
		CanSaveCredentials:     canSaveCredentials,
		NeedPassword:           needPassword,
	}

	return &paymentFormTemp
}

// GetPaymentForm Returns an invoice payment form. This method must be called when the user presses inlineKeyboardButtonBuy
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param theme Preferred payment form theme; pass null to use the default theme
func (client *Client) GetPaymentForm(chatId int64, messageId int64, theme *PaymentFormTheme) (*PaymentForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentForm",
		"chat_id":    chatId,
		"message_id": messageId,
		"theme":      theme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentForm PaymentForm
	err = json.Unmarshal(result.Raw, &paymentForm)
	return &paymentForm, err
}
