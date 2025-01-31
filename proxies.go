package tdlib

import (
	"encoding/json"
	"fmt"
)

// Proxies Represents a list of proxy servers
type Proxies struct {
	tdCommon
	Proxies []Proxy `json:"proxies"` // List of proxy servers
}

// MessageType return the string telegram-type of Proxies
func (proxies *Proxies) MessageType() string {
	return "proxies"
}

// NewProxies creates a new Proxies
//
// @param proxies List of proxy servers
func NewProxies(proxies []Proxy) *Proxies {
	proxiesTemp := Proxies{
		tdCommon: tdCommon{Type: "proxies"},
		Proxies:  proxies,
	}

	return &proxiesTemp
}

// GetProxies Returns list of proxies that are currently set up. Can be called before authorization
func (client *Client) GetProxies() (*Proxies, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getProxies",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxies Proxies
	err = json.Unmarshal(result.Raw, &proxies)
	return &proxies, err
}
