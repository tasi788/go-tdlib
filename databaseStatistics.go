package tdlib

import (
	"encoding/json"
	"fmt"
)

// DatabaseStatistics Contains database statistics
type DatabaseStatistics struct {
	tdCommon
	Statistics string `json:"statistics"` // Database statistics in an unspecified human-readable format
}

// MessageType return the string telegram-type of DatabaseStatistics
func (databaseStatistics *DatabaseStatistics) MessageType() string {
	return "databaseStatistics"
}

// NewDatabaseStatistics creates a new DatabaseStatistics
//
// @param statistics Database statistics in an unspecified human-readable format
func NewDatabaseStatistics(statistics string) *DatabaseStatistics {
	databaseStatisticsTemp := DatabaseStatistics{
		tdCommon:   tdCommon{Type: "databaseStatistics"},
		Statistics: statistics,
	}

	return &databaseStatisticsTemp
}

// GetDatabaseStatistics Returns database statistics
func (client *Client) GetDatabaseStatistics() (*DatabaseStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDatabaseStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var databaseStatistics DatabaseStatistics
	err = json.Unmarshal(result.Raw, &databaseStatistics)
	return &databaseStatistics, err
}
