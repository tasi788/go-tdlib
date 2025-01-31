package tdlib

import (
	"encoding/json"
	"fmt"
)

// LogTags Contains a list of available TDLib internal log tags
type LogTags struct {
	tdCommon
	Tags []string `json:"tags"` // List of log tags
}

// MessageType return the string telegram-type of LogTags
func (logTags *LogTags) MessageType() string {
	return "logTags"
}

// NewLogTags creates a new LogTags
//
// @param tags List of log tags
func NewLogTags(tags []string) *LogTags {
	logTagsTemp := LogTags{
		tdCommon: tdCommon{Type: "logTags"},
		Tags:     tags,
	}

	return &logTagsTemp
}

// GetLogTags Returns list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (client *Client) GetLogTags() (*LogTags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTags",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var logTags LogTags
	err = json.Unmarshal(result.Raw, &logTags)
	return &logTags, err
}
