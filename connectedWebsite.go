package tdlib

// ConnectedWebsite Contains information about one website the current user is logged in with Telegram
type ConnectedWebsite struct {
	tdCommon
	Id             JSONInt64 `json:"id"`               // Website identifier
	DomainName     string    `json:"domain_name"`      // The domain name of the website
	BotUserId      int64     `json:"bot_user_id"`      // User identifier of a bot linked with the website
	Browser        string    `json:"browser"`          // The version of a browser used to log in
	Platform       string    `json:"platform"`         // Operating system the browser is running on
	LogInDate      int32     `json:"log_in_date"`      // Point in time (Unix timestamp) when the user was logged in
	LastActiveDate int32     `json:"last_active_date"` // Point in time (Unix timestamp) when obtained authorization was last used
	Ip             string    `json:"ip"`               // IP address from which the user was logged in, in human-readable format
	Location       string    `json:"location"`         // Human-readable description of a country and a region, from which the user was logged in, based on the IP address
}

// MessageType return the string telegram-type of ConnectedWebsite
func (connectedWebsite *ConnectedWebsite) MessageType() string {
	return "connectedWebsite"
}

// NewConnectedWebsite creates a new ConnectedWebsite
//
// @param id Website identifier
// @param domainName The domain name of the website
// @param botUserId User identifier of a bot linked with the website
// @param browser The version of a browser used to log in
// @param platform Operating system the browser is running on
// @param logInDate Point in time (Unix timestamp) when the user was logged in
// @param lastActiveDate Point in time (Unix timestamp) when obtained authorization was last used
// @param ip IP address from which the user was logged in, in human-readable format
// @param location Human-readable description of a country and a region, from which the user was logged in, based on the IP address
func NewConnectedWebsite(id JSONInt64, domainName string, botUserId int64, browser string, platform string, logInDate int32, lastActiveDate int32, ip string, location string) *ConnectedWebsite {
	connectedWebsiteTemp := ConnectedWebsite{
		tdCommon:       tdCommon{Type: "connectedWebsite"},
		Id:             id,
		DomainName:     domainName,
		BotUserId:      botUserId,
		Browser:        browser,
		Platform:       platform,
		LogInDate:      logInDate,
		LastActiveDate: lastActiveDate,
		Ip:             ip,
		Location:       location,
	}

	return &connectedWebsiteTemp
}
