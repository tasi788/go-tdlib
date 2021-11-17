package tdlib

import (
	"encoding/json"
	"fmt"
)

// Update Contains notifications about data changes
type Update interface {
	GetUpdateEnum() UpdateEnum
}

// UpdateEnum Alias for abstract Update 'Sub-Classes', used as constant-enum here
type UpdateEnum string

// Update enums
const (
	UpdateAuthorizationStateType             UpdateEnum = "updateAuthorizationState"
	UpdateNewMessageType                     UpdateEnum = "updateNewMessage"
	UpdateMessageSendAcknowledgedType        UpdateEnum = "updateMessageSendAcknowledged"
	UpdateMessageSendSucceededType           UpdateEnum = "updateMessageSendSucceeded"
	UpdateMessageSendFailedType              UpdateEnum = "updateMessageSendFailed"
	UpdateMessageContentType                 UpdateEnum = "updateMessageContent"
	UpdateMessageEditedType                  UpdateEnum = "updateMessageEdited"
	UpdateMessageIsPinnedType                UpdateEnum = "updateMessageIsPinned"
	UpdateMessageInteractionInfoType         UpdateEnum = "updateMessageInteractionInfo"
	UpdateMessageContentOpenedType           UpdateEnum = "updateMessageContentOpened"
	UpdateMessageMentionReadType             UpdateEnum = "updateMessageMentionRead"
	UpdateMessageLiveLocationViewedType      UpdateEnum = "updateMessageLiveLocationViewed"
	UpdateNewChatType                        UpdateEnum = "updateNewChat"
	UpdateChatTitleType                      UpdateEnum = "updateChatTitle"
	UpdateChatPhotoType                      UpdateEnum = "updateChatPhoto"
	UpdateChatPermissionsType                UpdateEnum = "updateChatPermissions"
	UpdateChatLastMessageType                UpdateEnum = "updateChatLastMessage"
	UpdateChatPositionType                   UpdateEnum = "updateChatPosition"
	UpdateChatIsMarkedAsUnreadType           UpdateEnum = "updateChatIsMarkedAsUnread"
	UpdateChatIsBlockedType                  UpdateEnum = "updateChatIsBlocked"
	UpdateChatHasScheduledMessagesType       UpdateEnum = "updateChatHasScheduledMessages"
	UpdateChatVideoChatType                  UpdateEnum = "updateChatVideoChat"
	UpdateChatDefaultDisableNotificationType UpdateEnum = "updateChatDefaultDisableNotification"
	UpdateChatReadInboxType                  UpdateEnum = "updateChatReadInbox"
	UpdateChatReadOutboxType                 UpdateEnum = "updateChatReadOutbox"
	UpdateChatUnreadMentionCountType         UpdateEnum = "updateChatUnreadMentionCount"
	UpdateChatNotificationSettingsType       UpdateEnum = "updateChatNotificationSettings"
	UpdateScopeNotificationSettingsType      UpdateEnum = "updateScopeNotificationSettings"
	UpdateChatMessageTtlSettingType          UpdateEnum = "updateChatMessageTtlSetting"
	UpdateChatActionBarType                  UpdateEnum = "updateChatActionBar"
	UpdateChatThemeType                      UpdateEnum = "updateChatTheme"
	UpdateChatPendingJoinRequestsType        UpdateEnum = "updateChatPendingJoinRequests"
	UpdateChatReplyMarkupType                UpdateEnum = "updateChatReplyMarkup"
	UpdateChatDraftMessageType               UpdateEnum = "updateChatDraftMessage"
	UpdateChatFiltersType                    UpdateEnum = "updateChatFilters"
	UpdateChatOnlineMemberCountType          UpdateEnum = "updateChatOnlineMemberCount"
	UpdateNotificationType                   UpdateEnum = "updateNotification"
	UpdateNotificationGroupType              UpdateEnum = "updateNotificationGroup"
	UpdateActiveNotificationsType            UpdateEnum = "updateActiveNotifications"
	UpdateHavePendingNotificationsType       UpdateEnum = "updateHavePendingNotifications"
	UpdateDeleteMessagesType                 UpdateEnum = "updateDeleteMessages"
	UpdateUserChatActionType                 UpdateEnum = "updateUserChatAction"
	UpdateUserStatusType                     UpdateEnum = "updateUserStatus"
	UpdateUserType                           UpdateEnum = "updateUser"
	UpdateBasicGroupType                     UpdateEnum = "updateBasicGroup"
	UpdateSupergroupType                     UpdateEnum = "updateSupergroup"
	UpdateSecretChatType                     UpdateEnum = "updateSecretChat"
	UpdateUserFullInfoType                   UpdateEnum = "updateUserFullInfo"
	UpdateBasicGroupFullInfoType             UpdateEnum = "updateBasicGroupFullInfo"
	UpdateSupergroupFullInfoType             UpdateEnum = "updateSupergroupFullInfo"
	UpdateServiceNotificationType            UpdateEnum = "updateServiceNotification"
	UpdateFileType                           UpdateEnum = "updateFile"
	UpdateFileGenerationStartType            UpdateEnum = "updateFileGenerationStart"
	UpdateFileGenerationStopType             UpdateEnum = "updateFileGenerationStop"
	UpdateCallType                           UpdateEnum = "updateCall"
	UpdateGroupCallType                      UpdateEnum = "updateGroupCall"
	UpdateGroupCallParticipantType           UpdateEnum = "updateGroupCallParticipant"
	UpdateNewCallSignalingDataType           UpdateEnum = "updateNewCallSignalingData"
	UpdateUserPrivacySettingRulesType        UpdateEnum = "updateUserPrivacySettingRules"
	UpdateUnreadMessageCountType             UpdateEnum = "updateUnreadMessageCount"
	UpdateUnreadChatCountType                UpdateEnum = "updateUnreadChatCount"
	UpdateOptionType                         UpdateEnum = "updateOption"
	UpdateStickerSetType                     UpdateEnum = "updateStickerSet"
	UpdateInstalledStickerSetsType           UpdateEnum = "updateInstalledStickerSets"
	UpdateTrendingStickerSetsType            UpdateEnum = "updateTrendingStickerSets"
	UpdateRecentStickersType                 UpdateEnum = "updateRecentStickers"
	UpdateFavoriteStickersType               UpdateEnum = "updateFavoriteStickers"
	UpdateSavedAnimationsType                UpdateEnum = "updateSavedAnimations"
	UpdateSelectedBackgroundType             UpdateEnum = "updateSelectedBackground"
	UpdateChatThemesType                     UpdateEnum = "updateChatThemes"
	UpdateLanguagePackStringsType            UpdateEnum = "updateLanguagePackStrings"
	UpdateConnectionStateType                UpdateEnum = "updateConnectionState"
	UpdateTermsOfServiceType                 UpdateEnum = "updateTermsOfService"
	UpdateUsersNearbyType                    UpdateEnum = "updateUsersNearby"
	UpdateDiceEmojisType                     UpdateEnum = "updateDiceEmojis"
	UpdateAnimatedEmojiMessageClickedType    UpdateEnum = "updateAnimatedEmojiMessageClicked"
	UpdateAnimationSearchParametersType      UpdateEnum = "updateAnimationSearchParameters"
	UpdateSuggestedActionsType               UpdateEnum = "updateSuggestedActions"
	UpdateNewInlineQueryType                 UpdateEnum = "updateNewInlineQuery"
	UpdateNewChosenInlineResultType          UpdateEnum = "updateNewChosenInlineResult"
	UpdateNewCallbackQueryType               UpdateEnum = "updateNewCallbackQuery"
	UpdateNewInlineCallbackQueryType         UpdateEnum = "updateNewInlineCallbackQuery"
	UpdateNewShippingQueryType               UpdateEnum = "updateNewShippingQuery"
	UpdateNewPreCheckoutQueryType            UpdateEnum = "updateNewPreCheckoutQuery"
	UpdateNewCustomEventType                 UpdateEnum = "updateNewCustomEvent"
	UpdateNewCustomQueryType                 UpdateEnum = "updateNewCustomQuery"
	UpdatePollType                           UpdateEnum = "updatePoll"
	UpdatePollAnswerType                     UpdateEnum = "updatePollAnswer"
	UpdateChatMemberType                     UpdateEnum = "updateChatMember"
	UpdateNewChatJoinRequestType             UpdateEnum = "updateNewChatJoinRequest"
)

