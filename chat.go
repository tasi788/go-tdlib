package tdlib

import (
	"encoding/json"
	"fmt"
)

// Chat A chat. (Can be a private chat, basic group, supergroup, or secret chat)
type Chat struct {
	tdCommon
	Id                         int64                     `json:"id"`                           // Chat unique identifier
	Type                       ChatType                  `json:"type"`                         // Type of the chat
	Title                      string                    `json:"title"`                        // Chat title
	Photo                      *ChatPhotoInfo            `json:"photo"`                        // Chat photo; may be null
	Permissions                *ChatPermissions          `json:"permissions"`                  // Actions that non-administrator chat members are allowed to take in the chat
	LastMessage                *Message                  `json:"last_message"`                 // Last message in the chat; may be null
	Positions                  []ChatPosition            `json:"positions"`                    // Positions of the chat in chat lists
	MessageSenderId            MessageSender             `json:"message_sender_id"`            // Identifier of a user or chat that is selected to send messages in the chat; may be null if the user can't change message sender
	HasProtectedContent        bool                      `json:"has_protected_content"`        // True, if chat content can't be saved locally, forwarded, or copied
	IsMarkedAsUnread           bool                      `json:"is_marked_as_unread"`          // True, if the chat is marked as unread
	IsBlocked                  bool                      `json:"is_blocked"`                   // True, if the chat is blocked by the current user and private messages from the chat can't be received
	HasScheduledMessages       bool                      `json:"has_scheduled_messages"`       // True, if the chat has scheduled messages
	CanBeDeletedOnlyForSelf    bool                      `json:"can_be_deleted_only_for_self"` // True, if the chat messages can be deleted only for the current user while other users will continue to see the messages
	CanBeDeletedForAllUsers    bool                      `json:"can_be_deleted_for_all_users"` // True, if the chat messages can be deleted for all users
	CanBeReported              bool                      `json:"can_be_reported"`              // True, if the chat can be reported to Telegram moderators through reportChat or reportChatPhoto
	DefaultDisableNotification bool                      `json:"default_disable_notification"` // Default value of the disable_notification parameter, used when a message is sent to the chat
	UnreadCount                int32                     `json:"unread_count"`                 // Number of unread messages in the chat
	LastReadInboxMessageId     int64                     `json:"last_read_inbox_message_id"`   // Identifier of the last read incoming message
	LastReadOutboxMessageId    int64                     `json:"last_read_outbox_message_id"`  // Identifier of the last read outgoing message
	UnreadMentionCount         int32                     `json:"unread_mention_count"`         // Number of unread messages with a mention/reply in the chat
	NotificationSettings       *ChatNotificationSettings `json:"notification_settings"`        // Notification settings for this chat
	MessageTtl                 int32                     `json:"message_ttl"`                  // Current message Time To Live setting (self-destruct timer) for the chat; 0 if not defined. TTL is counted from the time message or its content is viewed in secret chats and from the send date in other chats
	ThemeName                  string                    `json:"theme_name"`                   // If non-empty, name of a theme, set for the chat
	ActionBar                  ChatActionBar             `json:"action_bar"`                   // Information about actions which must be possible to do through the chat action bar; may be null
	VideoChat                  *VideoChat                `json:"video_chat"`                   // Information about video chat of the chat
	PendingJoinRequests        *ChatJoinRequestsInfo     `json:"pending_join_requests"`        // Information about pending join requests; may be null
	ReplyMarkupMessageId       int64                     `json:"reply_markup_message_id"`      // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
	DraftMessage               *DraftMessage             `json:"draft_message"`                // A draft of a message in the chat; may be null
	ClientData                 string                    `json:"client_data"`                  // Application-specific data associated with the chat. (For example, the chat scroll position or local chat notification settings can be stored here.) Persistent if the message database is used
}

// MessageType return the string telegram-type of Chat
func (chat *Chat) MessageType() string {
	return "chat"
}

