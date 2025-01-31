package tdlib

import (
	"encoding/json"
	"fmt"
)

// RecommendedChatFilters Contains a list of recommended chat filters
type RecommendedChatFilters struct {
	tdCommon
	ChatFilters []RecommendedChatFilter `json:"chat_filters"` // List of recommended chat filters
}

// MessageType return the string telegram-type of RecommendedChatFilters
func (recommendedChatFilters *RecommendedChatFilters) MessageType() string {
	return "recommendedChatFilters"
}

// NewRecommendedChatFilters creates a new RecommendedChatFilters
//
// @param chatFilters List of recommended chat filters
func NewRecommendedChatFilters(chatFilters []RecommendedChatFilter) *RecommendedChatFilters {
	recommendedChatFiltersTemp := RecommendedChatFilters{
		tdCommon:    tdCommon{Type: "recommendedChatFilters"},
		ChatFilters: chatFilters,
	}

	return &recommendedChatFiltersTemp
}

// GetRecommendedChatFilters Returns recommended chat filters for the current user
func (client *Client) GetRecommendedChatFilters() (*RecommendedChatFilters, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecommendedChatFilters",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var recommendedChatFilters RecommendedChatFilters
	err = json.Unmarshal(result.Raw, &recommendedChatFilters)
	return &recommendedChatFilters, err
}