func unmarshalUpdate(rawMsg *json.RawMessage) (Update, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UpdateEnum(objMap["@type"].(string)) {
	case UpdateAuthorizationStateType:
		var updateAuthorizationState UpdateAuthorizationState
		err := json.Unmarshal(*rawMsg, &updateAuthorizationState)
		return &updateAuthorizationState, err

	case UpdateNewMessageType:
		var updateNewMessage UpdateNewMessage
		err := json.Unmarshal(*rawMsg, &updateNewMessage)
		return &updateNewMessage, err

	case UpdateMessageSendAcknowledgedType:
		var updateMessageSendAcknowledged UpdateMessageSendAcknowledged
		err := json.Unmarshal(*rawMsg, &updateMessageSendAcknowledged)
		return &updateMessageSendAcknowledged, err

	case UpdateMessageSendSucceededType:
		var updateMessageSendSucceeded UpdateMessageSendSucceeded
		err := json.Unmarshal(*rawMsg, &updateMessageSendSucceeded)
		return &updateMessageSendSucceeded, err

	case UpdateMessageSendFailedType:
		var updateMessageSendFailed UpdateMessageSendFailed
		err := json.Unmarshal(*rawMsg, &updateMessageSendFailed)
		return &updateMessageSendFailed, err

	case UpdateMessageContentType:
		var updateMessageContent UpdateMessageContent
		err := json.Unmarshal(*rawMsg, &updateMessageContent)
		return &updateMessageContent, err

	case UpdateMessageEditedType:
		var updateMessageEdited UpdateMessageEdited
		err := json.Unmarshal(*rawMsg, &updateMessageEdited)
		return &updateMessageEdited, err

	case UpdateMessageIsPinnedType:
		var updateMessageIsPinned UpdateMessageIsPinned
		err := json.Unmarshal(*rawMsg, &updateMessageIsPinned)
		return &updateMessageIsPinned, err

	case UpdateMessageInteractionInfoType:
		var updateMessageInteractionInfo UpdateMessageInteractionInfo
		err := json.Unmarshal(*rawMsg, &updateMessageInteractionInfo)
		return &updateMessageInteractionInfo, err

	case UpdateMessageContentOpenedType:
		var updateMessageContentOpened UpdateMessageContentOpened
		err := json.Unmarshal(*rawMsg, &updateMessageContentOpened)
		return &updateMessageContentOpened, err

	case UpdateMessageMentionReadType:
		var updateMessageMentionRead UpdateMessageMentionRead
		err := json.Unmarshal(*rawMsg, &updateMessageMentionRead)
		return &updateMessageMentionRead, err

	case UpdateMessageLiveLocationViewedType:
		var updateMessageLiveLocationViewed UpdateMessageLiveLocationViewed
		err := json.Unmarshal(*rawMsg, &updateMessageLiveLocationViewed)
		return &updateMessageLiveLocationViewed, err

	case UpdateNewChatType:
		var updateNewChat UpdateNewChat
		err := json.Unmarshal(*rawMsg, &updateNewChat)
		return &updateNewChat, err

	case UpdateChatTitleType:
		var updateChatTitle UpdateChatTitle
		err := json.Unmarshal(*rawMsg, &updateChatTitle)
		return &updateChatTitle, err

	case UpdateChatPhotoType:
		var updateChatPhoto UpdateChatPhoto
		err := json.Unmarshal(*rawMsg, &updateChatPhoto)
		return &updateChatPhoto, err

	case UpdateChatPermissionsType:
		var updateChatPermissions UpdateChatPermissions
		err := json.Unmarshal(*rawMsg, &updateChatPermissions)
		return &updateChatPermissions, err

	case UpdateChatLastMessageType:
		var updateChatLastMessage UpdateChatLastMessage
		err := json.Unmarshal(*rawMsg, &updateChatLastMessage)
		return &updateChatLastMessage, err

	case UpdateChatPositionType:
		var updateChatPosition UpdateChatPosition
		err := json.Unmarshal(*rawMsg, &updateChatPosition)
		return &updateChatPosition, err

	case UpdateChatIsMarkedAsUnreadType:
		var updateChatIsMarkedAsUnread UpdateChatIsMarkedAsUnread
		err := json.Unmarshal(*rawMsg, &updateChatIsMarkedAsUnread)
		return &updateChatIsMarkedAsUnread, err

	case UpdateChatIsBlockedType:
		var updateChatIsBlocked UpdateChatIsBlocked
		err := json.Unmarshal(*rawMsg, &updateChatIsBlocked)
		return &updateChatIsBlocked, err

	case UpdateChatHasScheduledMessagesType:
		var updateChatHasScheduledMessages UpdateChatHasScheduledMessages
		err := json.Unmarshal(*rawMsg, &updateChatHasScheduledMessages)
		return &updateChatHasScheduledMessages, err

	case UpdateChatVideoChatType:
		var updateChatVideoChat UpdateChatVideoChat
		err := json.Unmarshal(*rawMsg, &updateChatVideoChat)
		return &updateChatVideoChat, err

	case UpdateChatDefaultDisableNotificationType:
		var updateChatDefaultDisableNotification UpdateChatDefaultDisableNotification
		err := json.Unmarshal(*rawMsg, &updateChatDefaultDisableNotification)
		return &updateChatDefaultDisableNotification, err

	case UpdateChatReadInboxType:
		var updateChatReadInbox UpdateChatReadInbox
		err := json.Unmarshal(*rawMsg, &updateChatReadInbox)
		return &updateChatReadInbox, err

	case UpdateChatReadOutboxType:
		var updateChatReadOutbox UpdateChatReadOutbox
		err := json.Unmarshal(*rawMsg, &updateChatReadOutbox)
		return &updateChatReadOutbox, err

	case UpdateChatUnreadMentionCountType:
		var updateChatUnreadMentionCount UpdateChatUnreadMentionCount
		err := json.Unmarshal(*rawMsg, &updateChatUnreadMentionCount)
		return &updateChatUnreadMentionCount, err

	case UpdateChatNotificationSettingsType:
		var updateChatNotificationSettings UpdateChatNotificationSettings
		err := json.Unmarshal(*rawMsg, &updateChatNotificationSettings)
		return &updateChatNotificationSettings, err

	case UpdateScopeNotificationSettingsType:
		var updateScopeNotificationSettings UpdateScopeNotificationSettings
		err := json.Unmarshal(*rawMsg, &updateScopeNotificationSettings)
		return &updateScopeNotificationSettings, err

	case UpdateChatMessageTtlSettingType:
		var updateChatMessageTtlSetting UpdateChatMessageTtlSetting
		err := json.Unmarshal(*rawMsg, &updateChatMessageTtlSetting)
		return &updateChatMessageTtlSetting, err

	case UpdateChatActionBarType:
		var updateChatActionBar UpdateChatActionBar
		err := json.Unmarshal(*rawMsg, &updateChatActionBar)
		return &updateChatActionBar, err

	case UpdateChatThemeType:
		var updateChatTheme UpdateChatTheme
		err := json.Unmarshal(*rawMsg, &updateChatTheme)
		return &updateChatTheme, err

	case UpdateChatPendingJoinRequestsType:
		var updateChatPendingJoinRequests UpdateChatPendingJoinRequests
		err := json.Unmarshal(*rawMsg, &updateChatPendingJoinRequests)
		return &updateChatPendingJoinRequests, err

	case UpdateChatReplyMarkupType:
		var updateChatReplyMarkup UpdateChatReplyMarkup
		err := json.Unmarshal(*rawMsg, &updateChatReplyMarkup)
		return &updateChatReplyMarkup, err

	case UpdateChatDraftMessageType:
		var updateChatDraftMessage UpdateChatDraftMessage
		err := json.Unmarshal(*rawMsg, &updateChatDraftMessage)
		return &updateChatDraftMessage, err

	case UpdateChatFiltersType:
		var updateChatFilters UpdateChatFilters
		err := json.Unmarshal(*rawMsg, &updateChatFilters)
		return &updateChatFilters, err

	case UpdateChatOnlineMemberCountType:
		var updateChatOnlineMemberCount UpdateChatOnlineMemberCount
		err := json.Unmarshal(*rawMsg, &updateChatOnlineMemberCount)
		return &updateChatOnlineMemberCount, err

	case UpdateNotificationType:
		var updateNotification UpdateNotification
		err := json.Unmarshal(*rawMsg, &updateNotification)
		return &updateNotification, err

	case UpdateNotificationGroupType:
		var updateNotificationGroup UpdateNotificationGroup
		err := json.Unmarshal(*rawMsg, &updateNotificationGroup)
		return &updateNotificationGroup, err

	case UpdateActiveNotificationsType:
		var updateActiveNotifications UpdateActiveNotifications
		err := json.Unmarshal(*rawMsg, &updateActiveNotifications)
		return &updateActiveNotifications, err

	case UpdateHavePendingNotificationsType:
		var updateHavePendingNotifications UpdateHavePendingNotifications
		err := json.Unmarshal(*rawMsg, &updateHavePendingNotifications)
		return &updateHavePendingNotifications, err

	case UpdateDeleteMessagesType:
		var updateDeleteMessages UpdateDeleteMessages
		err := json.Unmarshal(*rawMsg, &updateDeleteMessages)
		return &updateDeleteMessages, err

	case UpdateUserChatActionType:
		var updateUserChatAction UpdateUserChatAction
		err := json.Unmarshal(*rawMsg, &updateUserChatAction)
		return &updateUserChatAction, err

	case UpdateUserStatusType:
		var updateUserStatus UpdateUserStatus
		err := json.Unmarshal(*rawMsg, &updateUserStatus)
		return &updateUserStatus, err

	case UpdateUserType:
		var updateUser UpdateUser
		err := json.Unmarshal(*rawMsg, &updateUser)
		return &updateUser, err

	case UpdateBasicGroupType:
		var updateBasicGroup UpdateBasicGroup
		err := json.Unmarshal(*rawMsg, &updateBasicGroup)
		return &updateBasicGroup, err

	case UpdateSupergroupType:
		var updateSupergroup UpdateSupergroup
		err := json.Unmarshal(*rawMsg, &updateSupergroup)
		return &updateSupergroup, err

	case UpdateSecretChatType:
		var updateSecretChat UpdateSecretChat
		err := json.Unmarshal(*rawMsg, &updateSecretChat)
		return &updateSecretChat, err

	case UpdateUserFullInfoType:
		var updateUserFullInfo UpdateUserFullInfo
		err := json.Unmarshal(*rawMsg, &updateUserFullInfo)
		return &updateUserFullInfo, err

	case UpdateBasicGroupFullInfoType:
		var updateBasicGroupFullInfo UpdateBasicGroupFullInfo
		err := json.Unmarshal(*rawMsg, &updateBasicGroupFullInfo)
		return &updateBasicGroupFullInfo, err

	case UpdateSupergroupFullInfoType:
		var updateSupergroupFullInfo UpdateSupergroupFullInfo
		err := json.Unmarshal(*rawMsg, &updateSupergroupFullInfo)
		return &updateSupergroupFullInfo, err

	case UpdateServiceNotificationType:
		var updateServiceNotification UpdateServiceNotification
		err := json.Unmarshal(*rawMsg, &updateServiceNotification)
		return &updateServiceNotification, err

	case UpdateFileType:
		var updateFile UpdateFile
		err := json.Unmarshal(*rawMsg, &updateFile)
		return &updateFile, err

	case UpdateFileGenerationStartType:
		var updateFileGenerationStart UpdateFileGenerationStart
		err := json.Unmarshal(*rawMsg, &updateFileGenerationStart)
		return &updateFileGenerationStart, err

	case UpdateFileGenerationStopType:
		var updateFileGenerationStop UpdateFileGenerationStop
		err := json.Unmarshal(*rawMsg, &updateFileGenerationStop)
		return &updateFileGenerationStop, err

	case UpdateCallType:
		var updateCall UpdateCall
		err := json.Unmarshal(*rawMsg, &updateCall)
		return &updateCall, err

	case UpdateGroupCallType:
		var updateGroupCall UpdateGroupCall
		err := json.Unmarshal(*rawMsg, &updateGroupCall)
		return &updateGroupCall, err

	case UpdateGroupCallParticipantType:
		var updateGroupCallParticipant UpdateGroupCallParticipant
		err := json.Unmarshal(*rawMsg, &updateGroupCallParticipant)
		return &updateGroupCallParticipant, err

	case UpdateNewCallSignalingDataType:
		var updateNewCallSignalingData UpdateNewCallSignalingData
		err := json.Unmarshal(*rawMsg, &updateNewCallSignalingData)
		return &updateNewCallSignalingData, err

	case UpdateUserPrivacySettingRulesType:
		var updateUserPrivacySettingRules UpdateUserPrivacySettingRules
		err := json.Unmarshal(*rawMsg, &updateUserPrivacySettingRules)
		return &updateUserPrivacySettingRules, err

	case UpdateUnreadMessageCountType:
		var updateUnreadMessageCount UpdateUnreadMessageCount
		err := json.Unmarshal(*rawMsg, &updateUnreadMessageCount)
		return &updateUnreadMessageCount, err

	case UpdateUnreadChatCountType:
		var updateUnreadChatCount UpdateUnreadChatCount
		err := json.Unmarshal(*rawMsg, &updateUnreadChatCount)
		return &updateUnreadChatCount, err

	case UpdateOptionType:
		var updateOption UpdateOption
		err := json.Unmarshal(*rawMsg, &updateOption)
		return &updateOption, err

	case UpdateStickerSetType:
		var updateStickerSet UpdateStickerSet
		err := json.Unmarshal(*rawMsg, &updateStickerSet)
		return &updateStickerSet, err

	case UpdateInstalledStickerSetsType:
		var updateInstalledStickerSets UpdateInstalledStickerSets
		err := json.Unmarshal(*rawMsg, &updateInstalledStickerSets)
		return &updateInstalledStickerSets, err

	case UpdateTrendingStickerSetsType:
		var updateTrendingStickerSets UpdateTrendingStickerSets
		err := json.Unmarshal(*rawMsg, &updateTrendingStickerSets)
		return &updateTrendingStickerSets, err

	case UpdateRecentStickersType:
		var updateRecentStickers UpdateRecentStickers
		err := json.Unmarshal(*rawMsg, &updateRecentStickers)
		return &updateRecentStickers, err

	case UpdateFavoriteStickersType:
		var updateFavoriteStickers UpdateFavoriteStickers
		err := json.Unmarshal(*rawMsg, &updateFavoriteStickers)
		return &updateFavoriteStickers, err

	case UpdateSavedAnimationsType:
		var updateSavedAnimations UpdateSavedAnimations
		err := json.Unmarshal(*rawMsg, &updateSavedAnimations)
		return &updateSavedAnimations, err

	case UpdateSelectedBackgroundType:
		var updateSelectedBackground UpdateSelectedBackground
		err := json.Unmarshal(*rawMsg, &updateSelectedBackground)
		return &updateSelectedBackground, err

	case UpdateChatThemesType:
		var updateChatThemes UpdateChatThemes
		err := json.Unmarshal(*rawMsg, &updateChatThemes)
		return &updateChatThemes, err

	case UpdateLanguagePackStringsType:
		var updateLanguagePackStrings UpdateLanguagePackStrings
		err := json.Unmarshal(*rawMsg, &updateLanguagePackStrings)
		return &updateLanguagePackStrings, err

	case UpdateConnectionStateType:
		var updateConnectionState UpdateConnectionState
		err := json.Unmarshal(*rawMsg, &updateConnectionState)
		return &updateConnectionState, err

	case UpdateTermsOfServiceType:
		var updateTermsOfService UpdateTermsOfService
		err := json.Unmarshal(*rawMsg, &updateTermsOfService)
		return &updateTermsOfService, err

	case UpdateUsersNearbyType:
		var updateUsersNearby UpdateUsersNearby
		err := json.Unmarshal(*rawMsg, &updateUsersNearby)
		return &updateUsersNearby, err

	case UpdateDiceEmojisType:
		var updateDiceEmojis UpdateDiceEmojis
		err := json.Unmarshal(*rawMsg, &updateDiceEmojis)
		return &updateDiceEmojis, err

	case UpdateAnimatedEmojiMessageClickedType:
		var updateAnimatedEmojiMessageClicked UpdateAnimatedEmojiMessageClicked
		err := json.Unmarshal(*rawMsg, &updateAnimatedEmojiMessageClicked)
		return &updateAnimatedEmojiMessageClicked, err

	case UpdateAnimationSearchParametersType:
		var updateAnimationSearchParameters UpdateAnimationSearchParameters
		err := json.Unmarshal(*rawMsg, &updateAnimationSearchParameters)
		return &updateAnimationSearchParameters, err

	case UpdateSuggestedActionsType:
		var updateSuggestedActions UpdateSuggestedActions
		err := json.Unmarshal(*rawMsg, &updateSuggestedActions)
		return &updateSuggestedActions, err

	case UpdateNewInlineQueryType:
		var updateNewInlineQuery UpdateNewInlineQuery
		err := json.Unmarshal(*rawMsg, &updateNewInlineQuery)
		return &updateNewInlineQuery, err

	case UpdateNewChosenInlineResultType:
		var updateNewChosenInlineResult UpdateNewChosenInlineResult
		err := json.Unmarshal(*rawMsg, &updateNewChosenInlineResult)
		return &updateNewChosenInlineResult, err

	case UpdateNewCallbackQueryType:
		var updateNewCallbackQuery UpdateNewCallbackQuery
		err := json.Unmarshal(*rawMsg, &updateNewCallbackQuery)
		return &updateNewCallbackQuery, err

	case UpdateNewInlineCallbackQueryType:
		var updateNewInlineCallbackQuery UpdateNewInlineCallbackQuery
		err := json.Unmarshal(*rawMsg, &updateNewInlineCallbackQuery)
		return &updateNewInlineCallbackQuery, err

	case UpdateNewShippingQueryType:
		var updateNewShippingQuery UpdateNewShippingQuery
		err := json.Unmarshal(*rawMsg, &updateNewShippingQuery)
		return &updateNewShippingQuery, err

	case UpdateNewPreCheckoutQueryType:
		var updateNewPreCheckoutQuery UpdateNewPreCheckoutQuery
		err := json.Unmarshal(*rawMsg, &updateNewPreCheckoutQuery)
		return &updateNewPreCheckoutQuery, err

	case UpdateNewCustomEventType:
		var updateNewCustomEvent UpdateNewCustomEvent
		err := json.Unmarshal(*rawMsg, &updateNewCustomEvent)
		return &updateNewCustomEvent, err

	case UpdateNewCustomQueryType:
		var updateNewCustomQuery UpdateNewCustomQuery
		err := json.Unmarshal(*rawMsg, &updateNewCustomQuery)
		return &updateNewCustomQuery, err

	case UpdatePollType:
		var updatePoll UpdatePoll
		err := json.Unmarshal(*rawMsg, &updatePoll)
		return &updatePoll, err

	case UpdatePollAnswerType:
		var updatePollAnswer UpdatePollAnswer
		err := json.Unmarshal(*rawMsg, &updatePollAnswer)
		return &updatePollAnswer, err

	case UpdateChatMemberType:
		var updateChatMember UpdateChatMember
		err := json.Unmarshal(*rawMsg, &updateChatMember)
		return &updateChatMember, err

	case UpdateNewChatJoinRequestType:
		var updateNewChatJoinRequest UpdateNewChatJoinRequest
		err := json.Unmarshal(*rawMsg, &updateNewChatJoinRequest)
		return &updateNewChatJoinRequest, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// UpdateAuthorizationState The user authorization state has changed
type UpdateAuthorizationState struct {
	tdCommon
	AuthorizationState AuthorizationState `json:"authorization_state"` // New authorization state
}

// MessageType return the string telegram-type of UpdateAuthorizationState
func (updateAuthorizationState *UpdateAuthorizationState) MessageType() string {
	return "updateAuthorizationState"
}

// NewUpdateAuthorizationState creates a new UpdateAuthorizationState
//
// @param authorizationState New authorization state
func NewUpdateAuthorizationState(authorizationState AuthorizationState) *UpdateAuthorizationState {
	updateAuthorizationStateTemp := UpdateAuthorizationState{
		tdCommon:           tdCommon{Type: "updateAuthorizationState"},
		AuthorizationState: authorizationState,
	}

	return &updateAuthorizationStateTemp
}

// UnmarshalJSON unmarshal to json
func (updateAuthorizationState *UpdateAuthorizationState) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateAuthorizationState.tdCommon = tempObj.tdCommon

	fieldAuthorizationState, _ := unmarshalAuthorizationState(objMap["authorization_state"])
	updateAuthorizationState.AuthorizationState = fieldAuthorizationState

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateAuthorizationState *UpdateAuthorizationState) GetUpdateEnum() UpdateEnum {
	return UpdateAuthorizationStateType
}

// UpdateNewMessage A new message was received; can also be an outgoing message
type UpdateNewMessage struct {
	tdCommon
	Message *Message `json:"message"` // The new message
}

// MessageType return the string telegram-type of UpdateNewMessage
func (updateNewMessage *UpdateNewMessage) MessageType() string {
	return "updateNewMessage"
}

// NewUpdateNewMessage creates a new UpdateNewMessage
//
// @param message The new message
func NewUpdateNewMessage(message *Message) *UpdateNewMessage {
	updateNewMessageTemp := UpdateNewMessage{
		tdCommon: tdCommon{Type: "updateNewMessage"},
		Message:  message,
	}

	return &updateNewMessageTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewMessage *UpdateNewMessage) GetUpdateEnum() UpdateEnum {
	return UpdateNewMessageType
}

// UpdateMessageSendAcknowledged A request to send a message has reached the Telegram server. This doesn't mean that the message will be sent successfully or even that the send message request will be processed. This update will be sent only if the option "use_quick_ack" is set to true. This update may be sent multiple times for the same message
type UpdateMessageSendAcknowledged struct {
	tdCommon
	ChatId    int64 `json:"chat_id"`    // The chat identifier of the sent message
	MessageId int64 `json:"message_id"` // A temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendAcknowledged
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) MessageType() string {
	return "updateMessageSendAcknowledged"
}

// NewUpdateMessageSendAcknowledged creates a new UpdateMessageSendAcknowledged
//
// @param chatId The chat identifier of the sent message
// @param messageId A temporary message identifier
func NewUpdateMessageSendAcknowledged(chatId int64, messageId int64) *UpdateMessageSendAcknowledged {
	updateMessageSendAcknowledgedTemp := UpdateMessageSendAcknowledged{
		tdCommon:  tdCommon{Type: "updateMessageSendAcknowledged"},
		ChatId:    chatId,
		MessageId: messageId,
	}

	return &updateMessageSendAcknowledgedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendAcknowledgedType
}

// UpdateMessageSendSucceeded A message has been successfully sent
type UpdateMessageSendSucceeded struct {
	tdCommon
	Message      *Message `json:"message"`        // The sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change
	OldMessageId int64    `json:"old_message_id"` // The previous temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendSucceeded
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) MessageType() string {
	return "updateMessageSendSucceeded"
}

// NewUpdateMessageSendSucceeded creates a new UpdateMessageSendSucceeded
//
// @param message The sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change
// @param oldMessageId The previous temporary message identifier
func NewUpdateMessageSendSucceeded(message *Message, oldMessageId int64) *UpdateMessageSendSucceeded {
	updateMessageSendSucceededTemp := UpdateMessageSendSucceeded{
		tdCommon:     tdCommon{Type: "updateMessageSendSucceeded"},
		Message:      message,
		OldMessageId: oldMessageId,
	}

	return &updateMessageSendSucceededTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendSucceededType
}

// UpdateMessageSendFailed A message failed to send. Be aware that some messages being sent can be irrecoverably deleted, in which case updateDeleteMessages will be received instead of this update
type UpdateMessageSendFailed struct {
	tdCommon
	Message      *Message `json:"message"`        // The failed to send message
	OldMessageId int64    `json:"old_message_id"` // The previous temporary message identifier
	ErrorCode    int32    `json:"error_code"`     // An error code
	ErrorMessage string   `json:"error_message"`  // Error message
}

// MessageType return the string telegram-type of UpdateMessageSendFailed
func (updateMessageSendFailed *UpdateMessageSendFailed) MessageType() string {
	return "updateMessageSendFailed"
}

// NewUpdateMessageSendFailed creates a new UpdateMessageSendFailed
//
// @param message The failed to send message
// @param oldMessageId The previous temporary message identifier
// @param errorCode An error code
// @param errorMessage Error message
func NewUpdateMessageSendFailed(message *Message, oldMessageId int64, errorCode int32, errorMessage string) *UpdateMessageSendFailed {
	updateMessageSendFailedTemp := UpdateMessageSendFailed{
		tdCommon:     tdCommon{Type: "updateMessageSendFailed"},
		Message:      message,
		OldMessageId: oldMessageId,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}

	return &updateMessageSendFailedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendFailed *UpdateMessageSendFailed) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendFailedType
}

// UpdateMessageContent The message content has changed
type UpdateMessageContent struct {
	tdCommon
	ChatId     int64          `json:"chat_id"`     // Chat identifier
	MessageId  int64          `json:"message_id"`  // Message identifier
	NewContent MessageContent `json:"new_content"` // New message content
}

// MessageType return the string telegram-type of UpdateMessageContent
func (updateMessageContent *UpdateMessageContent) MessageType() string {
	return "updateMessageContent"
}

// NewUpdateMessageContent creates a new UpdateMessageContent
//
// @param chatId Chat identifier
// @param messageId Message identifier
// @param newContent New message content
func NewUpdateMessageContent(chatId int64, messageId int64, newContent MessageContent) *UpdateMessageContent {
	updateMessageContentTemp := UpdateMessageContent{
		tdCommon:   tdCommon{Type: "updateMessageContent"},
		ChatId:     chatId,
		MessageId:  messageId,
		NewContent: newContent,
	}

	return &updateMessageContentTemp
}

// UnmarshalJSON unmarshal to json
func (updateMessageContent *UpdateMessageContent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatId    int64 `json:"chat_id"`    // Chat identifier
		MessageId int64 `json:"message_id"` // Message identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateMessageContent.tdCommon = tempObj.tdCommon
	updateMessageContent.ChatId = tempObj.ChatId
	updateMessageContent.MessageId = tempObj.MessageId

	fieldNewContent, _ := unmarshalMessageContent(objMap["new_content"])
	updateMessageContent.NewContent = fieldNewContent

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateMessageContent *UpdateMessageContent) GetUpdateEnum() UpdateEnum {
	return UpdateMessageContentType
}

// UpdateMessageEdited A message was edited. Changes in the message content will come in a separate updateMessageContent
type UpdateMessageEdited struct {
	tdCommon
	ChatId      int64       `json:"chat_id"`      // Chat identifier
	MessageId   int64       `json:"message_id"`   // Message identifier
	EditDate    int32       `json:"edit_date"`    // Point in time (Unix timestamp) when the message was edited
	ReplyMarkup ReplyMarkup `json:"reply_markup"` // New message reply markup; may be null
}

// MessageType return the string telegram-type of UpdateMessageEdited
func (updateMessageEdited *UpdateMessageEdited) MessageType() string {
	return "updateMessageEdited"
}

// NewUpdateMessageEdited creates a new UpdateMessageEdited
//
// @param chatId Chat identifier
// @param messageId Message identifier
// @param editDate Point in time (Unix timestamp) when the message was edited
// @param replyMarkup New message reply markup; may be null
func NewUpdateMessageEdited(chatId int64, messageId int64, editDate int32, replyMarkup ReplyMarkup) *UpdateMessageEdited {
	updateMessageEditedTemp := UpdateMessageEdited{
		tdCommon:    tdCommon{Type: "updateMessageEdited"},
		ChatId:      chatId,
		MessageId:   messageId,
		EditDate:    editDate,
		ReplyMarkup: replyMarkup,
	}

	return &updateMessageEditedTemp
}

// UnmarshalJSON unmarshal to json
func (updateMessageEdited *UpdateMessageEdited) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatId    int64 `json:"chat_id"`    // Chat identifier
		MessageId int64 `json:"message_id"` // Message identifier
		EditDate  int32 `json:"edit_date"`  // Point in time (Unix timestamp) when the message was edited

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateMessageEdited.tdCommon = tempObj.tdCommon
	updateMessageEdited.ChatId = tempObj.ChatId
	updateMessageEdited.MessageId = tempObj.MessageId
	updateMessageEdited.EditDate = tempObj.EditDate

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	updateMessageEdited.ReplyMarkup = fieldReplyMarkup

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateMessageEdited *UpdateMessageEdited) GetUpdateEnum() UpdateEnum {
	return UpdateMessageEditedType
}

// UpdateMessageIsPinned The message pinned state was changed
type UpdateMessageIsPinned struct {
	tdCommon
	ChatId    int64 `json:"chat_id"`    // Chat identifier
	MessageId int64 `json:"message_id"` // The message identifier
	IsPinned  bool  `json:"is_pinned"`  // True, if the message is pinned
}

// MessageType return the string telegram-type of UpdateMessageIsPinned
func (updateMessageIsPinned *UpdateMessageIsPinned) MessageType() string {
	return "updateMessageIsPinned"
}

// NewUpdateMessageIsPinned creates a new UpdateMessageIsPinned
//
// @param chatId Chat identifier
// @param messageId The message identifier
// @param isPinned True, if the message is pinned
func NewUpdateMessageIsPinned(chatId int64, messageId int64, isPinned bool) *UpdateMessageIsPinned {
	updateMessageIsPinnedTemp := UpdateMessageIsPinned{
		tdCommon:  tdCommon{Type: "updateMessageIsPinned"},
		ChatId:    chatId,
		MessageId: messageId,
		IsPinned:  isPinned,
	}

	return &updateMessageIsPinnedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageIsPinned *UpdateMessageIsPinned) GetUpdateEnum() UpdateEnum {
	return UpdateMessageIsPinnedType
}

// UpdateMessageInteractionInfo The information about interactions with a message has changed
type UpdateMessageInteractionInfo struct {
	tdCommon
	ChatId          int64                   `json:"chat_id"`          // Chat identifier
	MessageId       int64                   `json:"message_id"`       // Message identifier
	InteractionInfo *MessageInteractionInfo `json:"interaction_info"` // New information about interactions with the message; may be null
}

