package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatMember Describes a user or a chat as a member of another chat
type ChatMember struct {
	tdCommon
	MemberId       MessageSender    `json:"member_id"`        // Identifier of the chat member. Currently, other chats can be only Left or Banned. Only supergroups and channels can have other chats as Left or Banned members and these chats must be supergroups or channels
	InviterUserId  int64            `json:"inviter_user_id"`  // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
	JoinedChatDate int32            `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined the chat
	Status         ChatMemberStatus `json:"status"`           // Status of the member in the chat
}

// MessageType return the string telegram-type of ChatMember
func (chatMember *ChatMember) MessageType() string {
	return "chatMember"
}

// NewChatMember creates a new ChatMember
//
// @param memberId Identifier of the chat member. Currently, other chats can be only Left or Banned. Only supergroups and channels can have other chats as Left or Banned members and these chats must be supergroups or channels
// @param inviterUserId Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
// @param joinedChatDate Point in time (Unix timestamp) when the user joined the chat
// @param status Status of the member in the chat
func NewChatMember(memberId MessageSender, inviterUserId int64, joinedChatDate int32, status ChatMemberStatus) *ChatMember {
	chatMemberTemp := ChatMember{
		tdCommon:       tdCommon{Type: "chatMember"},
		MemberId:       memberId,
		InviterUserId:  inviterUserId,
		JoinedChatDate: joinedChatDate,
		Status:         status,
	}

	return &chatMemberTemp
}

// UnmarshalJSON unmarshal to json
func (chatMember *ChatMember) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		InviterUserId  int64 `json:"inviter_user_id"`  // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
		JoinedChatDate int32 `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined the chat

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatMember.tdCommon = tempObj.tdCommon
	chatMember.InviterUserId = tempObj.InviterUserId
	chatMember.JoinedChatDate = tempObj.JoinedChatDate

	fieldMemberId, _ := unmarshalMessageSender(objMap["member_id"])
	chatMember.MemberId = fieldMemberId

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	chatMember.Status = fieldStatus

	return nil
}

// GetChatMember Returns information about a single member of a chat
// @param chatId Chat identifier
// @param memberId Member identifier
func (client *Client) GetChatMember(chatId int64, memberId MessageSender) (*ChatMember, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getChatMember",
		"chat_id":   chatId,
		"member_id": memberId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMember ChatMember
	err = json.Unmarshal(result.Raw, &chatMember)
	return &chatMember, err
}