// NewChat creates a new Chat
//
// @param id Chat unique identifier
// @param typeParam Type of the chat
// @param title Chat title
// @param photo Chat photo; may be null
// @param permissions Actions that non-administrator chat members are allowed to take in the chat
// @param lastMessage Last message in the chat; may be null
// @param positions Positions of the chat in chat lists
// @param messageSenderId Identifier of a user or chat that is selected to send messages in the chat; may be null if the user can't change message sender
// @param hasProtectedContent True, if chat content can't be saved locally, forwarded, or copied
// @param isMarkedAsUnread True, if the chat is marked as unread
// @param isBlocked True, if the chat is blocked by the current user and private messages from the chat can't be received
// @param hasScheduledMessages True, if the chat has scheduled messages
// @param canBeDeletedOnlyForSelf True, if the chat messages can be deleted only for the current user while other users will continue to see the messages
// @param canBeDeletedForAllUsers True, if the chat messages can be deleted for all users
// @param canBeReported True, if the chat can be reported to Telegram moderators through reportChat or reportChatPhoto
// @param defaultDisableNotification Default value of the disable_notification parameter, used when a message is sent to the chat
// @param unreadCount Number of unread messages in the chat
// @param lastReadInboxMessageId Identifier of the last read incoming message
// @param lastReadOutboxMessageId Identifier of the last read outgoing message
// @param unreadMentionCount Number of unread messages with a mention/reply in the chat
// @param notificationSettings Notification settings for this chat
// @param messageTtl Current message Time To Live setting (self-destruct timer) for the chat; 0 if not defined. TTL is counted from the time message or its content is viewed in secret chats and from the send date in other chats
// @param themeName If non-empty, name of a theme, set for the chat
// @param actionBar Information about actions which must be possible to do through the chat action bar; may be null
// @param videoChat Information about video chat of the chat
// @param pendingJoinRequests Information about pending join requests; may be null
// @param replyMarkupMessageId Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
// @param draftMessage A draft of a message in the chat; may be null
// @param clientData Application-specific data associated with the chat. (For example, the chat scroll position or local chat notification settings can be stored here.) Persistent if the message database is used
func NewChat(id int64, typeParam ChatType, title string, photo *ChatPhotoInfo, permissions *ChatPermissions, lastMessage *Message, positions []ChatPosition, messageSenderId MessageSender, hasProtectedContent bool, isMarkedAsUnread bool, isBlocked bool, hasScheduledMessages bool, canBeDeletedOnlyForSelf bool, canBeDeletedForAllUsers bool, canBeReported bool, defaultDisableNotification bool, unreadCount int32, lastReadInboxMessageId int64, lastReadOutboxMessageId int64, unreadMentionCount int32, notificationSettings *ChatNotificationSettings, messageTtl int32, themeName string, actionBar ChatActionBar, videoChat *VideoChat, pendingJoinRequests *ChatJoinRequestsInfo, replyMarkupMessageId int64, draftMessage *DraftMessage, clientData string) *Chat {
	chatTemp := Chat{
		tdCommon:                   tdCommon{Type: "chat"},
		Id:                         id,
		Type:                       typeParam,
		Title:                      title,
		Photo:                      photo,
		Permissions:                permissions,
		LastMessage:                lastMessage,
		Positions:                  positions,
		MessageSenderId:            messageSenderId,
		HasProtectedContent:        hasProtectedContent,
		IsMarkedAsUnread:           isMarkedAsUnread,
		IsBlocked:                  isBlocked,
		HasScheduledMessages:       hasScheduledMessages,
		CanBeDeletedOnlyForSelf:    canBeDeletedOnlyForSelf,
		CanBeDeletedForAllUsers:    canBeDeletedForAllUsers,
		CanBeReported:              canBeReported,
		DefaultDisableNotification: defaultDisableNotification,
		UnreadCount:                unreadCount,
		LastReadInboxMessageId:     lastReadInboxMessageId,
		LastReadOutboxMessageId:    lastReadOutboxMessageId,
		UnreadMentionCount:         unreadMentionCount,
		NotificationSettings:       notificationSettings,
		MessageTtl:                 messageTtl,
		ThemeName:                  themeName,
		ActionBar:                  actionBar,
		VideoChat:                  videoChat,
		PendingJoinRequests:        pendingJoinRequests,
		ReplyMarkupMessageId:       replyMarkupMessageId,
		DraftMessage:               draftMessage,
		ClientData:                 clientData,
	}

	return &chatTemp
}

