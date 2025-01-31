package tdlib

import (
	"encoding/json"
	"fmt"
)

// Error An object of this type can be returned on every function call, in case of an error
type Error struct {
	tdCommon
	Code    int32  `json:"code"`    // Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
	Message string `json:"message"` // Error message; subject to future changes
}

// MessageType return the string telegram-type of Error
func (error *Error) MessageType() string {
	return "error"
}

// NewError creates a new Error
//
// @param code Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
// @param message Error message; subject to future changes
func NewError(code int32, message string) *Error {
	errorTemp := Error{
		tdCommon: tdCommon{Type: "error"},
		Code:     code,
		Message:  message,
	}

	return &errorTemp
}

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously
// @param error The error to be returned
func (client *Client) TestReturnError(error *Error) (*Error, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testReturnError",
		"error": error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var errorDummy Error
	err = json.Unmarshal(result.Raw, &errorDummy)
	return &errorDummy, err
}
