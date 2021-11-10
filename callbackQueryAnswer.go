package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallbackQueryAnswer Contains a bot's answer to a callback query
type CallbackQueryAnswer struct {
	tdCommon
	Text      string `json:"text"`       // Text of the answer
	ShowAlert bool   `json:"show_alert"` // True, if an alert must be shown to the user instead of a toast notification
	Url       string `json:"url"`        // URL to be opened
}

// MessageType return the string telegram-type of CallbackQueryAnswer
func (callbackQueryAnswer *CallbackQueryAnswer) MessageType() string {
	return "callbackQueryAnswer"
}

// NewCallbackQueryAnswer creates a new CallbackQueryAnswer
//
// @param text Text of the answer
// @param showAlert True, if an alert must be shown to the user instead of a toast notification
// @param url URL to be opened
func NewCallbackQueryAnswer(text string, showAlert bool, url string) *CallbackQueryAnswer {
	callbackQueryAnswerTemp := CallbackQueryAnswer{
		tdCommon:  tdCommon{Type: "callbackQueryAnswer"},
		Text:      text,
		ShowAlert: showAlert,
		Url:       url,
	}

	return &callbackQueryAnswerTemp
}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param chatId Identifier of the chat with the message
// @param messageId Identifier of the message from which the query originated
// @param payload Query payload
func (client *Client) GetCallbackQueryAnswer(chatId int64, messageId int64, payload CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"payload":    payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callbackQueryAnswer CallbackQueryAnswer
	err = json.Unmarshal(result.Raw, &callbackQueryAnswer)
	return &callbackQueryAnswer, err
}
