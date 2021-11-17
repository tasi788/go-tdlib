package tdlib

import (
	"encoding/json"
	"fmt"
)

// SecretChat Represents a secret chat
type SecretChat struct {
	tdCommon
	Id         int32           `json:"id"`          // Secret chat identifier
	UserId     int64           `json:"user_id"`     // Identifier of the chat partner
	State      SecretChatState `json:"state"`       // State of the secret chat
	IsOutbound bool            `json:"is_outbound"` // True, if the chat was created by the current user; otherwise false
	KeyHash    []byte          `json:"key_hash"`    // Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 little-endian bytes, which must be split into groups of 2 bits, each denoting a pixel of one of 4 colors FFFFFF, D5E6F3, 2D5775, and 2F99C9. The pixels must be used to make a 12x12 square image filled from left to right, top to bottom. Alternatively, the first 32 bytes of the hash can be converted to the hexadecimal format and printed as 32 2-digit hex numbers
	Layer      int32           `json:"layer"`       // Secret chat layer; determines features supported by the chat partner's application. Nested text entities and underline and strikethrough entities are supported if the layer >= 101
}

// MessageType return the string telegram-type of SecretChat
func (secretChat *SecretChat) MessageType() string {
	return "secretChat"
}

// NewSecretChat creates a new SecretChat
//
// @param id Secret chat identifier
// @param userId Identifier of the chat partner
// @param state State of the secret chat
// @param isOutbound True, if the chat was created by the current user; otherwise false
// @param keyHash Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 little-endian bytes, which must be split into groups of 2 bits, each denoting a pixel of one of 4 colors FFFFFF, D5E6F3, 2D5775, and 2F99C9. The pixels must be used to make a 12x12 square image filled from left to right, top to bottom. Alternatively, the first 32 bytes of the hash can be converted to the hexadecimal format and printed as 32 2-digit hex numbers
// @param layer Secret chat layer; determines features supported by the chat partner's application. Nested text entities and underline and strikethrough entities are supported if the layer >= 101
func NewSecretChat(id int32, userId int64, state SecretChatState, isOutbound bool, keyHash []byte, layer int32) *SecretChat {
	secretChatTemp := SecretChat{
		tdCommon:   tdCommon{Type: "secretChat"},
		Id:         id,
		UserId:     userId,
		State:      state,
		IsOutbound: isOutbound,
		KeyHash:    keyHash,
		Layer:      layer,
	}

	return &secretChatTemp
}

// UnmarshalJSON unmarshal to json
func (secretChat *SecretChat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id         int32  `json:"id"`          // Secret chat identifier
		UserId     int64  `json:"user_id"`     // Identifier of the chat partner
		IsOutbound bool   `json:"is_outbound"` // True, if the chat was created by the current user; otherwise false
		KeyHash    []byte `json:"key_hash"`    // Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 little-endian bytes, which must be split into groups of 2 bits, each denoting a pixel of one of 4 colors FFFFFF, D5E6F3, 2D5775, and 2F99C9. The pixels must be used to make a 12x12 square image filled from left to right, top to bottom. Alternatively, the first 32 bytes of the hash can be converted to the hexadecimal format and printed as 32 2-digit hex numbers
		Layer      int32  `json:"layer"`       // Secret chat layer; determines features supported by the chat partner's application. Nested text entities and underline and strikethrough entities are supported if the layer >= 101
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	secretChat.tdCommon = tempObj.tdCommon
	secretChat.Id = tempObj.Id
	secretChat.UserId = tempObj.UserId
	secretChat.IsOutbound = tempObj.IsOutbound
	secretChat.KeyHash = tempObj.KeyHash
	secretChat.Layer = tempObj.Layer

	fieldState, _ := unmarshalSecretChatState(objMap["state"])
	secretChat.State = fieldState

	return nil
}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline request
// @param secretChatId Secret chat identifier
func (client *Client) GetSecretChat(secretChatId int32) (*SecretChat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var secretChatDummy SecretChat
	err = json.Unmarshal(result.Raw, &secretChatDummy)
	return &secretChatDummy, err
}