// MessageType return the string telegram-type of UpdateMessageInteractionInfo
func (updateMessageInteractionInfo *UpdateMessageInteractionInfo) MessageType() string {
	return "updateMessageInteractionInfo"
}

// NewUpdateMessageInteractionInfo creates a new UpdateMessageInteractionInfo
//
// @param chatId Chat identifier
// @param messageId Message identifier
// @param interactionInfo New information about interactions with the message; may be null
func NewUpdateMessageInteractionInfo(chatId int64, messageId int64, interactionInfo *MessageInteractionInfo) *UpdateMessageInteractionInfo {
	updateMessageInteractionInfoTemp := UpdateMessageInteractionInfo{
		tdCommon:        tdCommon{Type: "updateMessageInteractionInfo"},
		ChatId:          chatId,
		MessageId:       messageId,
		InteractionInfo: interactionInfo,
	}

	return &updateMessageInteractionInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageInteractionInfo *UpdateMessageInteractionInfo) GetUpdateEnum() UpdateEnum {
	return UpdateMessageInteractionInfoType
}

// UpdateMessageContentOpened The message content was opened. Updates voice note messages to "listened", video note messages to "viewed" and starts the TTL timer for self-destructing messages
type UpdateMessageContentOpened struct {
	tdCommon
	ChatId    int64 `json:"chat_id"`    // Chat identifier
	MessageId int64 `json:"message_id"` // Message identifier
}

// MessageType return the string telegram-type of UpdateMessageContentOpened
func (updateMessageContentOpened *UpdateMessageContentOpened) MessageType() string {
	return "updateMessageContentOpened"
}

// NewUpdateMessageContentOpened creates a new UpdateMessageContentOpened
//
// @param chatId Chat identifier
// @param messageId Message identifier
func NewUpdateMessageContentOpened(chatId int64, messageId int64) *UpdateMessageContentOpened {
	updateMessageContentOpenedTemp := UpdateMessageContentOpened{
		tdCommon:  tdCommon{Type: "updateMessageContentOpened"},
		ChatId:    chatId,
		MessageId: messageId,
	}

	return &updateMessageContentOpenedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageContentOpened *UpdateMessageContentOpened) GetUpdateEnum() UpdateEnum {
	return UpdateMessageContentOpenedType
}

// UpdateMessageMentionRead A message with an unread mention was read
type UpdateMessageMentionRead struct {
	tdCommon
	ChatId             int64 `json:"chat_id"`              // Chat identifier
	MessageId          int64 `json:"message_id"`           // Message identifier
	UnreadMentionCount int32 `json:"unread_mention_count"` // The new number of unread mention messages left in the chat
}

// MessageType return the string telegram-type of UpdateMessageMentionRead
func (updateMessageMentionRead *UpdateMessageMentionRead) MessageType() string {
	return "updateMessageMentionRead"
}

// NewUpdateMessageMentionRead creates a new UpdateMessageMentionRead
//
// @param chatId Chat identifier
// @param messageId Message identifier
// @param unreadMentionCount The new number of unread mention messages left in the chat
func NewUpdateMessageMentionRead(chatId int64, messageId int64, unreadMentionCount int32) *UpdateMessageMentionRead {
	updateMessageMentionReadTemp := UpdateMessageMentionRead{
		tdCommon:           tdCommon{Type: "updateMessageMentionRead"},
		ChatId:             chatId,
		MessageId:          messageId,
		UnreadMentionCount: unreadMentionCount,
	}

	return &updateMessageMentionReadTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageMentionRead *UpdateMessageMentionRead) GetUpdateEnum() UpdateEnum {
	return UpdateMessageMentionReadType
}

// UpdateMessageLiveLocationViewed A message with a live location was viewed. When the update is received, the application is supposed to update the live location
type UpdateMessageLiveLocationViewed struct {
	tdCommon
	ChatId    int64 `json:"chat_id"`    // Identifier of the chat with the live location message
	MessageId int64 `json:"message_id"` // Identifier of the message with live location
}

// MessageType return the string telegram-type of UpdateMessageLiveLocationViewed
func (updateMessageLiveLocationViewed *UpdateMessageLiveLocationViewed) MessageType() string {
	return "updateMessageLiveLocationViewed"
}

// NewUpdateMessageLiveLocationViewed creates a new UpdateMessageLiveLocationViewed
//
// @param chatId Identifier of the chat with the live location message
// @param messageId Identifier of the message with live location
func NewUpdateMessageLiveLocationViewed(chatId int64, messageId int64) *UpdateMessageLiveLocationViewed {
	updateMessageLiveLocationViewedTemp := UpdateMessageLiveLocationViewed{
		tdCommon:  tdCommon{Type: "updateMessageLiveLocationViewed"},
		ChatId:    chatId,
		MessageId: messageId,
	}

	return &updateMessageLiveLocationViewedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageLiveLocationViewed *UpdateMessageLiveLocationViewed) GetUpdateEnum() UpdateEnum {
	return UpdateMessageLiveLocationViewedType
}

// UpdateNewChat A new chat has been loaded/created. This update is guaranteed to come before the chat identifier is returned to the application. The chat field changes will be reported through separate updates
type UpdateNewChat struct {
	tdCommon
	Chat *Chat `json:"chat"` // The chat
}

// MessageType return the string telegram-type of UpdateNewChat
func (updateNewChat *UpdateNewChat) MessageType() string {
	return "updateNewChat"
}

// NewUpdateNewChat creates a new UpdateNewChat
//
// @param chat The chat
func NewUpdateNewChat(chat *Chat) *UpdateNewChat {
	updateNewChatTemp := UpdateNewChat{
		tdCommon: tdCommon{Type: "updateNewChat"},
		Chat:     chat,
	}

	return &updateNewChatTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewChat *UpdateNewChat) GetUpdateEnum() UpdateEnum {
	return UpdateNewChatType
}

// UpdateChatTitle The title of a chat was changed
type UpdateChatTitle struct {
	tdCommon
	ChatId int64  `json:"chat_id"` // Chat identifier
	Title  string `json:"title"`   // The new chat title
}

// MessageType return the string telegram-type of UpdateChatTitle
func (updateChatTitle *UpdateChatTitle) MessageType() string {
	return "updateChatTitle"
}

// NewUpdateChatTitle creates a new UpdateChatTitle
//
// @param chatId Chat identifier
// @param title The new chat title
func NewUpdateChatTitle(chatId int64, title string) *UpdateChatTitle {
	updateChatTitleTemp := UpdateChatTitle{
		tdCommon: tdCommon{Type: "updateChatTitle"},
		ChatId:   chatId,
		Title:    title,
	}

	return &updateChatTitleTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatTitle *UpdateChatTitle) GetUpdateEnum() UpdateEnum {
	return UpdateChatTitleType
}

// UpdateChatPhoto A chat photo was changed
type UpdateChatPhoto struct {
	tdCommon
	ChatId int64          `json:"chat_id"` // Chat identifier
	Photo  *ChatPhotoInfo `json:"photo"`   // The new chat photo; may be null
}

// MessageType return the string telegram-type of UpdateChatPhoto
func (updateChatPhoto *UpdateChatPhoto) MessageType() string {
	return "updateChatPhoto"
}

// NewUpdateChatPhoto creates a new UpdateChatPhoto
//
// @param chatId Chat identifier
// @param photo The new chat photo; may be null
func NewUpdateChatPhoto(chatId int64, photo *ChatPhotoInfo) *UpdateChatPhoto {
	updateChatPhotoTemp := UpdateChatPhoto{
		tdCommon: tdCommon{Type: "updateChatPhoto"},
		ChatId:   chatId,
		Photo:    photo,
	}

	return &updateChatPhotoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatPhoto *UpdateChatPhoto) GetUpdateEnum() UpdateEnum {
	return UpdateChatPhotoType
}

// UpdateChatPermissions Chat permissions was changed
type UpdateChatPermissions struct {
	tdCommon
	ChatId      int64            `json:"chat_id"`     // Chat identifier
	Permissions *ChatPermissions `json:"permissions"` // The new chat permissions
}

// MessageType return the string telegram-type of UpdateChatPermissions
func (updateChatPermissions *UpdateChatPermissions) MessageType() string {
	return "updateChatPermissions"
}

// NewUpdateChatPermissions creates a new UpdateChatPermissions
//
// @param chatId Chat identifier
// @param permissions The new chat permissions
func NewUpdateChatPermissions(chatId int64, permissions *ChatPermissions) *UpdateChatPermissions {
	updateChatPermissionsTemp := UpdateChatPermissions{
		tdCommon:    tdCommon{Type: "updateChatPermissions"},
		ChatId:      chatId,
		Permissions: permissions,
	}

	return &updateChatPermissionsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatPermissions *UpdateChatPermissions) GetUpdateEnum() UpdateEnum {
	return UpdateChatPermissionsType
}

// UpdateChatLastMessage The last message of a chat was changed. If last_message is null, then the last message in the chat became unknown. Some new unknown messages might be added to the chat in this case
type UpdateChatLastMessage struct {
	tdCommon
	ChatId      int64          `json:"chat_id"`      // Chat identifier
	LastMessage *Message       `json:"last_message"` // The new last message in the chat; may be null
	Positions   []ChatPosition `json:"positions"`    // The new chat positions in the chat lists
}

// MessageType return the string telegram-type of UpdateChatLastMessage
func (updateChatLastMessage *UpdateChatLastMessage) MessageType() string {
	return "updateChatLastMessage"
}

// NewUpdateChatLastMessage creates a new UpdateChatLastMessage
//
// @param chatId Chat identifier
// @param lastMessage The new last message in the chat; may be null
// @param positions The new chat positions in the chat lists
func NewUpdateChatLastMessage(chatId int64, lastMessage *Message, positions []ChatPosition) *UpdateChatLastMessage {
	updateChatLastMessageTemp := UpdateChatLastMessage{
		tdCommon:    tdCommon{Type: "updateChatLastMessage"},
		ChatId:      chatId,
		LastMessage: lastMessage,
		Positions:   positions,
	}

	return &updateChatLastMessageTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatLastMessage *UpdateChatLastMessage) GetUpdateEnum() UpdateEnum {
	return UpdateChatLastMessageType
}

// UpdateChatPosition The position of a chat in a chat list has changed. Instead of this update updateChatLastMessage or updateChatDraftMessage might be sent
type UpdateChatPosition struct {
	tdCommon
	ChatId   int64         `json:"chat_id"`  // Chat identifier
	Position *ChatPosition `json:"position"` // New chat position. If new order is 0, then the chat needs to be removed from the list
}

// MessageType return the string telegram-type of UpdateChatPosition
func (updateChatPosition *UpdateChatPosition) MessageType() string {
	return "updateChatPosition"
}

// NewUpdateChatPosition creates a new UpdateChatPosition
//
// @param chatId Chat identifier
// @param position New chat position. If new order is 0, then the chat needs to be removed from the list
func NewUpdateChatPosition(chatId int64, position *ChatPosition) *UpdateChatPosition {
	updateChatPositionTemp := UpdateChatPosition{
		tdCommon: tdCommon{Type: "updateChatPosition"},
		ChatId:   chatId,
		Position: position,
	}

	return &updateChatPositionTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatPosition *UpdateChatPosition) GetUpdateEnum() UpdateEnum {
	return UpdateChatPositionType
}

// UpdateChatIsMarkedAsUnread A chat was marked as unread or was read
type UpdateChatIsMarkedAsUnread struct {
	tdCommon
	ChatId           int64 `json:"chat_id"`             // Chat identifier
	IsMarkedAsUnread bool  `json:"is_marked_as_unread"` // New value of is_marked_as_unread
}

// MessageType return the string telegram-type of UpdateChatIsMarkedAsUnread
func (updateChatIsMarkedAsUnread *UpdateChatIsMarkedAsUnread) MessageType() string {
	return "updateChatIsMarkedAsUnread"
}

// NewUpdateChatIsMarkedAsUnread creates a new UpdateChatIsMarkedAsUnread
//
// @param chatId Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func NewUpdateChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) *UpdateChatIsMarkedAsUnread {
	updateChatIsMarkedAsUnreadTemp := UpdateChatIsMarkedAsUnread{
		tdCommon:         tdCommon{Type: "updateChatIsMarkedAsUnread"},
		ChatId:           chatId,
		IsMarkedAsUnread: isMarkedAsUnread,
	}

	return &updateChatIsMarkedAsUnreadTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsMarkedAsUnread *UpdateChatIsMarkedAsUnread) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsMarkedAsUnreadType
}

// UpdateChatIsBlocked A chat was blocked or unblocked
type UpdateChatIsBlocked struct {
	tdCommon
	ChatId    int64 `json:"chat_id"`    // Chat identifier
	IsBlocked bool  `json:"is_blocked"` // New value of is_blocked
}

// MessageType return the string telegram-type of UpdateChatIsBlocked
func (updateChatIsBlocked *UpdateChatIsBlocked) MessageType() string {
	return "updateChatIsBlocked"
}

// NewUpdateChatIsBlocked creates a new UpdateChatIsBlocked
//
// @param chatId Chat identifier
// @param isBlocked New value of is_blocked
func NewUpdateChatIsBlocked(chatId int64, isBlocked bool) *UpdateChatIsBlocked {
	updateChatIsBlockedTemp := UpdateChatIsBlocked{
		tdCommon:  tdCommon{Type: "updateChatIsBlocked"},
		ChatId:    chatId,
		IsBlocked: isBlocked,
	}

	return &updateChatIsBlockedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsBlocked *UpdateChatIsBlocked) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsBlockedType
}

// UpdateChatHasScheduledMessages A chat's has_scheduled_messages field has changed
type UpdateChatHasScheduledMessages struct {
	tdCommon
	ChatId               int64 `json:"chat_id"`                // Chat identifier
	HasScheduledMessages bool  `json:"has_scheduled_messages"` // New value of has_scheduled_messages
}

// MessageType return the string telegram-type of UpdateChatHasScheduledMessages
func (updateChatHasScheduledMessages *UpdateChatHasScheduledMessages) MessageType() string {
	return "updateChatHasScheduledMessages"
}

// NewUpdateChatHasScheduledMessages creates a new UpdateChatHasScheduledMessages
//
// @param chatId Chat identifier
// @param hasScheduledMessages New value of has_scheduled_messages
func NewUpdateChatHasScheduledMessages(chatId int64, hasScheduledMessages bool) *UpdateChatHasScheduledMessages {
	updateChatHasScheduledMessagesTemp := UpdateChatHasScheduledMessages{
		tdCommon:             tdCommon{Type: "updateChatHasScheduledMessages"},
		ChatId:               chatId,
		HasScheduledMessages: hasScheduledMessages,
	}

	return &updateChatHasScheduledMessagesTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatHasScheduledMessages *UpdateChatHasScheduledMessages) GetUpdateEnum() UpdateEnum {
	return UpdateChatHasScheduledMessagesType
}

// UpdateChatVideoChat A chat video chat state has changed
type UpdateChatVideoChat struct {
	tdCommon
	ChatId    int64      `json:"chat_id"`    // Chat identifier
	VideoChat *VideoChat `json:"video_chat"` // New value of video_chat
}

// MessageType return the string telegram-type of UpdateChatVideoChat
func (updateChatVideoChat *UpdateChatVideoChat) MessageType() string {
	return "updateChatVideoChat"
}

// NewUpdateChatVideoChat creates a new UpdateChatVideoChat
//
// @param chatId Chat identifier
// @param videoChat New value of video_chat
func NewUpdateChatVideoChat(chatId int64, videoChat *VideoChat) *UpdateChatVideoChat {
	updateChatVideoChatTemp := UpdateChatVideoChat{
		tdCommon:  tdCommon{Type: "updateChatVideoChat"},
		ChatId:    chatId,
		VideoChat: videoChat,
	}

	return &updateChatVideoChatTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatVideoChat *UpdateChatVideoChat) GetUpdateEnum() UpdateEnum {
	return UpdateChatVideoChatType
}

// UpdateChatDefaultDisableNotification The value of the default disable_notification parameter, used when a message is sent to the chat, was changed
type UpdateChatDefaultDisableNotification struct {
	tdCommon
	ChatId                     int64 `json:"chat_id"`                      // Chat identifier
	DefaultDisableNotification bool  `json:"default_disable_notification"` // The new default_disable_notification value
}

// MessageType return the string telegram-type of UpdateChatDefaultDisableNotification
func (updateChatDefaultDisableNotification *UpdateChatDefaultDisableNotification) MessageType() string {
	return "updateChatDefaultDisableNotification"
}

// NewUpdateChatDefaultDisableNotification creates a new UpdateChatDefaultDisableNotification
//
// @param chatId Chat identifier
// @param defaultDisableNotification The new default_disable_notification value
func NewUpdateChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) *UpdateChatDefaultDisableNotification {
	updateChatDefaultDisableNotificationTemp := UpdateChatDefaultDisableNotification{
		tdCommon:                   tdCommon{Type: "updateChatDefaultDisableNotification"},
		ChatId:                     chatId,
		DefaultDisableNotification: defaultDisableNotification,
	}

	return &updateChatDefaultDisableNotificationTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatDefaultDisableNotification *UpdateChatDefaultDisableNotification) GetUpdateEnum() UpdateEnum {
	return UpdateChatDefaultDisableNotificationType
}

