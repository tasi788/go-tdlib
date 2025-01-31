package tdlib

import (
	"encoding/json"
	"fmt"
)

// CallProblem Describes the exact type of a problem with a call
type CallProblem interface {
	GetCallProblemEnum() CallProblemEnum
}

// CallProblemEnum Alias for abstract CallProblem 'Sub-Classes', used as constant-enum here
type CallProblemEnum string

// CallProblem enums
const (
	CallProblemEchoType            CallProblemEnum = "callProblemEcho"
	CallProblemNoiseType           CallProblemEnum = "callProblemNoise"
	CallProblemInterruptionsType   CallProblemEnum = "callProblemInterruptions"
	CallProblemDistortedSpeechType CallProblemEnum = "callProblemDistortedSpeech"
	CallProblemSilentLocalType     CallProblemEnum = "callProblemSilentLocal"
	CallProblemSilentRemoteType    CallProblemEnum = "callProblemSilentRemote"
	CallProblemDroppedType         CallProblemEnum = "callProblemDropped"
	CallProblemDistortedVideoType  CallProblemEnum = "callProblemDistortedVideo"
	CallProblemPixelatedVideoType  CallProblemEnum = "callProblemPixelatedVideo"
)

func unmarshalCallProblem(rawMsg *json.RawMessage) (CallProblem, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallProblemEnum(objMap["@type"].(string)) {
	case CallProblemEchoType:
		var callProblemEcho CallProblemEcho
		err := json.Unmarshal(*rawMsg, &callProblemEcho)
		return &callProblemEcho, err

	case CallProblemNoiseType:
		var callProblemNoise CallProblemNoise
		err := json.Unmarshal(*rawMsg, &callProblemNoise)
		return &callProblemNoise, err

	case CallProblemInterruptionsType:
		var callProblemInterruptions CallProblemInterruptions
		err := json.Unmarshal(*rawMsg, &callProblemInterruptions)
		return &callProblemInterruptions, err

	case CallProblemDistortedSpeechType:
		var callProblemDistortedSpeech CallProblemDistortedSpeech
		err := json.Unmarshal(*rawMsg, &callProblemDistortedSpeech)
		return &callProblemDistortedSpeech, err

	case CallProblemSilentLocalType:
		var callProblemSilentLocal CallProblemSilentLocal
		err := json.Unmarshal(*rawMsg, &callProblemSilentLocal)
		return &callProblemSilentLocal, err

	case CallProblemSilentRemoteType:
		var callProblemSilentRemote CallProblemSilentRemote
		err := json.Unmarshal(*rawMsg, &callProblemSilentRemote)
		return &callProblemSilentRemote, err

	case CallProblemDroppedType:
		var callProblemDropped CallProblemDropped
		err := json.Unmarshal(*rawMsg, &callProblemDropped)
		return &callProblemDropped, err

	case CallProblemDistortedVideoType:
		var callProblemDistortedVideo CallProblemDistortedVideo
		err := json.Unmarshal(*rawMsg, &callProblemDistortedVideo)
		return &callProblemDistortedVideo, err

	case CallProblemPixelatedVideoType:
		var callProblemPixelatedVideo CallProblemPixelatedVideo
		err := json.Unmarshal(*rawMsg, &callProblemPixelatedVideo)
		return &callProblemPixelatedVideo, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// CallProblemEcho The user heard their own voice
type CallProblemEcho struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemEcho
func (callProblemEcho *CallProblemEcho) MessageType() string {
	return "callProblemEcho"
}

// NewCallProblemEcho creates a new CallProblemEcho
//
func NewCallProblemEcho() *CallProblemEcho {
	callProblemEchoTemp := CallProblemEcho{
		tdCommon: tdCommon{Type: "callProblemEcho"},
	}

	return &callProblemEchoTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemEcho *CallProblemEcho) GetCallProblemEnum() CallProblemEnum {
	return CallProblemEchoType
}

// CallProblemNoise The user heard background noise
type CallProblemNoise struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemNoise
func (callProblemNoise *CallProblemNoise) MessageType() string {
	return "callProblemNoise"
}

// NewCallProblemNoise creates a new CallProblemNoise
//
func NewCallProblemNoise() *CallProblemNoise {
	callProblemNoiseTemp := CallProblemNoise{
		tdCommon: tdCommon{Type: "callProblemNoise"},
	}

	return &callProblemNoiseTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemNoise *CallProblemNoise) GetCallProblemEnum() CallProblemEnum {
	return CallProblemNoiseType
}

// CallProblemInterruptions The other side kept disappearing
type CallProblemInterruptions struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemInterruptions
func (callProblemInterruptions *CallProblemInterruptions) MessageType() string {
	return "callProblemInterruptions"
}

// NewCallProblemInterruptions creates a new CallProblemInterruptions
//
func NewCallProblemInterruptions() *CallProblemInterruptions {
	callProblemInterruptionsTemp := CallProblemInterruptions{
		tdCommon: tdCommon{Type: "callProblemInterruptions"},
	}

	return &callProblemInterruptionsTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemInterruptions *CallProblemInterruptions) GetCallProblemEnum() CallProblemEnum {
	return CallProblemInterruptionsType
}

// CallProblemDistortedSpeech The speech was distorted
type CallProblemDistortedSpeech struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemDistortedSpeech
func (callProblemDistortedSpeech *CallProblemDistortedSpeech) MessageType() string {
	return "callProblemDistortedSpeech"
}

// NewCallProblemDistortedSpeech creates a new CallProblemDistortedSpeech
//
func NewCallProblemDistortedSpeech() *CallProblemDistortedSpeech {
	callProblemDistortedSpeechTemp := CallProblemDistortedSpeech{
		tdCommon: tdCommon{Type: "callProblemDistortedSpeech"},
	}

	return &callProblemDistortedSpeechTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemDistortedSpeech *CallProblemDistortedSpeech) GetCallProblemEnum() CallProblemEnum {
	return CallProblemDistortedSpeechType
}

// CallProblemSilentLocal The user couldn't hear the other side
type CallProblemSilentLocal struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemSilentLocal
func (callProblemSilentLocal *CallProblemSilentLocal) MessageType() string {
	return "callProblemSilentLocal"
}

// NewCallProblemSilentLocal creates a new CallProblemSilentLocal
//
func NewCallProblemSilentLocal() *CallProblemSilentLocal {
	callProblemSilentLocalTemp := CallProblemSilentLocal{
		tdCommon: tdCommon{Type: "callProblemSilentLocal"},
	}

	return &callProblemSilentLocalTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemSilentLocal *CallProblemSilentLocal) GetCallProblemEnum() CallProblemEnum {
	return CallProblemSilentLocalType
}

// CallProblemSilentRemote The other side couldn't hear the user
type CallProblemSilentRemote struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemSilentRemote
func (callProblemSilentRemote *CallProblemSilentRemote) MessageType() string {
	return "callProblemSilentRemote"
}

// NewCallProblemSilentRemote creates a new CallProblemSilentRemote
//
func NewCallProblemSilentRemote() *CallProblemSilentRemote {
	callProblemSilentRemoteTemp := CallProblemSilentRemote{
		tdCommon: tdCommon{Type: "callProblemSilentRemote"},
	}

	return &callProblemSilentRemoteTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemSilentRemote *CallProblemSilentRemote) GetCallProblemEnum() CallProblemEnum {
	return CallProblemSilentRemoteType
}

// CallProblemDropped The call ended unexpectedly
type CallProblemDropped struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemDropped
func (callProblemDropped *CallProblemDropped) MessageType() string {
	return "callProblemDropped"
}

// NewCallProblemDropped creates a new CallProblemDropped
//
func NewCallProblemDropped() *CallProblemDropped {
	callProblemDroppedTemp := CallProblemDropped{
		tdCommon: tdCommon{Type: "callProblemDropped"},
	}

	return &callProblemDroppedTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemDropped *CallProblemDropped) GetCallProblemEnum() CallProblemEnum {
	return CallProblemDroppedType
}

// CallProblemDistortedVideo The video was distorted
type CallProblemDistortedVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemDistortedVideo
func (callProblemDistortedVideo *CallProblemDistortedVideo) MessageType() string {
	return "callProblemDistortedVideo"
}

// NewCallProblemDistortedVideo creates a new CallProblemDistortedVideo
//
func NewCallProblemDistortedVideo() *CallProblemDistortedVideo {
	callProblemDistortedVideoTemp := CallProblemDistortedVideo{
		tdCommon: tdCommon{Type: "callProblemDistortedVideo"},
	}

	return &callProblemDistortedVideoTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemDistortedVideo *CallProblemDistortedVideo) GetCallProblemEnum() CallProblemEnum {
	return CallProblemDistortedVideoType
}

// CallProblemPixelatedVideo The video was pixelated
type CallProblemPixelatedVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of CallProblemPixelatedVideo
func (callProblemPixelatedVideo *CallProblemPixelatedVideo) MessageType() string {
	return "callProblemPixelatedVideo"
}

// NewCallProblemPixelatedVideo creates a new CallProblemPixelatedVideo
//
func NewCallProblemPixelatedVideo() *CallProblemPixelatedVideo {
	callProblemPixelatedVideoTemp := CallProblemPixelatedVideo{
		tdCommon: tdCommon{Type: "callProblemPixelatedVideo"},
	}

	return &callProblemPixelatedVideoTemp
}

// GetCallProblemEnum return the enum type of this object
func (callProblemPixelatedVideo *CallProblemPixelatedVideo) GetCallProblemEnum() CallProblemEnum {
	return CallProblemPixelatedVideoType
}
