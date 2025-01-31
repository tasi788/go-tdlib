package tdlib

import (
	"encoding/json"
	"fmt"
)

// TestBytes A simple object containing a sequence of bytes; for testing only
type TestBytes struct {
	tdCommon
	Value []byte `json:"value"` // Bytes
}

// MessageType return the string telegram-type of TestBytes
func (testBytes *TestBytes) MessageType() string {
	return "testBytes"
}

// NewTestBytes creates a new TestBytes
//
// @param value Bytes
func NewTestBytes(value []byte) *TestBytes {
	testBytesTemp := TestBytes{
		tdCommon: tdCommon{Type: "testBytes"},
		Value:    value,
	}

	return &testBytesTemp
}

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization
// @param x Bytes to return
func (client *Client) TestCallBytes(x []byte) (*TestBytes, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallBytes",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var testBytes TestBytes
	err = json.Unmarshal(result.Raw, &testBytes)
	return &testBytes, err
}
