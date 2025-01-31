package tdlib

// RecommendedChatFilter Describes a recommended chat filter
type RecommendedChatFilter struct {
	tdCommon
	Filter      *ChatFilter `json:"filter"`      // The chat filter
	Description string      `json:"description"` // Chat filter description
}

// MessageType return the string telegram-type of RecommendedChatFilter
func (recommendedChatFilter *RecommendedChatFilter) MessageType() string {
	return "recommendedChatFilter"
}

// NewRecommendedChatFilter creates a new RecommendedChatFilter
//
// @param filter The chat filter
// @param description Chat filter description
func NewRecommendedChatFilter(filter *ChatFilter, description string) *RecommendedChatFilter {
	recommendedChatFilterTemp := RecommendedChatFilter{
		tdCommon:    tdCommon{Type: "recommendedChatFilter"},
		Filter:      filter,
		Description: description,
	}

	return &recommendedChatFilterTemp
}