// UpdateChatReadInbox Incoming messages were read or number of unread messages has been changed
type UpdateChatReadInbox struct {
	tdCommon
	ChatId                 int64 `json:"chat_id"`                    // Chat identifier
	LastReadInboxMessageId int64 `json:"last_read_inbox_message_id"` // Identifier of the last read incoming message
	UnreadCount            int32 `json:"unread_count"`               // The number of unread messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatReadInbox
func (updateChatReadInbox *UpdateChatReadInbox) MessageType() string {
	return "updateChatReadInbox"
}

// NewUpdateChatReadInbox creates a new UpdateChatReadInbox
//
// @param chatId Chat identifier
// @param lastReadInboxMessageId Identifier of the last read incoming message
// @param unreadCount The number of unread messages left in the chat
func NewUpdateChatReadInbox(chatId int64, lastReadInboxMessageId int64, unreadCount int32) *UpdateChatReadInbox {
	updateChatReadInboxTemp := UpdateChatReadInbox{
		tdCommon:               tdCommon{Type: "updateChatReadInbox"},
		ChatId:                 chatId,
		LastReadInboxMessageId: lastReadInboxMessageId,
		UnreadCount:            unreadCount,
	}

	return &updateChatReadInboxTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatReadInbox *UpdateChatReadInbox) GetUpdateEnum() UpdateEnum {
	return UpdateChatReadInboxType
}

// UpdateChatReadOutbox Outgoing messages were read
type UpdateChatReadOutbox struct {
	tdCommon
	ChatId                  int64 `json:"chat_id"`                     // Chat identifier
	LastReadOutboxMessageId int64 `json:"last_read_outbox_message_id"` // Identifier of last read outgoing message
}

// MessageType return the string telegram-type of UpdateChatReadOutbox
func (updateChatReadOutbox *UpdateChatReadOutbox) MessageType() string {
	return "updateChatReadOutbox"
}

// NewUpdateChatReadOutbox creates a new UpdateChatReadOutbox
//
// @param chatId Chat identifier
// @param lastReadOutboxMessageId Identifier of last read outgoing message
func NewUpdateChatReadOutbox(chatId int64, lastReadOutboxMessageId int64) *UpdateChatReadOutbox {
	updateChatReadOutboxTemp := UpdateChatReadOutbox{
		tdCommon:                tdCommon{Type: "updateChatReadOutbox"},
		ChatId:                  chatId,
		LastReadOutboxMessageId: lastReadOutboxMessageId,
	}

	return &updateChatReadOutboxTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatReadOutbox *UpdateChatReadOutbox) GetUpdateEnum() UpdateEnum {
	return UpdateChatReadOutboxType
}

// UpdateChatUnreadMentionCount The chat unread_mention_count has changed
type UpdateChatUnreadMentionCount struct {
	tdCommon
	ChatId             int64 `json:"chat_id"`              // Chat identifier
	UnreadMentionCount int32 `json:"unread_mention_count"` // The number of unread mention messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatUnreadMentionCount
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) MessageType() string {
	return "updateChatUnreadMentionCount"
}

// NewUpdateChatUnreadMentionCount creates a new UpdateChatUnreadMentionCount
//
// @param chatId Chat identifier
// @param unreadMentionCount The number of unread mention messages left in the chat
func NewUpdateChatUnreadMentionCount(chatId int64, unreadMentionCount int32) *UpdateChatUnreadMentionCount {
	updateChatUnreadMentionCountTemp := UpdateChatUnreadMentionCount{
		tdCommon:           tdCommon{Type: "updateChatUnreadMentionCount"},
		ChatId:             chatId,
		UnreadMentionCount: unreadMentionCount,
	}

	return &updateChatUnreadMentionCountTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) GetUpdateEnum() UpdateEnum {
	return UpdateChatUnreadMentionCountType
}

// UpdateChatNotificationSettings Notification settings for a chat were changed
type UpdateChatNotificationSettings struct {
	tdCommon
	ChatId               int64                     `json:"chat_id"`               // Chat identifier
	NotificationSettings *ChatNotificationSettings `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateChatNotificationSettings
func (updateChatNotificationSettings *UpdateChatNotificationSettings) MessageType() string {
	return "updateChatNotificationSettings"
}

// NewUpdateChatNotificationSettings creates a new UpdateChatNotificationSettings
//
// @param chatId Chat identifier
// @param notificationSettings The new notification settings
func NewUpdateChatNotificationSettings(chatId int64, notificationSettings *ChatNotificationSettings) *UpdateChatNotificationSettings {
	updateChatNotificationSettingsTemp := UpdateChatNotificationSettings{
		tdCommon:             tdCommon{Type: "updateChatNotificationSettings"},
		ChatId:               chatId,
		NotificationSettings: notificationSettings,
	}

	return &updateChatNotificationSettingsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatNotificationSettings *UpdateChatNotificationSettings) GetUpdateEnum() UpdateEnum {
	return UpdateChatNotificationSettingsType
}

// UpdateScopeNotificationSettings Notification settings for some type of chats were updated
type UpdateScopeNotificationSettings struct {
	tdCommon
	Scope                NotificationSettingsScope  `json:"scope"`                 // Types of chats for which notification settings were updated
	NotificationSettings *ScopeNotificationSettings `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateScopeNotificationSettings
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) MessageType() string {
	return "updateScopeNotificationSettings"
}

// NewUpdateScopeNotificationSettings creates a new UpdateScopeNotificationSettings
//
// @param scope Types of chats for which notification settings were updated
// @param notificationSettings The new notification settings
func NewUpdateScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) *UpdateScopeNotificationSettings {
	updateScopeNotificationSettingsTemp := UpdateScopeNotificationSettings{
		tdCommon:             tdCommon{Type: "updateScopeNotificationSettings"},
		Scope:                scope,
		NotificationSettings: notificationSettings,
	}

	return &updateScopeNotificationSettingsTemp
}

// UnmarshalJSON unmarshal to json
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NotificationSettings *ScopeNotificationSettings `json:"notification_settings"` // The new notification settings
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateScopeNotificationSettings.tdCommon = tempObj.tdCommon
	updateScopeNotificationSettings.NotificationSettings = tempObj.NotificationSettings

	fieldScope, _ := unmarshalNotificationSettingsScope(objMap["scope"])
	updateScopeNotificationSettings.Scope = fieldScope

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) GetUpdateEnum() UpdateEnum {
	return UpdateScopeNotificationSettingsType
}

// UpdateChatMessageTtlSetting The message Time To Live setting for a chat was changed
type UpdateChatMessageTtlSetting struct {
	tdCommon
	ChatId            int64 `json:"chat_id"`             // Chat identifier
	MessageTtlSetting int32 `json:"message_ttl_setting"` // New value of message_ttl_setting
}

// MessageType return the string telegram-type of UpdateChatMessageTtlSetting
func (updateChatMessageTtlSetting *UpdateChatMessageTtlSetting) MessageType() string {
	return "updateChatMessageTtlSetting"
}

// NewUpdateChatMessageTtlSetting creates a new UpdateChatMessageTtlSetting
//
// @param chatId Chat identifier
// @param messageTtlSetting New value of message_ttl_setting
func NewUpdateChatMessageTtlSetting(chatId int64, messageTtlSetting int32) *UpdateChatMessageTtlSetting {
	updateChatMessageTtlSettingTemp := UpdateChatMessageTtlSetting{
		tdCommon:          tdCommon{Type: "updateChatMessageTtlSetting"},
		ChatId:            chatId,
		MessageTtlSetting: messageTtlSetting,
	}

	return &updateChatMessageTtlSettingTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatMessageTtlSetting *UpdateChatMessageTtlSetting) GetUpdateEnum() UpdateEnum {
	return UpdateChatMessageTtlSettingType
}

// UpdateChatActionBar The chat action bar was changed
type UpdateChatActionBar struct {
	tdCommon
	ChatId    int64         `json:"chat_id"`    // Chat identifier
	ActionBar ChatActionBar `json:"action_bar"` // The new value of the action bar; may be null
}

// MessageType return the string telegram-type of UpdateChatActionBar
func (updateChatActionBar *UpdateChatActionBar) MessageType() string {
	return "updateChatActionBar"
}

// NewUpdateChatActionBar creates a new UpdateChatActionBar
//
// @param chatId Chat identifier
// @param actionBar The new value of the action bar; may be null
func NewUpdateChatActionBar(chatId int64, actionBar ChatActionBar) *UpdateChatActionBar {
	updateChatActionBarTemp := UpdateChatActionBar{
		tdCommon:  tdCommon{Type: "updateChatActionBar"},
		ChatId:    chatId,
		ActionBar: actionBar,
	}

	return &updateChatActionBarTemp
}

// UnmarshalJSON unmarshal to json
func (updateChatActionBar *UpdateChatActionBar) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatId int64 `json:"chat_id"` // Chat identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateChatActionBar.tdCommon = tempObj.tdCommon
	updateChatActionBar.ChatId = tempObj.ChatId

	fieldActionBar, _ := unmarshalChatActionBar(objMap["action_bar"])
	updateChatActionBar.ActionBar = fieldActionBar

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateChatActionBar *UpdateChatActionBar) GetUpdateEnum() UpdateEnum {
	return UpdateChatActionBarType
}

// UpdateChatTheme The chat theme was changed
type UpdateChatTheme struct {
	tdCommon
	ChatId    int64  `json:"chat_id"`    // Chat identifier
	ThemeName string `json:"theme_name"` // The new name of the chat theme; may be empty if theme was reset to default
}

// MessageType return the string telegram-type of UpdateChatTheme
func (updateChatTheme *UpdateChatTheme) MessageType() string {
	return "updateChatTheme"
}

// NewUpdateChatTheme creates a new UpdateChatTheme
//
// @param chatId Chat identifier
// @param themeName The new name of the chat theme; may be empty if theme was reset to default
func NewUpdateChatTheme(chatId int64, themeName string) *UpdateChatTheme {
	updateChatThemeTemp := UpdateChatTheme{
		tdCommon:  tdCommon{Type: "updateChatTheme"},
		ChatId:    chatId,
		ThemeName: themeName,
	}

	return &updateChatThemeTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatTheme *UpdateChatTheme) GetUpdateEnum() UpdateEnum {
	return UpdateChatThemeType
}

// UpdateChatPendingJoinRequests The chat pending join requests were changed
type UpdateChatPendingJoinRequests struct {
	tdCommon
	ChatId              int64                 `json:"chat_id"`               // Chat identifier
	PendingJoinRequests *ChatJoinRequestsInfo `json:"pending_join_requests"` // The new data about pending join requests; may be null
}

// MessageType return the string telegram-type of UpdateChatPendingJoinRequests
func (updateChatPendingJoinRequests *UpdateChatPendingJoinRequests) MessageType() string {
	return "updateChatPendingJoinRequests"
}

// NewUpdateChatPendingJoinRequests creates a new UpdateChatPendingJoinRequests
//
// @param chatId Chat identifier
// @param pendingJoinRequests The new data about pending join requests; may be null
func NewUpdateChatPendingJoinRequests(chatId int64, pendingJoinRequests *ChatJoinRequestsInfo) *UpdateChatPendingJoinRequests {
	updateChatPendingJoinRequestsTemp := UpdateChatPendingJoinRequests{
		tdCommon:            tdCommon{Type: "updateChatPendingJoinRequests"},
		ChatId:              chatId,
		PendingJoinRequests: pendingJoinRequests,
	}

	return &updateChatPendingJoinRequestsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatPendingJoinRequests *UpdateChatPendingJoinRequests) GetUpdateEnum() UpdateEnum {
	return UpdateChatPendingJoinRequestsType
}

// UpdateChatReplyMarkup The default chat reply markup was changed. Can occur because new messages with reply markup were received or because an old reply markup was hidden by the user
type UpdateChatReplyMarkup struct {
	tdCommon
	ChatId               int64 `json:"chat_id"`                 // Chat identifier
	ReplyMarkupMessageId int64 `json:"reply_markup_message_id"` // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
}

// MessageType return the string telegram-type of UpdateChatReplyMarkup
func (updateChatReplyMarkup *UpdateChatReplyMarkup) MessageType() string {
	return "updateChatReplyMarkup"
}

// NewUpdateChatReplyMarkup creates a new UpdateChatReplyMarkup
//
// @param chatId Chat identifier
// @param replyMarkupMessageId Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
func NewUpdateChatReplyMarkup(chatId int64, replyMarkupMessageId int64) *UpdateChatReplyMarkup {
	updateChatReplyMarkupTemp := UpdateChatReplyMarkup{
		tdCommon:             tdCommon{Type: "updateChatReplyMarkup"},
		ChatId:               chatId,
		ReplyMarkupMessageId: replyMarkupMessageId,
	}

	return &updateChatReplyMarkupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatReplyMarkup *UpdateChatReplyMarkup) GetUpdateEnum() UpdateEnum {
	return UpdateChatReplyMarkupType
}

// UpdateChatDraftMessage A chat draft has changed. Be aware that the update may come in the currently opened chat but with old content of the draft. If the user has changed the content of the draft, this update mustn't be applied
type UpdateChatDraftMessage struct {
	tdCommon
	ChatId       int64          `json:"chat_id"`       // Chat identifier
	DraftMessage *DraftMessage  `json:"draft_message"` // The new draft message; may be null
	Positions    []ChatPosition `json:"positions"`     // The new chat positions in the chat lists
}

// MessageType return the string telegram-type of UpdateChatDraftMessage
func (updateChatDraftMessage *UpdateChatDraftMessage) MessageType() string {
	return "updateChatDraftMessage"
}

// NewUpdateChatDraftMessage creates a new UpdateChatDraftMessage
//
// @param chatId Chat identifier
// @param draftMessage The new draft message; may be null
// @param positions The new chat positions in the chat lists
func NewUpdateChatDraftMessage(chatId int64, draftMessage *DraftMessage, positions []ChatPosition) *UpdateChatDraftMessage {
	updateChatDraftMessageTemp := UpdateChatDraftMessage{
		tdCommon:     tdCommon{Type: "updateChatDraftMessage"},
		ChatId:       chatId,
		DraftMessage: draftMessage,
		Positions:    positions,
	}

	return &updateChatDraftMessageTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatDraftMessage *UpdateChatDraftMessage) GetUpdateEnum() UpdateEnum {
	return UpdateChatDraftMessageType
}

// UpdateChatFilters The list of chat filters or a chat filter has changed
type UpdateChatFilters struct {
	tdCommon
	ChatFilters []ChatFilterInfo `json:"chat_filters"` // The new list of chat filters
}

// MessageType return the string telegram-type of UpdateChatFilters
func (updateChatFilters *UpdateChatFilters) MessageType() string {
	return "updateChatFilters"
}

// NewUpdateChatFilters creates a new UpdateChatFilters
//
// @param chatFilters The new list of chat filters
func NewUpdateChatFilters(chatFilters []ChatFilterInfo) *UpdateChatFilters {
	updateChatFiltersTemp := UpdateChatFilters{
		tdCommon:    tdCommon{Type: "updateChatFilters"},
		ChatFilters: chatFilters,
	}

	return &updateChatFiltersTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatFilters *UpdateChatFilters) GetUpdateEnum() UpdateEnum {
	return UpdateChatFiltersType
}

// UpdateChatOnlineMemberCount The number of online group members has changed. This update with non-zero count is sent only for currently opened chats. There is no guarantee that it will be sent just after the count has changed
type UpdateChatOnlineMemberCount struct {
	tdCommon
	ChatId            int64 `json:"chat_id"`             // Identifier of the chat
	OnlineMemberCount int32 `json:"online_member_count"` // New number of online members in the chat, or 0 if unknown
}

// MessageType return the string telegram-type of UpdateChatOnlineMemberCount
func (updateChatOnlineMemberCount *UpdateChatOnlineMemberCount) MessageType() string {
	return "updateChatOnlineMemberCount"
}

// NewUpdateChatOnlineMemberCount creates a new UpdateChatOnlineMemberCount
//
// @param chatId Identifier of the chat
// @param onlineMemberCount New number of online members in the chat, or 0 if unknown
func NewUpdateChatOnlineMemberCount(chatId int64, onlineMemberCount int32) *UpdateChatOnlineMemberCount {
	updateChatOnlineMemberCountTemp := UpdateChatOnlineMemberCount{
		tdCommon:          tdCommon{Type: "updateChatOnlineMemberCount"},
		ChatId:            chatId,
		OnlineMemberCount: onlineMemberCount,
	}

	return &updateChatOnlineMemberCountTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatOnlineMemberCount *UpdateChatOnlineMemberCount) GetUpdateEnum() UpdateEnum {
	return UpdateChatOnlineMemberCountType
}

// UpdateNotification A notification was changed
type UpdateNotification struct {
	tdCommon
	NotificationGroupId int32         `json:"notification_group_id"` // Unique notification group identifier
	Notification        *Notification `json:"notification"`          // Changed notification
}

// MessageType return the string telegram-type of UpdateNotification
func (updateNotification *UpdateNotification) MessageType() string {
	return "updateNotification"
}

// NewUpdateNotification creates a new UpdateNotification
//
// @param notificationGroupId Unique notification group identifier
// @param notification Changed notification
func NewUpdateNotification(notificationGroupId int32, notification *Notification) *UpdateNotification {
	updateNotificationTemp := UpdateNotification{
		tdCommon:            tdCommon{Type: "updateNotification"},
		NotificationGroupId: notificationGroupId,
		Notification:        notification,
	}

	return &updateNotificationTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNotification *UpdateNotification) GetUpdateEnum() UpdateEnum {
	return UpdateNotificationType
}

// UpdateNotificationGroup A list of active notifications in a notification group has changed
type UpdateNotificationGroup struct {
	tdCommon
	NotificationGroupId        int32                 `json:"notification_group_id"`         // Unique notification group identifier
	Type                       NotificationGroupType `json:"type"`                          // New type of the notification group
	ChatId                     int64                 `json:"chat_id"`                       // Identifier of a chat to which all notifications in the group belong
	NotificationSettingsChatId int64                 `json:"notification_settings_chat_id"` // Chat identifier, which notification settings must be applied to the added notifications
	IsSilent                   bool                  `json:"is_silent"`                     // True, if the notifications must be shown without sound
	TotalCount                 int32                 `json:"total_count"`                   // Total number of unread notifications in the group, can be bigger than number of active notifications
	AddedNotifications         []Notification        `json:"added_notifications"`           // List of added group notifications, sorted by notification ID
	RemovedNotificationIds     []int32               `json:"removed_notification_ids"`      // Identifiers of removed group notifications, sorted by notification ID
}

// MessageType return the string telegram-type of UpdateNotificationGroup
func (updateNotificationGroup *UpdateNotificationGroup) MessageType() string {
	return "updateNotificationGroup"
}

// NewUpdateNotificationGroup creates a new UpdateNotificationGroup
//
// @param notificationGroupId Unique notification group identifier
// @param typeParam New type of the notification group
// @param chatId Identifier of a chat to which all notifications in the group belong
// @param notificationSettingsChatId Chat identifier, which notification settings must be applied to the added notifications
// @param isSilent True, if the notifications must be shown without sound
// @param totalCount Total number of unread notifications in the group, can be bigger than number of active notifications
// @param addedNotifications List of added group notifications, sorted by notification ID
// @param removedNotificationIds Identifiers of removed group notifications, sorted by notification ID
func NewUpdateNotificationGroup(notificationGroupId int32, typeParam NotificationGroupType, chatId int64, notificationSettingsChatId int64, isSilent bool, totalCount int32, addedNotifications []Notification, removedNotificationIds []int32) *UpdateNotificationGroup {
	updateNotificationGroupTemp := UpdateNotificationGroup{
		tdCommon:                   tdCommon{Type: "updateNotificationGroup"},
		NotificationGroupId:        notificationGroupId,
		Type:                       typeParam,
		ChatId:                     chatId,
		NotificationSettingsChatId: notificationSettingsChatId,
		IsSilent:                   isSilent,
		TotalCount:                 totalCount,
		AddedNotifications:         addedNotifications,
		RemovedNotificationIds:     removedNotificationIds,
	}

	return &updateNotificationGroupTemp
}

// UnmarshalJSON unmarshal to json
func (updateNotificationGroup *UpdateNotificationGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NotificationGroupId        int32          `json:"notification_group_id"`         // Unique notification group identifier
		ChatId                     int64          `json:"chat_id"`                       // Identifier of a chat to which all notifications in the group belong
		NotificationSettingsChatId int64          `json:"notification_settings_chat_id"` // Chat identifier, which notification settings must be applied to the added notifications
		IsSilent                   bool           `json:"is_silent"`                     // True, if the notifications must be shown without sound
		TotalCount                 int32          `json:"total_count"`                   // Total number of unread notifications in the group, can be bigger than number of active notifications
		AddedNotifications         []Notification `json:"added_notifications"`           // List of added group notifications, sorted by notification ID
		RemovedNotificationIds     []int32        `json:"removed_notification_ids"`      // Identifiers of removed group notifications, sorted by notification ID
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNotificationGroup.tdCommon = tempObj.tdCommon
	updateNotificationGroup.NotificationGroupId = tempObj.NotificationGroupId
	updateNotificationGroup.ChatId = tempObj.ChatId
	updateNotificationGroup.NotificationSettingsChatId = tempObj.NotificationSettingsChatId
	updateNotificationGroup.IsSilent = tempObj.IsSilent
	updateNotificationGroup.TotalCount = tempObj.TotalCount
	updateNotificationGroup.AddedNotifications = tempObj.AddedNotifications
	updateNotificationGroup.RemovedNotificationIds = tempObj.RemovedNotificationIds

	fieldType, _ := unmarshalNotificationGroupType(objMap["type"])
	updateNotificationGroup.Type = fieldType

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNotificationGroup *UpdateNotificationGroup) GetUpdateEnum() UpdateEnum {
	return UpdateNotificationGroupType
}

