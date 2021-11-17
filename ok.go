package tdlib

import (
	"encoding/json"
	"fmt"
)

// Ok An object of this type is returned on a successful function call for certain functions
type Ok struct {
	tdCommon
}

// MessageType return the string telegram-type of Ok
func (ok *Ok) MessageType() string {
	return "ok"
}

// NewOk creates a new Ok
//
func NewOk() *Ok {
	okTemp := Ok{
		tdCommon: tdCommon{Type: "ok"},
	}

	return &okTemp
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
// @param parameters Parameters for TDLib initialization
func (client *Client) SetTdlibParameters(parameters *TdlibParameters) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setTdlibParameters",
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckDatabaseEncryptionKey Checks the database encryption key for correctness. Works only when the current authorization state is authorizationStateWaitEncryptionKey
// @param encryptionKey Encryption key to check or set up
func (client *Client) CheckDatabaseEncryptionKey(encryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "checkDatabaseEncryptionKey",
		"encryption_key": encryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber, or if there is no pending authentication query and the current authorization state is authorizationStateWaitCode, authorizationStateWaitRegistration, or authorizationStateWaitPassword
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number; pass null to use default settings
func (client *Client) SetAuthenticationPhoneNumber(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setAuthenticationPhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ResendAuthenticationCode Re-sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode, the next_code_type of the result is not null and the server-specified timeout has passed
func (client *Client) ResendAuthenticationCode() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendAuthenticationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode
// @param code The verification code received via SMS, Telegram message, phone call, or flash call
func (client *Client) CheckAuthenticationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RequestQrCodeAuthentication Requests QR code authentication by scanning a QR code on another logged in device. Works only when the current authorization state is authorizationStateWaitPhoneNumber, or if there is no pending authentication query and the current authorization state is authorizationStateWaitCode, authorizationStateWaitRegistration, or authorizationStateWaitPassword
// @param otherUserIds List of user identifiers of other users currently using the application
func (client *Client) RequestQrCodeAuthentication(otherUserIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "requestQrCodeAuthentication",
		"other_user_ids": otherUserIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RegisterUser Finishes user registration. Works only when the current authorization state is authorizationStateWaitRegistration
// @param firstName The first name of the user; 1-64 characters
// @param lastName The last name of the user; 0-64 characters
func (client *Client) RegisterUser(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "registerUser",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckAuthenticationPassword Checks the authentication password for correctness. Works only when the current authorization state is authorizationStateWaitPassword
// @param password The password to check
func (client *Client) CheckAuthenticationPassword(password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RequestAuthenticationPasswordRecovery Requests to send a password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (client *Client) RequestAuthenticationPasswordRecovery() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestAuthenticationPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckAuthenticationPasswordRecoveryCode Checks whether a password recovery code sent to an email address is valid. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
func (client *Client) CheckAuthenticationPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "checkAuthenticationPasswordRecoveryCode",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RecoverAuthenticationPassword Recovers the password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
func (client *Client) RecoverAuthenticationPassword(recoveryCode string, newPassword string, newHint string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverAuthenticationPassword",
		"recovery_code": recoveryCode,
		"new_password":  newPassword,
		"new_hint":      newHint,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in
// @param token The bot token
func (client *Client) CheckAuthenticationBotToken(token string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationBotToken",
		"token": token,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) LogOut() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "logOut",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (client *Client) Close() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "close",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed. After the destruction completes updateAuthorizationState with authorizationStateClosed will be sent. Can be called before authorization
func (client *Client) Destroy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "destroy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain
// @param newEncryptionKey New encryption key
func (client *Client) SetDatabaseEncryptionKey(newEncryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setDatabaseEncryptionKey",
		"new_encryption_key": newEncryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid
// @param recoveryCode Recovery code to check
func (client *Client) CheckPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "checkPasswordRecoveryCode",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CancelPasswordReset Cancels reset of 2-step verification password. The method can be called if passwordState.pending_reset_date > 0
func (client *Client) CancelPasswordReset() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "cancelPasswordReset",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// LoadChats Loads more chats from a chat list. The loaded chats and their positions in the chat list will be sent through updates. Chats are sorted by the pair (chat.position.order, chat.id) in descending order. Returns a 404 error if all chats has been loaded
// @param chatList The chat list in which to load chats; pass null to load chats from the main chat list
// @param limit The maximum number of chats to be loaded. For optimal performance, the number of loaded chats is chosen by TDLib and can be smaller than the specified limit, even if the end of the list is not reached
func (client *Client) LoadChats(chatList ChatList, limit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "loadChats",
		"chat_list": chatList,
		"limit":     limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of frequently used chats
// @param chatId Chat identifier
func (client *Client) RemoveTopChat(category TopChatCategory, chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeTopChat",
		"category": category,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first
// @param chatId Identifier of the chat to add
func (client *Client) AddRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats
// @param chatId Identifier of the chat to be removed
func (client *Client) RemoveRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (client *Client) ClearRecentlyFoundChats() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearRecentlyFoundChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckCreatedPublicChatsLimit Checks whether the maximum number of owned public chats has been reached. Returns corresponding error if the limit was reached
// @param typeParam Type of the public chats, for which to check the limit
func (client *Client) CheckCreatedPublicChatsLimit(typeParam PublicChatType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkCreatedPublicChatsLimit",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteChatHistory Deletes all messages in the chat. Use chat.can_be_deleted_only_for_self and chat.can_be_deleted_for_all_users fields to find whether and how the method can be applied to the chat
// @param chatId Chat identifier
// @param removeFromChatList Pass true if the chat needs to be removed from the chat list
// @param revoke Pass true to try to delete chat history for all users
func (client *Client) DeleteChatHistory(chatId int64, removeFromChatList bool, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "deleteChatHistory",
		"chat_id":               chatId,
		"remove_from_chat_list": removeFromChatList,
		"revoke":                revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// DeleteChat Deletes a chat along with all messages in the corresponding chat for all chat members; requires owner privileges. For group chats this will release the username and remove all members. Chats with more than 1000 members can't be deleted using this method
// @param chatId Chat identifier
func (client *Client) DeleteChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteAllCallMessages Deletes all call messages
// @param revoke Pass true to delete the messages for all users
func (client *Client) DeleteAllCallMessages(revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAllCallMessages",
		"revoke": revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// ViewSponsoredMessage Informs TDLib that a sponsored message was viewed by the user
// @param chatId Identifier of the chat with the sponsored message
// @param sponsoredMessageId The identifier of the sponsored message being viewed
func (client *Client) ViewSponsoredMessage(chatId int64, sponsoredMessageId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "viewSponsoredMessage",
		"chat_id":              chatId,
		"sponsored_message_id": sponsoredMessageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveNotification Removes an active notification from notification list. Needs to be called only if the notification is removed by the current user
// @param notificationGroupId Identifier of notification group to which the notification belongs
// @param notificationId Identifier of removed notification
func (client *Client) RemoveNotification(notificationGroupId int32, notificationId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotification",
		"notification_group_id": notificationGroupId,
		"notification_id":       notificationId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveNotificationGroup Removes a group of active notifications. Needs to be called only if the notification group is removed by the current user
// @param notificationGroupId Notification group identifier
// @param maxNotificationId The maximum identifier of removed notifications
func (client *Client) RemoveNotificationGroup(notificationGroupId int32, maxNotificationId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotificationGroup",
		"notification_group_id": notificationGroupId,
		"max_notification_id":   maxNotificationId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendChatScreenshotTakenNotification Sends a notification about a screenshot taken in a chat. Supported only in private and secret chats
// @param chatId Chat identifier
func (client *Client) SendChatScreenshotTakenNotification(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatScreenshotTakenNotification",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteMessages Deletes messages
// @param chatId Chat identifier
// @param messageIds Identifiers of the messages to be deleted
// @param revoke Pass true to try to delete messages for all chat members. Always true for supergroups, channels and secret chats
func (client *Client) DeleteMessages(chatId int64, messageIds []int64, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"revoke":      revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// DeleteChatMessagesFromUser Deletes all messages sent by the specified user to a chat. Supported only for supergroups; requires can_delete_messages administrator privileges
// @param chatId Chat identifier
// @param userId User identifier
func (client *Client) DeleteChatMessagesFromUser(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChatMessagesFromUser",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteChatMessagesByDate Deletes all messages between the specified dates in a chat. Supported only for private chats and basic groups. Messages sent in the last 30 seconds will not be deleted
// @param chatId Chat identifier
// @param minDate The minimum date of the messages to delete
// @param maxDate The maximum date of the messages to delete
// @param revoke Pass true to try to delete chat messages for all users; private chats only
func (client *Client) DeleteChatMessagesByDate(chatId int64, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "deleteChatMessagesByDate",
		"chat_id":  chatId,
		"min_date": minDate,
		"max_date": maxDate,
		"revoke":   revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param inputMessageContent New text content of the message. Must be of type inputMessageText
func (client *Client) EditInlineMessageText(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageText",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param location New location content of the message; pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditInlineMessageLiveLocation(inlineMessageId string, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editInlineMessageLiveLocation",
		"inline_message_id":      inlineMessageId,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditInlineMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: inputMessageAnimation, inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
func (client *Client) EditInlineMessageMedia(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageMedia",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
// @param caption New message content caption; pass null to remove caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditInlineMessageCaption(inlineMessageId string, replyMarkup ReplyMarkup, caption *FormattedText) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageCaption",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
		"caption":           caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; pass null if none
func (client *Client) EditInlineMessageReplyMarkup(inlineMessageId string, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageReplyMarkup",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditMessageSchedulingState Edits the time when a scheduled message will be sent. Scheduling state of all messages in the same album or forwarded together with the message will be also changed
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param schedulingState The new message scheduling state; pass null to send the message immediately
func (client *Client) EditMessageSchedulingState(chatId int64, messageId int64, schedulingState MessageSchedulingState) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "editMessageSchedulingState",
		"chat_id":          chatId,
		"message_id":       messageId,
		"scheduling_state": schedulingState,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetPollAnswer Changes the user answer to a poll. A poll in quiz mode can be answered only once
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param optionIds 0-based identifiers of answer options, chosen by the user. User can choose more than 1 answer option only is the poll allows multiple answers
func (client *Client) SetPollAnswer(chatId int64, messageId int64, optionIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setPollAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"option_ids": optionIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// StopPoll Stops a poll. A poll in a message can be stopped when the message has can_be_edited flag set
// @param chatId Identifier of the chat to which the poll belongs
// @param messageId Identifier of the message containing the poll
// @param replyMarkup The new message reply markup; pass null if none; for bots only
func (client *Client) StopPoll(chatId int64, messageId int64, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "stopPoll",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// HideSuggestedAction Hides a suggested action
// @param action Suggested action to hide
func (client *Client) HideSuggestedAction(action SuggestedAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "hideSuggestedAction",
		"action": action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AnswerInlineQuery Sets the result of an inline query; for bots only
// @param inlineQueryId Identifier of the inline query
// @param isPersonal True, if the result of the query can be cached for the specified user
// @param results The results of the query
// @param cacheTime Allowed time to cache the results of the query, in seconds
// @param nextOffset Offset for the next inline query; pass an empty string if there are no more results
// @param switchPmText If non-empty, this text must be shown on the button that opens a private chat with the bot and sends a start message to the bot with the parameter switch_pm_parameter
// @param switchPmParameter The parameter for the bot start message
func (client *Client) AnswerInlineQuery(inlineQueryId JSONInt64, isPersonal bool, results []InputInlineQueryResult, cacheTime int32, nextOffset string, switchPmText string, switchPmParameter string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "answerInlineQuery",
		"inline_query_id":     inlineQueryId,
		"is_personal":         isPersonal,
		"results":             results,
		"cache_time":          cacheTime,
		"next_offset":         nextOffset,
		"switch_pm_text":      switchPmText,
		"switch_pm_parameter": switchPmParameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
// @param callbackQueryId Identifier of the callback query
// @param text Text of the answer
// @param showAlert If true, an alert must be shown to the user instead of a toast notification
// @param url URL to be opened
// @param cacheTime Time during which the result of the query can be cached, in seconds
func (client *Client) AnswerCallbackQuery(callbackQueryId JSONInt64, text string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerCallbackQuery",
		"callback_query_id": callbackQueryId,
		"text":              text,
		"show_alert":        showAlert,
		"url":               url,
		"cache_time":        cacheTime,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AnswerShippingQuery Sets the result of a shipping query; for bots only
// @param shippingQueryId Identifier of the shipping query
// @param shippingOptions Available shipping options
// @param errorMessage An error message, empty on success
func (client *Client) AnswerShippingQuery(shippingQueryId JSONInt64, shippingOptions []ShippingOption, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerShippingQuery",
		"shipping_query_id": shippingQueryId,
		"shipping_options":  shippingOptions,
		"error_message":     errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only
// @param preCheckoutQueryId Identifier of the pre-checkout query
// @param errorMessage An error message, empty on success
func (client *Client) AnswerPreCheckoutQuery(preCheckoutQueryId JSONInt64, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "answerPreCheckoutQuery",
		"pre_checkout_query_id": preCheckoutQueryId,
		"error_message":         errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
// @param inlineMessageId Inline message identifier
// @param editMessage True, if the message needs to be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetInlineGameScore(inlineMessageId string, editMessage bool, userId int64, score int32, force bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setInlineGameScore",
		"inline_message_id": inlineMessageId,
		"edit_message":      editMessage,
		"user_id":           userId,
		"score":             score,
		"force":             force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a ForceReply reply markup has been used. UpdateChatReplyMarkup will be sent if the reply markup is changed
// @param chatId Chat identifier
// @param messageId The message identifier of the used keyboard
func (client *Client) DeleteChatReplyMarkup(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "deleteChatReplyMarkup",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendChatAction Sends a notification about user activity in a chat
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the action was performed
// @param action The action description; pass null to cancel the currently active action
func (client *Client) SendChatAction(chatId int64, messageThreadId int64, action ChatAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendChatAction",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"action":            action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// OpenChat Informs TDLib that the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats)
// @param chatId Chat identifier
func (client *Client) OpenChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "openChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CloseChat Informs TDLib that the chat is closed by the user. Many useful activities depend on the chat being opened or closed
// @param chatId Chat identifier
func (client *Client) CloseChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "closeChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ViewMessages Informs TDLib that messages are being viewed by the user. Many useful activities depend on whether the messages are currently being viewed or not (e.g., marking messages as read, incrementing a view counter, updating a view counter, removing deleted messages in supergroups and channels)
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the messages are being viewed
// @param messageIds The identifiers of the messages being viewed
// @param forceRead True, if messages in closed chats must be marked as read by the request
func (client *Client) ViewMessages(chatId int64, messageThreadId int64, messageIds []int64, forceRead bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "viewMessages",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"message_ids":       messageIds,
		"force_read":        forceRead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// OpenMessageContent Informs TDLib that the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message). An updateMessageContentOpened update will be generated if something has changed
// @param chatId Chat identifier of the message
// @param messageId Identifier of the message with the opened content
func (client *Client) OpenMessageContent(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "openMessageContent",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReadAllChatMentions Marks all mentions in a chat as read
// @param chatId Chat identifier
func (client *Client) ReadAllChatMentions(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readAllChatMentions",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddChatToList Adds a chat to a chat list. A chat can't be simultaneously in Main and Archive chat lists, so it is automatically removed from another one if needed
// @param chatId Chat identifier
// @param chatList The chat list. Use getChatListsToAddChat to get suitable chat lists
func (client *Client) AddChatToList(chatId int64, chatList ChatList) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addChatToList",
		"chat_id":   chatId,
		"chat_list": chatList,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteChatFilter Deletes existing chat filter
// @param chatFilterId Chat filter identifier
func (client *Client) DeleteChatFilter(chatFilterId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "deleteChatFilter",
		"chat_filter_id": chatFilterId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReorderChatFilters Changes the order of chat filters
// @param chatFilterIds Identifiers of chat filters in the new correct order
func (client *Client) ReorderChatFilters(chatFilterIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderChatFilters",
		"chat_filter_ids": chatFilterIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatId Chat identifier
// @param title New title of the chat; 1-128 characters
func (client *Client) SetChatTitle(chatId int64, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatTitle",
		"chat_id": chatId,
		"title":   title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatId Chat identifier
// @param photo New chat photo; pass null to delete the chat photo
func (client *Client) SetChatPhoto(chatId int64, photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatPhoto",
		"chat_id": chatId,
		"photo":   photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatMessageTtlSetting Changes the message TTL setting (sets a new self-destruct timer) in a chat. Requires can_delete_messages administrator right in basic groups, supergroups and channels Message TTL setting of a chat with the current user (Saved Messages) and the chat 777000 (Telegram) can't be changed
// @param chatId Chat identifier
// @param ttl New TTL value, in seconds; must be one of 0, 86400, 7 * 86400, or 31 * 86400 unless the chat is secret
func (client *Client) SetChatMessageTtlSetting(chatId int64, ttl int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatMessageTtlSetting",
		"chat_id": chatId,
		"ttl":     ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatPermissions Changes the chat members permissions. Supported only for basic groups and supergroups. Requires can_restrict_members administrator right
// @param chatId Chat identifier
// @param permissions New non-administrator members permissions in the chat
func (client *Client) SetChatPermissions(chatId int64, permissions *ChatPermissions) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatPermissions",
		"chat_id":     chatId,
		"permissions": permissions,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatTheme Changes the chat theme. Supported only in private and secret chats
// @param chatId Chat identifier
// @param themeName Name of the new chat theme; pass an empty string to return the default theme
func (client *Client) SetChatTheme(chatId int64, themeName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setChatTheme",
		"chat_id":    chatId,
		"theme_name": themeName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatDraftMessage Changes the draft message in a chat
// @param chatId Chat identifier
// @param messageThreadId If not 0, a message thread identifier in which the draft was changed
// @param draftMessage New draft message; pass null to remove the draft
func (client *Client) SetChatDraftMessage(chatId int64, messageThreadId int64, draftMessage *DraftMessage) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setChatDraftMessage",
		"chat_id":           chatId,
		"message_thread_id": messageThreadId,
		"draft_message":     draftMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatNotificationSettings Changes the notification settings of a chat. Notification settings of a chat with the current user (Saved Messages) can't be changed
// @param chatId Chat identifier
// @param notificationSettings New notification settings for the chat. If the chat is muted for more than 1 week, it is considered to be muted forever
func (client *Client) SetChatNotificationSettings(chatId int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setChatNotificationSettings",
		"chat_id":               chatId,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat
// @param chatId Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func (client *Client) ToggleChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleChatIsMarkedAsUnread",
		"chat_id":             chatId,
		"is_marked_as_unread": isMarkedAsUnread,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat
// @param chatId Chat identifier
// @param defaultDisableNotification New value of default_disable_notification
func (client *Client) ToggleChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                        "toggleChatDefaultDisableNotification",
		"chat_id":                      chatId,
		"default_disable_notification": defaultDisableNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatClientData Changes application-specific data associated with a chat
// @param chatId Chat identifier
// @param clientData New value of client_data
func (client *Client) SetChatClientData(chatId int64, clientData string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatClientData",
		"chat_id":     chatId,
		"client_data": clientData,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatDescription Changes information about a chat. Available for basic groups, supergroups, and channels. Requires can_change_info administrator right
// @param chatId Identifier of the chat
// @param description New chat description; 0-255 characters
func (client *Client) SetChatDescription(chatId int64, description string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatDescription",
		"chat_id":     chatId,
		"description": description,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatDiscussionGroup Changes the discussion group of a channel chat; requires can_change_info administrator right in the channel if it is specified
// @param chatId Identifier of the channel chat. Pass 0 to remove a link from the supergroup passed in the second argument to a linked channel chat (requires can_pin_messages rights in the supergroup)
// @param discussionChatId Identifier of a new channel's discussion group. Use 0 to remove the discussion group. Use the method getSuitableDiscussionChats to find all suitable groups. Basic group chats must be first upgraded to supergroup chats. If new chat members don't have access to old messages in the supergroup, then toggleSupergroupIsAllHistoryAvailable must be used first to change that
func (client *Client) SetChatDiscussionGroup(chatId int64, discussionChatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setChatDiscussionGroup",
		"chat_id":            chatId,
		"discussion_chat_id": discussionChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatLocation Changes the location of a chat. Available only for some location-based supergroups, use supergroupFullInfo.can_set_location to check whether the method is allowed to use
// @param chatId Chat identifier
// @param location New location for the chat; must be valid and not null
func (client *Client) SetChatLocation(chatId int64, location *ChatLocation) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setChatLocation",
		"chat_id":  chatId,
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatSlowModeDelay Changes the slow mode delay of a chat. Available only for supergroups; requires can_restrict_members rights
// @param chatId Chat identifier
// @param slowModeDelay New slow mode delay for the chat, in seconds; must be one of 0, 10, 30, 60, 300, 900, 3600
func (client *Client) SetChatSlowModeDelay(chatId int64, slowModeDelay int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "setChatSlowModeDelay",
		"chat_id":         chatId,
		"slow_mode_delay": slowModeDelay,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// PinChatMessage Pins a message in a chat; requires can_pin_messages rights or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
// @param messageId Identifier of the new pinned message
// @param disableNotification True, if there must be no notification about the pinned message. Notifications are always disabled in channels and private chats
// @param onlyForSelf True, if the message needs to be pinned for one side only; private chats only
func (client *Client) PinChatMessage(chatId int64, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "pinChatMessage",
		"chat_id":              chatId,
		"message_id":           messageId,
		"disable_notification": disableNotification,
		"only_for_self":        onlyForSelf,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// UnpinChatMessage Removes a pinned message from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
// @param messageId Identifier of the removed pinned message
func (client *Client) UnpinChatMessage(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "unpinChatMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// UnpinAllChatMessages Removes all pinned messages from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatId Identifier of the chat
func (client *Client) UnpinAllChatMessages(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "unpinAllChatMessages",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// JoinChat Adds the current user as a new member to a chat. Private and secret chats can't be joined using this method
// @param chatId Chat identifier
func (client *Client) JoinChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "joinChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// LeaveChat Removes the current user from chat members. Private and secret chats can't be left using this method
// @param chatId Chat identifier
func (client *Client) LeaveChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "leaveChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddChatMember Adds a new member to a chat. Members can't be added to private or secret chats
// @param chatId Chat identifier
// @param userId Identifier of the user
// @param forwardLimit The number of earlier messages from the chat to be forwarded to the new member; up to 100. Ignored for supergroups and channels, or if the added user is a bot
func (client *Client) AddChatMember(chatId int64, userId int64, forwardLimit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "addChatMember",
		"chat_id":       chatId,
		"user_id":       userId,
		"forward_limit": forwardLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddChatMembers Adds multiple new members to a chat. Currently this method is only available for supergroups and channels. This method can't be used to join a chat. Members can't be added to a channel if it has more than 200 members
// @param chatId Chat identifier
// @param userIds Identifiers of the users to be added to the chat. The maximum number of added users is 20 for supergroups and 100 for channels
func (client *Client) AddChatMembers(chatId int64, userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "addChatMembers",
		"chat_id":  chatId,
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetChatMemberStatus Changes the status of a chat member, needs appropriate privileges. This function is currently not suitable for transferring chat ownership; use transferChatOwnership instead. Use addChatMember or banChatMember if some additional parameters needs to be passed
// @param chatId Chat identifier
// @param memberId Member identifier. Chats can be only banned and unbanned in supergroups and channels
// @param status The new status of the member in the chat
func (client *Client) SetChatMemberStatus(chatId int64, memberId MessageSender, status ChatMemberStatus) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setChatMemberStatus",
		"chat_id":   chatId,
		"member_id": memberId,
		"status":    status,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// BanChatMember Bans a member in a chat. Members can't be banned in private or secret chats. In supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first
// @param chatId Chat identifier
// @param memberId Member identifier
// @param bannedUntilDate Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever. Ignored in basic groups
// @param revokeMessages Pass true to delete all messages in the chat for the user that is being removed. Always true for supergroups and channels
func (client *Client) BanChatMember(chatId int64, memberId MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "banChatMember",
		"chat_id":           chatId,
		"member_id":         memberId,
		"banned_until_date": bannedUntilDate,
		"revoke_messages":   revokeMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err
}

// TransferChatOwnership Changes the owner of a chat. The current user must be a current owner of the chat. Use the method canTransferOwnership to check whether the ownership can be transferred from the current session. Available only for supergroups and channel chats
// @param chatId Chat identifier
// @param userId Identifier of the user to which transfer the ownership. The ownership can't be transferred to a bot or to a deleted user
// @param password The password of the current user
func (client *Client) TransferChatOwnership(chatId int64, userId int64, password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "transferChatOwnership",
		"chat_id":  chatId,
		"user_id":  userId,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ClearAllDraftMessages Clears draft messages in all chats
// @param excludeSecretChats If true, local draft messages in secret chats will not be cleared
func (client *Client) ClearAllDraftMessages(excludeSecretChats bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "clearAllDraftMessages",
		"exclude_secret_chats": excludeSecretChats,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetScopeNotificationSettings Changes notification settings for chats of a given type
// @param scope Types of chats for which to change the notification settings
// @param notificationSettings The new notification settings for the given scope
func (client *Client) SetScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setScopeNotificationSettings",
		"scope":                 scope,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ResetAllNotificationSettings Resets all notification settings to their default values. By default, all chats are unmuted, the sound is set to "default" and message previews are shown
func (client *Client) ResetAllNotificationSettings() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetAllNotificationSettings",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleChatIsPinned Changes the pinned state of a chat. There can be up to GetOption("pinned_chat_count_max")/GetOption("pinned_archived_chat_count_max") pinned non-secret chats and the same number of secret chats in the main/arhive chat list
// @param chatList Chat list in which to change the pinned state of the chat
// @param chatId Chat identifier
// @param isPinned True, if the chat is pinned
func (client *Client) ToggleChatIsPinned(chatList ChatList, chatId int64, isPinned bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "toggleChatIsPinned",
		"chat_list": chatList,
		"chat_id":   chatId,
		"is_pinned": isPinned,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetPinnedChats Changes the order of pinned chats
// @param chatList Chat list in which to change the order of pinned chats
// @param chatIds The new list of pinned chats
func (client *Client) SetPinnedChats(chatList ChatList, chatIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setPinnedChats",
		"chat_list": chatList,
		"chat_ids":  chatIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing
// @param fileId Identifier of a file to stop downloading
// @param onlyIfPending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (client *Client) CancelDownloadFile(fileId int32, onlyIfPending bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "cancelDownloadFile",
		"file_id":         fileId,
		"only_if_pending": onlyIfPending,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CancelUploadFile Stops the uploading of a file. Supported only for files uploaded by using uploadFile. For other files the behavior is undefined
// @param fileId Identifier of the file to stop uploading
func (client *Client) CancelUploadFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "cancelUploadFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// WriteGeneratedFilePart Writes a part of a generated file. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct write to the destination file
// @param generationId The identifier of the generation process
// @param offset The offset from which to write the data to the file
// @param data The data to write
func (client *Client) WriteGeneratedFilePart(generationId JSONInt64, offset int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "writeGeneratedFilePart",
		"generation_id": generationId,
		"offset":        offset,
		"data":          data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetFileGenerationProgress Informs TDLib on a file generation progress
// @param generationId The identifier of the generation process
// @param expectedSize Expected size of the generated file, in bytes; 0 if unknown
// @param localPrefixSize The number of bytes already generated
func (client *Client) SetFileGenerationProgress(generationId JSONInt64, expectedSize int32, localPrefixSize int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setFileGenerationProgress",
		"generation_id":     generationId,
		"expected_size":     expectedSize,
		"local_prefix_size": localPrefixSize,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// FinishFileGeneration Finishes the file generation
// @param generationId The identifier of the generation process
// @param error If passed, the file generation has failed and must be terminated; pass null if the file generation succeeded
func (client *Client) FinishFileGeneration(generationId JSONInt64, error *Error) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "finishFileGeneration",
		"generation_id": generationId,
		"error":         error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteFile Deletes a file from the TDLib file cache
// @param fileId Identifier of the file to delete
func (client *Client) DeleteFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ImportMessages Imports messages exported from another app
// @param chatId Identifier of a chat to which the messages will be imported. It must be an identifier of a private chat with a mutual contact or an identifier of a supergroup chat with can_change_info administrator right
// @param messageFile File with messages to import. Only inputFileLocal and inputFileGenerated are supported. The file must not be previously uploaded
// @param attachedFiles Files used in the imported messages. Only inputFileLocal and inputFileGenerated are supported. The files must not be previously uploaded
func (client *Client) ImportMessages(chatId int64, messageFile InputFile, attachedFiles []InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "importMessages",
		"chat_id":        chatId,
		"message_file":   messageFile,
		"attached_files": attachedFiles,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteRevokedChatInviteLink Deletes revoked chat invite links. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param inviteLink Invite link to revoke
func (client *Client) DeleteRevokedChatInviteLink(chatId int64, inviteLink string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteRevokedChatInviteLink",
		"chat_id":     chatId,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteAllRevokedChatInviteLinks Deletes all revoked chat invite links created by a given chat administrator. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatId Chat identifier
// @param creatorUserId User identifier of a chat administrator, which links will be deleted. Must be an identifier of the current user for non-owner
func (client *Client) DeleteAllRevokedChatInviteLinks(chatId int64, creatorUserId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "deleteAllRevokedChatInviteLinks",
		"chat_id":         chatId,
		"creator_user_id": creatorUserId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ApproveChatJoinRequest Approves pending join request in a chat
// @param chatId Chat identifier
// @param userId Identifier of the user, which request will be approved
func (client *Client) ApproveChatJoinRequest(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "approveChatJoinRequest",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeclineChatJoinRequest Declines pending join request in a chat
// @param chatId Chat identifier
// @param userId Identifier of the user, which request will be declined
func (client *Client) DeclineChatJoinRequest(chatId int64, userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "declineChatJoinRequest",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AcceptCall Accepts an incoming call
// @param callId Call identifier
// @param protocol The call protocols supported by the application
func (client *Client) AcceptCall(callId int32, protocol *CallProtocol) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "acceptCall",
		"call_id":  callId,
		"protocol": protocol,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendCallSignalingData Sends call signaling data
// @param callId Call identifier
// @param data The data
func (client *Client) SendCallSignalingData(callId int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendCallSignalingData",
		"call_id": callId,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DiscardCall Discards a call
// @param callId Call identifier
// @param isDisconnected True, if the user was disconnected
// @param duration The call duration, in seconds
// @param isVideo True, if the call was a video call
// @param connectionId Identifier of the connection used during the call
func (client *Client) DiscardCall(callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "discardCall",
		"call_id":         callId,
		"is_disconnected": isDisconnected,
		"duration":        duration,
		"is_video":        isVideo,
		"connection_id":   connectionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendCallRating Sends a call rating
// @param callId Call identifier
// @param rating Call rating; 1-5
// @param comment An optional user comment if the rating is less than 5
// @param problems List of the exact types of problems with the call, specified by the user
func (client *Client) SendCallRating(callId int32, rating int32, comment string, problems []CallProblem) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "sendCallRating",
		"call_id":  callId,
		"rating":   rating,
		"comment":  comment,
		"problems": problems,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendCallDebugInformation Sends debug information for a call
// @param callId Call identifier
// @param debugInformation Debug information in application-specific format
func (client *Client) SendCallDebugInformation(callId int32, debugInformation string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendCallDebugInformation",
		"call_id":           callId,
		"debug_information": debugInformation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetVideoChatDefaultParticipant Changes default participant identifier, which can be used to join video chats in a chat
// @param chatId Chat identifier
// @param defaultParticipantId Default group call participant identifier to join the video chats
func (client *Client) SetVideoChatDefaultParticipant(chatId int64, defaultParticipantId MessageSender) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "setVideoChatDefaultParticipant",
		"chat_id":                chatId,
		"default_participant_id": defaultParticipantId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// StartScheduledGroupCall Starts a scheduled group call
// @param groupCallId Group call identifier
func (client *Client) StartScheduledGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "startScheduledGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallEnabledStartNotification Toggles whether the current user will receive a notification when the group call will start; scheduled group calls only
// @param groupCallId Group call identifier
// @param enabledStartNotification New value of the enabled_start_notification setting
func (client *Client) ToggleGroupCallEnabledStartNotification(groupCallId int32, enabledStartNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "toggleGroupCallEnabledStartNotification",
		"group_call_id":              groupCallId,
		"enabled_start_notification": enabledStartNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallScreenSharingIsPaused Pauses or unpauses screen sharing in a joined group call
// @param groupCallId Group call identifier
// @param isPaused True if screen sharing is paused
func (client *Client) ToggleGroupCallScreenSharingIsPaused(groupCallId int32, isPaused bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleGroupCallScreenSharingIsPaused",
		"group_call_id": groupCallId,
		"is_paused":     isPaused,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EndGroupCallScreenSharing Ends screen sharing in a joined group call
// @param groupCallId Group call identifier
func (client *Client) EndGroupCallScreenSharing(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "endGroupCallScreenSharing",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetGroupCallTitle Sets group call title. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
// @param title New group call title; 1-64 characters
func (client *Client) SetGroupCallTitle(groupCallId int32, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallTitle",
		"group_call_id": groupCallId,
		"title":         title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallMuteNewParticipants Toggles whether new participants of a group call can be unmuted only by administrators of the group call. Requires groupCall.can_toggle_mute_new_participants group call flag
// @param groupCallId Group call identifier
// @param muteNewParticipants New value of the mute_new_participants setting
func (client *Client) ToggleGroupCallMuteNewParticipants(groupCallId int32, muteNewParticipants bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "toggleGroupCallMuteNewParticipants",
		"group_call_id":         groupCallId,
		"mute_new_participants": muteNewParticipants,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RevokeGroupCallInviteLink Revokes invite link for a group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
func (client *Client) RevokeGroupCallInviteLink(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "revokeGroupCallInviteLink",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// InviteGroupCallParticipants Invites users to an active group call. Sends a service message of type messageInviteToGroupCall for video chats
// @param groupCallId Group call identifier
// @param userIds User identifiers. At most 10 users can be invited simultaneously
func (client *Client) InviteGroupCallParticipants(groupCallId int32, userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "inviteGroupCallParticipants",
		"group_call_id": groupCallId,
		"user_ids":      userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// StartGroupCallRecording Starts recording of an active group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
// @param title Group call recording title; 0-64 characters
// @param recordVideo Pass true to record a video file instead of an audio file
// @param usePortraitOrientation Pass true to use portrait orientation for video instead of landscape one
func (client *Client) StartGroupCallRecording(groupCallId int32, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "startGroupCallRecording",
		"group_call_id":            groupCallId,
		"title":                    title,
		"record_video":             recordVideo,
		"use_portrait_orientation": usePortraitOrientation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EndGroupCallRecording Ends recording of an active group call. Requires groupCall.can_be_managed group call flag
// @param groupCallId Group call identifier
func (client *Client) EndGroupCallRecording(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "endGroupCallRecording",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallIsMyVideoPaused Toggles whether current user's video is paused
// @param groupCallId Group call identifier
// @param isMyVideoPaused Pass true if the current user's video is paused
func (client *Client) ToggleGroupCallIsMyVideoPaused(groupCallId int32, isMyVideoPaused bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "toggleGroupCallIsMyVideoPaused",
		"group_call_id":      groupCallId,
		"is_my_video_paused": isMyVideoPaused,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallIsMyVideoEnabled Toggles whether current user's video is enabled
// @param groupCallId Group call identifier
// @param isMyVideoEnabled Pass true if the current user's video is enabled
func (client *Client) ToggleGroupCallIsMyVideoEnabled(groupCallId int32, isMyVideoEnabled bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleGroupCallIsMyVideoEnabled",
		"group_call_id":       groupCallId,
		"is_my_video_enabled": isMyVideoEnabled,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetGroupCallParticipantIsSpeaking Informs TDLib that speaking state of a participant of an active group has changed
// @param groupCallId Group call identifier
// @param audioSource Group call participant's synchronization audio source identifier, or 0 for the current user
// @param isSpeaking True, if the user is speaking
func (client *Client) SetGroupCallParticipantIsSpeaking(groupCallId int32, audioSource int32, isSpeaking bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallParticipantIsSpeaking",
		"group_call_id": groupCallId,
		"audio_source":  audioSource,
		"is_speaking":   isSpeaking,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallParticipantIsMuted Toggles whether a participant of an active group call is muted, unmuted, or allowed to unmute themselves
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param isMuted Pass true if the user must be muted and false otherwise
func (client *Client) ToggleGroupCallParticipantIsMuted(groupCallId int32, participantId MessageSender, isMuted bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "toggleGroupCallParticipantIsMuted",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"is_muted":       isMuted,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetGroupCallParticipantVolumeLevel Changes volume level of a participant of an active group call. If the current user can manage the group call, then the participant's volume level will be changed for all users with the default volume level
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param volumeLevel New participant's volume level; 1-20000 in hundreds of percents
func (client *Client) SetGroupCallParticipantVolumeLevel(groupCallId int32, participantId MessageSender, volumeLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setGroupCallParticipantVolumeLevel",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"volume_level":   volumeLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleGroupCallParticipantIsHandRaised Toggles whether a group call participant hand is rased
// @param groupCallId Group call identifier
// @param participantId Participant identifier
// @param isHandRaised Pass true if the user's hand needs to be raised. Only self hand can be raised. Requires groupCall.can_be_managed group call flag to lower other's hand
func (client *Client) ToggleGroupCallParticipantIsHandRaised(groupCallId int32, participantId MessageSender, isHandRaised bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "toggleGroupCallParticipantIsHandRaised",
		"group_call_id":  groupCallId,
		"participant_id": participantId,
		"is_hand_raised": isHandRaised,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// LoadGroupCallParticipants Loads more participants of a group call. The loaded participants will be received through updates. Use the field groupCall.loaded_all_participants to check whether all participants has already been loaded
// @param groupCallId Group call identifier. The group call must be previously received through getGroupCall and must be joined or being joined
// @param limit The maximum number of participants to load; up to 100
func (client *Client) LoadGroupCallParticipants(groupCallId int32, limit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "loadGroupCallParticipants",
		"group_call_id": groupCallId,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// LeaveGroupCall Leaves a group call
// @param groupCallId Group call identifier
func (client *Client) LeaveGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "leaveGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DiscardGroupCall Discards a group call. Requires groupCall.can_be_managed
// @param groupCallId Group call identifier
func (client *Client) DiscardGroupCall(groupCallId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "discardGroupCall",
		"group_call_id": groupCallId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleMessageSenderIsBlocked Changes the block state of a message sender. Currently, only users and supergroup chats can be blocked
// @param sender Message Sender
// @param isBlocked New value of is_blocked
func (client *Client) ToggleMessageSenderIsBlocked(sender MessageSender, isBlocked bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "toggleMessageSenderIsBlocked",
		"sender":     sender,
		"is_blocked": isBlocked,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// BlockMessageSenderFromReplies Blocks an original sender of a message in the Replies chat
// @param messageId The identifier of an incoming message in the Replies chat
// @param deleteMessage Pass true if the message must be deleted
// @param deleteAllMessages Pass true if all messages from the same sender must be deleted
// @param reportSpam Pass true if the sender must be reported to the Telegram moderators
func (client *Client) BlockMessageSenderFromReplies(messageId int64, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "blockMessageSenderFromReplies",
		"message_id":          messageId,
		"delete_message":      deleteMessage,
		"delete_all_messages": deleteAllMessages,
		"report_spam":         reportSpam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddContact Adds a user to the contact list or edits an existing contact by their user identifier
// @param contact The contact to add or edit; phone number can be empty and needs to be specified only if known, vCard is ignored
// @param sharePhoneNumber True, if the new contact needs to be allowed to see current user's phone number. A corresponding rule to userPrivacySettingShowPhoneNumber will be added if needed. Use the field userFullInfo.need_phone_number_privacy_exception to check whether the current user needs to be asked to share their phone number
func (client *Client) AddContact(contact *Contact, sharePhoneNumber bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "addContact",
		"contact":            contact,
		"share_phone_number": sharePhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveContacts Removes users from the contact list
// @param userIds Identifiers of users to be deleted
func (client *Client) RemoveContacts(userIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeContacts",
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ClearImportedContacts Clears all imported contacts, contact list remains unchanged
func (client *Client) ClearImportedContacts() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearImportedContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SharePhoneNumber Shares the phone number of the current user with a mutual contact. Supposed to be called when the user clicks on chatActionBarSharePhoneNumber
// @param userId Identifier of the user with whom to share the phone number. The user must be a mutual contact
func (client *Client) SharePhoneNumber(userId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sharePhoneNumber",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set
// @param setId Identifier of the sticker set
// @param isInstalled The new value of is_installed
// @param isArchived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (client *Client) ChangeStickerSet(setId JSONInt64, isInstalled bool, isArchived bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changeStickerSet",
		"set_id":       setId,
		"is_installed": isInstalled,
		"is_archived":  isArchived,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user
// @param stickerSetIds Identifiers of viewed trending sticker sets
func (client *Client) ViewTrendingStickerSets(stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "viewTrendingStickerSets",
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReorderInstalledStickerSets Changes the order of installed sticker sets
// @param isMasks Pass true to change the order of mask sticker sets; pass false to change the order of ordinary sticker sets
// @param stickerSetIds Identifiers of installed sticker sets in the new correct order
func (client *Client) ReorderInstalledStickerSets(isMasks bool, stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderInstalledStickerSets",
		"is_masks":        isMasks,
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers
// @param isAttached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers
// @param sticker Sticker file to delete
func (client *Client) RemoveRecentSticker(isAttached bool, sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "removeRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ClearRecentStickers Clears the list of recently used stickers
// @param isAttached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (client *Client) ClearRecentStickers(isAttached bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "clearRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param sticker Sticker file to add
func (client *Client) AddFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers
// @param sticker Sticker file to delete from the list
func (client *Client) RemoveFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first. Only non-secret video animations with MIME type "video/mp4" can be added to the list
// @param animation The animation file to be added. Only animations known to the server (i.e., successfully sent via a message) can be added to the list
func (client *Client) AddSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveSavedAnimation Removes an animation from the list of saved animations
// @param animation Animation file to be removed
func (client *Client) RemoveSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "removeSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags
// @param hashtag Hashtag to delete
func (client *Client) RemoveRecentHashtag(hashtag string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentHashtag",
		"hashtag": hashtag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetProfilePhoto Changes a profile photo for the current user
// @param photo Profile photo to set
func (client *Client) SetProfilePhoto(photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setProfilePhoto",
		"photo": photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteProfilePhoto Deletes a profile photo
// @param profilePhotoId Identifier of the profile photo to delete
func (client *Client) DeleteProfilePhoto(profilePhotoId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteProfilePhoto",
		"profile_photo_id": profilePhotoId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetName Changes the first and last name of the current user
// @param firstName The new value of the first name for the current user; 1-64 characters
// @param lastName The new value of the optional last name for the current user; 0-64 characters
func (client *Client) SetName(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setName",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetBio Changes the bio of the current user
// @param bio The new value of the user bio; 0-70 characters without line feeds
func (client *Client) SetBio(bio string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setBio",
		"bio":   bio,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetUsername Changes the username of the current user
// @param username The new value of the username. Use an empty string to remove the username
func (client *Client) SetUsername(username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setUsername",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetLocation Changes the location of the current user. Needs to be called if GetOption("is_location_visible") is true and location changes for more than 1 kilometer
// @param location The new location of the user
func (client *Client) SetLocation(location *Location) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setLocation",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckChangePhoneNumberCode Checks the authentication code sent to confirm a new phone number of the user
// @param code Verification code received by SMS, phone call or flash call
func (client *Client) CheckChangePhoneNumberCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkChangePhoneNumberCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetCommands Sets the list of commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to change commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code. If empty, the commands will be applied to all users from the given scope, for which language there are no dedicated commands
// @param commands List of the bot's commands
func (client *Client) SetCommands(scope BotCommandScope, languageCode string, commands []BotCommand) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setCommands",
		"scope":         scope,
		"language_code": languageCode,
		"commands":      commands,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteCommands Deletes commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant; pass null to delete commands in the default bot command scope
// @param languageCode A two-letter ISO 639-1 country code or an empty string
func (client *Client) DeleteCommands(scope BotCommandScope, languageCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "deleteCommands",
		"scope":         scope,
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TerminateSession Terminates a session of the current user
// @param sessionId Session identifier
func (client *Client) TerminateSession(sessionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "terminateSession",
		"session_id": sessionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (client *Client) TerminateAllOtherSessions() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "terminateAllOtherSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DisconnectWebsite Disconnects website from the current user's Telegram account
// @param websiteId Website identifier
func (client *Client) DisconnectWebsite(websiteId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "disconnectWebsite",
		"website_id": websiteId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (client *Client) DisconnectAllWebsites() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disconnectAllWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetSupergroupUsername Changes the username of a supergroup or channel, requires owner privileges in the supergroup or channel
// @param supergroupId Identifier of the supergroup or channel
// @param username New value of the username. Use an empty string to remove the username
func (client *Client) SetSupergroupUsername(supergroupId int64, username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setSupergroupUsername",
		"supergroup_id": supergroupId,
		"username":      username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires can_change_info administrator right
// @param supergroupId Identifier of the supergroup
// @param stickerSetId New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (client *Client) SetSupergroupStickerSet(supergroupId int64, stickerSetId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setSupergroupStickerSet",
		"supergroup_id":  supergroupId,
		"sticker_set_id": stickerSetId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleSupergroupSignMessages Toggles sender signatures messages sent in a channel; requires can_change_info administrator right
// @param supergroupId Identifier of the channel
// @param signMessages New value of sign_messages
func (client *Client) ToggleSupergroupSignMessages(supergroupId int64, signMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupSignMessages",
		"supergroup_id": supergroupId,
		"sign_messages": signMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires can_change_info administrator right
// @param supergroupId The identifier of the supergroup
// @param isAllHistoryAvailable The new value of is_all_history_available
func (client *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupId int64, isAllHistoryAvailable bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "toggleSupergroupIsAllHistoryAvailable",
		"supergroup_id":            supergroupId,
		"is_all_history_available": isAllHistoryAvailable,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ToggleSupergroupIsBroadcastGroup Upgrades supergroup to a broadcast group; requires owner privileges in the supergroup
// @param supergroupId Identifier of the supergroup
func (client *Client) ToggleSupergroupIsBroadcastGroup(supergroupId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupIsBroadcastGroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReportSupergroupSpam Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
// @param supergroupId Supergroup identifier
// @param userId User identifier
// @param messageIds Identifiers of messages sent in the supergroup by the user. This list must be non-empty
func (client *Client) ReportSupergroupSpam(supergroupId int64, userId int64, messageIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "reportSupergroupSpam",
		"supergroup_id": supergroupId,
		"user_id":       userId,
		"message_ids":   messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CloseSecretChat Closes a secret chat, effectively transferring its state to secretChatStateClosed
// @param secretChatId Secret chat identifier
func (client *Client) CloseSecretChat(secretChatId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "closeSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteSavedOrderInfo Deletes saved order info
func (client *Client) DeleteSavedOrderInfo() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (client *Client) DeleteSavedCredentials() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedCredentials",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveBackground Removes background from the list of installed backgrounds
// @param backgroundId The background identifier
func (client *Client) RemoveBackground(backgroundId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "removeBackground",
		"background_id": backgroundId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ResetBackgrounds Resets list of installed backgrounds to its default value
func (client *Client) ResetBackgrounds() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetBackgrounds",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SynchronizeLanguagePack Fetches the latest versions of all strings from a language pack in the current localization target from the server. This method doesn't need to be called explicitly for the current used/base language packs. Can be called before authorization
// @param languagePackId Language pack identifier
func (client *Client) SynchronizeLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "synchronizeLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddCustomServerLanguagePack Adds a custom server language pack to the list of installed language packs in current localization target. Can be called before authorization
// @param languagePackId Identifier of a language pack to be added; may be different from a name that is used in an "https://t.me/setlanguage/" link
func (client *Client) AddCustomServerLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "addCustomServerLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetCustomLanguagePack Adds or changes a custom local language pack to the current localization target
// @param info Information about the language pack. Language pack ID must start with 'X', consist only of English letters, digits and hyphens, and must not exceed 64 characters. Can be called before authorization
// @param strings Strings of the new language pack
func (client *Client) SetCustomLanguagePack(info *LanguagePackInfo, strings []LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setCustomLanguagePack",
		"info":    info,
		"strings": strings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EditCustomLanguagePackInfo Edits information about a custom local language pack in the current localization target. Can be called before authorization
// @param info New information about the custom local language pack
func (client *Client) EditCustomLanguagePackInfo(info *LanguagePackInfo) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "editCustomLanguagePackInfo",
		"info":  info,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom local language pack. Can be called before authorization
// @param languagePackId Identifier of a previously added custom local language pack in the current localization target
// @param newString New language pack string
func (client *Client) SetCustomLanguagePackString(languagePackId string, newString *LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "setCustomLanguagePackString",
		"language_pack_id": languagePackId,
		"new_string":       newString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack which is currently in use (including base language pack) or is being synchronized can't be deleted. Can be called before authorization
// @param languagePackId Identifier of the language pack to delete
func (client *Client) DeleteLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ProcessPushNotification Handles a push notification. Returns error with code 406 if the push notification is not supported and connection to the server is required to fetch new data. Can be called before authorization
// @param payload JSON-encoded push notification payload with all fields sent by the server, and "google.sent_time" and "google.notification.sound" fields added
func (client *Client) ProcessPushNotification(payload string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "processPushNotification",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetUserPrivacySettingRules Changes user privacy settings
// @param setting The privacy setting
// @param rules The new privacy rules
func (client *Client) SetUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setUserPrivacySettingRules",
		"setting": setting,
		"rules":   rules,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
// @param name The name of the option
// @param value The new value of the option; pass null to reset option value to a default value
func (client *Client) SetOption(name string, value OptionValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setOption",
		"name":  name,
		"value": value,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetAccountTtl Changes the period of inactivity after which the account of the current user will automatically be deleted
// @param ttl New account TTL
func (client *Client) SetAccountTtl(ttl *AccountTtl) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setAccountTtl",
		"ttl":   ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account. Can be called before authorization when the current authorization state is authorizationStateWaitPassword
// @param reason The reason why the account was deleted; optional
func (client *Client) DeleteAccount(reason string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAccount",
		"reason": reason,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveChatActionBar Removes a chat action bar without any other action
// @param chatId Chat identifier
func (client *Client) RemoveChatActionBar(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeChatActionBar",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReportChat Reports a chat to the Telegram moderators. A chat can be reported only from the chat action bar, or if this is a private chat with a bot, a private chat with a user sharing their location, a supergroup, or a channel, since other chats can't be checked by moderators
// @param chatId Chat identifier
// @param messageIds Identifiers of reported messages, if any
// @param reason The reason for reporting the chat
// @param text Additional report details; 0-1024 characters
func (client *Client) ReportChat(chatId int64, messageIds []int64, reason ChatReportReason, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "reportChat",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"reason":      reason,
		"text":        text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ReportChatPhoto Reports a chat photo to the Telegram moderators. A chat photo can be reported only if this is a private chat with a bot, a private chat with a user sharing their location, a supergroup, or a channel, since other chats can't be checked by moderators
// @param chatId Chat identifier
// @param fileId Identifier of the photo to report. Only full photos from chatPhoto can be reported
// @param reason The reason for reporting the chat photo
// @param text Additional report details; 0-1024 characters
func (client *Client) ReportChatPhoto(chatId int64, fileId int32, reason ChatReportReason, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "reportChatPhoto",
		"chat_id": chatId,
		"file_id": fileId,
		"reason":  reason,
		"text":    text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks, so it must be called whenever the network is changed, even if the network type remains the same. Network type is used to check whether the library can use the network at all and also for collecting detailed network data usage statistics
// @param typeParam The new network type; pass null to set network type to networkTypeOther
func (client *Client) SetNetworkType(typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setNetworkType",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization
// @param entry The network statistics entry with the data to be added to statistics
func (client *Client) AddNetworkStatistics(entry NetworkStatisticsEntry) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "addNetworkStatistics",
		"entry": entry,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (client *Client) ResetNetworkStatistics() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetNetworkStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetAutoDownloadSettings Sets auto-download settings
// @param settings New user auto-download settings
// @param typeParam Type of the network for which the new settings are relevant
func (client *Client) SetAutoDownloadSettings(settings *AutoDownloadSettings, typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setAutoDownloadSettings",
		"settings": settings,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DeletePassportElement Deletes a Telegram Passport element
// @param typeParam Element type
func (client *Client) DeletePassportElement(typeParam PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deletePassportElement",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed
// @param userId User identifier
// @param errors The errors
func (client *Client) SetPassportElementErrors(userId int64, errors []InputPassportElementError) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setPassportElementErrors",
		"user_id": userId,
		"errors":  errors,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckPhoneNumberVerificationCode Checks the phone number verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckPhoneNumberVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckEmailAddressVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkEmailAddressVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service. This method must be called after getPassportAuthorizationFormAvailableElements if some previously available elements are going to be reused
// @param autorizationFormId Authorization form identifier
// @param typeParams Types of Telegram Passport elements chosen by user to complete the authorization form
func (client *Client) SendPassportAuthorizationForm(autorizationFormId int32, typeParams []PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "sendPassportAuthorizationForm",
		"autorization_form_id": autorizationFormId,
		"types":                typeParams,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// CheckPhoneNumberConfirmationCode Checks phone number confirmation code
// @param code The phone number confirmation code
func (client *Client) CheckPhoneNumberConfirmationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberConfirmationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
// @param pendingUpdateCount The number of pending updates
// @param errorMessage The last error message
func (client *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "setBotUpdatesStatus",
		"pending_update_count": pendingUpdateCount,
		"error_message":        errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
// @param position New position of the sticker in the set, zero-based
func (client *Client) SetStickerPositionInSet(sticker InputFile, position int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setStickerPositionInSet",
		"sticker":  sticker,
		"position": position,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
func (client *Client) RemoveStickerFromSet(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeStickerFromSet",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AcceptTermsOfService Accepts Telegram terms of services
// @param termsOfServiceId Terms of service identifier
func (client *Client) AcceptTermsOfService(termsOfServiceId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "acceptTermsOfService",
		"terms_of_service_id": termsOfServiceId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AnswerCustomQuery Answers a custom query; for bots only
// @param customQueryId Identifier of a custom query
// @param data JSON-serialized answer to the query
func (client *Client) AnswerCustomQuery(customQueryId JSONInt64, data string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "answerCustomQuery",
		"custom_query_id": customQueryId,
		"data":            data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before initialization
// @param seconds Number of seconds before the function returns
func (client *Client) SetAlarm(seconds float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setAlarm",
		"seconds": seconds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SaveApplicationLogEvent Saves application log event on the server. Can be called before authorization
// @param typeParam Event type
// @param chatId Optional chat identifier, associated with the event
// @param data The log event data
func (client *Client) SaveApplicationLogEvent(typeParam string, chatId int64, data JsonValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "saveApplicationLogEvent",
		"type":    typeParam,
		"chat_id": chatId,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) EnableProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "enableProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (client *Client) DisableProxy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disableProxy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// RemoveProxy Removes a proxy server. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) RemoveProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetLogStream Sets new log stream for internal logging of TDLib. Can be called synchronously
// @param logStream New log stream
func (client *Client) SetLogStream(logStream LogStream) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setLogStream",
		"log_stream": logStream,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib. Can be called synchronously
// @param newVerbosityLevel New value of the verbosity level for logging. Value 0 corresponds to fatal errors, value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings, value 3 corresponds to informational, value 4 corresponds to debug, value 5 corresponds to verbose debug, value greater than 5 and up to 1023 can be used to enable even more logging
func (client *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// SetLogTagVerbosityLevel Sets the verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
// @param newVerbosityLevel New verbosity level; 1-1024
func (client *Client) SetLogTagVerbosityLevel(tag string, newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogTagVerbosityLevel",
		"tag":                 tag,
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// AddLogMessage Adds a message to TDLib internal log. Can be called synchronously
// @param verbosityLevel The minimum verbosity level needed for the message to be logged; 0-1023
// @param text Text of a message to log
func (client *Client) AddLogMessage(verbosityLevel int32, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "addLogMessage",
		"verbosity_level": verbosityLevel,
		"text":            text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TestCallEmpty Does nothing; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestCallEmpty() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallEmpty",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only. Can be called before authorization
func (client *Client) TestNetwork() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testNetwork",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TestProxy Sends a simple network request to the Telegram servers via proxy; for testing only. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param typeParam Proxy type
// @param dcId Identifier of a datacenter, with which to test connection
// @param timeout The maximum overall timeout for the request
func (client *Client) TestProxy(server string, port int32, typeParam ProxyType, dcId int32, timeout float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "testProxy",
		"server":  server,
		"port":    port,
		"type":    typeParam,
		"dc_id":   dcId,
		"timeout": timeout,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (client *Client) TestGetDifference() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testGetDifference",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err
}
