package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestVectorInt A simple object containing a vector of numbers; for testing only
type TestVectorInt struct {
	tdCommon
	Value []int32 `json:"value"` // Vector of numbers
}

// MessageType return the string telegram-type of TestVectorInt
func (testVectorInt *TestVectorInt) MessageType() string {
	return "testVectorInt"
}

// NewTestVectorInt creates a new TestVectorInt
//
// @param value Vector of numbers
func NewTestVectorInt(value []int32) *TestVectorInt {
	testVectorIntTemp := TestVectorInt{
		tdCommon: tdCommon{Type: "testVectorInt"},
		Value:    value,
	}

	return &testVectorIntTemp
}

// TestCallVectorInt Returns the received vector of numbers; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of numbers to return
func (client *Client) TestCallVectorInt(x []int32) (*TestVectorInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorInt TestVectorInt
	err = json.Unmarshal(result.Raw, &testVectorInt)
	return &testVectorInt, err
}
