package tdlib

import (
	"encoding/json"
	"fmt"
)

// DeviceToken Represents a data needed to subscribe for push notifications through registerDevice method. To use specific push notification service, the correct application platform must be specified and a valid server authentication data must be uploaded at https://my.telegram.org
type DeviceToken interface {
	GetDeviceTokenEnum() DeviceTokenEnum
}

// DeviceTokenEnum Alias for abstract DeviceToken 'Sub-Classes', used as constant-enum here
type DeviceTokenEnum string

// DeviceToken enums
const (
	DeviceTokenFirebaseCloudMessagingType DeviceTokenEnum = "deviceTokenFirebaseCloudMessaging"
	DeviceTokenApplePushType              DeviceTokenEnum = "deviceTokenApplePush"
	DeviceTokenApplePushVoIPType          DeviceTokenEnum = "deviceTokenApplePushVoIP"
	DeviceTokenWindowsPushType            DeviceTokenEnum = "deviceTokenWindowsPush"
	DeviceTokenMicrosoftPushType          DeviceTokenEnum = "deviceTokenMicrosoftPush"
	DeviceTokenMicrosoftPushVoIPType      DeviceTokenEnum = "deviceTokenMicrosoftPushVoIP"
	DeviceTokenWebPushType                DeviceTokenEnum = "deviceTokenWebPush"
	DeviceTokenSimplePushType             DeviceTokenEnum = "deviceTokenSimplePush"
	DeviceTokenUbuntuPushType             DeviceTokenEnum = "deviceTokenUbuntuPush"
	DeviceTokenBlackBerryPushType         DeviceTokenEnum = "deviceTokenBlackBerryPush"
	DeviceTokenTizenPushType              DeviceTokenEnum = "deviceTokenTizenPush"
)

