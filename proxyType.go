package tdlib

import (
	"encoding/json"
	"fmt"
)

// ProxyType Describes the type of a proxy server
type ProxyType interface {
	GetProxyTypeEnum() ProxyTypeEnum
}

// ProxyTypeEnum Alias for abstract ProxyType 'Sub-Classes', used as constant-enum here
type ProxyTypeEnum string

// ProxyType enums
const (
	ProxyTypeSocks5Type  ProxyTypeEnum = "proxyTypeSocks5"
	ProxyTypeHttpType    ProxyTypeEnum = "proxyTypeHttp"
	ProxyTypeMtprotoType ProxyTypeEnum = "proxyTypeMtproto"
)

func unmarshalProxyType(rawMsg *json.RawMessage) (ProxyType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ProxyTypeEnum(objMap["@type"].(string)) {
	case ProxyTypeSocks5Type:
		var proxyTypeSocks5 ProxyTypeSocks5
		err := json.Unmarshal(*rawMsg, &proxyTypeSocks5)
		return &proxyTypeSocks5, err

	case ProxyTypeHttpType:
		var proxyTypeHttp ProxyTypeHttp
		err := json.Unmarshal(*rawMsg, &proxyTypeHttp)
		return &proxyTypeHttp, err

	case ProxyTypeMtprotoType:
		var proxyTypeMtproto ProxyTypeMtproto
		err := json.Unmarshal(*rawMsg, &proxyTypeMtproto)
		return &proxyTypeMtproto, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ProxyTypeSocks5 A SOCKS5 proxy server
type ProxyTypeSocks5 struct {
	tdCommon
	Username string `json:"username"` // Username for logging in; may be empty
	Password string `json:"password"` // Password for logging in; may be empty
}

// MessageType return the string telegram-type of ProxyTypeSocks5
func (proxyTypeSocks5 *ProxyTypeSocks5) MessageType() string {
	return "proxyTypeSocks5"
}

// NewProxyTypeSocks5 creates a new ProxyTypeSocks5
//
// @param username Username for logging in; may be empty
// @param password Password for logging in; may be empty
func NewProxyTypeSocks5(username string, password string) *ProxyTypeSocks5 {
	proxyTypeSocks5Temp := ProxyTypeSocks5{
		tdCommon: tdCommon{Type: "proxyTypeSocks5"},
		Username: username,
		Password: password,
	}

	return &proxyTypeSocks5Temp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeSocks5 *ProxyTypeSocks5) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeSocks5Type
}

// ProxyTypeHttp A HTTP transparent proxy server
type ProxyTypeHttp struct {
	tdCommon
	Username string `json:"username"`  // Username for logging in; may be empty
	Password string `json:"password"`  // Password for logging in; may be empty
	HttpOnly bool   `json:"http_only"` // Pass true if the proxy supports only HTTP requests and doesn't support transparent TCP connections via HTTP CONNECT method
}

// MessageType return the string telegram-type of ProxyTypeHttp
func (proxyTypeHttp *ProxyTypeHttp) MessageType() string {
	return "proxyTypeHttp"
}

// NewProxyTypeHttp creates a new ProxyTypeHttp
//
// @param username Username for logging in; may be empty
// @param password Password for logging in; may be empty
// @param httpOnly Pass true if the proxy supports only HTTP requests and doesn't support transparent TCP connections via HTTP CONNECT method
func NewProxyTypeHttp(username string, password string, httpOnly bool) *ProxyTypeHttp {
	proxyTypeHttpTemp := ProxyTypeHttp{
		tdCommon: tdCommon{Type: "proxyTypeHttp"},
		Username: username,
		Password: password,
		HttpOnly: httpOnly,
	}

	return &proxyTypeHttpTemp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeHttp *ProxyTypeHttp) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeHttpType
}

// ProxyTypeMtproto An MTProto proxy server
type ProxyTypeMtproto struct {
	tdCommon
	Secret string `json:"secret"` // The proxy's secret in hexadecimal encoding
}

// MessageType return the string telegram-type of ProxyTypeMtproto
func (proxyTypeMtproto *ProxyTypeMtproto) MessageType() string {
	return "proxyTypeMtproto"
}

// NewProxyTypeMtproto creates a new ProxyTypeMtproto
//
// @param secret The proxy's secret in hexadecimal encoding
func NewProxyTypeMtproto(secret string) *ProxyTypeMtproto {
	proxyTypeMtprotoTemp := ProxyTypeMtproto{
		tdCommon: tdCommon{Type: "proxyTypeMtproto"},
		Secret:   secret,
	}

	return &proxyTypeMtprotoTemp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeMtproto *ProxyTypeMtproto) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeMtprotoType
}
