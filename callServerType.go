package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallServerType Describes the type of a call server
type CallServerType interface {
	GetCallServerTypeEnum() CallServerTypeEnum
}

// CallServerTypeEnum Alias for abstract CallServerType 'Sub-Classes', used as constant-enum here
type CallServerTypeEnum string

// CallServerType enums
const (
	CallServerTypeTelegramReflectorType CallServerTypeEnum = "callServerTypeTelegramReflector"
	CallServerTypeWebrtcType            CallServerTypeEnum = "callServerTypeWebrtc"
)

func unmarshalCallServerType(rawMsg *json.RawMessage) (CallServerType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallServerTypeEnum(objMap["@type"].(string)) {
	case CallServerTypeTelegramReflectorType:
		var callServerTypeTelegramReflector CallServerTypeTelegramReflector
		err := json.Unmarshal(*rawMsg, &callServerTypeTelegramReflector)
		return &callServerTypeTelegramReflector, err

	case CallServerTypeWebrtcType:
		var callServerTypeWebrtc CallServerTypeWebrtc
		err := json.Unmarshal(*rawMsg, &callServerTypeWebrtc)
		return &callServerTypeWebrtc, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CallServerTypeTelegramReflector A Telegram call reflector
type CallServerTypeTelegramReflector struct {
	tdCommon
	PeerTag []byte `json:"peer_tag"` // A peer tag to be used with the reflector
}

// MessageType return the string telegram-type of CallServerTypeTelegramReflector
func (callServerTypeTelegramReflector *CallServerTypeTelegramReflector) MessageType() string {
	return "callServerTypeTelegramReflector"
}

// NewCallServerTypeTelegramReflector creates a new CallServerTypeTelegramReflector
//
// @param peerTag A peer tag to be used with the reflector
func NewCallServerTypeTelegramReflector(peerTag []byte) *CallServerTypeTelegramReflector {
	callServerTypeTelegramReflectorTemp := CallServerTypeTelegramReflector{
		tdCommon: tdCommon{Type: "callServerTypeTelegramReflector"},
		PeerTag:  peerTag,
	}

	return &callServerTypeTelegramReflectorTemp
}

// GetCallServerTypeEnum return the enum type of this object
func (callServerTypeTelegramReflector *CallServerTypeTelegramReflector) GetCallServerTypeEnum() CallServerTypeEnum {
	return CallServerTypeTelegramReflectorType
}

// CallServerTypeWebrtc A WebRTC server
type CallServerTypeWebrtc struct {
	tdCommon
	Username     string `json:"username"`      // Username to be used for authentication
	Password     string `json:"password"`      // Authentication password
	SupportsTurn bool   `json:"supports_turn"` // True, if the server supports TURN
	SupportsStun bool   `json:"supports_stun"` // True, if the server supports STUN
}

// MessageType return the string telegram-type of CallServerTypeWebrtc
func (callServerTypeWebrtc *CallServerTypeWebrtc) MessageType() string {
	return "callServerTypeWebrtc"
}

// NewCallServerTypeWebrtc creates a new CallServerTypeWebrtc
//
// @param username Username to be used for authentication
// @param password Authentication password
// @param supportsTurn True, if the server supports TURN
// @param supportsStun True, if the server supports STUN
func NewCallServerTypeWebrtc(username string, password string, supportsTurn bool, supportsStun bool) *CallServerTypeWebrtc {
	callServerTypeWebrtcTemp := CallServerTypeWebrtc{
		tdCommon:     tdCommon{Type: "callServerTypeWebrtc"},
		Username:     username,
		Password:     password,
		SupportsTurn: supportsTurn,
		SupportsStun: supportsStun,
	}

	return &callServerTypeWebrtcTemp
}

// GetCallServerTypeEnum return the enum type of this object
func (callServerTypeWebrtc *CallServerTypeWebrtc) GetCallServerTypeEnum() CallServerTypeEnum {
	return CallServerTypeWebrtcType
}
