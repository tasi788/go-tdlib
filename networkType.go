package tdlib

import (
	"encoding/json"
	"fmt"
)

// NetworkType Represents the type of a network
type NetworkType interface {
	GetNetworkTypeEnum() NetworkTypeEnum
}

// NetworkTypeEnum Alias for abstract NetworkType 'Sub-Classes', used as constant-enum here
type NetworkTypeEnum string

// NetworkType enums
const (
	NetworkTypeNoneType          NetworkTypeEnum = "networkTypeNone"
	NetworkTypeMobileType        NetworkTypeEnum = "networkTypeMobile"
	NetworkTypeMobileRoamingType NetworkTypeEnum = "networkTypeMobileRoaming"
	NetworkTypeWiFiType          NetworkTypeEnum = "networkTypeWiFi"
	NetworkTypeOtherType         NetworkTypeEnum = "networkTypeOther"
)

func unmarshalNetworkType(rawMsg *json.RawMessage) (NetworkType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NetworkTypeEnum(objMap["@type"].(string)) {
	case NetworkTypeNoneType:
		var networkTypeNone NetworkTypeNone
		err := json.Unmarshal(*rawMsg, &networkTypeNone)
		return &networkTypeNone, err

	case NetworkTypeMobileType:
		var networkTypeMobile NetworkTypeMobile
		err := json.Unmarshal(*rawMsg, &networkTypeMobile)
		return &networkTypeMobile, err

	case NetworkTypeMobileRoamingType:
		var networkTypeMobileRoaming NetworkTypeMobileRoaming
		err := json.Unmarshal(*rawMsg, &networkTypeMobileRoaming)
		return &networkTypeMobileRoaming, err

	case NetworkTypeWiFiType:
		var networkTypeWiFi NetworkTypeWiFi
		err := json.Unmarshal(*rawMsg, &networkTypeWiFi)
		return &networkTypeWiFi, err

	case NetworkTypeOtherType:
		var networkTypeOther NetworkTypeOther
		err := json.Unmarshal(*rawMsg, &networkTypeOther)
		return &networkTypeOther, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// NetworkTypeNone The network is not available
type NetworkTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeNone
func (networkTypeNone *NetworkTypeNone) MessageType() string {
	return "networkTypeNone"
}

// NewNetworkTypeNone creates a new NetworkTypeNone
//
func NewNetworkTypeNone() *NetworkTypeNone {
	networkTypeNoneTemp := NetworkTypeNone{
		tdCommon: tdCommon{Type: "networkTypeNone"},
	}

	return &networkTypeNoneTemp
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeNone *NetworkTypeNone) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeNoneType
}

// NetworkTypeMobile A mobile network
type NetworkTypeMobile struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeMobile
func (networkTypeMobile *NetworkTypeMobile) MessageType() string {
	return "networkTypeMobile"
}

// NewNetworkTypeMobile creates a new NetworkTypeMobile
//
func NewNetworkTypeMobile() *NetworkTypeMobile {
	networkTypeMobileTemp := NetworkTypeMobile{
		tdCommon: tdCommon{Type: "networkTypeMobile"},
	}

	return &networkTypeMobileTemp
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeMobile *NetworkTypeMobile) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeMobileType
}

// NetworkTypeMobileRoaming A mobile roaming network
type NetworkTypeMobileRoaming struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeMobileRoaming
func (networkTypeMobileRoaming *NetworkTypeMobileRoaming) MessageType() string {
	return "networkTypeMobileRoaming"
}

// NewNetworkTypeMobileRoaming creates a new NetworkTypeMobileRoaming
//
func NewNetworkTypeMobileRoaming() *NetworkTypeMobileRoaming {
	networkTypeMobileRoamingTemp := NetworkTypeMobileRoaming{
		tdCommon: tdCommon{Type: "networkTypeMobileRoaming"},
	}

	return &networkTypeMobileRoamingTemp
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeMobileRoaming *NetworkTypeMobileRoaming) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeMobileRoamingType
}

// NetworkTypeWiFi A Wi-Fi network
type NetworkTypeWiFi struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeWiFi
func (networkTypeWiFi *NetworkTypeWiFi) MessageType() string {
	return "networkTypeWiFi"
}

// NewNetworkTypeWiFi creates a new NetworkTypeWiFi
//
func NewNetworkTypeWiFi() *NetworkTypeWiFi {
	networkTypeWiFiTemp := NetworkTypeWiFi{
		tdCommon: tdCommon{Type: "networkTypeWiFi"},
	}

	return &networkTypeWiFiTemp
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeWiFi *NetworkTypeWiFi) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeWiFiType
}

// NetworkTypeOther A different network type (e.g., Ethernet network)
type NetworkTypeOther struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeOther
func (networkTypeOther *NetworkTypeOther) MessageType() string {
	return "networkTypeOther"
}

// NewNetworkTypeOther creates a new NetworkTypeOther
//
func NewNetworkTypeOther() *NetworkTypeOther {
	networkTypeOtherTemp := NetworkTypeOther{
		tdCommon: tdCommon{Type: "networkTypeOther"},
	}

	return &networkTypeOtherTemp
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeOther *NetworkTypeOther) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeOtherType
}
