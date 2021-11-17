package tdlib

import (
	"encoding/json"
	"fmt"
)

// Text Contains some text
type Text struct {
	tdCommon
	Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of Text
func (text *Text) MessageType() string {
	return "text"
}

// NewText creates a new Text
//
// @param text Text
func NewText(text string) *Text {
	textTemp := Text{
		tdCommon: tdCommon{Type: "text"},
		Text:     text,
	}

	return &textTemp
}

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only for messages in supergroups and channels with a username
// @param chatId Identifier of the chat to which the message belongs
// @param messageId Identifier of the message
// @param forAlbum Pass true to return an HTML code for embedding of the whole media album
func (client *Client) GetMessageEmbeddingCode(chatId int64, messageId int64, forAlbum bool) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageEmbeddingCode",
		"chat_id":    chatId,
		"message_id": messageId,
		"for_album":  forAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously
// @param fileName The name of the file or path to the file
func (client *Client) GetFileMimeType(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileMimeType",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously
// @param mimeType The MIME type of the file
func (client *Client) GetFileExtension(mimeType string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileExtension",
		"mime_type": mimeType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// CleanFileName Removes potentially dangerous characters from the name of a file. The encoding of the file name is supposed to be UTF-8. Returns an empty string on failure. Can be called synchronously
// @param fileName File name or path to the file
func (client *Client) CleanFileName(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "cleanFileName",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously
// @param jsonStringValue The JsonValue object
func (client *Client) GetJsonString(jsonStringValue JsonValue) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getJsonString",
		"json_value": jsonStringValue,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetChatFilterDefaultIconName Returns default icon name for a filter. Can be called synchronously
// @param filter Chat filter
func (client *Client) GetChatFilterDefaultIconName(filter *ChatFilter) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getChatFilterDefaultIconName",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetSuggestedFileName Returns suggested name for saving a file in a given directory
// @param fileId Identifier of the file
// @param directory Directory in which the file is supposed to be saved
func (client *Client) GetSuggestedFileName(fileId int32, directory string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getSuggestedFileName",
		"file_id":   fileId,
		"directory": directory,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetMessageImportConfirmationText Returns a confirmation text to be shown to the user before starting message import
// @param chatId Identifier of a chat to which the messages will be imported. It must be an identifier of a private chat with a mutual contact or an identifier of a supergroup chat with can_change_info administrator right
func (client *Client) GetMessageImportConfirmationText(chatId int64) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getMessageImportConfirmationText",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// JoinGroupCall Joins an active group call. Returns join response payload for tgcalls
// @param groupCallId Group call identifier
// @param participantId Identifier of a group call participant, which will be used to join the call; pass null to join as self; video chats only
// @param audioSourceId Caller audio channel synchronization source identifier; received from tgcalls
// @param payload Group call join payload; received from tgcalls
// @param isMuted True, if the user's microphone is muted
// @param isMyVideoEnabled True, if the user's video is enabled
// @param inviteHash If non-empty, invite hash to be used to join the group call without being muted by administrators
func (client *Client) JoinGroupCall(groupCallId int32, participantId MessageSender, audioSourceId int32, payload string, isMuted bool, isMyVideoEnabled bool, inviteHash string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "joinGroupCall",
		"group_call_id":       groupCallId,
		"participant_id":      participantId,
		"audio_source_id":     audioSourceId,
		"payload":             payload,
		"is_muted":            isMuted,
		"is_my_video_enabled": isMyVideoEnabled,
		"invite_hash":         inviteHash,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// StartGroupCallScreenSharing Starts screen sharing in a joined group call. Returns join response payload for tgcalls
// @param groupCallId Group call identifier
// @param audioSourceId Screen sharing audio channel synchronization source identifier; received from tgcalls
// @param payload Group call join payload; received from tgcalls
func (client *Client) StartGroupCallScreenSharing(groupCallId int32, audioSourceId int32, payload string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "startGroupCallScreenSharing",
		"group_call_id":   groupCallId,
		"audio_source_id": audioSourceId,
		"payload":         payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which must be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
func (client *Client) GetPreferredCountryLanguage(countryCode string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getPreferredCountryLanguage",
		"country_code": countryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetSuggestedStickerSetName Returns a suggested name for a new sticker set with a given title
// @param title Sticker set title; 1-64 characters
func (client *Client) GetSuggestedStickerSetName(title string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSuggestedStickerSetName",
		"title": title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}

// GetCountryCode Uses the current IP address to find the current country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (client *Client) GetCountryCode() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountryCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err
}
