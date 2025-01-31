package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatStatistics Contains a detailed statistics about a chat
type ChatStatistics interface {
	GetChatStatisticsEnum() ChatStatisticsEnum
}

// ChatStatisticsEnum Alias for abstract ChatStatistics 'Sub-Classes', used as constant-enum here
type ChatStatisticsEnum string

// ChatStatistics enums
const (
	ChatStatisticsSupergroupType ChatStatisticsEnum = "chatStatisticsSupergroup"
	ChatStatisticsChannelType    ChatStatisticsEnum = "chatStatisticsChannel"
)

func unmarshalChatStatistics(rawMsg *json.RawMessage) (ChatStatistics, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatStatisticsEnum(objMap["@type"].(string)) {
	case ChatStatisticsSupergroupType:
		var chatStatisticsSupergroup ChatStatisticsSupergroup
		err := json.Unmarshal(*rawMsg, &chatStatisticsSupergroup)
		return &chatStatisticsSupergroup, err

	case ChatStatisticsChannelType:
		var chatStatisticsChannel ChatStatisticsChannel
		err := json.Unmarshal(*rawMsg, &chatStatisticsChannel)
		return &chatStatisticsChannel, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatStatisticsSupergroup A detailed statistics about a supergroup chat
type ChatStatisticsSupergroup struct {
	tdCommon
	Period              *DateRange                               `json:"period"`                // A period to which the statistics applies
	MemberCount         *StatisticalValue                        `json:"member_count"`          // Number of members in the chat
	MessageCount        *StatisticalValue                        `json:"message_count"`         // Number of messages sent to the chat
	ViewerCount         *StatisticalValue                        `json:"viewer_count"`          // Number of users who viewed messages in the chat
	SenderCount         *StatisticalValue                        `json:"sender_count"`          // Number of users who sent messages to the chat
	MemberCountGraph    StatisticalGraph                         `json:"member_count_graph"`    // A graph containing number of members in the chat
	JoinGraph           StatisticalGraph                         `json:"join_graph"`            // A graph containing number of members joined and left the chat
	JoinBySourceGraph   StatisticalGraph                         `json:"join_by_source_graph"`  // A graph containing number of new member joins per source
	LanguageGraph       StatisticalGraph                         `json:"language_graph"`        // A graph containing distribution of active users per language
	MessageContentGraph StatisticalGraph                         `json:"message_content_graph"` // A graph containing distribution of sent messages by content type
	ActionGraph         StatisticalGraph                         `json:"action_graph"`          // A graph containing number of different actions in the chat
	DayGraph            StatisticalGraph                         `json:"day_graph"`             // A graph containing distribution of message views per hour
	WeekGraph           StatisticalGraph                         `json:"week_graph"`            // A graph containing distribution of message views per day of week
	TopSenders          []ChatStatisticsMessageSenderInfo        `json:"top_senders"`           // List of users sent most messages in the last week
	TopAdministrators   []ChatStatisticsAdministratorActionsInfo `json:"top_administrators"`    // List of most active administrators in the last week
	TopInviters         []ChatStatisticsInviterInfo              `json:"top_inviters"`          // List of most active inviters of new members in the last week
}

// MessageType return the string telegram-type of ChatStatisticsSupergroup
func (chatStatisticsSupergroup *ChatStatisticsSupergroup) MessageType() string {
	return "chatStatisticsSupergroup"
}

// NewChatStatisticsSupergroup creates a new ChatStatisticsSupergroup
//
// @param period A period to which the statistics applies
// @param memberCount Number of members in the chat
// @param messageCount Number of messages sent to the chat
// @param viewerCount Number of users who viewed messages in the chat
// @param senderCount Number of users who sent messages to the chat
// @param memberCountGraph A graph containing number of members in the chat
// @param joinGraph A graph containing number of members joined and left the chat
// @param joinBySourceGraph A graph containing number of new member joins per source
// @param languageGraph A graph containing distribution of active users per language
// @param messageContentGraph A graph containing distribution of sent messages by content type
// @param actionGraph A graph containing number of different actions in the chat
// @param dayGraph A graph containing distribution of message views per hour
// @param weekGraph A graph containing distribution of message views per day of week
// @param topSenders List of users sent most messages in the last week
// @param topAdministrators List of most active administrators in the last week
// @param topInviters List of most active inviters of new members in the last week
func NewChatStatisticsSupergroup(period *DateRange, memberCount *StatisticalValue, messageCount *StatisticalValue, viewerCount *StatisticalValue, senderCount *StatisticalValue, memberCountGraph StatisticalGraph, joinGraph StatisticalGraph, joinBySourceGraph StatisticalGraph, languageGraph StatisticalGraph, messageContentGraph StatisticalGraph, actionGraph StatisticalGraph, dayGraph StatisticalGraph, weekGraph StatisticalGraph, topSenders []ChatStatisticsMessageSenderInfo, topAdministrators []ChatStatisticsAdministratorActionsInfo, topInviters []ChatStatisticsInviterInfo) *ChatStatisticsSupergroup {
	chatStatisticsSupergroupTemp := ChatStatisticsSupergroup{
		tdCommon:            tdCommon{Type: "chatStatisticsSupergroup"},
		Period:              period,
		MemberCount:         memberCount,
		MessageCount:        messageCount,
		ViewerCount:         viewerCount,
		SenderCount:         senderCount,
		MemberCountGraph:    memberCountGraph,
		JoinGraph:           joinGraph,
		JoinBySourceGraph:   joinBySourceGraph,
		LanguageGraph:       languageGraph,
		MessageContentGraph: messageContentGraph,
		ActionGraph:         actionGraph,
		DayGraph:            dayGraph,
		WeekGraph:           weekGraph,
		TopSenders:          topSenders,
		TopAdministrators:   topAdministrators,
		TopInviters:         topInviters,
	}

	return &chatStatisticsSupergroupTemp
}

// UnmarshalJSON unmarshal to json
func (chatStatisticsSupergroup *ChatStatisticsSupergroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Period            *DateRange                               `json:"period"`             // A period to which the statistics applies
		MemberCount       *StatisticalValue                        `json:"member_count"`       // Number of members in the chat
		MessageCount      *StatisticalValue                        `json:"message_count"`      // Number of messages sent to the chat
		ViewerCount       *StatisticalValue                        `json:"viewer_count"`       // Number of users who viewed messages in the chat
		SenderCount       *StatisticalValue                        `json:"sender_count"`       // Number of users who sent messages to the chat
		TopSenders        []ChatStatisticsMessageSenderInfo        `json:"top_senders"`        // List of users sent most messages in the last week
		TopAdministrators []ChatStatisticsAdministratorActionsInfo `json:"top_administrators"` // List of most active administrators in the last week
		TopInviters       []ChatStatisticsInviterInfo              `json:"top_inviters"`       // List of most active inviters of new members in the last week
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatStatisticsSupergroup.tdCommon = tempObj.tdCommon
	chatStatisticsSupergroup.Period = tempObj.Period
	chatStatisticsSupergroup.MemberCount = tempObj.MemberCount
	chatStatisticsSupergroup.MessageCount = tempObj.MessageCount
	chatStatisticsSupergroup.ViewerCount = tempObj.ViewerCount
	chatStatisticsSupergroup.SenderCount = tempObj.SenderCount
	chatStatisticsSupergroup.TopSenders = tempObj.TopSenders
	chatStatisticsSupergroup.TopAdministrators = tempObj.TopAdministrators
	chatStatisticsSupergroup.TopInviters = tempObj.TopInviters

	fieldMemberCountGraph, _ := unmarshalStatisticalGraph(objMap["member_count_graph"])
	chatStatisticsSupergroup.MemberCountGraph = fieldMemberCountGraph

	fieldJoinGraph, _ := unmarshalStatisticalGraph(objMap["join_graph"])
	chatStatisticsSupergroup.JoinGraph = fieldJoinGraph

	fieldJoinBySourceGraph, _ := unmarshalStatisticalGraph(objMap["join_by_source_graph"])
	chatStatisticsSupergroup.JoinBySourceGraph = fieldJoinBySourceGraph

	fieldLanguageGraph, _ := unmarshalStatisticalGraph(objMap["language_graph"])
	chatStatisticsSupergroup.LanguageGraph = fieldLanguageGraph

	fieldMessageContentGraph, _ := unmarshalStatisticalGraph(objMap["message_content_graph"])
	chatStatisticsSupergroup.MessageContentGraph = fieldMessageContentGraph

	fieldActionGraph, _ := unmarshalStatisticalGraph(objMap["action_graph"])
	chatStatisticsSupergroup.ActionGraph = fieldActionGraph

	fieldDayGraph, _ := unmarshalStatisticalGraph(objMap["day_graph"])
	chatStatisticsSupergroup.DayGraph = fieldDayGraph

	fieldWeekGraph, _ := unmarshalStatisticalGraph(objMap["week_graph"])
	chatStatisticsSupergroup.WeekGraph = fieldWeekGraph

	return nil
}

// GetChatStatisticsEnum return the enum type of this object
func (chatStatisticsSupergroup *ChatStatisticsSupergroup) GetChatStatisticsEnum() ChatStatisticsEnum {
	return ChatStatisticsSupergroupType
}

// ChatStatisticsChannel A detailed statistics about a channel chat
type ChatStatisticsChannel struct {
	tdCommon
	Period                         *DateRange                             `json:"period"`                           // A period to which the statistics applies
	MemberCount                    *StatisticalValue                      `json:"member_count"`                     // Number of members in the chat
	MeanViewCount                  *StatisticalValue                      `json:"mean_view_count"`                  // Mean number of times the recently sent messages was viewed
	MeanShareCount                 *StatisticalValue                      `json:"mean_share_count"`                 // Mean number of times the recently sent messages was shared
	EnabledNotificationsPercentage float64                                `json:"enabled_notifications_percentage"` // A percentage of users with enabled notifications for the chat
	MemberCountGraph               StatisticalGraph                       `json:"member_count_graph"`               // A graph containing number of members in the chat
	JoinGraph                      StatisticalGraph                       `json:"join_graph"`                       // A graph containing number of members joined and left the chat
	MuteGraph                      StatisticalGraph                       `json:"mute_graph"`                       // A graph containing number of members muted and unmuted the chat
	ViewCountByHourGraph           StatisticalGraph                       `json:"view_count_by_hour_graph"`         // A graph containing number of message views in a given hour in the last two weeks
	ViewCountBySourceGraph         StatisticalGraph                       `json:"view_count_by_source_graph"`       // A graph containing number of message views per source
	JoinBySourceGraph              StatisticalGraph                       `json:"join_by_source_graph"`             // A graph containing number of new member joins per source
	LanguageGraph                  StatisticalGraph                       `json:"language_graph"`                   // A graph containing number of users viewed chat messages per language
	MessageInteractionGraph        StatisticalGraph                       `json:"message_interaction_graph"`        // A graph containing number of chat message views and shares
	InstantViewInteractionGraph    StatisticalGraph                       `json:"instant_view_interaction_graph"`   // A graph containing number of views of associated with the chat instant views
	RecentMessageInteractions      []ChatStatisticsMessageInteractionInfo `json:"recent_message_interactions"`      // Detailed statistics about number of views and shares of recently sent messages
}

// MessageType return the string telegram-type of ChatStatisticsChannel
func (chatStatisticsChannel *ChatStatisticsChannel) MessageType() string {
	return "chatStatisticsChannel"
}

// NewChatStatisticsChannel creates a new ChatStatisticsChannel
//
// @param period A period to which the statistics applies
// @param memberCount Number of members in the chat
// @param meanViewCount Mean number of times the recently sent messages was viewed
// @param meanShareCount Mean number of times the recently sent messages was shared
// @param enabledNotificationsPercentage A percentage of users with enabled notifications for the chat
// @param memberCountGraph A graph containing number of members in the chat
// @param joinGraph A graph containing number of members joined and left the chat
// @param muteGraph A graph containing number of members muted and unmuted the chat
// @param viewCountByHourGraph A graph containing number of message views in a given hour in the last two weeks
// @param viewCountBySourceGraph A graph containing number of message views per source
// @param joinBySourceGraph A graph containing number of new member joins per source
// @param languageGraph A graph containing number of users viewed chat messages per language
// @param messageInteractionGraph A graph containing number of chat message views and shares
// @param instantViewInteractionGraph A graph containing number of views of associated with the chat instant views
// @param recentMessageInteractions Detailed statistics about number of views and shares of recently sent messages
func NewChatStatisticsChannel(period *DateRange, memberCount *StatisticalValue, meanViewCount *StatisticalValue, meanShareCount *StatisticalValue, enabledNotificationsPercentage float64, memberCountGraph StatisticalGraph, joinGraph StatisticalGraph, muteGraph StatisticalGraph, viewCountByHourGraph StatisticalGraph, viewCountBySourceGraph StatisticalGraph, joinBySourceGraph StatisticalGraph, languageGraph StatisticalGraph, messageInteractionGraph StatisticalGraph, instantViewInteractionGraph StatisticalGraph, recentMessageInteractions []ChatStatisticsMessageInteractionInfo) *ChatStatisticsChannel {
	chatStatisticsChannelTemp := ChatStatisticsChannel{
		tdCommon:                       tdCommon{Type: "chatStatisticsChannel"},
		Period:                         period,
		MemberCount:                    memberCount,
		MeanViewCount:                  meanViewCount,
		MeanShareCount:                 meanShareCount,
		EnabledNotificationsPercentage: enabledNotificationsPercentage,
		MemberCountGraph:               memberCountGraph,
		JoinGraph:                      joinGraph,
		MuteGraph:                      muteGraph,
		ViewCountByHourGraph:           viewCountByHourGraph,
		ViewCountBySourceGraph:         viewCountBySourceGraph,
		JoinBySourceGraph:              joinBySourceGraph,
		LanguageGraph:                  languageGraph,
		MessageInteractionGraph:        messageInteractionGraph,
		InstantViewInteractionGraph:    instantViewInteractionGraph,
		RecentMessageInteractions:      recentMessageInteractions,
	}

	return &chatStatisticsChannelTemp
}

// UnmarshalJSON unmarshal to json
func (chatStatisticsChannel *ChatStatisticsChannel) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Period                         *DateRange                             `json:"period"`                           // A period to which the statistics applies
		MemberCount                    *StatisticalValue                      `json:"member_count"`                     // Number of members in the chat
		MeanViewCount                  *StatisticalValue                      `json:"mean_view_count"`                  // Mean number of times the recently sent messages was viewed
		MeanShareCount                 *StatisticalValue                      `json:"mean_share_count"`                 // Mean number of times the recently sent messages was shared
		EnabledNotificationsPercentage float64                                `json:"enabled_notifications_percentage"` // A percentage of users with enabled notifications for the chat
		RecentMessageInteractions      []ChatStatisticsMessageInteractionInfo `json:"recent_message_interactions"`      // Detailed statistics about number of views and shares of recently sent messages
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatStatisticsChannel.tdCommon = tempObj.tdCommon
	chatStatisticsChannel.Period = tempObj.Period
	chatStatisticsChannel.MemberCount = tempObj.MemberCount
	chatStatisticsChannel.MeanViewCount = tempObj.MeanViewCount
	chatStatisticsChannel.MeanShareCount = tempObj.MeanShareCount
	chatStatisticsChannel.EnabledNotificationsPercentage = tempObj.EnabledNotificationsPercentage
	chatStatisticsChannel.RecentMessageInteractions = tempObj.RecentMessageInteractions

	fieldMemberCountGraph, _ := unmarshalStatisticalGraph(objMap["member_count_graph"])
	chatStatisticsChannel.MemberCountGraph = fieldMemberCountGraph

	fieldJoinGraph, _ := unmarshalStatisticalGraph(objMap["join_graph"])
	chatStatisticsChannel.JoinGraph = fieldJoinGraph

	fieldMuteGraph, _ := unmarshalStatisticalGraph(objMap["mute_graph"])
	chatStatisticsChannel.MuteGraph = fieldMuteGraph

	fieldViewCountByHourGraph, _ := unmarshalStatisticalGraph(objMap["view_count_by_hour_graph"])
	chatStatisticsChannel.ViewCountByHourGraph = fieldViewCountByHourGraph

	fieldViewCountBySourceGraph, _ := unmarshalStatisticalGraph(objMap["view_count_by_source_graph"])
	chatStatisticsChannel.ViewCountBySourceGraph = fieldViewCountBySourceGraph

	fieldJoinBySourceGraph, _ := unmarshalStatisticalGraph(objMap["join_by_source_graph"])
	chatStatisticsChannel.JoinBySourceGraph = fieldJoinBySourceGraph

	fieldLanguageGraph, _ := unmarshalStatisticalGraph(objMap["language_graph"])
	chatStatisticsChannel.LanguageGraph = fieldLanguageGraph

	fieldMessageInteractionGraph, _ := unmarshalStatisticalGraph(objMap["message_interaction_graph"])
	chatStatisticsChannel.MessageInteractionGraph = fieldMessageInteractionGraph

	fieldInstantViewInteractionGraph, _ := unmarshalStatisticalGraph(objMap["instant_view_interaction_graph"])
	chatStatisticsChannel.InstantViewInteractionGraph = fieldInstantViewInteractionGraph

	return nil
}

// GetChatStatisticsEnum return the enum type of this object
func (chatStatisticsChannel *ChatStatisticsChannel) GetChatStatisticsEnum() ChatStatisticsEnum {
	return ChatStatisticsChannelType
}

// GetChatStatistics Returns detailed statistics about a chat. Currently, this method can be used only for supergroups and channels. Can be used only if supergroupFullInfo.can_get_statistics == true
// @param chatId Chat identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetChatStatistics(chatId int64, isDark bool) (ChatStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatStatistics",
		"chat_id": chatId,
		"is_dark": isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch ChatStatisticsEnum(result.Data["@type"].(string)) {

	case ChatStatisticsSupergroupType:
		var chatStatistics ChatStatisticsSupergroup
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	case ChatStatisticsChannelType:
		var chatStatistics ChatStatisticsChannel
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
