package tdlib

import (
	"encoding/json"
)

// Poll Describes a poll
type Poll struct {
	tdCommon
	Id                 JSONInt64    `json:"id"`                    // Unique poll identifier
	Question           string       `json:"question"`              // Poll question; 1-300 characters
	Options            []PollOption `json:"options"`               // List of poll answer options
	TotalVoterCount    int32        `json:"total_voter_count"`     // Total number of voters, participating in the poll
	RecentVoterUserIds []int64      `json:"recent_voter_user_ids"` // User identifiers of recent voters, if the poll is non-anonymous
	IsAnonymous        bool         `json:"is_anonymous"`          // True, if the poll is anonymous
	Type               PollType     `json:"type"`                  // Type of the poll
	OpenPeriod         int32        `json:"open_period"`           // Amount of time the poll will be active after creation, in seconds
	CloseDate          int32        `json:"close_date"`            // Point in time (Unix timestamp) when the poll will be automatically closed
	IsClosed           bool         `json:"is_closed"`             // True, if the poll is closed
}

// MessageType return the string telegram-type of Poll
func (poll *Poll) MessageType() string {
	return "poll"
}

// NewPoll creates a new Poll
//
// @param id Unique poll identifier
// @param question Poll question; 1-300 characters
// @param options List of poll answer options
// @param totalVoterCount Total number of voters, participating in the poll
// @param recentVoterUserIds User identifiers of recent voters, if the poll is non-anonymous
// @param isAnonymous True, if the poll is anonymous
// @param typeParam Type of the poll
// @param openPeriod Amount of time the poll will be active after creation, in seconds
// @param closeDate Point in time (Unix timestamp) when the poll will be automatically closed
// @param isClosed True, if the poll is closed
func NewPoll(id JSONInt64, question string, options []PollOption, totalVoterCount int32, recentVoterUserIds []int64, isAnonymous bool, typeParam PollType, openPeriod int32, closeDate int32, isClosed bool) *Poll {
	pollTemp := Poll{
		tdCommon:           tdCommon{Type: "poll"},
		Id:                 id,
		Question:           question,
		Options:            options,
		TotalVoterCount:    totalVoterCount,
		RecentVoterUserIds: recentVoterUserIds,
		IsAnonymous:        isAnonymous,
		Type:               typeParam,
		OpenPeriod:         openPeriod,
		CloseDate:          closeDate,
		IsClosed:           isClosed,
	}

	return &pollTemp
}

// UnmarshalJSON unmarshal to json
func (poll *Poll) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                 JSONInt64    `json:"id"`                    // Unique poll identifier
		Question           string       `json:"question"`              // Poll question; 1-300 characters
		Options            []PollOption `json:"options"`               // List of poll answer options
		TotalVoterCount    int32        `json:"total_voter_count"`     // Total number of voters, participating in the poll
		RecentVoterUserIds []int64      `json:"recent_voter_user_ids"` // User identifiers of recent voters, if the poll is non-anonymous
		IsAnonymous        bool         `json:"is_anonymous"`          // True, if the poll is anonymous
		OpenPeriod         int32        `json:"open_period"`           // Amount of time the poll will be active after creation, in seconds
		CloseDate          int32        `json:"close_date"`            // Point in time (Unix timestamp) when the poll will be automatically closed
		IsClosed           bool         `json:"is_closed"`             // True, if the poll is closed
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	poll.tdCommon = tempObj.tdCommon
	poll.Id = tempObj.Id
	poll.Question = tempObj.Question
	poll.Options = tempObj.Options
	poll.TotalVoterCount = tempObj.TotalVoterCount
	poll.RecentVoterUserIds = tempObj.RecentVoterUserIds
	poll.IsAnonymous = tempObj.IsAnonymous
	poll.OpenPeriod = tempObj.OpenPeriod
	poll.CloseDate = tempObj.CloseDate
	poll.IsClosed = tempObj.IsClosed

	fieldType, _ := unmarshalPollType(objMap["type"])
	poll.Type = fieldType

	return nil
}
