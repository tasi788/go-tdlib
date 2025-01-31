package tdlib

// PollOption Describes one answer option of a poll
type PollOption struct {
	tdCommon
	Text           string `json:"text"`            // Option text; 1-100 characters
	VoterCount     int32  `json:"voter_count"`     // Number of voters for this option, available only for closed or voted polls
	VotePercentage int32  `json:"vote_percentage"` // The percentage of votes for this option; 0-100
	IsChosen       bool   `json:"is_chosen"`       // True, if the option was chosen by the user
	IsBeingChosen  bool   `json:"is_being_chosen"` // True, if the option is being chosen by a pending setPollAnswer request
}

// MessageType return the string telegram-type of PollOption
func (pollOption *PollOption) MessageType() string {
	return "pollOption"
}

// NewPollOption creates a new PollOption
//
// @param text Option text; 1-100 characters
// @param voterCount Number of voters for this option, available only for closed or voted polls
// @param votePercentage The percentage of votes for this option; 0-100
// @param isChosen True, if the option was chosen by the user
// @param isBeingChosen True, if the option is being chosen by a pending setPollAnswer request
func NewPollOption(text string, voterCount int32, votePercentage int32, isChosen bool, isBeingChosen bool) *PollOption {
	pollOptionTemp := PollOption{
		tdCommon:       tdCommon{Type: "pollOption"},
		Text:           text,
		VoterCount:     voterCount,
		VotePercentage: votePercentage,
		IsChosen:       isChosen,
		IsBeingChosen:  isBeingChosen,
	}

	return &pollOptionTemp
}
