package tdlib

import (
	"encoding/json"
	"fmt"
)

// CustomRequestResult Contains the result of a custom request
type CustomRequestResult struct {
	tdCommon
	Result string `json:"result"` // A JSON-serialized result
}

// MessageType return the string telegram-type of CustomRequestResult
func (customRequestResult *CustomRequestResult) MessageType() string {
	return "customRequestResult"
}

// NewCustomRequestResult creates a new CustomRequestResult
//
// @param result A JSON-serialized result
func NewCustomRequestResult(result string) *CustomRequestResult {
	customRequestResultTemp := CustomRequestResult{
		tdCommon: tdCommon{Type: "customRequestResult"},
		Result:   result,
	}

	return &customRequestResultTemp
}

// SendCustomRequest Sends a custom request; for bots only
// @param method The method name
// @param parameters JSON-serialized method parameters
func (client *Client) SendCustomRequest(method string, parameters string) (*CustomRequestResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "sendCustomRequest",
		"method":     method,
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var customRequestResult CustomRequestResult
	err = json.Unmarshal(result.Raw, &customRequestResult)
	return &customRequestResult, err
}
