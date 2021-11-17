package tdlib

import (
	"encoding/json"
	"fmt"
)

// PushReceiverId Contains a globally unique push receiver identifier, which can be used to identify which account has received a push notification
type PushReceiverId struct {
	tdCommon
	Id JSONInt64 `json:"id"` // The globally unique identifier of push notification subscription
}

// MessageType return the string telegram-type of PushReceiverId
func (pushReceiverId *PushReceiverId) MessageType() string {
	return "pushReceiverId"
}

// NewPushReceiverId creates a new PushReceiverId
//
// @param id The globally unique identifier of push notification subscription
func NewPushReceiverId(id JSONInt64) *PushReceiverId {
	pushReceiverIdTemp := PushReceiverId{
		tdCommon: tdCommon{Type: "pushReceiverId"},
		Id:       id,
	}

	return &pushReceiverIdTemp
}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription
// @param deviceToken Device token
// @param otherUserIds List of user identifiers of other users currently using the application
func (client *Client) RegisterDevice(deviceToken DeviceToken, otherUserIds []int64) (*PushReceiverId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverId PushReceiverId
	err = json.Unmarshal(result.Raw, &pushReceiverId)
	return &pushReceiverId, err
}

// GetPushReceiverId Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously
// @param payload JSON-encoded push notification payload
func (client *Client) GetPushReceiverId(payload string) (*PushReceiverId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getPushReceiverId",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverId PushReceiverId
	err = json.Unmarshal(result.Raw, &pushReceiverId)
	return &pushReceiverId, err
}
