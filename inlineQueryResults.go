package tdlib

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResults Represents the results of the inline query. Use sendInlineQueryResultMessage to send the result of the query
type InlineQueryResults struct {
	tdCommon
	InlineQueryId     JSONInt64           `json:"inline_query_id"`     // Unique identifier of the inline query
	NextOffset        string              `json:"next_offset"`         // The offset for the next request. If empty, there are no more results
	Results           []InlineQueryResult `json:"results"`             // Results of the query
	SwitchPmText      string              `json:"switch_pm_text"`      // If non-empty, this text must be shown on the button, which opens a private chat with the bot and sends the bot a start message with the switch_pm_parameter
	SwitchPmParameter string              `json:"switch_pm_parameter"` // Parameter for the bot start message
}

// MessageType return the string telegram-type of InlineQueryResults
func (inlineQueryResults *InlineQueryResults) MessageType() string {
	return "inlineQueryResults"
}

// NewInlineQueryResults creates a new InlineQueryResults
//
// @param inlineQueryId Unique identifier of the inline query
// @param nextOffset The offset for the next request. If empty, there are no more results
// @param results Results of the query
// @param switchPmText If non-empty, this text must be shown on the button, which opens a private chat with the bot and sends the bot a start message with the switch_pm_parameter
// @param switchPmParameter Parameter for the bot start message
func NewInlineQueryResults(inlineQueryId JSONInt64, nextOffset string, results []InlineQueryResult, switchPmText string, switchPmParameter string) *InlineQueryResults {
	inlineQueryResultsTemp := InlineQueryResults{
		tdCommon:          tdCommon{Type: "inlineQueryResults"},
		InlineQueryId:     inlineQueryId,
		NextOffset:        nextOffset,
		Results:           results,
		SwitchPmText:      switchPmText,
		SwitchPmParameter: switchPmParameter,
	}

	return &inlineQueryResultsTemp
}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param botUserId The identifier of the target bot
// @param chatId Identifier of the chat where the query was sent
// @param userLocation Location of the user; pass null if unknown or the bot doesn't need user's location
// @param query Text of the query
// @param offset Offset of the first entry to return
func (client *Client) GetInlineQueryResults(botUserId int64, chatId int64, userLocation *Location, query string, offset string) (*InlineQueryResults, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getInlineQueryResults",
		"bot_user_id":   botUserId,
		"chat_id":       chatId,
		"user_location": userLocation,
		"query":         query,
		"offset":        offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var inlineQueryResults InlineQueryResults
	err = json.Unmarshal(result.Raw, &inlineQueryResults)
	return &inlineQueryResults, err
}
