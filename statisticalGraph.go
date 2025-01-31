package tdlib

import (
	"encoding/json"
	"fmt"
)

// StatisticalGraph Describes a statistical graph
type StatisticalGraph interface {
	GetStatisticalGraphEnum() StatisticalGraphEnum
}

// StatisticalGraphEnum Alias for abstract StatisticalGraph 'Sub-Classes', used as constant-enum here
type StatisticalGraphEnum string

// StatisticalGraph enums
const (
	StatisticalGraphDataType  StatisticalGraphEnum = "statisticalGraphData"
	StatisticalGraphAsyncType StatisticalGraphEnum = "statisticalGraphAsync"
	StatisticalGraphErrorType StatisticalGraphEnum = "statisticalGraphError"
)

func unmarshalStatisticalGraph(rawMsg *json.RawMessage) (StatisticalGraph, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch StatisticalGraphEnum(objMap["@type"].(string)) {
	case StatisticalGraphDataType:
		var statisticalGraphData StatisticalGraphData
		err := json.Unmarshal(*rawMsg, &statisticalGraphData)
		return &statisticalGraphData, err

	case StatisticalGraphAsyncType:
		var statisticalGraphAsync StatisticalGraphAsync
		err := json.Unmarshal(*rawMsg, &statisticalGraphAsync)
		return &statisticalGraphAsync, err

	case StatisticalGraphErrorType:
		var statisticalGraphError StatisticalGraphError
		err := json.Unmarshal(*rawMsg, &statisticalGraphError)
		return &statisticalGraphError, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// StatisticalGraphData A graph data
type StatisticalGraphData struct {
	tdCommon
	JsonData  string `json:"json_data"`  // Graph data in JSON format
	ZoomToken string `json:"zoom_token"` // If non-empty, a token which can be used to receive a zoomed in graph
}

// MessageType return the string telegram-type of StatisticalGraphData
func (statisticalGraphData *StatisticalGraphData) MessageType() string {
	return "statisticalGraphData"
}

// NewStatisticalGraphData creates a new StatisticalGraphData
//
// @param jsonStringData Graph data in JSON format
// @param zoomToken If non-empty, a token which can be used to receive a zoomed in graph
func NewStatisticalGraphData(jsonStringData string, zoomToken string) *StatisticalGraphData {
	statisticalGraphDataTemp := StatisticalGraphData{
		tdCommon:  tdCommon{Type: "statisticalGraphData"},
		JsonData:  jsonStringData,
		ZoomToken: zoomToken,
	}

	return &statisticalGraphDataTemp
}

// GetStatisticalGraphEnum return the enum type of this object
func (statisticalGraphData *StatisticalGraphData) GetStatisticalGraphEnum() StatisticalGraphEnum {
	return StatisticalGraphDataType
}

// StatisticalGraphAsync The graph data to be asynchronously loaded through getStatisticalGraph
type StatisticalGraphAsync struct {
	tdCommon
	Token string `json:"token"` // The token to use for data loading
}

// MessageType return the string telegram-type of StatisticalGraphAsync
func (statisticalGraphAsync *StatisticalGraphAsync) MessageType() string {
	return "statisticalGraphAsync"
}

// NewStatisticalGraphAsync creates a new StatisticalGraphAsync
//
// @param token The token to use for data loading
func NewStatisticalGraphAsync(token string) *StatisticalGraphAsync {
	statisticalGraphAsyncTemp := StatisticalGraphAsync{
		tdCommon: tdCommon{Type: "statisticalGraphAsync"},
		Token:    token,
	}

	return &statisticalGraphAsyncTemp
}

// GetStatisticalGraphEnum return the enum type of this object
func (statisticalGraphAsync *StatisticalGraphAsync) GetStatisticalGraphEnum() StatisticalGraphEnum {
	return StatisticalGraphAsyncType
}

// StatisticalGraphError An error message to be shown to the user instead of the graph
type StatisticalGraphError struct {
	tdCommon
	ErrorMessage string `json:"error_message"` // The error message
}

// MessageType return the string telegram-type of StatisticalGraphError
func (statisticalGraphError *StatisticalGraphError) MessageType() string {
	return "statisticalGraphError"
}

// NewStatisticalGraphError creates a new StatisticalGraphError
//
// @param errorMessage The error message
func NewStatisticalGraphError(errorMessage string) *StatisticalGraphError {
	statisticalGraphErrorTemp := StatisticalGraphError{
		tdCommon:     tdCommon{Type: "statisticalGraphError"},
		ErrorMessage: errorMessage,
	}

	return &statisticalGraphErrorTemp
}

// GetStatisticalGraphEnum return the enum type of this object
func (statisticalGraphError *StatisticalGraphError) GetStatisticalGraphEnum() StatisticalGraphEnum {
	return StatisticalGraphErrorType
}

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph
// @param chatId Chat identifier
// @param token The token for graph loading
// @param x X-value for zoomed in graph or 0 otherwise
func (client *Client) GetStatisticalGraph(chatId int64, token string, x int64) (StatisticalGraph, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStatisticalGraph",
		"chat_id": chatId,
		"token":   token,
		"x":       x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch StatisticalGraphEnum(result.Data["@type"].(string)) {

	case StatisticalGraphDataType:
		var statisticalGraph StatisticalGraphData
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphAsyncType:
		var statisticalGraph StatisticalGraphAsync
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphErrorType:
		var statisticalGraph StatisticalGraphError
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