func unmarshalDeviceToken(rawMsg *json.RawMessage) (DeviceToken, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch DeviceTokenEnum(objMap["@type"].(string)) {
	case DeviceTokenFirebaseCloudMessagingType:
		var deviceTokenFirebaseCloudMessaging DeviceTokenFirebaseCloudMessaging
		err := json.Unmarshal(*rawMsg, &deviceTokenFirebaseCloudMessaging)
		return &deviceTokenFirebaseCloudMessaging, err

	case DeviceTokenApplePushType:
		var deviceTokenApplePush DeviceTokenApplePush
		err := json.Unmarshal(*rawMsg, &deviceTokenApplePush)
		return &deviceTokenApplePush, err

	case DeviceTokenApplePushVoIPType:
		var deviceTokenApplePushVoIP DeviceTokenApplePushVoIP
		err := json.Unmarshal(*rawMsg, &deviceTokenApplePushVoIP)
		return &deviceTokenApplePushVoIP, err

	case DeviceTokenWindowsPushType:
		var deviceTokenWindowsPush DeviceTokenWindowsPush
		err := json.Unmarshal(*rawMsg, &deviceTokenWindowsPush)
		return &deviceTokenWindowsPush, err

	case DeviceTokenMicrosoftPushType:
		var deviceTokenMicrosoftPush DeviceTokenMicrosoftPush
		err := json.Unmarshal(*rawMsg, &deviceTokenMicrosoftPush)
		return &deviceTokenMicrosoftPush, err

	case DeviceTokenMicrosoftPushVoIPType:
		var deviceTokenMicrosoftPushVoIP DeviceTokenMicrosoftPushVoIP
		err := json.Unmarshal(*rawMsg, &deviceTokenMicrosoftPushVoIP)
		return &deviceTokenMicrosoftPushVoIP, err

	case DeviceTokenWebPushType:
		var deviceTokenWebPush DeviceTokenWebPush
		err := json.Unmarshal(*rawMsg, &deviceTokenWebPush)
		return &deviceTokenWebPush, err

	case DeviceTokenSimplePushType:
		var deviceTokenSimplePush DeviceTokenSimplePush
		err := json.Unmarshal(*rawMsg, &deviceTokenSimplePush)
		return &deviceTokenSimplePush, err

	case DeviceTokenUbuntuPushType:
		var deviceTokenUbuntuPush DeviceTokenUbuntuPush
		err := json.Unmarshal(*rawMsg, &deviceTokenUbuntuPush)
		return &deviceTokenUbuntuPush, err

	case DeviceTokenBlackBerryPushType:
		var deviceTokenBlackBerryPush DeviceTokenBlackBerryPush
		err := json.Unmarshal(*rawMsg, &deviceTokenBlackBerryPush)
		return &deviceTokenBlackBerryPush, err

	case DeviceTokenTizenPushType:
		var deviceTokenTizenPush DeviceTokenTizenPush
		err := json.Unmarshal(*rawMsg, &deviceTokenTizenPush)
		return &deviceTokenTizenPush, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// DeviceTokenFirebaseCloudMessaging A token for Firebase Cloud Messaging
type DeviceTokenFirebaseCloudMessaging struct {
	tdCommon
	Token   string `json:"token"`   // Device registration token; may be empty to deregister a device
	Encrypt bool   `json:"encrypt"` // True, if push notifications must be additionally encrypted
}

// MessageType return the string telegram-type of DeviceTokenFirebaseCloudMessaging
func (deviceTokenFirebaseCloudMessaging *DeviceTokenFirebaseCloudMessaging) MessageType() string {
	return "deviceTokenFirebaseCloudMessaging"
}

// NewDeviceTokenFirebaseCloudMessaging creates a new DeviceTokenFirebaseCloudMessaging
//
// @param token Device registration token; may be empty to deregister a device
// @param encrypt True, if push notifications must be additionally encrypted
func NewDeviceTokenFirebaseCloudMessaging(token string, encrypt bool) *DeviceTokenFirebaseCloudMessaging {
	deviceTokenFirebaseCloudMessagingTemp := DeviceTokenFirebaseCloudMessaging{
		tdCommon: tdCommon{Type: "deviceTokenFirebaseCloudMessaging"},
		Token:    token,
		Encrypt:  encrypt,
	}

	return &deviceTokenFirebaseCloudMessagingTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenFirebaseCloudMessaging *DeviceTokenFirebaseCloudMessaging) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenFirebaseCloudMessagingType
}

// DeviceTokenApplePush A token for Apple Push Notification service
type DeviceTokenApplePush struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token; may be empty to deregister a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePush
func (deviceTokenApplePush *DeviceTokenApplePush) MessageType() string {
	return "deviceTokenApplePush"
}

// NewDeviceTokenApplePush creates a new DeviceTokenApplePush
//
// @param deviceToken Device token; may be empty to deregister a device
// @param isAppSandbox True, if App Sandbox is enabled
func NewDeviceTokenApplePush(deviceToken string, isAppSandbox bool) *DeviceTokenApplePush {
	deviceTokenApplePushTemp := DeviceTokenApplePush{
		tdCommon:     tdCommon{Type: "deviceTokenApplePush"},
		DeviceToken:  deviceToken,
		IsAppSandbox: isAppSandbox,
	}

	return &deviceTokenApplePushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePush *DeviceTokenApplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushType
}

// DeviceTokenApplePushVoIP A token for Apple Push Notification service VoIP notifications
type DeviceTokenApplePushVoIP struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token; may be empty to deregister a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
	Encrypt      bool   `json:"encrypt"`        // True, if push notifications must be additionally encrypted
}

// MessageType return the string telegram-type of DeviceTokenApplePushVoIP
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) MessageType() string {
	return "deviceTokenApplePushVoIP"
}

// NewDeviceTokenApplePushVoIP creates a new DeviceTokenApplePushVoIP
//
// @param deviceToken Device token; may be empty to deregister a device
// @param isAppSandbox True, if App Sandbox is enabled
// @param encrypt True, if push notifications must be additionally encrypted
func NewDeviceTokenApplePushVoIP(deviceToken string, isAppSandbox bool, encrypt bool) *DeviceTokenApplePushVoIP {
	deviceTokenApplePushVoIPTemp := DeviceTokenApplePushVoIP{
		tdCommon:     tdCommon{Type: "deviceTokenApplePushVoIP"},
		DeviceToken:  deviceToken,
		IsAppSandbox: isAppSandbox,
		Encrypt:      encrypt,
	}

	return &deviceTokenApplePushVoIPTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushVoIPType
}

