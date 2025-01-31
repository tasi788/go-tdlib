package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageStatistics A detailed statistics about a message
type MessageStatistics struct {
	tdCommon
	MessageInteractionGraph StatisticalGraph `json:"message_interaction_graph"` // A graph containing number of message views and shares
}

// MessageType return the string telegram-type of MessageStatistics
func (messageStatistics *MessageStatistics) MessageType() string {
	return "messageStatistics"
}

// NewMessageStatistics creates a new MessageStatistics
//
// @param messageInteractionGraph A graph containing number of message views and shares
func NewMessageStatistics(messageInteractionGraph StatisticalGraph) *MessageStatistics {
	messageStatisticsTemp := MessageStatistics{
		tdCommon:                tdCommon{Type: "messageStatistics"},
		MessageInteractionGraph: messageInteractionGraph,
	}

	return &messageStatisticsTemp
}

// UnmarshalJSON unmarshal to json
func (messageStatistics *MessageStatistics) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageStatistics.tdCommon = tempObj.tdCommon

	fieldMessageInteractionGraph, _ := unmarshalStatisticalGraph(objMap["message_interaction_graph"])
	messageStatistics.MessageInteractionGraph = fieldMessageInteractionGraph

	return nil
}

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if message.can_get_statistics == true
// @param chatId Chat identifier
// @param messageId Message identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetMessageStatistics(chatId int64, messageId int64, isDark bool) (*MessageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageStatistics",
		"chat_id":    chatId,
		"message_id": messageId,
		"is_dark":    isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageStatistics MessageStatistics
	err = json.Unmarshal(result.Raw, &messageStatistics)
	return &messageStatistics, err
}