// UpdateActiveNotifications Contains active notifications that was shown on previous application launches. This update is sent only if the message database is used. In that case it comes once before any updateNotification and updateNotificationGroup update
type UpdateActiveNotifications struct {
	tdCommon
	Groups []NotificationGroup `json:"groups"` // Lists of active notification groups
}

// MessageType return the string telegram-type of UpdateActiveNotifications
func (updateActiveNotifications *UpdateActiveNotifications) MessageType() string {
	return "updateActiveNotifications"
}

// NewUpdateActiveNotifications creates a new UpdateActiveNotifications
//
// @param groups Lists of active notification groups
func NewUpdateActiveNotifications(groups []NotificationGroup) *UpdateActiveNotifications {
	updateActiveNotificationsTemp := UpdateActiveNotifications{
		tdCommon: tdCommon{Type: "updateActiveNotifications"},
		Groups:   groups,
	}

	return &updateActiveNotificationsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateActiveNotifications *UpdateActiveNotifications) GetUpdateEnum() UpdateEnum {
	return UpdateActiveNotificationsType
}

// UpdateHavePendingNotifications Describes whether there are some pending notification updates. Can be used to prevent application from killing, while there are some pending notifications
type UpdateHavePendingNotifications struct {
	tdCommon
	HaveDelayedNotifications    bool `json:"have_delayed_notifications"`    // True, if there are some delayed notification updates, which will be sent soon
	HaveUnreceivedNotifications bool `json:"have_unreceived_notifications"` // True, if there can be some yet unreceived notifications, which are being fetched from the server
}

// MessageType return the string telegram-type of UpdateHavePendingNotifications
func (updateHavePendingNotifications *UpdateHavePendingNotifications) MessageType() string {
	return "updateHavePendingNotifications"
}

// NewUpdateHavePendingNotifications creates a new UpdateHavePendingNotifications
//
// @param haveDelayedNotifications True, if there are some delayed notification updates, which will be sent soon
// @param haveUnreceivedNotifications True, if there can be some yet unreceived notifications, which are being fetched from the server
func NewUpdateHavePendingNotifications(haveDelayedNotifications bool, haveUnreceivedNotifications bool) *UpdateHavePendingNotifications {
	updateHavePendingNotificationsTemp := UpdateHavePendingNotifications{
		tdCommon:                    tdCommon{Type: "updateHavePendingNotifications"},
		HaveDelayedNotifications:    haveDelayedNotifications,
		HaveUnreceivedNotifications: haveUnreceivedNotifications,
	}

	return &updateHavePendingNotificationsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateHavePendingNotifications *UpdateHavePendingNotifications) GetUpdateEnum() UpdateEnum {
	return UpdateHavePendingNotificationsType
}

// UpdateDeleteMessages Some messages were deleted
type UpdateDeleteMessages struct {
	tdCommon
	ChatId      int64   `json:"chat_id"`      // Chat identifier
	MessageIds  []int64 `json:"message_ids"`  // Identifiers of the deleted messages
	IsPermanent bool    `json:"is_permanent"` // True, if the messages are permanently deleted by a user (as opposed to just becoming inaccessible)
	FromCache   bool    `json:"from_cache"`   // True, if the messages are deleted only from the cache and can possibly be retrieved again in the future
}

// MessageType return the string telegram-type of UpdateDeleteMessages
func (updateDeleteMessages *UpdateDeleteMessages) MessageType() string {
	return "updateDeleteMessages"
}

// NewUpdateDeleteMessages creates a new UpdateDeleteMessages
//
// @param chatId Chat identifier
// @param messageIds Identifiers of the deleted messages
// @param isPermanent True, if the messages are permanently deleted by a user (as opposed to just becoming inaccessible)
// @param fromCache True, if the messages are deleted only from the cache and can possibly be retrieved again in the future
func NewUpdateDeleteMessages(chatId int64, messageIds []int64, isPermanent bool, fromCache bool) *UpdateDeleteMessages {
	updateDeleteMessagesTemp := UpdateDeleteMessages{
		tdCommon:    tdCommon{Type: "updateDeleteMessages"},
		ChatId:      chatId,
		MessageIds:  messageIds,
		IsPermanent: isPermanent,
		FromCache:   fromCache,
	}

	return &updateDeleteMessagesTemp
}

// GetUpdateEnum return the enum type of this object
func (updateDeleteMessages *UpdateDeleteMessages) GetUpdateEnum() UpdateEnum {
	return UpdateDeleteMessagesType
}

// UpdateUserChatAction User activity in the chat has changed
type UpdateUserChatAction struct {
	tdCommon
	ChatId          int64      `json:"chat_id"`           // Chat identifier
	MessageThreadId int64      `json:"message_thread_id"` // If not 0, a message thread identifier in which the action was performed
	UserId          int64      `json:"user_id"`           // Identifier of a user performing an action
	Action          ChatAction `json:"action"`            // The action description
}

// MessageType return the string telegram-type of UpdateUserChatAction
func (updateUserChatAction *UpdateUserChatAction) MessageType() string {
	return "updateUserChatAction"
}

// NewUpdateUserChatAction creates a new UpdateUserChatAction
//
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the action was performed
// @param userId Identifier of a user performing an action
// @param action The action description
func NewUpdateUserChatAction(chatId int64, messageThreadId int64, userId int64, action ChatAction) *UpdateUserChatAction {
	updateUserChatActionTemp := UpdateUserChatAction{
		tdCommon:        tdCommon{Type: "updateUserChatAction"},
		ChatId:          chatId,
		MessageThreadId: messageThreadId,
		UserId:          userId,
		Action:          action,
	}

	return &updateUserChatActionTemp
}

// UnmarshalJSON unmarshal to json
func (updateUserChatAction *UpdateUserChatAction) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatId          int64 `json:"chat_id"`           // Chat identifier
		MessageThreadId int64 `json:"message_thread_id"` // If not 0, a message thread identifier in which the action was performed
		UserId          int64 `json:"user_id"`           // Identifier of a user performing an action

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserChatAction.tdCommon = tempObj.tdCommon
	updateUserChatAction.ChatId = tempObj.ChatId
	updateUserChatAction.MessageThreadId = tempObj.MessageThreadId
	updateUserChatAction.UserId = tempObj.UserId

	fieldAction, _ := unmarshalChatAction(objMap["action"])
	updateUserChatAction.Action = fieldAction

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserChatAction *UpdateUserChatAction) GetUpdateEnum() UpdateEnum {
	return UpdateUserChatActionType
}

// UpdateUserStatus The user went online or offline
type UpdateUserStatus struct {
	tdCommon
	UserId int64      `json:"user_id"` // User identifier
	Status UserStatus `json:"status"`  // New status of the user
}

// MessageType return the string telegram-type of UpdateUserStatus
func (updateUserStatus *UpdateUserStatus) MessageType() string {
	return "updateUserStatus"
}

// NewUpdateUserStatus creates a new UpdateUserStatus
//
// @param userId User identifier
// @param status New status of the user
func NewUpdateUserStatus(userId int64, status UserStatus) *UpdateUserStatus {
	updateUserStatusTemp := UpdateUserStatus{
		tdCommon: tdCommon{Type: "updateUserStatus"},
		UserId:   userId,
		Status:   status,
	}

	return &updateUserStatusTemp
}

// UnmarshalJSON unmarshal to json
func (updateUserStatus *UpdateUserStatus) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserId int64 `json:"user_id"` // User identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserStatus.tdCommon = tempObj.tdCommon
	updateUserStatus.UserId = tempObj.UserId

	fieldStatus, _ := unmarshalUserStatus(objMap["status"])
	updateUserStatus.Status = fieldStatus

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserStatus *UpdateUserStatus) GetUpdateEnum() UpdateEnum {
	return UpdateUserStatusType
}

// UpdateUser Some data of a user has changed. This update is guaranteed to come before the user identifier is returned to the application
type UpdateUser struct {
	tdCommon
	User *User `json:"user"` // New data about the user
}

// MessageType return the string telegram-type of UpdateUser
func (updateUser *UpdateUser) MessageType() string {
	return "updateUser"
}

// NewUpdateUser creates a new UpdateUser
//
// @param user New data about the user
func NewUpdateUser(user *User) *UpdateUser {
	updateUserTemp := UpdateUser{
		tdCommon: tdCommon{Type: "updateUser"},
		User:     user,
	}

	return &updateUserTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUser *UpdateUser) GetUpdateEnum() UpdateEnum {
	return UpdateUserType
}

// UpdateBasicGroup Some data of a basic group has changed. This update is guaranteed to come before the basic group identifier is returned to the application
type UpdateBasicGroup struct {
	tdCommon
	BasicGroup *BasicGroup `json:"basic_group"` // New data about the group
}

// MessageType return the string telegram-type of UpdateBasicGroup
func (updateBasicGroup *UpdateBasicGroup) MessageType() string {
	return "updateBasicGroup"
}

// NewUpdateBasicGroup creates a new UpdateBasicGroup
//
// @param basicGroup New data about the group
func NewUpdateBasicGroup(basicGroup *BasicGroup) *UpdateBasicGroup {
	updateBasicGroupTemp := UpdateBasicGroup{
		tdCommon:   tdCommon{Type: "updateBasicGroup"},
		BasicGroup: basicGroup,
	}

	return &updateBasicGroupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroup *UpdateBasicGroup) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupType
}

// UpdateSupergroup Some data of a supergroup or a channel has changed. This update is guaranteed to come before the supergroup identifier is returned to the application
type UpdateSupergroup struct {
	tdCommon
	Supergroup *Supergroup `json:"supergroup"` // New data about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroup
func (updateSupergroup *UpdateSupergroup) MessageType() string {
	return "updateSupergroup"
}

// NewUpdateSupergroup creates a new UpdateSupergroup
//
// @param supergroup New data about the supergroup
func NewUpdateSupergroup(supergroup *Supergroup) *UpdateSupergroup {
	updateSupergroupTemp := UpdateSupergroup{
		tdCommon:   tdCommon{Type: "updateSupergroup"},
		Supergroup: supergroup,
	}

	return &updateSupergroupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroup *UpdateSupergroup) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupType
}

// UpdateSecretChat Some data of a secret chat has changed. This update is guaranteed to come before the secret chat identifier is returned to the application
type UpdateSecretChat struct {
	tdCommon
	SecretChat *SecretChat `json:"secret_chat"` // New data about the secret chat
}

// MessageType return the string telegram-type of UpdateSecretChat
func (updateSecretChat *UpdateSecretChat) MessageType() string {
	return "updateSecretChat"
}

// NewUpdateSecretChat creates a new UpdateSecretChat
//
// @param secretChat New data about the secret chat
func NewUpdateSecretChat(secretChat *SecretChat) *UpdateSecretChat {
	updateSecretChatTemp := UpdateSecretChat{
		tdCommon:   tdCommon{Type: "updateSecretChat"},
		SecretChat: secretChat,
	}

	return &updateSecretChatTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSecretChat *UpdateSecretChat) GetUpdateEnum() UpdateEnum {
	return UpdateSecretChatType
}

// UpdateUserFullInfo Some data from userFullInfo has been changed
type UpdateUserFullInfo struct {
	tdCommon
	UserId       int64         `json:"user_id"`        // User identifier
	UserFullInfo *UserFullInfo `json:"user_full_info"` // New full information about the user
}

// MessageType return the string telegram-type of UpdateUserFullInfo
func (updateUserFullInfo *UpdateUserFullInfo) MessageType() string {
	return "updateUserFullInfo"
}

// NewUpdateUserFullInfo creates a new UpdateUserFullInfo
//
// @param userId User identifier
// @param userFullInfo New full information about the user
func NewUpdateUserFullInfo(userId int64, userFullInfo *UserFullInfo) *UpdateUserFullInfo {
	updateUserFullInfoTemp := UpdateUserFullInfo{
		tdCommon:     tdCommon{Type: "updateUserFullInfo"},
		UserId:       userId,
		UserFullInfo: userFullInfo,
	}

	return &updateUserFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUserFullInfo *UpdateUserFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateUserFullInfoType
}

// UpdateBasicGroupFullInfo Some data from basicGroupFullInfo has been changed
type UpdateBasicGroupFullInfo struct {
	tdCommon
	BasicGroupId       int64               `json:"basic_group_id"`        // Identifier of a basic group
	BasicGroupFullInfo *BasicGroupFullInfo `json:"basic_group_full_info"` // New full information about the group
}

// MessageType return the string telegram-type of UpdateBasicGroupFullInfo
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) MessageType() string {
	return "updateBasicGroupFullInfo"
}

// NewUpdateBasicGroupFullInfo creates a new UpdateBasicGroupFullInfo
//
// @param basicGroupId Identifier of a basic group
// @param basicGroupFullInfo New full information about the group
func NewUpdateBasicGroupFullInfo(basicGroupId int64, basicGroupFullInfo *BasicGroupFullInfo) *UpdateBasicGroupFullInfo {
	updateBasicGroupFullInfoTemp := UpdateBasicGroupFullInfo{
		tdCommon:           tdCommon{Type: "updateBasicGroupFullInfo"},
		BasicGroupId:       basicGroupId,
		BasicGroupFullInfo: basicGroupFullInfo,
	}

	return &updateBasicGroupFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupFullInfoType
}

// UpdateSupergroupFullInfo Some data from supergroupFullInfo has been changed
type UpdateSupergroupFullInfo struct {
	tdCommon
	SupergroupId       int64               `json:"supergroup_id"`        // Identifier of the supergroup or channel
	SupergroupFullInfo *SupergroupFullInfo `json:"supergroup_full_info"` // New full information about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroupFullInfo
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) MessageType() string {
	return "updateSupergroupFullInfo"
}

// NewUpdateSupergroupFullInfo creates a new UpdateSupergroupFullInfo
//
// @param supergroupId Identifier of the supergroup or channel
// @param supergroupFullInfo New full information about the supergroup
func NewUpdateSupergroupFullInfo(supergroupId int64, supergroupFullInfo *SupergroupFullInfo) *UpdateSupergroupFullInfo {
	updateSupergroupFullInfoTemp := UpdateSupergroupFullInfo{
		tdCommon:           tdCommon{Type: "updateSupergroupFullInfo"},
		SupergroupId:       supergroupId,
		SupergroupFullInfo: supergroupFullInfo,
	}

	return &updateSupergroupFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupFullInfoType
}

// UpdateServiceNotification Service notification from the server. Upon receiving this the application must show a popup with the content of the notification
type UpdateServiceNotification struct {
	tdCommon
	Type    string         `json:"type"`    // Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" must be shown under notification; if user presses the second, all local data must be destroyed using Destroy method
	Content MessageContent `json:"content"` // Notification content
}

// MessageType return the string telegram-type of UpdateServiceNotification
func (updateServiceNotification *UpdateServiceNotification) MessageType() string {
	return "updateServiceNotification"
}

// NewUpdateServiceNotification creates a new UpdateServiceNotification
//
// @param typeParam Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" must be shown under notification; if user presses the second, all local data must be destroyed using Destroy method
// @param content Notification content
func NewUpdateServiceNotification(typeParam string, content MessageContent) *UpdateServiceNotification {
	updateServiceNotificationTemp := UpdateServiceNotification{
		tdCommon: tdCommon{Type: "updateServiceNotification"},
		Type:     typeParam,
		Content:  content,
	}

	return &updateServiceNotificationTemp
}

// UnmarshalJSON unmarshal to json
func (updateServiceNotification *UpdateServiceNotification) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Type string `json:"type"` // Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" must be shown under notification; if user presses the second, all local data must be destroyed using Destroy method

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateServiceNotification.tdCommon = tempObj.tdCommon
	updateServiceNotification.Type = tempObj.Type

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	updateServiceNotification.Content = fieldContent

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateServiceNotification *UpdateServiceNotification) GetUpdateEnum() UpdateEnum {
	return UpdateServiceNotificationType
}

// UpdateFile Information about a file was updated
type UpdateFile struct {
	tdCommon
	File *File `json:"file"` // New data about the file
}

// MessageType return the string telegram-type of UpdateFile
func (updateFile *UpdateFile) MessageType() string {
	return "updateFile"
}

// NewUpdateFile creates a new UpdateFile
//
// @param file New data about the file
func NewUpdateFile(file *File) *UpdateFile {
	updateFileTemp := UpdateFile{
		tdCommon: tdCommon{Type: "updateFile"},
		File:     file,
	}

	return &updateFileTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFile *UpdateFile) GetUpdateEnum() UpdateEnum {
	return UpdateFileType
}

// UpdateFileGenerationStart The file generation process needs to be started by the application
type UpdateFileGenerationStart struct {
	tdCommon
	GenerationId    JSONInt64 `json:"generation_id"`    // Unique identifier for the generation process
	OriginalPath    string    `json:"original_path"`    // The path to a file from which a new file is generated; may be empty
	DestinationPath string    `json:"destination_path"` // The path to a file that must be created and where the new file is generated
	Conversion      string    `json:"conversion"`       // String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains an HTTP/HTTPS URL of a file, which must be downloaded by the application
}

// MessageType return the string telegram-type of UpdateFileGenerationStart
func (updateFileGenerationStart *UpdateFileGenerationStart) MessageType() string {
	return "updateFileGenerationStart"
}

// NewUpdateFileGenerationStart creates a new UpdateFileGenerationStart
//
// @param generationId Unique identifier for the generation process
// @param originalPath The path to a file from which a new file is generated; may be empty
// @param destinationPath The path to a file that must be created and where the new file is generated
// @param conversion String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains an HTTP/HTTPS URL of a file, which must be downloaded by the application
func NewUpdateFileGenerationStart(generationId JSONInt64, originalPath string, destinationPath string, conversion string) *UpdateFileGenerationStart {
	updateFileGenerationStartTemp := UpdateFileGenerationStart{
		tdCommon:        tdCommon{Type: "updateFileGenerationStart"},
		GenerationId:    generationId,
		OriginalPath:    originalPath,
		DestinationPath: destinationPath,
		Conversion:      conversion,
	}

	return &updateFileGenerationStartTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFileGenerationStart *UpdateFileGenerationStart) GetUpdateEnum() UpdateEnum {
	return UpdateFileGenerationStartType
}

// UpdateFileGenerationStop File generation is no longer needed
type UpdateFileGenerationStop struct {
	tdCommon
	GenerationId JSONInt64 `json:"generation_id"` // Unique identifier for the generation process
}

// MessageType return the string telegram-type of UpdateFileGenerationStop
func (updateFileGenerationStop *UpdateFileGenerationStop) MessageType() string {
	return "updateFileGenerationStop"
}