// DeviceTokenWindowsPush A token for Windows Push Notification Services
type DeviceTokenWindowsPush struct {
	tdCommon
	AccessToken string `json:"access_token"` // The access token that will be used to send notifications; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenWindowsPush
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) MessageType() string {
	return "deviceTokenWindowsPush"
}

// NewDeviceTokenWindowsPush creates a new DeviceTokenWindowsPush
//
// @param accessToken The access token that will be used to send notifications; may be empty to deregister a device
func NewDeviceTokenWindowsPush(accessToken string) *DeviceTokenWindowsPush {
	deviceTokenWindowsPushTemp := DeviceTokenWindowsPush{
		tdCommon:    tdCommon{Type: "deviceTokenWindowsPush"},
		AccessToken: accessToken,
	}

	return &deviceTokenWindowsPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWindowsPushType
}

// DeviceTokenMicrosoftPush A token for Microsoft Push Notification Service
type DeviceTokenMicrosoftPush struct {
	tdCommon
	ChannelUri string `json:"channel_uri"` // Push notification channel URI; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPush
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) MessageType() string {
	return "deviceTokenMicrosoftPush"
}

// NewDeviceTokenMicrosoftPush creates a new DeviceTokenMicrosoftPush
//
// @param channelUri Push notification channel URI; may be empty to deregister a device
func NewDeviceTokenMicrosoftPush(channelUri string) *DeviceTokenMicrosoftPush {
	deviceTokenMicrosoftPushTemp := DeviceTokenMicrosoftPush{
		tdCommon:   tdCommon{Type: "deviceTokenMicrosoftPush"},
		ChannelUri: channelUri,
	}

	return &deviceTokenMicrosoftPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushType
}

// DeviceTokenMicrosoftPushVoIP A token for Microsoft Push Notification Service VoIP channel
type DeviceTokenMicrosoftPushVoIP struct {
	tdCommon
	ChannelUri string `json:"channel_uri"` // Push notification channel URI; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPushVoIP
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) MessageType() string {
	return "deviceTokenMicrosoftPushVoIP"
}

// NewDeviceTokenMicrosoftPushVoIP creates a new DeviceTokenMicrosoftPushVoIP
//
// @param channelUri Push notification channel URI; may be empty to deregister a device
func NewDeviceTokenMicrosoftPushVoIP(channelUri string) *DeviceTokenMicrosoftPushVoIP {
	deviceTokenMicrosoftPushVoIPTemp := DeviceTokenMicrosoftPushVoIP{
		tdCommon:   tdCommon{Type: "deviceTokenMicrosoftPushVoIP"},
		ChannelUri: channelUri,
	}

	return &deviceTokenMicrosoftPushVoIPTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushVoIPType
}

// DeviceTokenWebPush A token for web Push API
type DeviceTokenWebPush struct {
	tdCommon
	Endpoint        string `json:"endpoint"`         // Absolute URL exposed by the push service where the application server can send push messages; may be empty to deregister a device
	P256dhBase64url string `json:"p256dh_base64url"` // Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
	AuthBase64url   string `json:"auth_base64url"`   // Base64url-encoded authentication secret
}

// MessageType return the string telegram-type of DeviceTokenWebPush
func (deviceTokenWebPush *DeviceTokenWebPush) MessageType() string {
	return "deviceTokenWebPush"
}