// UnmarshalJSON unmarshal to json
func (chat *Chat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Id                         int64                     `json:"id"`                           // Chat unique identifier
		Title                      string                    `json:"title"`                        // Chat title
		Photo                      *ChatPhotoInfo            `json:"photo"`                        // Chat photo; may be null
		Permissions                *ChatPermissions          `json:"permissions"`                  // Actions that non-administrator chat members are allowed to take in the chat
		LastMessage                *Message                  `json:"last_message"`                 // Last message in the chat; may be null
		Positions                  []ChatPosition            `json:"positions"`                    // Positions of the chat in chat lists
		HasProtectedContent        bool                      `json:"has_protected_content"`        // True, if chat content can't be saved locally, forwarded, or copied
		IsMarkedAsUnread           bool                      `json:"is_marked_as_unread"`          // True, if the chat is marked as unread
		IsBlocked                  bool                      `json:"is_blocked"`                   // True, if the chat is blocked by the current user and private messages from the chat can't be received
		HasScheduledMessages       bool                      `json:"has_scheduled_messages"`       // True, if the chat has scheduled messages
		CanBeDeletedOnlyForSelf    bool                      `json:"can_be_deleted_only_for_self"` // True, if the chat messages can be deleted only for the current user while other users will continue to see the messages
		CanBeDeletedForAllUsers    bool                      `json:"can_be_deleted_for_all_users"` // True, if the chat messages can be deleted for all users
		CanBeReported              bool                      `json:"can_be_reported"`              // True, if the chat can be reported to Telegram moderators through reportChat or reportChatPhoto
		DefaultDisableNotification bool                      `json:"default_disable_notification"` // Default value of the disable_notification parameter, used when a message is sent to the chat
		UnreadCount                int32                     `json:"unread_count"`                 // Number of unread messages in the chat
		LastReadInboxMessageId     int64                     `json:"last_read_inbox_message_id"`   // Identifier of the last read incoming message
		LastReadOutboxMessageId    int64                     `json:"last_read_outbox_message_id"`  // Identifier of the last read outgoing message
		UnreadMentionCount         int32                     `json:"unread_mention_count"`         // Number of unread messages with a mention/reply in the chat
		NotificationSettings       *ChatNotificationSettings `json:"notification_settings"`        // Notification settings for this chat
		MessageTtl                 int32                     `json:"message_ttl"`                  // Current message Time To Live setting (self-destruct timer) for the chat; 0 if not defined. TTL is counted from the time message or its content is viewed in secret chats and from the send date in other chats
		ThemeName                  string                    `json:"theme_name"`                   // If non-empty, name of a theme, set for the chat
		VideoChat                  *VideoChat                `json:"video_chat"`                   // Information about video chat of the chat
		PendingJoinRequests        *ChatJoinRequestsInfo     `json:"pending_join_requests"`        // Information about pending join requests; may be null
		ReplyMarkupMessageId       int64                     `json:"reply_markup_message_id"`      // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
		DraftMessage               *DraftMessage             `json:"draft_message"`                // A draft of a message in the chat; may be null
		ClientData                 string                    `json:"client_data"`                  // Application-specific data associated with the chat. (For example, the chat scroll position or local chat notification settings can be stored here.) Persistent if the message database is used
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chat.tdCommon = tempObj.tdCommon
	chat.Id = tempObj.Id
	chat.Title = tempObj.Title
	chat.Photo = tempObj.Photo
	chat.Permissions = tempObj.Permissions
	chat.LastMessage = tempObj.LastMessage
	chat.Positions = tempObj.Positions
	chat.HasProtectedContent = tempObj.HasProtectedContent
	chat.IsMarkedAsUnread = tempObj.IsMarkedAsUnread
	chat.IsBlocked = tempObj.IsBlocked
	chat.HasScheduledMessages = tempObj.HasScheduledMessages
	chat.CanBeDeletedOnlyForSelf = tempObj.CanBeDeletedOnlyForSelf
	chat.CanBeDeletedForAllUsers = tempObj.CanBeDeletedForAllUsers
	chat.CanBeReported = tempObj.CanBeReported
	chat.DefaultDisableNotification = tempObj.DefaultDisableNotification
	chat.UnreadCount = tempObj.UnreadCount
	chat.LastReadInboxMessageId = tempObj.LastReadInboxMessageId
	chat.LastReadOutboxMessageId = tempObj.LastReadOutboxMessageId
	chat.UnreadMentionCount = tempObj.UnreadMentionCount
	chat.NotificationSettings = tempObj.NotificationSettings
	chat.MessageTtl = tempObj.MessageTtl
	chat.ThemeName = tempObj.ThemeName
	chat.VideoChat = tempObj.VideoChat
	chat.PendingJoinRequests = tempObj.PendingJoinRequests
	chat.ReplyMarkupMessageId = tempObj.ReplyMarkupMessageId
	chat.DraftMessage = tempObj.DraftMessage
	chat.ClientData = tempObj.ClientData

	fieldType, _ := unmarshalChatType(objMap["type"])
	chat.Type = fieldType

	fieldMessageSenderId, _ := unmarshalMessageSender(objMap["message_sender_id"])
	chat.MessageSenderId = fieldMessageSenderId

	fieldActionBar, _ := unmarshalChatActionBar(objMap["action_bar"])
	chat.ActionBar = fieldActionBar

	return nil
}

