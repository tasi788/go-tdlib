package tdlib

import (
	"encoding/json"
	"fmt"
)

// Sessions Contains a list of sessions
type Sessions struct {
	tdCommon
	Sessions               []Session `json:"sessions"`                  // List of sessions
	InactiveSessionTtlDays int32     `json:"inactive_session_ttl_days"` // Number of days of inactivity before sessions will automatically be terminated; 1-366 days
}

// MessageType return the string telegram-type of Sessions
func (sessions *Sessions) MessageType() string {
	return "sessions"
}

// NewSessions creates a new Sessions
//
// @param sessions List of sessions
// @param inactiveSessionTtlDays Number of days of inactivity before sessions will automatically be terminated; 1-366 days
func NewSessions(sessions []Session, inactiveSessionTtlDays int32) *Sessions {
	sessionsTemp := Sessions{
		tdCommon:               tdCommon{Type: "sessions"},
		Sessions:               sessions,
		InactiveSessionTtlDays: inactiveSessionTtlDays,
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
