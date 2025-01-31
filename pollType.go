package tdlib

import (
	"encoding/json"
	"fmt"
)

// PollType Describes the type of a poll
type PollType interface {
	GetPollTypeEnum() PollTypeEnum
}

// PollTypeEnum Alias for abstract PollType 'Sub-Classes', used as constant-enum here
type PollTypeEnum string

// PollType enums
const (
	PollTypeRegularType PollTypeEnum = "pollTypeRegular"
	PollTypeQuizType    PollTypeEnum = "pollTypeQuiz"
)

func unmarshalPollType(rawMsg *json.RawMessage) (PollType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PollTypeEnum(objMap["@type"].(string)) {
	case PollTypeRegularType:
		var pollTypeRegular PollTypeRegular
		err := json.Unmarshal(*rawMsg, &pollTypeRegular)
		return &pollTypeRegular, err

	case PollTypeQuizType:
		var pollTypeQuiz PollTypeQuiz
		err := json.Unmarshal(*rawMsg, &pollTypeQuiz)
		return &pollTypeQuiz, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PollTypeRegular A regular poll
type PollTypeRegular struct {
	tdCommon
	AllowMultipleAnswers bool `json:"allow_multiple_answers"` // True, if multiple answer options can be chosen simultaneously
}

// MessageType return the string telegram-type of PollTypeRegular
func (pollTypeRegular *PollTypeRegular) MessageType() string {
	return "pollTypeRegular"
}

// NewPollTypeRegular creates a new PollTypeRegular
//
// @param allowMultipleAnswers True, if multiple answer options can be chosen simultaneously
func NewPollTypeRegular(allowMultipleAnswers bool) *PollTypeRegular {
	pollTypeRegularTemp := PollTypeRegular{
		tdCommon:             tdCommon{Type: "pollTypeRegular"},
		AllowMultipleAnswers: allowMultipleAnswers,
	}

	return &pollTypeRegularTemp
}

// GetPollTypeEnum return the enum type of this object
func (pollTypeRegular *PollTypeRegular) GetPollTypeEnum() PollTypeEnum {
	return PollTypeRegularType
}

// PollTypeQuiz A poll in quiz mode, which has exactly one correct answer option and can be answered only once
type PollTypeQuiz struct {
	tdCommon
	CorrectOptionId int32          `json:"correct_option_id"` // 0-based identifier of the correct answer option; -1 for a yet unanswered poll
	Explanation     *FormattedText `json:"explanation"`       // Text that is shown when the user chooses an incorrect answer or taps on the lamp icon; 0-200 characters with at most 2 line feeds; empty for a yet unanswered poll
}

// MessageType return the string telegram-type of PollTypeQuiz
func (pollTypeQuiz *PollTypeQuiz) MessageType() string {
	return "pollTypeQuiz"
}

// NewPollTypeQuiz creates a new PollTypeQuiz
//
// @param correctOptionId 0-based identifier of the correct answer option; -1 for a yet unanswered poll
// @param explanation Text that is shown when the user chooses an incorrect answer or taps on the lamp icon; 0-200 characters with at most 2 line feeds; empty for a yet unanswered poll
func NewPollTypeQuiz(correctOptionId int32, explanation *FormattedText) *PollTypeQuiz {
	pollTypeQuizTemp := PollTypeQuiz{
		tdCommon:        tdCommon{Type: "pollTypeQuiz"},
		CorrectOptionId: correctOptionId,
		Explanation:     explanation,
	}

	return &pollTypeQuizTemp
}

// GetPollTypeEnum return the enum type of this object
func (pollTypeQuiz *PollTypeQuiz) GetPollTypeEnum() PollTypeEnum {
	return PollTypeQuizType
}