// GetChat Returns information about a chat by its identifier, this is an offline request if the current user is not a bot
// @param chatId Chat identifier
func (client *Client) GetChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err
}

// SearchPublicChat Searches a public chat by its username. Currently, only private chats, supergroups and channels can be public. Returns the chat if found; otherwise an error is returned
// @param username Username to be resolved
func (client *Client) SearchPublicChat(username string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchPublicChat",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreatePrivateChat Returns an existing chat corresponding to a given user
// @param userId User identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreatePrivateChat(userId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createPrivateChat",
		"user_id": userId,
		"force":   force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group
// @param basicGroupId Basic group identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateBasicGroupChat(basicGroupId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createBasicGroupChat",
		"basic_group_id": basicGroupId,
		"force":          force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel
// @param supergroupId Supergroup or channel identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateSupergroupChat(supergroupId int64, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "createSupergroupChat",
		"supergroup_id": supergroupId,
		"force":         force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat
// @param secretChatId Secret chat identifier
func (client *Client) CreateSecretChat(secretChatId int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err
}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns the newly created chat
// @param userIds Identifiers of users to be added to the basic group
// @param title Title of the new basic group; 1-128 characters
func (client *Client) CreateNewBasicGroupChat(userIds []int64, title string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewBasicGroupChat",
		"user_ids": userIds,
		"title":    title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
// @param title Title of the new chat; 1-128 characters
// @param isChannel True, if a channel chat needs to be created
// @param description Chat description; 0-255 characters
// @param location Chat location if a location-based supergroup is being created; pass null to create an ordinary supergroup chat
// @param forImport True, if the supergroup is created for importing messages using importMessage
func (client *Client) CreateNewSupergroupChat(title string, isChannel bool, description string, location *ChatLocation, forImport bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "createNewSupergroupChat",
		"title":       title,
		"is_channel":  isChannel,
		"description": description,
		"location":    location,
		"for_import":  forImport,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat
// @param userId Identifier of the target user
func (client *Client) CreateNewSecretChat(userId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createNewSecretChat",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires creator privileges. Deactivates the original basic group
// @param chatId Identifier of the chat to upgrade
func (client *Client) UpgradeBasicGroupChatToSupergroupChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err
}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible
// @param inviteLink Invite link to use
func (client *Client) JoinChatByInviteLink(inviteLink string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "joinChatByInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err
}