// NewUpdateFileGenerationStop creates a new UpdateFileGenerationStop
//
// @param generationId Unique identifier for the generation process
func NewUpdateFileGenerationStop(generationId JSONInt64) *UpdateFileGenerationStop {
	updateFileGenerationStopTemp := UpdateFileGenerationStop{
		tdCommon:     tdCommon{Type: "updateFileGenerationStop"},
		GenerationId: generationId,
	}

	return &updateFileGenerationStopTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFileGenerationStop *UpdateFileGenerationStop) GetUpdateEnum() UpdateEnum {
	return UpdateFileGenerationStopType
}

// UpdateCall New call was created or information about a call was updated
type UpdateCall struct {
	tdCommon
	Call *Call `json:"call"` // New data about a call
}

// MessageType return the string telegram-type of UpdateCall
func (updateCall *UpdateCall) MessageType() string {
	return "updateCall"
}

// NewUpdateCall creates a new UpdateCall
//
// @param call New data about a call
func NewUpdateCall(call *Call) *UpdateCall {
	updateCallTemp := UpdateCall{
		tdCommon: tdCommon{Type: "updateCall"},
		Call:     call,
	}

	return &updateCallTemp
}

// GetUpdateEnum return the enum type of this object
func (updateCall *UpdateCall) GetUpdateEnum() UpdateEnum {
	return UpdateCallType
}

// UpdateGroupCall Information about a group call was updated
type UpdateGroupCall struct {
	tdCommon
	GroupCall *GroupCall `json:"group_call"` // New data about a group call
}

// MessageType return the string telegram-type of UpdateGroupCall
func (updateGroupCall *UpdateGroupCall) MessageType() string {
	return "updateGroupCall"
}

// NewUpdateGroupCall creates a new UpdateGroupCall
//
// @param groupCall New data about a group call
func NewUpdateGroupCall(groupCall *GroupCall) *UpdateGroupCall {
	updateGroupCallTemp := UpdateGroupCall{
		tdCommon:  tdCommon{Type: "updateGroupCall"},
		GroupCall: groupCall,
	}

	return &updateGroupCallTemp
}

// GetUpdateEnum return the enum type of this object
func (updateGroupCall *UpdateGroupCall) GetUpdateEnum() UpdateEnum {
	return UpdateGroupCallType
}

// UpdateGroupCallParticipant Information about a group call participant was changed. The updates are sent only after the group call is received through getGroupCall and only if the call is joined or being joined
type UpdateGroupCallParticipant struct {
	tdCommon
	GroupCallId int32                 `json:"group_call_id"` // Identifier of group call
	Participant *GroupCallParticipant `json:"participant"`   // New data about a participant
}

// MessageType return the string telegram-type of UpdateGroupCallParticipant
func (updateGroupCallParticipant *UpdateGroupCallParticipant) MessageType() string {
	return "updateGroupCallParticipant"
}

// NewUpdateGroupCallParticipant creates a new UpdateGroupCallParticipant
//
// @param groupCallId Identifier of group call
// @param participant New data about a participant
func NewUpdateGroupCallParticipant(groupCallId int32, participant *GroupCallParticipant) *UpdateGroupCallParticipant {
	updateGroupCallParticipantTemp := UpdateGroupCallParticipant{
		tdCommon:    tdCommon{Type: "updateGroupCallParticipant"},
		GroupCallId: groupCallId,
		Participant: participant,
	}

	return &updateGroupCallParticipantTemp
}

// GetUpdateEnum return the enum type of this object
func (updateGroupCallParticipant *UpdateGroupCallParticipant) GetUpdateEnum() UpdateEnum {
	return UpdateGroupCallParticipantType
}

// UpdateNewCallSignalingData New call signaling data arrived
type UpdateNewCallSignalingData struct {
	tdCommon
	CallId int32  `json:"call_id"` // The call identifier
	Data   []byte `json:"data"`    // The data
}

// MessageType return the string telegram-type of UpdateNewCallSignalingData
func (updateNewCallSignalingData *UpdateNewCallSignalingData) MessageType() string {
	return "updateNewCallSignalingData"
}

// NewUpdateNewCallSignalingData creates a new UpdateNewCallSignalingData
//
// @param callId The call identifier
// @param data The data
func NewUpdateNewCallSignalingData(callId int32, data []byte) *UpdateNewCallSignalingData {
	updateNewCallSignalingDataTemp := UpdateNewCallSignalingData{
		tdCommon: tdCommon{Type: "updateNewCallSignalingData"},
		CallId:   callId,
		Data:     data,
	}

	return &updateNewCallSignalingDataTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewCallSignalingData *UpdateNewCallSignalingData) GetUpdateEnum() UpdateEnum {
	return UpdateNewCallSignalingDataType
}

// UpdateUserPrivacySettingRules Some privacy setting rules have been changed
type UpdateUserPrivacySettingRules struct {
	tdCommon
	Setting UserPrivacySetting       `json:"setting"` // The privacy setting
	Rules   *UserPrivacySettingRules `json:"rules"`   // New privacy rules
}

// MessageType return the string telegram-type of UpdateUserPrivacySettingRules
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) MessageType() string {
	return "updateUserPrivacySettingRules"
}

// NewUpdateUserPrivacySettingRules creates a new UpdateUserPrivacySettingRules
//
// @param setting The privacy setting
// @param rules New privacy rules
func NewUpdateUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) *UpdateUserPrivacySettingRules {
	updateUserPrivacySettingRulesTemp := UpdateUserPrivacySettingRules{
		tdCommon: tdCommon{Type: "updateUserPrivacySettingRules"},
		Setting:  setting,
		Rules:    rules,
	}

	return &updateUserPrivacySettingRulesTemp
}

// UnmarshalJSON unmarshal to json
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Rules *UserPrivacySettingRules `json:"rules"` // New privacy rules
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserPrivacySettingRules.tdCommon = tempObj.tdCommon
	updateUserPrivacySettingRules.Rules = tempObj.Rules

	fieldSetting, _ := unmarshalUserPrivacySetting(objMap["setting"])
	updateUserPrivacySettingRules.Setting = fieldSetting

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) GetUpdateEnum() UpdateEnum {
	return UpdateUserPrivacySettingRulesType
}

// UpdateUnreadMessageCount Number of unread messages in a chat list has changed. This update is sent only if the message database is used
type UpdateUnreadMessageCount struct {
	tdCommon
	ChatList           ChatList `json:"chat_list"`            // The chat list with changed number of unread messages
	UnreadCount        int32    `json:"unread_count"`         // Total number of unread messages
	UnreadUnmutedCount int32    `json:"unread_unmuted_count"` // Total number of unread messages in unmuted chats
}

// MessageType return the string telegram-type of UpdateUnreadMessageCount
func (updateUnreadMessageCount *UpdateUnreadMessageCount) MessageType() string {
	return "updateUnreadMessageCount"
}

// NewUpdateUnreadMessageCount creates a new UpdateUnreadMessageCount
//
// @param chatList The chat list with changed number of unread messages
// @param unreadCount Total number of unread messages
// @param unreadUnmutedCount Total number of unread messages in unmuted chats
func NewUpdateUnreadMessageCount(chatList ChatList, unreadCount int32, unreadUnmutedCount int32) *UpdateUnreadMessageCount {
	updateUnreadMessageCountTemp := UpdateUnreadMessageCount{
		tdCommon:           tdCommon{Type: "updateUnreadMessageCount"},
		ChatList:           chatList,
		UnreadCount:        unreadCount,
		UnreadUnmutedCount: unreadUnmutedCount,
	}

	return &updateUnreadMessageCountTemp
}

// UnmarshalJSON unmarshal to json
func (updateUnreadMessageCount *UpdateUnreadMessageCount) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UnreadCount        int32 `json:"unread_count"`         // Total number of unread messages
		UnreadUnmutedCount int32 `json:"unread_unmuted_count"` // Total number of unread messages in unmuted chats
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUnreadMessageCount.tdCommon = tempObj.tdCommon
	updateUnreadMessageCount.UnreadCount = tempObj.UnreadCount
	updateUnreadMessageCount.UnreadUnmutedCount = tempObj.UnreadUnmutedCount

	fieldChatList, _ := unmarshalChatList(objMap["chat_list"])
	updateUnreadMessageCount.ChatList = fieldChatList

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUnreadMessageCount *UpdateUnreadMessageCount) GetUpdateEnum() UpdateEnum {
	return UpdateUnreadMessageCountType
}

// UpdateUnreadChatCount Number of unread chats, i.e. with unread messages or marked as unread, has changed. This update is sent only if the message database is used
type UpdateUnreadChatCount struct {
	tdCommon
	ChatList                   ChatList `json:"chat_list"`                      // The chat list with changed number of unread messages
	TotalCount                 int32    `json:"total_count"`                    // Approximate total number of chats in the chat list
	UnreadCount                int32    `json:"unread_count"`                   // Total number of unread chats
	UnreadUnmutedCount         int32    `json:"unread_unmuted_count"`           // Total number of unread unmuted chats
	MarkedAsUnreadCount        int32    `json:"marked_as_unread_count"`         // Total number of chats marked as unread
	MarkedAsUnreadUnmutedCount int32    `json:"marked_as_unread_unmuted_count"` // Total number of unmuted chats marked as unread
}

// MessageType return the string telegram-type of UpdateUnreadChatCount
func (updateUnreadChatCount *UpdateUnreadChatCount) MessageType() string {
	return "updateUnreadChatCount"
}

// NewUpdateUnreadChatCount creates a new UpdateUnreadChatCount
//
// @param chatList The chat list with changed number of unread messages
// @param totalCount Approximate total number of chats in the chat list
// @param unreadCount Total number of unread chats
// @param unreadUnmutedCount Total number of unread unmuted chats
// @param markedAsUnreadCount Total number of chats marked as unread
// @param markedAsUnreadUnmutedCount Total number of unmuted chats marked as unread
func NewUpdateUnreadChatCount(chatList ChatList, totalCount int32, unreadCount int32, unreadUnmutedCount int32, markedAsUnreadCount int32, markedAsUnreadUnmutedCount int32) *UpdateUnreadChatCount {
	updateUnreadChatCountTemp := UpdateUnreadChatCount{
		tdCommon:                   tdCommon{Type: "updateUnreadChatCount"},
		ChatList:                   chatList,
		TotalCount:                 totalCount,
		UnreadCount:                unreadCount,
		UnreadUnmutedCount:         unreadUnmutedCount,
		MarkedAsUnreadCount:        markedAsUnreadCount,
		MarkedAsUnreadUnmutedCount: markedAsUnreadUnmutedCount,
	}

	return &updateUnreadChatCountTemp
}

// UnmarshalJSON unmarshal to json
func (updateUnreadChatCount *UpdateUnreadChatCount) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		TotalCount                 int32 `json:"total_count"`                    // Approximate total number of chats in the chat list
		UnreadCount                int32 `json:"unread_count"`                   // Total number of unread chats
		UnreadUnmutedCount         int32 `json:"unread_unmuted_count"`           // Total number of unread unmuted chats
		MarkedAsUnreadCount        int32 `json:"marked_as_unread_count"`         // Total number of chats marked as unread
		MarkedAsUnreadUnmutedCount int32 `json:"marked_as_unread_unmuted_count"` // Total number of unmuted chats marked as unread
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUnreadChatCount.tdCommon = tempObj.tdCommon
	updateUnreadChatCount.TotalCount = tempObj.TotalCount
	updateUnreadChatCount.UnreadCount = tempObj.UnreadCount
	updateUnreadChatCount.UnreadUnmutedCount = tempObj.UnreadUnmutedCount
	updateUnreadChatCount.MarkedAsUnreadCount = tempObj.MarkedAsUnreadCount
	updateUnreadChatCount.MarkedAsUnreadUnmutedCount = tempObj.MarkedAsUnreadUnmutedCount

	fieldChatList, _ := unmarshalChatList(objMap["chat_list"])
	updateUnreadChatCount.ChatList = fieldChatList

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUnreadChatCount *UpdateUnreadChatCount) GetUpdateEnum() UpdateEnum {
	return UpdateUnreadChatCountType
}

// UpdateOption An option changed its value
type UpdateOption struct {
	tdCommon
	Name  string      `json:"name"`  // The option name
	Value OptionValue `json:"value"` // The new option value
}

// MessageType return the string telegram-type of UpdateOption
func (updateOption *UpdateOption) MessageType() string {
	return "updateOption"
}

// NewUpdateOption creates a new UpdateOption
//
// @param name The option name
// @param value The new option value
func NewUpdateOption(name string, value OptionValue) *UpdateOption {
	updateOptionTemp := UpdateOption{
		tdCommon: tdCommon{Type: "updateOption"},
		Name:     name,
		Value:    value,
	}

	return &updateOptionTemp
}

// UnmarshalJSON unmarshal to json
func (updateOption *UpdateOption) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Name string `json:"name"` // The option name

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateOption.tdCommon = tempObj.tdCommon
	updateOption.Name = tempObj.Name

	fieldValue, _ := unmarshalOptionValue(objMap["value"])
	updateOption.Value = fieldValue

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateOption *UpdateOption) GetUpdateEnum() UpdateEnum {
	return UpdateOptionType
}

// UpdateStickerSet A sticker set has changed
type UpdateStickerSet struct {
	tdCommon
	StickerSet *StickerSet `json:"sticker_set"` // The sticker set
}

// MessageType return the string telegram-type of UpdateStickerSet
func (updateStickerSet *UpdateStickerSet) MessageType() string {
	return "updateStickerSet"
}

// NewUpdateStickerSet creates a new UpdateStickerSet
//
// @param stickerSet The sticker set
func NewUpdateStickerSet(stickerSet *StickerSet) *UpdateStickerSet {
	updateStickerSetTemp := UpdateStickerSet{
		tdCommon:   tdCommon{Type: "updateStickerSet"},
		StickerSet: stickerSet,
	}

	return &updateStickerSetTemp
}

// GetUpdateEnum return the enum type of this object
func (updateStickerSet *UpdateStickerSet) GetUpdateEnum() UpdateEnum {
	return UpdateStickerSetType
}

// UpdateInstalledStickerSets The list of installed sticker sets was updated
type UpdateInstalledStickerSets struct {
	tdCommon
	IsMasks       bool        `json:"is_masks"`        // True, if the list of installed mask sticker sets was updated
	StickerSetIds []JSONInt64 `json:"sticker_set_ids"` // The new list of installed ordinary sticker sets
}

// MessageType return the string telegram-type of UpdateInstalledStickerSets
func (updateInstalledStickerSets *UpdateInstalledStickerSets) MessageType() string {
	return "updateInstalledStickerSets"
}

// NewUpdateInstalledStickerSets creates a new UpdateInstalledStickerSets
//
// @param isMasks True, if the list of installed mask sticker sets was updated
// @param stickerSetIds The new list of installed ordinary sticker sets
func NewUpdateInstalledStickerSets(isMasks bool, stickerSetIds []JSONInt64) *UpdateInstalledStickerSets {
	updateInstalledStickerSetsTemp := UpdateInstalledStickerSets{
		tdCommon:      tdCommon{Type: "updateInstalledStickerSets"},
		IsMasks:       isMasks,
		StickerSetIds: stickerSetIds,
	}

	return &updateInstalledStickerSetsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateInstalledStickerSets *UpdateInstalledStickerSets) GetUpdateEnum() UpdateEnum {
	return UpdateInstalledStickerSetsType
}

// UpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed
type UpdateTrendingStickerSets struct {
	tdCommon
	StickerSets *StickerSets `json:"sticker_sets"` // The prefix of the list of trending sticker sets with the newest trending sticker sets
}

// MessageType return the string telegram-type of UpdateTrendingStickerSets
func (updateTrendingStickerSets *UpdateTrendingStickerSets) MessageType() string {
	return "updateTrendingStickerSets"
}

// NewUpdateTrendingStickerSets creates a new UpdateTrendingStickerSets
//
// @param stickerSets The prefix of the list of trending sticker sets with the newest trending sticker sets
func NewUpdateTrendingStickerSets(stickerSets *StickerSets) *UpdateTrendingStickerSets {
	updateTrendingStickerSetsTemp := UpdateTrendingStickerSets{
		tdCommon:    tdCommon{Type: "updateTrendingStickerSets"},
		StickerSets: stickerSets,
	}

	return &updateTrendingStickerSetsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateTrendingStickerSets *UpdateTrendingStickerSets) GetUpdateEnum() UpdateEnum {
	return UpdateTrendingStickerSetsType
}

// UpdateRecentStickers The list of recently used stickers was updated
type UpdateRecentStickers struct {
	tdCommon
	IsAttached bool    `json:"is_attached"` // True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated
	StickerIds []int32 `json:"sticker_ids"` // The new list of file identifiers of recently used stickers
}

// MessageType return the string telegram-type of UpdateRecentStickers
func (updateRecentStickers *UpdateRecentStickers) MessageType() string {
	return "updateRecentStickers"
}

// NewUpdateRecentStickers creates a new UpdateRecentStickers
//
// @param isAttached True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated
// @param stickerIds The new list of file identifiers of recently used stickers
func NewUpdateRecentStickers(isAttached bool, stickerIds []int32) *UpdateRecentStickers {
	updateRecentStickersTemp := UpdateRecentStickers{
		tdCommon:   tdCommon{Type: "updateRecentStickers"},
		IsAttached: isAttached,
		StickerIds: stickerIds,
	}

	return &updateRecentStickersTemp
}

// GetUpdateEnum return the enum type of this object
func (updateRecentStickers *UpdateRecentStickers) GetUpdateEnum() UpdateEnum {
	return UpdateRecentStickersType
}

// UpdateFavoriteStickers The list of favorite stickers was updated
type UpdateFavoriteStickers struct {
	tdCommon
	StickerIds []int32 `json:"sticker_ids"` // The new list of file identifiers of favorite stickers
}

// MessageType return the string telegram-type of UpdateFavoriteStickers
func (updateFavoriteStickers *UpdateFavoriteStickers) MessageType() string {
	return "updateFavoriteStickers"
}

// NewUpdateFavoriteStickers creates a new UpdateFavoriteStickers
//
// @param stickerIds The new list of file identifiers of favorite stickers
func NewUpdateFavoriteStickers(stickerIds []int32) *UpdateFavoriteStickers {
	updateFavoriteStickersTemp := UpdateFavoriteStickers{
		tdCommon:   tdCommon{Type: "updateFavoriteStickers"},
		StickerIds: stickerIds,
	}

	return &updateFavoriteStickersTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFavoriteStickers *UpdateFavoriteStickers) GetUpdateEnum() UpdateEnum {
	return UpdateFavoriteStickersType
}

// UpdateSavedAnimations The list of saved animations was updated
type UpdateSavedAnimations struct {
	tdCommon
	AnimationIds []int32 `json:"animation_ids"` // The new list of file identifiers of saved animations
}

// MessageType return the string telegram-type of UpdateSavedAnimations
func (updateSavedAnimations *UpdateSavedAnimations) MessageType() string {
	return "updateSavedAnimations"
}

// NewUpdateSavedAnimations creates a new UpdateSavedAnimations
//
// @param animationIds The new list of file identifiers of saved animations
func NewUpdateSavedAnimations(animationIds []int32) *UpdateSavedAnimations {
	updateSavedAnimationsTemp := UpdateSavedAnimations{
		tdCommon:     tdCommon{Type: "updateSavedAnimations"},
		AnimationIds: animationIds,
	}

	return &updateSavedAnimationsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSavedAnimations *UpdateSavedAnimations) GetUpdateEnum() UpdateEnum {
	return UpdateSavedAnimationsType
}