// NewDeviceTokenWebPush creates a new DeviceTokenWebPush
//
// @param endpoint Absolute URL exposed by the push service where the application server can send push messages; may be empty to deregister a device
// @param p256dhBase64url Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
// @param authBase64url Base64url-encoded authentication secret
func NewDeviceTokenWebPush(endpoint string, p256dhBase64url string, authBase64url string) *DeviceTokenWebPush {
	deviceTokenWebPushTemp := DeviceTokenWebPush{
		tdCommon:        tdCommon{Type: "deviceTokenWebPush"},
		Endpoint:        endpoint,
		P256dhBase64url: p256dhBase64url,
		AuthBase64url:   authBase64url,
	}

	return &deviceTokenWebPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWebPush *DeviceTokenWebPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWebPushType
}

// DeviceTokenSimplePush A token for Simple Push API for Firefox OS
type DeviceTokenSimplePush struct {
	tdCommon
	Endpoint string `json:"endpoint"` // Absolute URL exposed by the push service where the application server can send push messages; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenSimplePush
func (deviceTokenSimplePush *DeviceTokenSimplePush) MessageType() string {
	return "deviceTokenSimplePush"
}

// NewDeviceTokenSimplePush creates a new DeviceTokenSimplePush
//
// @param endpoint Absolute URL exposed by the push service where the application server can send push messages; may be empty to deregister a device
func NewDeviceTokenSimplePush(endpoint string) *DeviceTokenSimplePush {
	deviceTokenSimplePushTemp := DeviceTokenSimplePush{
		tdCommon: tdCommon{Type: "deviceTokenSimplePush"},
		Endpoint: endpoint,
	}

	return &deviceTokenSimplePushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenSimplePush *DeviceTokenSimplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenSimplePushType
}

// DeviceTokenUbuntuPush A token for Ubuntu Push Client service
type DeviceTokenUbuntuPush struct {
	tdCommon
	Token string `json:"token"` // Token; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenUbuntuPush
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) MessageType() string {
	return "deviceTokenUbuntuPush"
}

// NewDeviceTokenUbuntuPush creates a new DeviceTokenUbuntuPush
//
// @param token Token; may be empty to deregister a device
func NewDeviceTokenUbuntuPush(token string) *DeviceTokenUbuntuPush {
	deviceTokenUbuntuPushTemp := DeviceTokenUbuntuPush{
		tdCommon: tdCommon{Type: "deviceTokenUbuntuPush"},
		Token:    token,
	}

	return &deviceTokenUbuntuPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenUbuntuPushType
}

// DeviceTokenBlackBerryPush A token for BlackBerry Push Service
type DeviceTokenBlackBerryPush struct {
	tdCommon
	Token string `json:"token"` // Token; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenBlackBerryPush
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) MessageType() string {
	return "deviceTokenBlackBerryPush"
}

// NewDeviceTokenBlackBerryPush creates a new DeviceTokenBlackBerryPush
//
// @param token Token; may be empty to deregister a device
func NewDeviceTokenBlackBerryPush(token string) *DeviceTokenBlackBerryPush {
	deviceTokenBlackBerryPushTemp := DeviceTokenBlackBerryPush{
		tdCommon: tdCommon{Type: "deviceTokenBlackBerryPush"},
		Token:    token,
	}

	return &deviceTokenBlackBerryPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenBlackBerryPushType
}

// DeviceTokenTizenPush A token for Tizen Push Service
type DeviceTokenTizenPush struct {
	tdCommon
	RegId string `json:"reg_id"` // Push service registration identifier; may be empty to deregister a device
}

// MessageType return the string telegram-type of DeviceTokenTizenPush
func (deviceTokenTizenPush *DeviceTokenTizenPush) MessageType() string {
	return "deviceTokenTizenPush"
}

// NewDeviceTokenTizenPush creates a new DeviceTokenTizenPush
//
// @param regId Push service registration identifier; may be empty to deregister a device
func NewDeviceTokenTizenPush(regId string) *DeviceTokenTizenPush {
	deviceTokenTizenPushTemp := DeviceTokenTizenPush{
		tdCommon: tdCommon{Type: "deviceTokenTizenPush"},
		RegId:    regId,
	}

	return &deviceTokenTizenPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenTizenPush *DeviceTokenTizenPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenTizenPushType
}
