package tdlib

import (
	"encoding/json"
)

// JsonObjectMember Represents one member of a JSON object
type JsonObjectMember struct {
	tdCommon
	Key   string    `json:"key"`   // Member's key
	Value JsonValue `json:"value"` // Member's value
}

// MessageType return the string telegram-type of JsonObjectMember
func (jsonObjectMember *JsonObjectMember) MessageType() string {
	return "jsonObjectMember"
}

// NewJsonObjectMember creates a new JsonObjectMember
//
// @param key Member's key
// @param value Member's value
func NewJsonObjectMember(key string, value JsonValue) *JsonObjectMember {
	jsonObjectMemberTemp := JsonObjectMember{
		tdCommon: tdCommon{Type: "jsonObjectMember"},
		Key:      key,
		Value:    value,
	}

	return &jsonObjectMemberTemp
}

// UnmarshalJSON unmarshal to json
func (jsonObjectMember *JsonObjectMember) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Key string `json:"key"` // Member's key

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	jsonObjectMember.tdCommon = tempObj.tdCommon
	jsonObjectMember.Key = tempObj.Key

	fieldValue, _ := unmarshalJsonValue(objMap["value"])
	jsonObjectMember.Value = fieldValue

	return nil
}
