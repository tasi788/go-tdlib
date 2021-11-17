package tdlib

import (
	"encoding/json"
	"fmt"
)

// Sessions Contains a list of sessions
type Sessions struct {
	tdCommon
	Sessions []Session `json:"sessions"` // List of sessions
}

// MessageType return the string telegram-type of Sessions
func (sessions *Sessions) MessageType() string {
	return "sessions"
}

// NewSessions creates a new Sessions
//
// @param sessions List of sessions
func NewSessions(sessions []Session) *Sessions {
	sessionsTemp := Sessions{
		tdCommon: tdCommon{Type: "sessions"},
		Sessions: sessions,
	}

	return &sessionsTemp
}

// GetActiveSessions Returns all active sessions of the current user
func (client *Client) GetActiveSessions() (*Sessions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var sessions Sessions
	err = json.Unmarshal(result.Raw, &sessions)
	return &sessions, err
}