// UpdateSelectedBackground The selected background has changed
type UpdateSelectedBackground struct {
	tdCommon
	ForDarkTheme bool        `json:"for_dark_theme"` // True, if background for dark theme has changed
	Background   *Background `json:"background"`     // The new selected background; may be null
}

// MessageType return the string telegram-type of UpdateSelectedBackground
func (updateSelectedBackground *UpdateSelectedBackground) MessageType() string {
	return "updateSelectedBackground"
}

// NewUpdateSelectedBackground creates a new UpdateSelectedBackground
//
// @param forDarkTheme True, if background for dark theme has changed
// @param background The new selected background; may be null
func NewUpdateSelectedBackground(forDarkTheme bool, background *Background) *UpdateSelectedBackground {
	updateSelectedBackgroundTemp := UpdateSelectedBackground{
		tdCommon:     tdCommon{Type: "updateSelectedBackground"},
		ForDarkTheme: forDarkTheme,
		Background:   background,
	}

	return &updateSelectedBackgroundTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSelectedBackground *UpdateSelectedBackground) GetUpdateEnum() UpdateEnum {
	return UpdateSelectedBackgroundType
}

// UpdateChatThemes The list of available chat themes has changed
type UpdateChatThemes struct {
	tdCommon
	ChatThemes []ChatTheme `json:"chat_themes"` // The new list of chat themes
}

// MessageType return the string telegram-type of UpdateChatThemes
func (updateChatThemes *UpdateChatThemes) MessageType() string {
	return "updateChatThemes"
}

// NewUpdateChatThemes creates a new UpdateChatThemes
//
// @param chatThemes The new list of chat themes
func NewUpdateChatThemes(chatThemes []ChatTheme) *UpdateChatThemes {
	updateChatThemesTemp := UpdateChatThemes{
		tdCommon:   tdCommon{Type: "updateChatThemes"},
		ChatThemes: chatThemes,
	}

	return &updateChatThemesTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatThemes *UpdateChatThemes) GetUpdateEnum() UpdateEnum {
	return UpdateChatThemesType
}

// UpdateLanguagePackStrings Some language pack strings have been updated
type UpdateLanguagePackStrings struct {
	tdCommon
	LocalizationTarget string               `json:"localization_target"` // Localization target to which the language pack belongs
	LanguagePackId     string               `json:"language_pack_id"`    // Identifier of the updated language pack
	Strings            []LanguagePackString `json:"strings"`             // List of changed language pack strings
}

// MessageType return the string telegram-type of UpdateLanguagePackStrings
func (updateLanguagePackStrings *UpdateLanguagePackStrings) MessageType() string {
	return "updateLanguagePackStrings"
}

// NewUpdateLanguagePackStrings creates a new UpdateLanguagePackStrings
//
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackId Identifier of the updated language pack
// @param strings List of changed language pack strings
func NewUpdateLanguagePackStrings(localizationTarget string, languagePackId string, strings []LanguagePackString) *UpdateLanguagePackStrings {
	updateLanguagePackStringsTemp := UpdateLanguagePackStrings{
		tdCommon:           tdCommon{Type: "updateLanguagePackStrings"},
		LocalizationTarget: localizationTarget,
		LanguagePackId:     languagePackId,
		Strings:            strings,
	}

	return &updateLanguagePackStringsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateLanguagePackStrings *UpdateLanguagePackStrings) GetUpdateEnum() UpdateEnum {
	return UpdateLanguagePackStringsType
}

// UpdateConnectionState The connection state has changed. This update must be used only to show a human-readable description of the connection state
type UpdateConnectionState struct {
	tdCommon
	State ConnectionState `json:"state"` // The new connection state
}

// MessageType return the string telegram-type of UpdateConnectionState
func (updateConnectionState *UpdateConnectionState) MessageType() string {
	return "updateConnectionState"
}

// NewUpdateConnectionState creates a new UpdateConnectionState
//
// @param state The new connection state
func NewUpdateConnectionState(state ConnectionState) *UpdateConnectionState {
	updateConnectionStateTemp := UpdateConnectionState{
		tdCommon: tdCommon{Type: "updateConnectionState"},
		State:    state,
	}

	return &updateConnectionStateTemp
}

// UnmarshalJSON unmarshal to json
func (updateConnectionState *UpdateConnectionState) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateConnectionState.tdCommon = tempObj.tdCommon

	fieldState, _ := unmarshalConnectionState(objMap["state"])
	updateConnectionState.State = fieldState

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateConnectionState *UpdateConnectionState) GetUpdateEnum() UpdateEnum {
	return UpdateConnectionStateType
}

// UpdateTermsOfService New terms of service must be accepted by the user. If the terms of service are declined, then the deleteAccount method must be called with the reason "Decline ToS update"
type UpdateTermsOfService struct {
	tdCommon
	TermsOfServiceId string          `json:"terms_of_service_id"` // Identifier of the terms of service
	TermsOfService   *TermsOfService `json:"terms_of_service"`    // The new terms of service
}

// MessageType return the string telegram-type of UpdateTermsOfService
func (updateTermsOfService *UpdateTermsOfService) MessageType() string {
	return "updateTermsOfService"
}

// NewUpdateTermsOfService creates a new UpdateTermsOfService
//
// @param termsOfServiceId Identifier of the terms of service
// @param termsOfService The new terms of service
func NewUpdateTermsOfService(termsOfServiceId string, termsOfService *TermsOfService) *UpdateTermsOfService {
	updateTermsOfServiceTemp := UpdateTermsOfService{
		tdCommon:         tdCommon{Type: "updateTermsOfService"},
		TermsOfServiceId: termsOfServiceId,
		TermsOfService:   termsOfService,
	}

	return &updateTermsOfServiceTemp
}

// GetUpdateEnum return the enum type of this object
func (updateTermsOfService *UpdateTermsOfService) GetUpdateEnum() UpdateEnum {
	return UpdateTermsOfServiceType
}

// UpdateUsersNearby The list of users nearby has changed. The update is guaranteed to be sent only 60 seconds after a successful searchChatsNearby request
type UpdateUsersNearby struct {
	tdCommon
	UsersNearby []ChatNearby `json:"users_nearby"` // The new list of users nearby
}

// MessageType return the string telegram-type of UpdateUsersNearby
func (updateUsersNearby *UpdateUsersNearby) MessageType() string {
	return "updateUsersNearby"
}

// NewUpdateUsersNearby creates a new UpdateUsersNearby
//
// @param usersNearby The new list of users nearby
func NewUpdateUsersNearby(usersNearby []ChatNearby) *UpdateUsersNearby {
	updateUsersNearbyTemp := UpdateUsersNearby{
		tdCommon:    tdCommon{Type: "updateUsersNearby"},
		UsersNearby: usersNearby,
	}

	return &updateUsersNearbyTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUsersNearby *UpdateUsersNearby) GetUpdateEnum() UpdateEnum {
	return UpdateUsersNearbyType
}

// UpdateDiceEmojis The list of supported dice emojis has changed
type UpdateDiceEmojis struct {
	tdCommon
	Emojis []string `json:"emojis"` // The new list of supported dice emojis
}

// MessageType return the string telegram-type of UpdateDiceEmojis
func (updateDiceEmojis *UpdateDiceEmojis) MessageType() string {
	return "updateDiceEmojis"
}

// NewUpdateDiceEmojis creates a new UpdateDiceEmojis
//
// @param emojis The new list of supported dice emojis
func NewUpdateDiceEmojis(emojis []string) *UpdateDiceEmojis {
	updateDiceEmojisTemp := UpdateDiceEmojis{
		tdCommon: tdCommon{Type: "updateDiceEmojis"},
		Emojis:   emojis,
	}

	return &updateDiceEmojisTemp
}

// GetUpdateEnum return the enum type of this object
func (updateDiceEmojis *UpdateDiceEmojis) GetUpdateEnum() UpdateEnum {
	return UpdateDiceEmojisType
}

// UpdateAnimatedEmojiMessageClicked Some animated emoji message was clicked and a big animated sticker must be played if the message is visible on the screen. chatActionWatchingAnimations with the text of the message needs to be sent if the sticker is played
type UpdateAnimatedEmojiMessageClicked struct {
	tdCommon
	ChatId    int64    `json:"chat_id"`    // Chat identifier
	MessageId int64    `json:"message_id"` // Message identifier
	Sticker   *Sticker `json:"sticker"`    // The animated sticker to be played
}

// MessageType return the string telegram-type of UpdateAnimatedEmojiMessageClicked
func (updateAnimatedEmojiMessageClicked *UpdateAnimatedEmojiMessageClicked) MessageType() string {
	return "updateAnimatedEmojiMessageClicked"
}

// NewUpdateAnimatedEmojiMessageClicked creates a new UpdateAnimatedEmojiMessageClicked
//
// @param chatId Chat identifier
// @param messageId Message identifier
// @param sticker The animated sticker to be played
func NewUpdateAnimatedEmojiMessageClicked(chatId int64, messageId int64, sticker *Sticker) *UpdateAnimatedEmojiMessageClicked {
	updateAnimatedEmojiMessageClickedTemp := UpdateAnimatedEmojiMessageClicked{
		tdCommon:  tdCommon{Type: "updateAnimatedEmojiMessageClicked"},
		ChatId:    chatId,
		MessageId: messageId,
		Sticker:   sticker,
	}

	return &updateAnimatedEmojiMessageClickedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateAnimatedEmojiMessageClicked *UpdateAnimatedEmojiMessageClicked) GetUpdateEnum() UpdateEnum {
	return UpdateAnimatedEmojiMessageClickedType
}

// UpdateAnimationSearchParameters The parameters of animation search through GetOption("animation_search_bot_username") bot has changed
type UpdateAnimationSearchParameters struct {
	tdCommon
	Provider string   `json:"provider"` // Name of the animation search provider
	Emojis   []string `json:"emojis"`   // The new list of emojis suggested for searching
}

// MessageType return the string telegram-type of UpdateAnimationSearchParameters
func (updateAnimationSearchParameters *UpdateAnimationSearchParameters) MessageType() string {
	return "updateAnimationSearchParameters"
}

// NewUpdateAnimationSearchParameters creates a new UpdateAnimationSearchParameters
//
// @param provider Name of the animation search provider
// @param emojis The new list of emojis suggested for searching
func NewUpdateAnimationSearchParameters(provider string, emojis []string) *UpdateAnimationSearchParameters {
	updateAnimationSearchParametersTemp := UpdateAnimationSearchParameters{
		tdCommon: tdCommon{Type: "updateAnimationSearchParameters"},
		Provider: provider,
		Emojis:   emojis,
	}

	return &updateAnimationSearchParametersTemp
}

// GetUpdateEnum return the enum type of this object
func (updateAnimationSearchParameters *UpdateAnimationSearchParameters) GetUpdateEnum() UpdateEnum {
	return UpdateAnimationSearchParametersType
}

// UpdateSuggestedActions The list of suggested to the user actions has changed
type UpdateSuggestedActions struct {
	tdCommon
	AddedActions   []SuggestedAction `json:"added_actions"`   // Added suggested actions
	RemovedActions []SuggestedAction `json:"removed_actions"` // Removed suggested actions
}

// MessageType return the string telegram-type of UpdateSuggestedActions
func (updateSuggestedActions *UpdateSuggestedActions) MessageType() string {
	return "updateSuggestedActions"
}

// NewUpdateSuggestedActions creates a new UpdateSuggestedActions
//
// @param addedActions Added suggested actions
// @param removedActions Removed suggested actions
func NewUpdateSuggestedActions(addedActions []SuggestedAction, removedActions []SuggestedAction) *UpdateSuggestedActions {
	updateSuggestedActionsTemp := UpdateSuggestedActions{
		tdCommon:       tdCommon{Type: "updateSuggestedActions"},
		AddedActions:   addedActions,
		RemovedActions: removedActions,
	}

	return &updateSuggestedActionsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSuggestedActions *UpdateSuggestedActions) GetUpdateEnum() UpdateEnum {
	return UpdateSuggestedActionsType
}

// UpdateNewInlineQuery A new incoming inline query; for bots only
type UpdateNewInlineQuery struct {
	tdCommon
	Id           JSONInt64 `json:"id"`             // Unique query identifier
	SenderUserId int64     `json:"sender_user_id"` // Identifier of the user who sent the query
	UserLocation *Location `json:"user_location"`  // User location; may be null
	ChatType     ChatType  `json:"chat_type"`      // The type of the chat, from which the query originated; may be null if unknown
	Query        string    `json:"query"`          // Text of the query
	Offset       string    `json:"offset"`         // Offset of the first entry to return
}

// MessageType return the string telegram-type of UpdateNewInlineQuery
func (updateNewInlineQuery *UpdateNewInlineQuery) MessageType() string {
	return "updateNewInlineQuery"
}

// NewUpdateNewInlineQuery creates a new UpdateNewInlineQuery
//
// @param id Unique query identifier
// @param senderUserId Identifier of the user who sent the query
// @param userLocation User location; may be null
// @param chatType The type of the chat, from which the query originated; may be null if unknown
// @param query Text of the query
// @param offset Offset of the first entry to return
func NewUpdateNewInlineQuery(id JSONInt64, senderUserId int64, userLocation *Location, chatType ChatType, query string, offset string) *UpdateNewInlineQuery {
	updateNewInlineQueryTemp := UpdateNewInlineQuery{
		tdCommon:     tdCommon{Type: "updateNewInlineQuery"},
		Id:           id,
		SenderUserId: senderUserId,
		UserLocation: userLocation,
		ChatType:     chatType,
		Query:        query,
		Offset:       offset,
	}

	return &updateNewInlineQueryTemp
}

// UnmarshalJSON unmarshal to json
func (updateNewInlineQuery *UpdateNewInlineQuery) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id           JSONInt64 `json:"id"`             // Unique query identifier
		SenderUserId int64     `json:"sender_user_id"` // Identifier of the user who sent the query
		UserLocation *Location `json:"user_location"`  // User location; may be null
		Query        string    `json:"query"`          // Text of the query
		Offset       string    `json:"offset"`         // Offset of the first entry to return
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNewInlineQuery.tdCommon = tempObj.tdCommon
	updateNewInlineQuery.Id = tempObj.Id
	updateNewInlineQuery.SenderUserId = tempObj.SenderUserId
	updateNewInlineQuery.UserLocation = tempObj.UserLocation
	updateNewInlineQuery.Query = tempObj.Query
	updateNewInlineQuery.Offset = tempObj.Offset

	fieldChatType, _ := unmarshalChatType(objMap["chat_type"])
	updateNewInlineQuery.ChatType = fieldChatType

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNewInlineQuery *UpdateNewInlineQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewInlineQueryType
}

// UpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only
type UpdateNewChosenInlineResult struct {
	tdCommon
	SenderUserId    int64     `json:"sender_user_id"`    // Identifier of the user who sent the query
	UserLocation    *Location `json:"user_location"`     // User location; may be null
	Query           string    `json:"query"`             // Text of the query
	ResultId        string    `json:"result_id"`         // Identifier of the chosen result
	InlineMessageId string    `json:"inline_message_id"` // Identifier of the sent inline message, if known
}

// MessageType return the string telegram-type of UpdateNewChosenInlineResult
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) MessageType() string {
	return "updateNewChosenInlineResult"
}

// NewUpdateNewChosenInlineResult creates a new UpdateNewChosenInlineResult
//
// @param senderUserId Identifier of the user who sent the query
// @param userLocation User location; may be null
// @param query Text of the query
// @param resultId Identifier of the chosen result
// @param inlineMessageId Identifier of the sent inline message, if known
func NewUpdateNewChosenInlineResult(senderUserId int64, userLocation *Location, query string, resultId string, inlineMessageId string) *UpdateNewChosenInlineResult {
	updateNewChosenInlineResultTemp := UpdateNewChosenInlineResult{
		tdCommon:        tdCommon{Type: "updateNewChosenInlineResult"},
		SenderUserId:    senderUserId,
		UserLocation:    userLocation,
		Query:           query,
		ResultId:        resultId,
		InlineMessageId: inlineMessageId,
	}

	return &updateNewChosenInlineResultTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) GetUpdateEnum() UpdateEnum {
	return UpdateNewChosenInlineResultType
}

// UpdateNewCallbackQuery A new incoming callback query; for bots only
type UpdateNewCallbackQuery struct {
	tdCommon
	Id           JSONInt64            `json:"id"`             // Unique query identifier
	SenderUserId int64                `json:"sender_user_id"` // Identifier of the user who sent the query
	ChatId       int64                `json:"chat_id"`        // Identifier of the chat where the query was sent
	MessageId    int64                `json:"message_id"`     // Identifier of the message, from which the query originated
	ChatInstance JSONInt64            `json:"chat_instance"`  // Identifier that uniquely corresponds to the chat to which the message was sent
	Payload      CallbackQueryPayload `json:"payload"`        // Query payload
}

// MessageType return the string telegram-type of UpdateNewCallbackQuery
func (updateNewCallbackQuery *UpdateNewCallbackQuery) MessageType() string {
	return "updateNewCallbackQuery"
}

// NewUpdateNewCallbackQuery creates a new UpdateNewCallbackQuery
//
// @param id Unique query identifier
// @param senderUserId Identifier of the user who sent the query
// @param chatId Identifier of the chat where the query was sent
// @param messageId Identifier of the message, from which the query originated
// @param chatInstance Identifier that uniquely corresponds to the chat to which the message was sent
// @param payload Query payload
func NewUpdateNewCallbackQuery(id JSONInt64, senderUserId int64, chatId int64, messageId int64, chatInstance JSONInt64, payload CallbackQueryPayload) *UpdateNewCallbackQuery {
	updateNewCallbackQueryTemp := UpdateNewCallbackQuery{
		tdCommon:     tdCommon{Type: "updateNewCallbackQuery"},
		Id:           id,
		SenderUserId: senderUserId,
		ChatId:       chatId,
		MessageId:    messageId,
		ChatInstance: chatInstance,
		Payload:      payload,
	}

	return &updateNewCallbackQueryTemp
}

// UnmarshalJSON unmarshal to json
func (updateNewCallbackQuery *UpdateNewCallbackQuery) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id           JSONInt64 `json:"id"`             // Unique query identifier
		SenderUserId int64     `json:"sender_user_id"` // Identifier of the user who sent the query
		ChatId       int64     `json:"chat_id"`        // Identifier of the chat where the query was sent
		MessageId    int64     `json:"message_id"`     // Identifier of the message, from which the query originated
		ChatInstance JSONInt64 `json:"chat_instance"`  // Identifier that uniquely corresponds to the chat to which the message was sent

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNewCallbackQuery.tdCommon = tempObj.tdCommon
	updateNewCallbackQuery.Id = tempObj.Id
	updateNewCallbackQuery.SenderUserId = tempObj.SenderUserId
	updateNewCallbackQuery.ChatId = tempObj.ChatId
	updateNewCallbackQuery.MessageId = tempObj.MessageId
	updateNewCallbackQuery.ChatInstance = tempObj.ChatInstance

	fieldPayload, _ := unmarshalCallbackQueryPayload(objMap["payload"])
	updateNewCallbackQuery.Payload = fieldPayload

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNewCallbackQuery *UpdateNewCallbackQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewCallbackQueryType
}

