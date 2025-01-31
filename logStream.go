package tdlib

import (
	"encoding/json"
	"fmt"
)

// LogStream Describes a stream to which TDLib internal log is written
type LogStream interface {
	GetLogStreamEnum() LogStreamEnum
}

// LogStreamEnum Alias for abstract LogStream 'Sub-Classes', used as constant-enum here
type LogStreamEnum string

// LogStream enums
const (
	LogStreamDefaultType LogStreamEnum = "logStreamDefault"
	LogStreamFileType    LogStreamEnum = "logStreamFile"
	LogStreamEmptyType   LogStreamEnum = "logStreamEmpty"
)

func unmarshalLogStream(rawMsg *json.RawMessage) (LogStream, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LogStreamEnum(objMap["@type"].(string)) {
	case LogStreamDefaultType:
		var logStreamDefault LogStreamDefault
		err := json.Unmarshal(*rawMsg, &logStreamDefault)
		return &logStreamDefault, err

	case LogStreamFileType:
		var logStreamFile LogStreamFile
		err := json.Unmarshal(*rawMsg, &logStreamFile)
		return &logStreamFile, err

	case LogStreamEmptyType:
		var logStreamEmpty LogStreamEmpty
		err := json.Unmarshal(*rawMsg, &logStreamEmpty)
		return &logStreamEmpty, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// LogStreamDefault The log is written to stderr or an OS specific log
type LogStreamDefault struct {
	tdCommon
}

// MessageType return the string telegram-type of LogStreamDefault
func (logStreamDefault *LogStreamDefault) MessageType() string {
	return "logStreamDefault"
}

// NewLogStreamDefault creates a new LogStreamDefault
//
func NewLogStreamDefault() *LogStreamDefault {
	logStreamDefaultTemp := LogStreamDefault{
		tdCommon: tdCommon{Type: "logStreamDefault"},
	}

	return &logStreamDefaultTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamDefault *LogStreamDefault) GetLogStreamEnum() LogStreamEnum {
	return LogStreamDefaultType
}

// LogStreamFile The log is written to a file
type LogStreamFile struct {
	tdCommon
	Path           string `json:"path"`            // Path to the file to where the internal TDLib log will be written
	MaxFileSize    int64  `json:"max_file_size"`   // The maximum size of the file to where the internal TDLib log is written before the file will automatically be rotated, in bytes
	RedirectStderr bool   `json:"redirect_stderr"` // Pass true to additionally redirect stderr to the log file. Ignored on Windows
}

// MessageType return the string telegram-type of LogStreamFile
func (logStreamFile *LogStreamFile) MessageType() string {
	return "logStreamFile"
}

// NewLogStreamFile creates a new LogStreamFile
//
// @param path Path to the file to where the internal TDLib log will be written
// @param maxFileSize The maximum size of the file to where the internal TDLib log is written before the file will automatically be rotated, in bytes
// @param redirectStderr Pass true to additionally redirect stderr to the log file. Ignored on Windows
func NewLogStreamFile(path string, maxFileSize int64, redirectStderr bool) *LogStreamFile {
	logStreamFileTemp := LogStreamFile{
		tdCommon:       tdCommon{Type: "logStreamFile"},
		Path:           path,
		MaxFileSize:    maxFileSize,
		RedirectStderr: redirectStderr,
	}

	return &logStreamFileTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamFile *LogStreamFile) GetLogStreamEnum() LogStreamEnum {
	return LogStreamFileType
}

// LogStreamEmpty The log is written nowhere
type LogStreamEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of LogStreamEmpty
func (logStreamEmpty *LogStreamEmpty) MessageType() string {
	return "logStreamEmpty"
}

// NewLogStreamEmpty creates a new LogStreamEmpty
//
func NewLogStreamEmpty() *LogStreamEmpty {
	logStreamEmptyTemp := LogStreamEmpty{
		tdCommon: tdCommon{Type: "logStreamEmpty"},
	}

	return &logStreamEmptyTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamEmpty *LogStreamEmpty) GetLogStreamEnum() LogStreamEnum {
	return LogStreamEmptyType
}

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogStream() (LogStream, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogStream",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LogStreamEnum(result.Data["@type"].(string)) {

	case LogStreamDefaultType:
		var logStream LogStreamDefault
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamFileType:
		var logStream LogStreamFile
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamEmptyType:
		var logStream LogStreamEmpty
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
