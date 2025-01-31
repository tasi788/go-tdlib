package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestVectorString A simple object containing a vector of strings; for testing only
type TestVectorString struct {
	tdCommon
	Value []string `json:"value"` // Vector of strings
}

// MessageType return the string telegram-type of TestVectorString
func (testVectorString *TestVectorString) MessageType() string {
	return "testVectorString"
}

// NewTestVectorString creates a new TestVectorString
//
// @param value Vector of strings
func NewTestVectorString(value []string) *TestVectorString {
	testVectorStringTemp := TestVectorString{
		tdCommon: tdCommon{Type: "testVectorString"},
		Value:    value,
	}

	return &testVectorStringTemp
}

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of strings to return
func (client *Client) TestCallVectorString(x []string) (*TestVectorString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorString TestVectorString
	err = json.Unmarshal(result.Raw, &testVectorString)
	return &testVectorString, err
}