// UpdateNewInlineCallbackQuery A new incoming callback query from a message sent via a bot; for bots only
type UpdateNewInlineCallbackQuery struct {
	tdCommon
	Id              JSONInt64            `json:"id"`                // Unique query identifier
	SenderUserId    int64                `json:"sender_user_id"`    // Identifier of the user who sent the query
	InlineMessageId string               `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
	ChatInstance    JSONInt64            `json:"chat_instance"`     // An identifier uniquely corresponding to the chat a message was sent to
	Payload         CallbackQueryPayload `json:"payload"`           // Query payload
}

// MessageType return the string telegram-type of UpdateNewInlineCallbackQuery
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) MessageType() string {
	return "updateNewInlineCallbackQuery"
}

// NewUpdateNewInlineCallbackQuery creates a new UpdateNewInlineCallbackQuery
//
// @param id Unique query identifier
// @param senderUserId Identifier of the user who sent the query
// @param inlineMessageId Identifier of the inline message, from which the query originated
// @param chatInstance An identifier uniquely corresponding to the chat a message was sent to
// @param payload Query payload
func NewUpdateNewInlineCallbackQuery(id JSONInt64, senderUserId int64, inlineMessageId string, chatInstance JSONInt64, payload CallbackQueryPayload) *UpdateNewInlineCallbackQuery {
	updateNewInlineCallbackQueryTemp := UpdateNewInlineCallbackQuery{
		tdCommon:        tdCommon{Type: "updateNewInlineCallbackQuery"},
		Id:              id,
		SenderUserId:    senderUserId,
		InlineMessageId: inlineMessageId,
		ChatInstance:    chatInstance,
		Payload:         payload,
	}

	return &updateNewInlineCallbackQueryTemp
}

// UnmarshalJSON unmarshal to json
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id              JSONInt64 `json:"id"`                // Unique query identifier
		SenderUserId    int64     `json:"sender_user_id"`    // Identifier of the user who sent the query
		InlineMessageId string    `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
		ChatInstance    JSONInt64 `json:"chat_instance"`     // An identifier uniquely corresponding to the chat a message was sent to

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNewInlineCallbackQuery.tdCommon = tempObj.tdCommon
	updateNewInlineCallbackQuery.Id = tempObj.Id
	updateNewInlineCallbackQuery.SenderUserId = tempObj.SenderUserId
	updateNewInlineCallbackQuery.InlineMessageId = tempObj.InlineMessageId
	updateNewInlineCallbackQuery.ChatInstance = tempObj.ChatInstance

	fieldPayload, _ := unmarshalCallbackQueryPayload(objMap["payload"])
	updateNewInlineCallbackQuery.Payload = fieldPayload

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewInlineCallbackQueryType
}

// UpdateNewShippingQuery A new incoming shipping query; for bots only. Only for invoices with flexible price
type UpdateNewShippingQuery struct {
	tdCommon
	Id              JSONInt64 `json:"id"`               // Unique query identifier
	SenderUserId    int64     `json:"sender_user_id"`   // Identifier of the user who sent the query
	InvoicePayload  string    `json:"invoice_payload"`  // Invoice payload
	ShippingAddress *Address  `json:"shipping_address"` // User shipping address
}

// MessageType return the string telegram-type of UpdateNewShippingQuery
func (updateNewShippingQuery *UpdateNewShippingQuery) MessageType() string {
	return "updateNewShippingQuery"
}

// NewUpdateNewShippingQuery creates a new UpdateNewShippingQuery
//
// @param id Unique query identifier
// @param senderUserId Identifier of the user who sent the query
// @param invoicePayload Invoice payload
// @param shippingAddress User shipping address
func NewUpdateNewShippingQuery(id JSONInt64, senderUserId int64, invoicePayload string, shippingAddress *Address) *UpdateNewShippingQuery {
	updateNewShippingQueryTemp := UpdateNewShippingQuery{
		tdCommon:        tdCommon{Type: "updateNewShippingQuery"},
		Id:              id,
		SenderUserId:    senderUserId,
		InvoicePayload:  invoicePayload,
		ShippingAddress: shippingAddress,
	}

	return &updateNewShippingQueryTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewShippingQuery *UpdateNewShippingQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewShippingQueryType
}

// UpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout
type UpdateNewPreCheckoutQuery struct {
	tdCommon
	Id               JSONInt64  `json:"id"`                 // Unique query identifier
	SenderUserId     int64      `json:"sender_user_id"`     // Identifier of the user who sent the query
	Currency         string     `json:"currency"`           // Currency for the product price
	TotalAmount      int64      `json:"total_amount"`       // Total price for the product, in the smallest units of the currency
	InvoicePayload   []byte     `json:"invoice_payload"`    // Invoice payload
	ShippingOptionId string     `json:"shipping_option_id"` // Identifier of a shipping option chosen by the user; may be empty if not applicable
	OrderInfo        *OrderInfo `json:"order_info"`         // Information about the order; may be null
}

// MessageType return the string telegram-type of UpdateNewPreCheckoutQuery
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) MessageType() string {
	return "updateNewPreCheckoutQuery"
}

// NewUpdateNewPreCheckoutQuery creates a new UpdateNewPreCheckoutQuery
//
// @param id Unique query identifier
// @param senderUserId Identifier of the user who sent the query
// @param currency Currency for the product price
// @param totalAmount Total price for the product, in the smallest units of the currency
// @param invoicePayload Invoice payload
// @param shippingOptionId Identifier of a shipping option chosen by the user; may be empty if not applicable
// @param orderInfo Information about the order; may be null
func NewUpdateNewPreCheckoutQuery(id JSONInt64, senderUserId int64, currency string, totalAmount int64, invoicePayload []byte, shippingOptionId string, orderInfo *OrderInfo) *UpdateNewPreCheckoutQuery {
	updateNewPreCheckoutQueryTemp := UpdateNewPreCheckoutQuery{
		tdCommon:         tdCommon{Type: "updateNewPreCheckoutQuery"},
		Id:               id,
		SenderUserId:     senderUserId,
		Currency:         currency,
		TotalAmount:      totalAmount,
		InvoicePayload:   invoicePayload,
		ShippingOptionId: shippingOptionId,
		OrderInfo:        orderInfo,
	}

	return &updateNewPreCheckoutQueryTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewPreCheckoutQueryType
}

// UpdateNewCustomEvent A new incoming event; for bots only
type UpdateNewCustomEvent struct {
	tdCommon
	Event string `json:"event"` // A JSON-serialized event
}

// MessageType return the string telegram-type of UpdateNewCustomEvent
func (updateNewCustomEvent *UpdateNewCustomEvent) MessageType() string {
	return "updateNewCustomEvent"
}

// NewUpdateNewCustomEvent creates a new UpdateNewCustomEvent
//
// @param event A JSON-serialized event
func NewUpdateNewCustomEvent(event string) *UpdateNewCustomEvent {
	updateNewCustomEventTemp := UpdateNewCustomEvent{
		tdCommon: tdCommon{Type: "updateNewCustomEvent"},
		Event:    event,
	}

	return &updateNewCustomEventTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewCustomEvent *UpdateNewCustomEvent) GetUpdateEnum() UpdateEnum {
	return UpdateNewCustomEventType
}

// UpdateNewCustomQuery A new incoming query; for bots only
type UpdateNewCustomQuery struct {
	tdCommon
	Id      JSONInt64 `json:"id"`      // The query identifier
	Data    string    `json:"data"`    // JSON-serialized query data
	Timeout int32     `json:"timeout"` // Query timeout
}

// MessageType return the string telegram-type of UpdateNewCustomQuery
func (updateNewCustomQuery *UpdateNewCustomQuery) MessageType() string {
	return "updateNewCustomQuery"
}

// NewUpdateNewCustomQuery creates a new UpdateNewCustomQuery
//
// @param id The query identifier
// @param data JSON-serialized query data
// @param timeout Query timeout
func NewUpdateNewCustomQuery(id JSONInt64, data string, timeout int32) *UpdateNewCustomQuery {
	updateNewCustomQueryTemp := UpdateNewCustomQuery{
		tdCommon: tdCommon{Type: "updateNewCustomQuery"},
		Id:       id,
		Data:     data,
		Timeout:  timeout,
	}

	return &updateNewCustomQueryTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewCustomQuery *UpdateNewCustomQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewCustomQueryType
}

// UpdatePoll A poll was updated; for bots only
type UpdatePoll struct {
	tdCommon
	Poll *Poll `json:"poll"` // New data about the poll
}

// MessageType return the string telegram-type of UpdatePoll
func (updatePoll *UpdatePoll) MessageType() string {
	return "updatePoll"
}

// NewUpdatePoll creates a new UpdatePoll
//
// @param poll New data about the poll
func NewUpdatePoll(poll *Poll) *UpdatePoll {
	updatePollTemp := UpdatePoll{
		tdCommon: tdCommon{Type: "updatePoll"},
		Poll:     poll,
	}

	return &updatePollTemp
}

// GetUpdateEnum return the enum type of this object
func (updatePoll *UpdatePoll) GetUpdateEnum() UpdateEnum {
	return UpdatePollType
}

// UpdatePollAnswer A user changed the answer to a poll; for bots only
type UpdatePollAnswer struct {
	tdCommon
	PollId    JSONInt64 `json:"poll_id"`    // Unique poll identifier
	UserId    int64     `json:"user_id"`    // The user, who changed the answer to the poll
	OptionIds []int32   `json:"option_ids"` // 0-based identifiers of answer options, chosen by the user
}

// MessageType return the string telegram-type of UpdatePollAnswer
func (updatePollAnswer *UpdatePollAnswer) MessageType() string {
	return "updatePollAnswer"
}

// NewUpdatePollAnswer creates a new UpdatePollAnswer
//
// @param pollId Unique poll identifier
// @param userId The user, who changed the answer to the poll
// @param optionIds 0-based identifiers of answer options, chosen by the user
func NewUpdatePollAnswer(pollId JSONInt64, userId int64, optionIds []int32) *UpdatePollAnswer {
	updatePollAnswerTemp := UpdatePollAnswer{
		tdCommon:  tdCommon{Type: "updatePollAnswer"},
		PollId:    pollId,
		UserId:    userId,
		OptionIds: optionIds,
	}

	return &updatePollAnswerTemp
}

// GetUpdateEnum return the enum type of this object
func (updatePollAnswer *UpdatePollAnswer) GetUpdateEnum() UpdateEnum {
	return UpdatePollAnswerType
}

// UpdateChatMember User rights changed in a chat; for bots only
type UpdateChatMember struct {
	tdCommon
	ChatId        int64           `json:"chat_id"`         // Chat identifier
	ActorUserId   int64           `json:"actor_user_id"`   // Identifier of the user, changing the rights
	Date          int32           `json:"date"`            // Point in time (Unix timestamp) when the user rights was changed
	InviteLink    *ChatInviteLink `json:"invite_link"`     // If user has joined the chat using an invite link, the invite link; may be null
	OldChatMember *ChatMember     `json:"old_chat_member"` // Previous chat member
	NewChatMember *ChatMember     `json:"new_chat_member"` // New chat member
}

// MessageType return the string telegram-type of UpdateChatMember
func (updateChatMember *UpdateChatMember) MessageType() string {
	return "updateChatMember"
}

// NewUpdateChatMember creates a new UpdateChatMember
//
// @param chatId Chat identifier
// @param actorUserId Identifier of the user, changing the rights
// @param date Point in time (Unix timestamp) when the user rights was changed
// @param inviteLink If user has joined the chat using an invite link, the invite link; may be null
// @param oldChatMember Previous chat member
// @param newChatMember New chat member
func NewUpdateChatMember(chatId int64, actorUserId int64, date int32, inviteLink *ChatInviteLink, oldChatMember *ChatMember, newChatMember *ChatMember) *UpdateChatMember {
	updateChatMemberTemp := UpdateChatMember{
		tdCommon:      tdCommon{Type: "updateChatMember"},
		ChatId:        chatId,
		ActorUserId:   actorUserId,
		Date:          date,
		InviteLink:    inviteLink,
		OldChatMember: oldChatMember,
		NewChatMember: newChatMember,
	}

	return &updateChatMemberTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatMember *UpdateChatMember) GetUpdateEnum() UpdateEnum {
	return UpdateChatMemberType
}

// UpdateNewChatJoinRequest A user sent a join request to a chat; for bots only
type UpdateNewChatJoinRequest struct {
	tdCommon
	ChatId     int64            `json:"chat_id"`     // Chat identifier
	Request    *ChatJoinRequest `json:"request"`     // Join request
	InviteLink *ChatInviteLink  `json:"invite_link"` // The invite link, which was used to send join request; may be null
}

// MessageType return the string telegram-type of UpdateNewChatJoinRequest
func (updateNewChatJoinRequest *UpdateNewChatJoinRequest) MessageType() string {
	return "updateNewChatJoinRequest"
}

// NewUpdateNewChatJoinRequest creates a new UpdateNewChatJoinRequest
//
// @param chatId Chat identifier
// @param request Join request
// @param inviteLink The invite link, which was used to send join request; may be null
func NewUpdateNewChatJoinRequest(chatId int64, request *ChatJoinRequest, inviteLink *ChatInviteLink) *UpdateNewChatJoinRequest {
	updateNewChatJoinRequestTemp := UpdateNewChatJoinRequest{
		tdCommon:   tdCommon{Type: "updateNewChatJoinRequest"},
		ChatId:     chatId,
		Request:    request,
		InviteLink: inviteLink,
	}

	return &updateNewChatJoinRequestTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewChatJoinRequest *UpdateNewChatJoinRequest) GetUpdateEnum() UpdateEnum {
	return UpdateNewChatJoinRequestType
}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestUseUpdate() (Update, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testUseUpdate",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch UpdateEnum(result.Data["@type"].(string)) {

	case UpdateAuthorizationStateType:
		var update UpdateAuthorizationState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewMessageType:
		var update UpdateNewMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendAcknowledgedType:
		var update UpdateMessageSendAcknowledged
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendSucceededType:
		var update UpdateMessageSendSucceeded
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendFailedType:
		var update UpdateMessageSendFailed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentType:
		var update UpdateMessageContent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageEditedType:
		var update UpdateMessageEdited
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageIsPinnedType:
		var update UpdateMessageIsPinned
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageInteractionInfoType:
		var update UpdateMessageInteractionInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentOpenedType:
		var update UpdateMessageContentOpened
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageMentionReadType:
		var update UpdateMessageMentionRead
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageLiveLocationViewedType:
		var update UpdateMessageLiveLocationViewed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatType:
		var update UpdateNewChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatTitleType:
		var update UpdateChatTitle
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPhotoType:
		var update UpdateChatPhoto
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPermissionsType:
		var update UpdateChatPermissions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatLastMessageType:
		var update UpdateChatLastMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPositionType:
		var update UpdateChatPosition
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsMarkedAsUnreadType:
		var update UpdateChatIsMarkedAsUnread
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsBlockedType:
		var update UpdateChatIsBlocked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatHasScheduledMessagesType:
		var update UpdateChatHasScheduledMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatVideoChatType:
		var update UpdateChatVideoChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDefaultDisableNotificationType:
		var update UpdateChatDefaultDisableNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadInboxType:
		var update UpdateChatReadInbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadOutboxType:
		var update UpdateChatReadOutbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatUnreadMentionCountType:
		var update UpdateChatUnreadMentionCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatNotificationSettingsType:
		var update UpdateChatNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateScopeNotificationSettingsType:
		var update UpdateScopeNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatMessageTtlSettingType:
		var update UpdateChatMessageTtlSetting
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatActionBarType:
		var update UpdateChatActionBar
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatThemeType:
		var update UpdateChatTheme
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPendingJoinRequestsType:
		var update UpdateChatPendingJoinRequests
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReplyMarkupType:
		var update UpdateChatReplyMarkup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDraftMessageType:
		var update UpdateChatDraftMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatFiltersType:
		var update UpdateChatFilters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatOnlineMemberCountType:
		var update UpdateChatOnlineMemberCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationType:
		var update UpdateNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationGroupType:
		var update UpdateNotificationGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateActiveNotificationsType:
		var update UpdateActiveNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateHavePendingNotificationsType:
		var update UpdateHavePendingNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDeleteMessagesType:
		var update UpdateDeleteMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserChatActionType:
		var update UpdateUserChatAction
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserStatusType:
		var update UpdateUserStatus
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserType:
		var update UpdateUser
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupType:
		var update UpdateBasicGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupType:
		var update UpdateSupergroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSecretChatType:
		var update UpdateSecretChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserFullInfoType:
		var update UpdateUserFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupFullInfoType:
		var update UpdateBasicGroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupFullInfoType:
		var update UpdateSupergroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateServiceNotificationType:
		var update UpdateServiceNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileType:
		var update UpdateFile
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStartType:
		var update UpdateFileGenerationStart
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStopType:
		var update UpdateFileGenerationStop
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateCallType:
		var update UpdateCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallType:
		var update UpdateGroupCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallParticipantType:
		var update UpdateGroupCallParticipant
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallSignalingDataType:
		var update UpdateNewCallSignalingData
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserPrivacySettingRulesType:
		var update UpdateUserPrivacySettingRules
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadMessageCountType:
		var update UpdateUnreadMessageCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadChatCountType:
		var update UpdateUnreadChatCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateOptionType:
		var update UpdateOption
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateStickerSetType:
		var update UpdateStickerSet
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateInstalledStickerSetsType:
		var update UpdateInstalledStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTrendingStickerSetsType:
		var update UpdateTrendingStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateRecentStickersType:
		var update UpdateRecentStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFavoriteStickersType:
		var update UpdateFavoriteStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSavedAnimationsType:
		var update UpdateSavedAnimations
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSelectedBackgroundType:
		var update UpdateSelectedBackground
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatThemesType:
		var update UpdateChatThemes
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateLanguagePackStringsType:
		var update UpdateLanguagePackStrings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateConnectionStateType:
		var update UpdateConnectionState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTermsOfServiceType:
		var update UpdateTermsOfService
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUsersNearbyType:
		var update UpdateUsersNearby
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDiceEmojisType:
		var update UpdateDiceEmojis
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateAnimatedEmojiMessageClickedType:
		var update UpdateAnimatedEmojiMessageClicked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateAnimationSearchParametersType:
		var update UpdateAnimationSearchParameters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSuggestedActionsType:
		var update UpdateSuggestedActions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineQueryType:
		var update UpdateNewInlineQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChosenInlineResultType:
		var update UpdateNewChosenInlineResult
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallbackQueryType:
		var update UpdateNewCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineCallbackQueryType:
		var update UpdateNewInlineCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewShippingQueryType:
		var update UpdateNewShippingQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewPreCheckoutQueryType:
		var update UpdateNewPreCheckoutQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomEventType:
		var update UpdateNewCustomEvent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomQueryType:
		var update UpdateNewCustomQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollType:
		var update UpdatePoll
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollAnswerType:
		var update UpdatePollAnswer
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatMemberType:
		var update UpdateChatMember
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatJoinRequestType:
		var update UpdateNewChatJoinRequest
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
