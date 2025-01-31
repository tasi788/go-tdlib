package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestVectorStringObject A simple object containing a vector of objects that hold a string; for testing only
type TestVectorStringObject struct {
	tdCommon
	Value []TestString `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorStringObject
func (testVectorStringObject *TestVectorStringObject) MessageType() string {
	return "testVectorStringObject"
}

// NewTestVectorStringObject creates a new TestVectorStringObject
//
// @param value Vector of objects
func NewTestVectorStringObject(value []TestString) *TestVectorStringObject {
	testVectorStringObjectTemp := TestVectorStringObject{
		tdCommon: tdCommon{Type: "testVectorStringObject"},
		Value:    value,
	}

	return &testVectorStringObjectTemp
}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorStringObject(x []TestString) (*TestVectorStringObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorStringObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorStringObject TestVectorStringObject
	err = json.Unmarshal(result.Raw, &testVectorStringObject)
	return &testVectorStringObject, err
}
