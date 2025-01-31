package tdlib

import (
	"encoding/json"
	"fmt"
)

// OptionValue Represents the value of an option
type OptionValue interface {
	GetOptionValueEnum() OptionValueEnum
}

// OptionValueEnum Alias for abstract OptionValue 'Sub-Classes', used as constant-enum here
type OptionValueEnum string

// OptionValue enums
const (
	OptionValueBooleanType OptionValueEnum = "optionValueBoolean"
	OptionValueEmptyType   OptionValueEnum = "optionValueEmpty"
	OptionValueIntegerType OptionValueEnum = "optionValueInteger"
	OptionValueStringType  OptionValueEnum = "optionValueString"
)

func unmarshalOptionValue(rawMsg *json.RawMessage) (OptionValue, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch OptionValueEnum(objMap["@type"].(string)) {
	case OptionValueBooleanType:
		var optionValueBoolean OptionValueBoolean
		err := json.Unmarshal(*rawMsg, &optionValueBoolean)
		return &optionValueBoolean, err

	case OptionValueEmptyType:
		var optionValueEmpty OptionValueEmpty
		err := json.Unmarshal(*rawMsg, &optionValueEmpty)
		return &optionValueEmpty, err

	case OptionValueIntegerType:
		var optionValueInteger OptionValueInteger
		err := json.Unmarshal(*rawMsg, &optionValueInteger)
		return &optionValueInteger, err

	case OptionValueStringType:
		var optionValueString OptionValueString
		err := json.Unmarshal(*rawMsg, &optionValueString)
		return &optionValueString, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// OptionValueBoolean Represents a boolean option
type OptionValueBoolean struct {
	tdCommon
	Value bool `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueBoolean
func (optionValueBoolean *OptionValueBoolean) MessageType() string {
	return "optionValueBoolean"
}

// NewOptionValueBoolean creates a new OptionValueBoolean
//
// @param value The value of the option
func NewOptionValueBoolean(value bool) *OptionValueBoolean {
	optionValueBooleanTemp := OptionValueBoolean{
		tdCommon: tdCommon{Type: "optionValueBoolean"},
		Value:    value,
	}

	return &optionValueBooleanTemp
}

// GetOptionValueEnum return the enum type of this object
func (optionValueBoolean *OptionValueBoolean) GetOptionValueEnum() OptionValueEnum {
	return OptionValueBooleanType
}

// OptionValueEmpty Represents an unknown option or an option which has a default value
type OptionValueEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of OptionValueEmpty
func (optionValueEmpty *OptionValueEmpty) MessageType() string {
	return "optionValueEmpty"
}

// NewOptionValueEmpty creates a new OptionValueEmpty
//
func NewOptionValueEmpty() *OptionValueEmpty {
	optionValueEmptyTemp := OptionValueEmpty{
		tdCommon: tdCommon{Type: "optionValueEmpty"},
	}

	return &optionValueEmptyTemp
}

// GetOptionValueEnum return the enum type of this object
func (optionValueEmpty *OptionValueEmpty) GetOptionValueEnum() OptionValueEnum {
	return OptionValueEmptyType
}

// OptionValueInteger Represents an integer option
type OptionValueInteger struct {
	tdCommon
	Value JSONInt64 `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueInteger
func (optionValueInteger *OptionValueInteger) MessageType() string {
	return "optionValueInteger"
}

// NewOptionValueInteger creates a new OptionValueInteger
//
// @param value The value of the option
func NewOptionValueInteger(value JSONInt64) *OptionValueInteger {
	optionValueIntegerTemp := OptionValueInteger{
		tdCommon: tdCommon{Type: "optionValueInteger"},
		Value:    value,
	}

	return &optionValueIntegerTemp
}

// GetOptionValueEnum return the enum type of this object
func (optionValueInteger *OptionValueInteger) GetOptionValueEnum() OptionValueEnum {
	return OptionValueIntegerType
}

// OptionValueString Represents a string option
type OptionValueString struct {
	tdCommon
	Value string `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueString
func (optionValueString *OptionValueString) MessageType() string {
	return "optionValueString"
}

// NewOptionValueString creates a new OptionValueString
//
// @param value The value of the option
func NewOptionValueString(value string) *OptionValueString {
	optionValueStringTemp := OptionValueString{
		tdCommon: tdCommon{Type: "optionValueString"},
		Value:    value,
	}

	return &optionValueStringTemp
}

// GetOptionValueEnum return the enum type of this object
func (optionValueString *OptionValueString) GetOptionValueEnum() OptionValueEnum {
	return OptionValueStringType
}

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization
// @param name The name of the option
func (client *Client) GetOption(name string) (OptionValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getOption",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch OptionValueEnum(result.Data["@type"].(string)) {

	case OptionValueBooleanType:
		var optionValue OptionValueBoolean
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueEmptyType:
		var optionValue OptionValueEmpty
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueIntegerType:
		var optionValue OptionValueInteger
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueStringType:
		var optionValue OptionValueString
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
