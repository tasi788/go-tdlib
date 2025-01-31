package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestInt A simple object containing a number; for testing only
type TestInt struct {
	tdCommon
	Value int32 `json:"value"` // Number
}

// MessageType return the string telegram-type of TestInt
func (testInt *TestInt) MessageType() string {
	return "testInt"
}

// NewTestInt creates a new TestInt
//
// @param value Number
func NewTestInt(value int32) *TestInt {
	testIntTemp := TestInt{
		tdCommon: tdCommon{Type: "testInt"},
		Value:    value,
	}

	return &testIntTemp
}

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization
// @param x Number to square
func (client *Client) TestSquareInt(x int32) (*TestInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testSquareInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testInt TestInt
	err = json.Unmarshal(result.Raw, &testInt)
	return &testInt, err
}
