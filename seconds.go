package tdlib

import (
	"encoding/json"
	"fmt"
)

// Seconds Contains a value representing a number of seconds
type Seconds struct {
	tdCommon
	Seconds float64 `json:"seconds"` // Number of seconds
}

// MessageType return the string telegram-type of Seconds
func (seconds *Seconds) MessageType() string {
	return "seconds"
}

// NewSeconds creates a new Seconds
//
// @param seconds Number of seconds
func NewSeconds(seconds float64) *Seconds {
	secondsTemp := Seconds{
		tdCommon: tdCommon{Type: "seconds"},
		Seconds:  seconds,
	}

	return &secondsTemp
}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization
// @param proxyId Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (client *Client) PingProxy(proxyId int32) (*Seconds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "pingProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var seconds Seconds
	err = json.Unmarshal(result.Raw, &seconds)
	return &seconds, err
}
