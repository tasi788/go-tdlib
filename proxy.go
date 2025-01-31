package tdlib

import (
	"encoding/json"
	"fmt"
)

// Proxy Contains information about a proxy server
type Proxy struct {
	tdCommon
	Id           int32     `json:"id"`             // Unique identifier of the proxy
	Server       string    `json:"server"`         // Proxy server IP address
	Port         int32     `json:"port"`           // Proxy server port
	LastUsedDate int32     `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
	IsEnabled    bool      `json:"is_enabled"`     // True, if the proxy is enabled now
	Type         ProxyType `json:"type"`           // Type of the proxy
}

// MessageType return the string telegram-type of Proxy
func (proxy *Proxy) MessageType() string {
	return "proxy"
}

// NewProxy creates a new Proxy
//
// @param id Unique identifier of the proxy
// @param server Proxy server IP address
// @param port Proxy server port
// @param lastUsedDate Point in time (Unix timestamp) when the proxy was last used; 0 if never
// @param isEnabled True, if the proxy is enabled now
// @param typeParam Type of the proxy
func NewProxy(id int32, server string, port int32, lastUsedDate int32, isEnabled bool, typeParam ProxyType) *Proxy {
	proxyTemp := Proxy{
		tdCommon:     tdCommon{Type: "proxy"},
		Id:           id,
		Server:       server,
		Port:         port,
		LastUsedDate: lastUsedDate,
		IsEnabled:    isEnabled,
		Type:         typeParam,
	}

	return &proxyTemp
}

// UnmarshalJSON unmarshal to json
func (proxy *Proxy) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id           int32  `json:"id"`             // Unique identifier of the proxy
		Server       string `json:"server"`         // Proxy server IP address
		Port         int32  `json:"port"`           // Proxy server port
		LastUsedDate int32  `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
		IsEnabled    bool   `json:"is_enabled"`     // True, if the proxy is enabled now

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	proxy.tdCommon = tempObj.tdCommon
	proxy.Id = tempObj.Id
	proxy.Server = tempObj.Server
	proxy.Port = tempObj.Port
	proxy.LastUsedDate = tempObj.LastUsedDate
	proxy.IsEnabled = tempObj.IsEnabled

	fieldType, _ := unmarshalProxyType(objMap["type"])
	proxy.Type = fieldType

	return nil
}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy needs to be enabled
// @param typeParam Proxy type
func (client *Client) AddProxy(server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "addProxy",
		"server": server,
		"port":   port,
		"enable": enable,
		"type":   typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxy Proxy
	err = json.Unmarshal(result.Raw, &proxy)
	return &proxy, err
}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
// @param proxyId Proxy identifier
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy needs to be enabled
// @param typeParam Proxy type
func (client *Client) EditProxy(proxyId int32, server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "editProxy",
		"proxy_id": proxyId,
		"server":   server,
		"port":     port,
		"enable":   enable,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxyDummy Proxy
	err = json.Unmarshal(result.Raw, &proxyDummy)
	return &proxyDummy, err
}
