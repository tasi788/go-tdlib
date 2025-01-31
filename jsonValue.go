package tdlib

import (
	"encoding/json"
	"fmt"
)

// JsonValue Represents a JSON value
type JsonValue interface {
	GetJsonValueEnum() JsonValueEnum
}

// JsonValueEnum Alias for abstract JsonValue 'Sub-Classes', used as constant-enum here
type JsonValueEnum string

// JsonValue enums
const (
	JsonValueNullType    JsonValueEnum = "jsonValueNull"
	JsonValueBooleanType JsonValueEnum = "jsonValueBoolean"
	JsonValueNumberType  JsonValueEnum = "jsonValueNumber"
	JsonValueStringType  JsonValueEnum = "jsonValueString"
	JsonValueArrayType   JsonValueEnum = "jsonValueArray"
	JsonValueObjectType  JsonValueEnum = "jsonValueObject"
)

func unmarshalJsonValue(rawMsg *json.RawMessage) (JsonValue, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch JsonValueEnum(objMap["@type"].(string)) {
	case JsonValueNullType:
		var jsonValueNull JsonValueNull
		err := json.Unmarshal(*rawMsg, &jsonValueNull)
		return &jsonValueNull, err

	case JsonValueBooleanType:
		var jsonValueBoolean JsonValueBoolean
		err := json.Unmarshal(*rawMsg, &jsonValueBoolean)
		return &jsonValueBoolean, err

	case JsonValueNumberType:
		var jsonValueNumber JsonValueNumber
		err := json.Unmarshal(*rawMsg, &jsonValueNumber)
		return &jsonValueNumber, err

	case JsonValueStringType:
		var jsonValueString JsonValueString
		err := json.Unmarshal(*rawMsg, &jsonValueString)
		return &jsonValueString, err

	case JsonValueArrayType:
		var jsonValueArray JsonValueArray
		err := json.Unmarshal(*rawMsg, &jsonValueArray)
		return &jsonValueArray, err

	case JsonValueObjectType:
		var jsonValueObject JsonValueObject
		err := json.Unmarshal(*rawMsg, &jsonValueObject)
		return &jsonValueObject, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// JsonValueNull Represents a null JSON value
type JsonValueNull struct {
	tdCommon
}

// MessageType return the string telegram-type of JsonValueNull
func (jsonValueNull *JsonValueNull) MessageType() string {
	return "jsonValueNull"
}

// NewJsonValueNull creates a new JsonValueNull
//
func NewJsonValueNull() *JsonValueNull {
	jsonValueNullTemp := JsonValueNull{
		tdCommon: tdCommon{Type: "jsonValueNull"},
	}

	return &jsonValueNullTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueNull *JsonValueNull) GetJsonValueEnum() JsonValueEnum {
	return JsonValueNullType
}

// JsonValueBoolean Represents a boolean JSON value
type JsonValueBoolean struct {
	tdCommon
	Value bool `json:"value"` // The value
}

// MessageType return the string telegram-type of JsonValueBoolean
func (jsonValueBoolean *JsonValueBoolean) MessageType() string {
	return "jsonValueBoolean"
}

// NewJsonValueBoolean creates a new JsonValueBoolean
//
// @param value The value
func NewJsonValueBoolean(value bool) *JsonValueBoolean {
	jsonValueBooleanTemp := JsonValueBoolean{
		tdCommon: tdCommon{Type: "jsonValueBoolean"},
		Value:    value,
	}

	return &jsonValueBooleanTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueBoolean *JsonValueBoolean) GetJsonValueEnum() JsonValueEnum {
	return JsonValueBooleanType
}

// JsonValueNumber Represents a numeric JSON value
type JsonValueNumber struct {
	tdCommon
	Value float64 `json:"value"` // The value
}

// MessageType return the string telegram-type of JsonValueNumber
func (jsonValueNumber *JsonValueNumber) MessageType() string {
	return "jsonValueNumber"
}

// NewJsonValueNumber creates a new JsonValueNumber
//
// @param value The value
func NewJsonValueNumber(value float64) *JsonValueNumber {
	jsonValueNumberTemp := JsonValueNumber{
		tdCommon: tdCommon{Type: "jsonValueNumber"},
		Value:    value,
	}

	return &jsonValueNumberTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueNumber *JsonValueNumber) GetJsonValueEnum() JsonValueEnum {
	return JsonValueNumberType
}

// JsonValueString Represents a string JSON value
type JsonValueString struct {
	tdCommon
	Value string `json:"value"` // The value
}

// MessageType return the string telegram-type of JsonValueString
func (jsonValueString *JsonValueString) MessageType() string {
	return "jsonValueString"
}

// NewJsonValueString creates a new JsonValueString
//
// @param value The value
func NewJsonValueString(value string) *JsonValueString {
	jsonValueStringTemp := JsonValueString{
		tdCommon: tdCommon{Type: "jsonValueString"},
		Value:    value,
	}

	return &jsonValueStringTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueString *JsonValueString) GetJsonValueEnum() JsonValueEnum {
	return JsonValueStringType
}

// JsonValueArray Represents a JSON array
type JsonValueArray struct {
	tdCommon
	Values []JsonValue `json:"values"` // The list of array elements
}

// MessageType return the string telegram-type of JsonValueArray
func (jsonValueArray *JsonValueArray) MessageType() string {
	return "jsonValueArray"
}

// NewJsonValueArray creates a new JsonValueArray
//
// @param values The list of array elements
func NewJsonValueArray(values []JsonValue) *JsonValueArray {
	jsonValueArrayTemp := JsonValueArray{
		tdCommon: tdCommon{Type: "jsonValueArray"},
		Values:   values,
	}

	return &jsonValueArrayTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueArray *JsonValueArray) GetJsonValueEnum() JsonValueEnum {
	return JsonValueArrayType
}

// JsonValueObject Represents a JSON object
type JsonValueObject struct {
	tdCommon
	Members []JsonObjectMember `json:"members"` // The list of object members
}

// MessageType return the string telegram-type of JsonValueObject
func (jsonValueObject *JsonValueObject) MessageType() string {
	return "jsonValueObject"
}

// NewJsonValueObject creates a new JsonValueObject
//
// @param members The list of object members
func NewJsonValueObject(members []JsonObjectMember) *JsonValueObject {
	jsonValueObjectTemp := JsonValueObject{
		tdCommon: tdCommon{Type: "jsonValueObject"},
		Members:  members,
	}

	return &jsonValueObjectTemp
}

// GetJsonValueEnum return the enum type of this object
func (jsonValueObject *JsonValueObject) GetJsonValueEnum() JsonValueEnum {
	return JsonValueObjectType
}

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously
// @param jsonString The JSON-serialized string
func (client *Client) GetJsonValue(jsonString string) (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getJsonValue",
		"json":  jsonString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (client *Client) GetApplicationConfig() (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getApplicationConfig",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
