package tdlib

import (
	"encoding/json"
	"fmt"
)

// NetworkStatisticsEntry Contains statistics about network usage
type NetworkStatisticsEntry interface {
	GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum
}

// NetworkStatisticsEntryEnum Alias for abstract NetworkStatisticsEntry 'Sub-Classes', used as constant-enum here
type NetworkStatisticsEntryEnum string

// NetworkStatisticsEntry enums
const (
	NetworkStatisticsEntryFileType NetworkStatisticsEntryEnum = "networkStatisticsEntryFile"
	NetworkStatisticsEntryCallType NetworkStatisticsEntryEnum = "networkStatisticsEntryCall"
)

func unmarshalNetworkStatisticsEntry(rawMsg *json.RawMessage) (NetworkStatisticsEntry, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NetworkStatisticsEntryEnum(objMap["@type"].(string)) {
	case NetworkStatisticsEntryFileType:
		var networkStatisticsEntryFile NetworkStatisticsEntryFile
		err := json.Unmarshal(*rawMsg, &networkStatisticsEntryFile)
		return &networkStatisticsEntryFile, err

	case NetworkStatisticsEntryCallType:
		var networkStatisticsEntryCall NetworkStatisticsEntryCall
		err := json.Unmarshal(*rawMsg, &networkStatisticsEntryCall)
		return &networkStatisticsEntryCall, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// NetworkStatisticsEntryFile Contains information about the total amount of data that was used to send and receive files
type NetworkStatisticsEntryFile struct {
	tdCommon
	FileType      FileType    `json:"file_type"`      // Type of the file the data is part of; pass null if the data isn't related to files
	NetworkType   NetworkType `json:"network_type"`   // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
	SentBytes     int64       `json:"sent_bytes"`     // Total number of bytes sent
	ReceivedBytes int64       `json:"received_bytes"` // Total number of bytes received
}

// MessageType return the string telegram-type of NetworkStatisticsEntryFile
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) MessageType() string {
	return "networkStatisticsEntryFile"
}

// NewNetworkStatisticsEntryFile creates a new NetworkStatisticsEntryFile
//
// @param fileType Type of the file the data is part of; pass null if the data isn't related to files
// @param networkType Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
// @param sentBytes Total number of bytes sent
// @param receivedBytes Total number of bytes received
func NewNetworkStatisticsEntryFile(fileType FileType, networkType NetworkType, sentBytes int64, receivedBytes int64) *NetworkStatisticsEntryFile {
	networkStatisticsEntryFileTemp := NetworkStatisticsEntryFile{
		tdCommon:      tdCommon{Type: "networkStatisticsEntryFile"},
		FileType:      fileType,
		NetworkType:   networkType,
		SentBytes:     sentBytes,
		ReceivedBytes: receivedBytes,
	}

	return &networkStatisticsEntryFileTemp
}

// UnmarshalJSON unmarshal to json
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		SentBytes     int64 `json:"sent_bytes"`     // Total number of bytes sent
		ReceivedBytes int64 `json:"received_bytes"` // Total number of bytes received
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	networkStatisticsEntryFile.tdCommon = tempObj.tdCommon
	networkStatisticsEntryFile.SentBytes = tempObj.SentBytes
	networkStatisticsEntryFile.ReceivedBytes = tempObj.ReceivedBytes

	fieldFileType, _ := unmarshalFileType(objMap["file_type"])
	networkStatisticsEntryFile.FileType = fieldFileType

	fieldNetworkType, _ := unmarshalNetworkType(objMap["network_type"])
	networkStatisticsEntryFile.NetworkType = fieldNetworkType

	return nil
}

// GetNetworkStatisticsEntryEnum return the enum type of this object
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum {
	return NetworkStatisticsEntryFileType
}

// NetworkStatisticsEntryCall Contains information about the total amount of data that was used for calls
type NetworkStatisticsEntryCall struct {
	tdCommon
	NetworkType   NetworkType `json:"network_type"`   // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
	SentBytes     int64       `json:"sent_bytes"`     // Total number of bytes sent
	ReceivedBytes int64       `json:"received_bytes"` // Total number of bytes received
	Duration      float64     `json:"duration"`       // Total call duration, in seconds
}

// MessageType return the string telegram-type of NetworkStatisticsEntryCall
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) MessageType() string {
	return "networkStatisticsEntryCall"
}

// NewNetworkStatisticsEntryCall creates a new NetworkStatisticsEntryCall
//
// @param networkType Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
// @param sentBytes Total number of bytes sent
// @param receivedBytes Total number of bytes received
// @param duration Total call duration, in seconds
func NewNetworkStatisticsEntryCall(networkType NetworkType, sentBytes int64, receivedBytes int64, duration float64) *NetworkStatisticsEntryCall {
	networkStatisticsEntryCallTemp := NetworkStatisticsEntryCall{
		tdCommon:      tdCommon{Type: "networkStatisticsEntryCall"},
		NetworkType:   networkType,
		SentBytes:     sentBytes,
		ReceivedBytes: receivedBytes,
		Duration:      duration,
	}

	return &networkStatisticsEntryCallTemp
}

// UnmarshalJSON unmarshal to json
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		SentBytes     int64   `json:"sent_bytes"`     // Total number of bytes sent
		ReceivedBytes int64   `json:"received_bytes"` // Total number of bytes received
		Duration      float64 `json:"duration"`       // Total call duration, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	networkStatisticsEntryCall.tdCommon = tempObj.tdCommon
	networkStatisticsEntryCall.SentBytes = tempObj.SentBytes
	networkStatisticsEntryCall.ReceivedBytes = tempObj.ReceivedBytes
	networkStatisticsEntryCall.Duration = tempObj.Duration

	fieldNetworkType, _ := unmarshalNetworkType(objMap["network_type"])
	networkStatisticsEntryCall.NetworkType = fieldNetworkType

	return nil
}

// GetNetworkStatisticsEntryEnum return the enum type of this object
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum {
	return NetworkStatisticsEntryCallType
}
