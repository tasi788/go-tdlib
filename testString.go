package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestString A simple object containing a string; for testing only
type TestString struct {
	tdCommon
	Value string `json:"value"` // String
}

// MessageType return the string telegram-type of TestString
func (testString *TestString) MessageType() string {
	return "testString"
}

// NewTestString creates a new TestString
//
// @param value String
func NewTestString(value string) *TestString {
	testStringTemp := TestString{
		tdCommon: tdCommon{Type: "testString"},
		Value:    value,
	}

	return &testStringTemp
}

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization
// @param x String to return
func (client *Client) TestCallString(x string) (*TestString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testString TestString
	err = json.Unmarshal(result.Raw, &testString)
	return &testString, err
}
