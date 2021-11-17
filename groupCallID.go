package tdlib

import (
	"encoding/json"
	"fmt"
)

// GroupCallId Contains the group call identifier
type GroupCallId struct {
	tdCommon
	Id int32 `json:"id"` // Group call identifier
}

// MessageType return the string telegram-type of GroupCallId
func (groupCallId *GroupCallId) MessageType() string {
	return "groupCallId"
}

// NewGroupCallId creates a new GroupCallId
//
// @param id Group call identifier
func NewGroupCallId(id int32) *GroupCallId {
	groupCallIdTemp := GroupCallId{
		tdCommon: tdCommon{Type: "groupCallId"},
		Id:       id,
	}

	return &groupCallIdTemp
}

// CreateVideoChat Creates a video chat (a group call bound to a chat). Available only for basic groups, supergroups and channels; requires can_manage_video_chats rights
// @param chatId Chat identifier, in which the video chat will be created
// @param title Group call title; if empty, chat title will be used
// @param startDate Point in time (Unix timestamp) when the group call is supposed to be started by an administrator; 0 to start the video chat immediately. The date must be at least 10 seconds and at most 8 days in the future
func (client *Client) CreateVideoChat(chatId int64, title string, startDate int32) (*GroupCallId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "createVideoChat",
		"chat_id":    chatId,
		"title":      title,
		"start_date": startDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallId GroupCallId
	err = json.Unmarshal(result.Raw, &groupCallId)
	return &groupCallId, err
}
