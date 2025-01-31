package tdlib

import (
	"encoding/json"
	"fmt"
)

// Session Contains information about one session in a Telegram application used by the current user. Sessions must be shown to the user in the returned order
type Session struct {
	tdCommon
	Id                    JSONInt64 `json:"id"`                      // Session identifier
	IsCurrent             bool      `json:"is_current"`              // True, if this session is the current session
	IsPasswordPending     bool      `json:"is_password_pending"`     // True, if a password is needed to complete authorization of the session
	CanAcceptSecretChats  bool      `json:"can_accept_secret_chats"` // True, if incoming secret chats can be accepted by the session
	CanAcceptCalls        bool      `json:"can_accept_calls"`        // True, if incoming calls can be accepted by the session
	ApiId                 int32     `json:"api_id"`                  // Telegram API identifier, as provided by the application
	ApplicationName       string    `json:"application_name"`        // Name of the application, as provided by the application
	ApplicationVersion    string    `json:"application_version"`     // The version of the application, as provided by the application
	IsOfficialApplication bool      `json:"is_official_application"` // True, if the application is an official application or uses the api_id of an official application
	DeviceModel           string    `json:"device_model"`            // Model of the device the application has been run or is running on, as provided by the application
	Platform              string    `json:"platform"`                // Operating system the application has been run or is running on, as provided by the application
	SystemVersion         string    `json:"system_version"`          // Version of the operating system the application has been run or is running on, as provided by the application
	LogInDate             int32     `json:"log_in_date"`             // Point in time (Unix timestamp) when the user has logged in
	LastActiveDate        int32     `json:"last_active_date"`        // Point in time (Unix timestamp) when the session was last used
	Ip                    string    `json:"ip"`                      // IP address from which the session was created, in human-readable format
	Country               string    `json:"country"`                 // A two-letter country code for the country from which the session was created, based on the IP address
	Region                string    `json:"region"`                  // Region code from which the session was created, based on the IP address
}

// MessageType return the string telegram-type of Session
func (session *Session) MessageType() string {
	return "session"
}

// NewSession creates a new Session
//
// @param id Session identifier
// @param isCurrent True, if this session is the current session
// @param isPasswordPending True, if a password is needed to complete authorization of the session
// @param canAcceptSecretChats True, if incoming secret chats can be accepted by the session
// @param canAcceptCalls True, if incoming calls can be accepted by the session
// @param apiId Telegram API identifier, as provided by the application
// @param applicationName Name of the application, as provided by the application
// @param applicationVersion The version of the application, as provided by the application
// @param isOfficialApplication True, if the application is an official application or uses the api_id of an official application
// @param deviceModel Model of the device the application has been run or is running on, as provided by the application
// @param platform Operating system the application has been run or is running on, as provided by the application
// @param systemVersion Version of the operating system the application has been run or is running on, as provided by the application
// @param logInDate Point in time (Unix timestamp) when the user has logged in
// @param lastActiveDate Point in time (Unix timestamp) when the session was last used
// @param ip IP address from which the session was created, in human-readable format
// @param country A two-letter country code for the country from which the session was created, based on the IP address
// @param region Region code from which the session was created, based on the IP address
func NewSession(id JSONInt64, isCurrent bool, isPasswordPending bool, canAcceptSecretChats bool, canAcceptCalls bool, apiId int32, applicationName string, applicationVersion string, isOfficialApplication bool, deviceModel string, platform string, systemVersion string, logInDate int32, lastActiveDate int32, ip string, country string, region string) *Session {
	sessionTemp := Session{
		tdCommon:              tdCommon{Type: "session"},
		Id:                    id,
		IsCurrent:             isCurrent,
		IsPasswordPending:     isPasswordPending,
		CanAcceptSecretChats:  canAcceptSecretChats,
		CanAcceptCalls:        canAcceptCalls,
		ApiId:                 apiId,
		ApplicationName:       applicationName,
		ApplicationVersion:    applicationVersion,
		IsOfficialApplication: isOfficialApplication,
		DeviceModel:           deviceModel,
		Platform:              platform,
		SystemVersion:         systemVersion,
		LogInDate:             logInDate,
		LastActiveDate:        lastActiveDate,
		Ip:                    ip,
		Country:               country,
		Region:                region,
	}

	return &sessionTemp
}

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success
// @param link A link from a QR code. The link must be scanned by the in-app camera
func (client *Client) ConfirmQrCodeAuthentication(link string) (*Session, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "confirmQrCodeAuthentication",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var session Session
	err = json.Unmarshal(result.Raw, &session)
	return &session, err
}
