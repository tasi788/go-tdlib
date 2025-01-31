package tdlib

import (
	"encoding/json"
	"fmt"
)

// NetworkStatistics A full list of available network statistic entries
type NetworkStatistics struct {
	tdCommon
	SinceDate int32                    `json:"since_date"` // Point in time (Unix timestamp) from which the statistics are collected
	Entries   []NetworkStatisticsEntry `json:"entries"`    // Network statistics entries
}

// MessageType return the string telegram-type of NetworkStatistics
func (networkStatistics *NetworkStatistics) MessageType() string {
	return "networkStatistics"
}

// NewNetworkStatistics creates a new NetworkStatistics
//
// @param sinceDate Point in time (Unix timestamp) from which the statistics are collected
// @param entries Network statistics entries
func NewNetworkStatistics(sinceDate int32, entries []NetworkStatisticsEntry) *NetworkStatistics {
	networkStatisticsTemp := NetworkStatistics{
		tdCommon:  tdCommon{Type: "networkStatistics"},
		SinceDate: sinceDate,
		Entries:   entries,
	}

	return &networkStatisticsTemp
}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization
// @param onlyCurrent If true, returns only data for the current library launch
func (client *Client) GetNetworkStatistics(onlyCurrent bool) (*NetworkStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getNetworkStatistics",
		"only_current": onlyCurrent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var networkStatistics NetworkStatistics
	err = json.Unmarshal(result.Raw, &networkStatistics)
	return &networkStatistics, err
}
