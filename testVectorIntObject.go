package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestVectorIntObject A simple object containing a vector of objects that hold a number; for testing only
type TestVectorIntObject struct {
	tdCommon
	Value []TestInt `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorIntObject
func (testVectorIntObject *TestVectorIntObject) MessageType() string {
	return "testVectorIntObject"
}

// NewTestVectorIntObject creates a new TestVectorIntObject
//
// @param value Vector of objects
func NewTestVectorIntObject(value []TestInt) *TestVectorIntObject {
	testVectorIntObjectTemp := TestVectorIntObject{
		tdCommon: tdCommon{Type: "testVectorIntObject"},
		Value:    value,
	}

	return &testVectorIntObjectTemp
}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorIntObject(x []TestInt) (*TestVectorIntObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorIntObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorIntObject TestVectorIntObject
	err = json.Unmarshal(result.Raw, &testVectorIntObject)
	return &testVectorIntObject, err
}
