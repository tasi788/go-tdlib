package tdlib

import (
	"encoding/json"
	"fmt"
)

// ConnectionState Describes the current state of the connection to Telegram servers
type ConnectionState interface {
	GetConnectionStateEnum() ConnectionStateEnum
}

// ConnectionStateEnum Alias for abstract ConnectionState 'Sub-Classes', used as constant-enum here
type ConnectionStateEnum string

// ConnectionState enums
const (
	ConnectionStateWaitingForNetworkType ConnectionStateEnum = "connectionStateWaitingForNetwork"
	ConnectionStateConnectingToProxyType ConnectionStateEnum = "connectionStateConnectingToProxy"
	ConnectionStateConnectingType        ConnectionStateEnum = "connectionStateConnecting"
	ConnectionStateUpdatingType          ConnectionStateEnum = "connectionStateUpdating"
	ConnectionStateReadyType             ConnectionStateEnum = "connectionStateReady"
)

func unmarshalConnectionState(rawMsg *json.RawMessage) (ConnectionState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ConnectionStateEnum(objMap["@type"].(string)) {
	case ConnectionStateWaitingForNetworkType:
		var connectionStateWaitingForNetwork ConnectionStateWaitingForNetwork
		err := json.Unmarshal(*rawMsg, &connectionStateWaitingForNetwork)
		return &connectionStateWaitingForNetwork, err

	case ConnectionStateConnectingToProxyType:
		var connectionStateConnectingToProxy ConnectionStateConnectingToProxy
		err := json.Unmarshal(*rawMsg, &connectionStateConnectingToProxy)
		return &connectionStateConnectingToProxy, err

	case ConnectionStateConnectingType:
		var connectionStateConnecting ConnectionStateConnecting
		err := json.Unmarshal(*rawMsg, &connectionStateConnecting)
		return &connectionStateConnecting, err

	case ConnectionStateUpdatingType:
		var connectionStateUpdating ConnectionStateUpdating
		err := json.Unmarshal(*rawMsg, &connectionStateUpdating)
		return &connectionStateUpdating, err

	case ConnectionStateReadyType:
		var connectionStateReady ConnectionStateReady
		err := json.Unmarshal(*rawMsg, &connectionStateReady)
		return &connectionStateReady, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ConnectionStateWaitingForNetwork Currently waiting for the network to become available. Use setNetworkType to change the available network type
type ConnectionStateWaitingForNetwork struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateWaitingForNetwork
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) MessageType() string {
	return "connectionStateWaitingForNetwork"
}

// NewConnectionStateWaitingForNetwork creates a new ConnectionStateWaitingForNetwork
//
func NewConnectionStateWaitingForNetwork() *ConnectionStateWaitingForNetwork {
	connectionStateWaitingForNetworkTemp := ConnectionStateWaitingForNetwork{
		tdCommon: tdCommon{Type: "connectionStateWaitingForNetwork"},
	}

	return &connectionStateWaitingForNetworkTemp
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateWaitingForNetworkType
}

// ConnectionStateConnectingToProxy Currently establishing a connection with a proxy server
type ConnectionStateConnectingToProxy struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateConnectingToProxy
func (connectionStateConnectingToProxy *ConnectionStateConnectingToProxy) MessageType() string {
	return "connectionStateConnectingToProxy"
}

// NewConnectionStateConnectingToProxy creates a new ConnectionStateConnectingToProxy
//
func NewConnectionStateConnectingToProxy() *ConnectionStateConnectingToProxy {
	connectionStateConnectingToProxyTemp := ConnectionStateConnectingToProxy{
		tdCommon: tdCommon{Type: "connectionStateConnectingToProxy"},
	}

	return &connectionStateConnectingToProxyTemp
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateConnectingToProxy *ConnectionStateConnectingToProxy) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateConnectingToProxyType
}

// ConnectionStateConnecting Currently establishing a connection to the Telegram servers
type ConnectionStateConnecting struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateConnecting
func (connectionStateConnecting *ConnectionStateConnecting) MessageType() string {
	return "connectionStateConnecting"
}

// NewConnectionStateConnecting creates a new ConnectionStateConnecting
//
func NewConnectionStateConnecting() *ConnectionStateConnecting {
	connectionStateConnectingTemp := ConnectionStateConnecting{
		tdCommon: tdCommon{Type: "connectionStateConnecting"},
	}

	return &connectionStateConnectingTemp
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateConnecting *ConnectionStateConnecting) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateConnectingType
}

// ConnectionStateUpdating Downloading data received while the application was offline
type ConnectionStateUpdating struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateUpdating
func (connectionStateUpdating *ConnectionStateUpdating) MessageType() string {
	return "connectionStateUpdating"
}

// NewConnectionStateUpdating creates a new ConnectionStateUpdating
//
func NewConnectionStateUpdating() *ConnectionStateUpdating {
	connectionStateUpdatingTemp := ConnectionStateUpdating{
		tdCommon: tdCommon{Type: "connectionStateUpdating"},
	}

	return &connectionStateUpdatingTemp
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateUpdating *ConnectionStateUpdating) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateUpdatingType
}

// ConnectionStateReady There is a working connection to the Telegram servers
type ConnectionStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateReady
func (connectionStateReady *ConnectionStateReady) MessageType() string {
	return "connectionStateReady"
}

// NewConnectionStateReady creates a new ConnectionStateReady
//
func NewConnectionStateReady() *ConnectionStateReady {
	connectionStateReadyTemp := ConnectionStateReady{
		tdCommon: tdCommon{Type: "connectionStateReady"},
	}

	return &connectionStateReadyTemp
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateReady *ConnectionStateReady) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateReadyType
}
