package tdlib

import (
	"encoding/json"
)

// ChatPosition Describes a position of a chat in a chat list
type ChatPosition struct {
	tdCommon
	List     ChatList   `json:"list"`      // The chat list
	Order    JSONInt64  `json:"order"`     // A parameter used to determine order of the chat in the chat list. Chats must be sorted by the pair (order, chat.id) in descending order
	IsPinned bool       `json:"is_pinned"` // True, if the chat is pinned in the chat list
	Source   ChatSource `json:"source"`    // Source of the chat in the chat list; may be null
}

// MessageType return the string telegram-type of ChatPosition
func (chatPosition *ChatPosition) MessageType() string {
	return "chatPosition"
}

// NewChatPosition creates a new ChatPosition
//
// @param list The chat list
// @param order A parameter used to determine order of the chat in the chat list. Chats must be sorted by the pair (order, chat.id) in descending order
// @param isPinned True, if the chat is pinned in the chat list
// @param source Source of the chat in the chat list; may be null
func NewChatPosition(list ChatList, order JSONInt64, isPinned bool, source ChatSource) *ChatPosition {
	chatPositionTemp := ChatPosition{
		tdCommon: tdCommon{Type: "chatPosition"},
		List:     list,
		Order:    order,
		IsPinned: isPinned,
		Source:   source,
	}

	return &chatPositionTemp
}

// UnmarshalJSON unmarshal to json
func (chatPosition *ChatPosition) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Order    JSONInt64 `json:"order"`     // A parameter used to determine order of the chat in the chat list. Chats must be sorted by the pair (order, chat.id) in descending order
		IsPinned bool      `json:"is_pinned"` // True, if the chat is pinned in the chat list

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatPosition.tdCommon = tempObj.tdCommon
	chatPosition.Order = tempObj.Order
	chatPosition.IsPinned = tempObj.IsPinned

	fieldList, _ := unmarshalChatList(objMap["list"])
	chatPosition.List = fieldList

	fieldSource, _ := unmarshalChatSource(objMap["source"])
	chatPosition.Source = fieldSource

	return nil
}
