package tdlib

import (
	"encoding/json"
	"fmt"
)

// ConnectedWebsites Contains a list of websites the current user is logged in with Telegram
type ConnectedWebsites struct {
	tdCommon
	Websites []ConnectedWebsite `json:"websites"` // List of connected websites
}

// MessageType return the string telegram-type of ConnectedWebsites
func (connectedWebsites *ConnectedWebsites) MessageType() string {
	return "connectedWebsites"
}

// NewConnectedWebsites creates a new ConnectedWebsites
//
// @param websites List of connected websites
func NewConnectedWebsites(websites []ConnectedWebsite) *ConnectedWebsites {
	connectedWebsitesTemp := ConnectedWebsites{
		tdCommon: tdCommon{Type: "connectedWebsites"},
		Websites: websites,
	}

	return &connectedWebsitesTemp
}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (client *Client) GetConnectedWebsites() (*ConnectedWebsites, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getConnectedWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var connectedWebsites ConnectedWebsites
	err = json.Unmarshal(result.Raw, &connectedWebsites)
	return &connectedWebsites, err
}
