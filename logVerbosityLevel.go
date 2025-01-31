package tdlib

import (
	"encoding/json"
	"fmt"
)

// LogVerbosityLevel Contains a TDLib internal log verbosity level
type LogVerbosityLevel struct {
	tdCommon
	VerbosityLevel int32 `json:"verbosity_level"` // Log verbosity level
}

// MessageType return the string telegram-type of LogVerbosityLevel
func (logVerbosityLevel *LogVerbosityLevel) MessageType() string {
	return "logVerbosityLevel"
}

// NewLogVerbosityLevel creates a new LogVerbosityLevel
//
// @param verbosityLevel Log verbosity level
func NewLogVerbosityLevel(verbosityLevel int32) *LogVerbosityLevel {
	logVerbosityLevelTemp := LogVerbosityLevel{
		tdCommon:       tdCommon{Type: "logVerbosityLevel"},
		VerbosityLevel: verbosityLevel,
	}

	return &logVerbosityLevelTemp
}

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogVerbosityLevel() (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogVerbosityLevel",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err
}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) GetLogTagVerbosityLevel(tag string) (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTagVerbosityLevel",
		"tag":   tag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err
}
