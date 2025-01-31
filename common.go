package tdlib

import (
	"strconv"
	"strings"
)

type tdCommon struct {
	Type  string `json:"@type"`
	Extra string `json:"@extra"`
}

// TdMessage is the interface for all messages send and received to/from tdlib
type TdMessage interface {
	MessageType() string
}

// JSONInt64 alias for int64, in order to deal with json big number problem
type JSONInt64 int64

// UpdateData alias for use in UpdateMsg
type UpdateData map[string]interface{}

// UpdateMsg is used to unmarshal received json strings into
type UpdateMsg struct {
	Data UpdateData
	Raw  []byte
}

// MarshalJSON marshals to json
func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
	intStr := strconv.FormatInt(int64(*jsonInt), 10)
	return []byte(intStr), nil
}

// UnmarshalJSON unmarshals from json
func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
	intStr := string(b)
	intStr = strings.Replace(intStr, "\"", "", 2)
	jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return err
	}
	*jsonInt = JSONInt64(jsonBigInt)
	return nil
}
