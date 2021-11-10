package tdlib

import (
	"encoding/json"
)

// CallServer Describes a server for relaying call data
type CallServer struct {
	tdCommon
	Id          JSONInt64      `json:"id"`           // Server identifier
	IpAddress   string         `json:"ip_address"`   // Server IPv4 address
	Ipv6Address string         `json:"ipv6_address"` // Server IPv6 address
	Port        int32          `json:"port"`         // Server port number
	Type        CallServerType `json:"type"`         // Server type
}

// MessageType return the string telegram-type of CallServer
func (callServer *CallServer) MessageType() string {
	return "callServer"
}

// NewCallServer creates a new CallServer
//
// @param id Server identifier
// @param ipAddress Server IPv4 address
// @param ipv6Address Server IPv6 address
// @param port Server port number
// @param typeParam Server type
func NewCallServer(id JSONInt64, ipAddress string, ipv6Address string, port int32, typeParam CallServerType) *CallServer {
	callServerTemp := CallServer{
		tdCommon:    tdCommon{Type: "callServer"},
		Id:          id,
		IpAddress:   ipAddress,
		Ipv6Address: ipv6Address,
		Port:        port,
		Type:        typeParam,
	}

	return &callServerTemp
}

// UnmarshalJSON unmarshal to json
func (callServer *CallServer) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id          JSONInt64 `json:"id"`           // Server identifier
		IpAddress   string    `json:"ip_address"`   // Server IPv4 address
		Ipv6Address string    `json:"ipv6_address"` // Server IPv6 address
		Port        int32     `json:"port"`         // Server port number

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	callServer.tdCommon = tempObj.tdCommon
	callServer.Id = tempObj.Id
	callServer.IpAddress = tempObj.IpAddress
	callServer.Ipv6Address = tempObj.Ipv6Address
	callServer.Port = tempObj.Port

	fieldType, _ := unmarshalCallServerType(objMap["type"])
	callServer.Type = fieldType

	return nil
}
