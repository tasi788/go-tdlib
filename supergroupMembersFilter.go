package tdlib

import (
	"encoding/json"
	"fmt"
)

// SupergroupMembersFilter Specifies the kind of chat members to return in getSupergroupMembers
type SupergroupMembersFilter interface {
	GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum
}

// SupergroupMembersFilterEnum Alias for abstract SupergroupMembersFilter 'Sub-Classes', used as constant-enum here
type SupergroupMembersFilterEnum string

// SupergroupMembersFilter enums
const (
	SupergroupMembersFilterRecentType         SupergroupMembersFilterEnum = "supergroupMembersFilterRecent"
	SupergroupMembersFilterContactsType       SupergroupMembersFilterEnum = "supergroupMembersFilterContacts"
	SupergroupMembersFilterAdministratorsType SupergroupMembersFilterEnum = "supergroupMembersFilterAdministrators"
	SupergroupMembersFilterSearchType         SupergroupMembersFilterEnum = "supergroupMembersFilterSearch"
	SupergroupMembersFilterRestrictedType     SupergroupMembersFilterEnum = "supergroupMembersFilterRestricted"
	SupergroupMembersFilterBannedType         SupergroupMembersFilterEnum = "supergroupMembersFilterBanned"
	SupergroupMembersFilterMentionType        SupergroupMembersFilterEnum = "supergroupMembersFilterMention"
	SupergroupMembersFilterBotsType           SupergroupMembersFilterEnum = "supergroupMembersFilterBots"
)

func unmarshalSupergroupMembersFilter(rawMsg *json.RawMessage) (SupergroupMembersFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SupergroupMembersFilterEnum(objMap["@type"].(string)) {
	case SupergroupMembersFilterRecentType:
		var supergroupMembersFilterRecent SupergroupMembersFilterRecent
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterRecent)
		return &supergroupMembersFilterRecent, err

	case SupergroupMembersFilterContactsType:
		var supergroupMembersFilterContacts SupergroupMembersFilterContacts
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterContacts)
		return &supergroupMembersFilterContacts, err

	case SupergroupMembersFilterAdministratorsType:
		var supergroupMembersFilterAdministrators SupergroupMembersFilterAdministrators
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterAdministrators)
		return &supergroupMembersFilterAdministrators, err

	case SupergroupMembersFilterSearchType:
		var supergroupMembersFilterSearch SupergroupMembersFilterSearch
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterSearch)
		return &supergroupMembersFilterSearch, err

	case SupergroupMembersFilterRestrictedType:
		var supergroupMembersFilterRestricted SupergroupMembersFilterRestricted
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterRestricted)
		return &supergroupMembersFilterRestricted, err

	case SupergroupMembersFilterBannedType:
		var supergroupMembersFilterBanned SupergroupMembersFilterBanned
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterBanned)
		return &supergroupMembersFilterBanned, err

	case SupergroupMembersFilterMentionType:
		var supergroupMembersFilterMention SupergroupMembersFilterMention
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterMention)
		return &supergroupMembersFilterMention, err

	case SupergroupMembersFilterBotsType:
		var supergroupMembersFilterBots SupergroupMembersFilterBots
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterBots)
		return &supergroupMembersFilterBots, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// SupergroupMembersFilterRecent Returns recently active users in reverse chronological order
type SupergroupMembersFilterRecent struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterRecent
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) MessageType() string {
	return "supergroupMembersFilterRecent"
}

// NewSupergroupMembersFilterRecent creates a new SupergroupMembersFilterRecent
//
func NewSupergroupMembersFilterRecent() *SupergroupMembersFilterRecent {
	supergroupMembersFilterRecentTemp := SupergroupMembersFilterRecent{
		tdCommon: tdCommon{Type: "supergroupMembersFilterRecent"},
	}

	return &supergroupMembersFilterRecentTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterRecentType
}

// SupergroupMembersFilterContacts Returns contacts of the user, which are members of the supergroup or channel
type SupergroupMembersFilterContacts struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterContacts
func (supergroupMembersFilterContacts *SupergroupMembersFilterContacts) MessageType() string {
	return "supergroupMembersFilterContacts"
}

// NewSupergroupMembersFilterContacts creates a new SupergroupMembersFilterContacts
//
// @param query Query to search for
func NewSupergroupMembersFilterContacts(query string) *SupergroupMembersFilterContacts {
	supergroupMembersFilterContactsTemp := SupergroupMembersFilterContacts{
		tdCommon: tdCommon{Type: "supergroupMembersFilterContacts"},
		Query:    query,
	}

	return &supergroupMembersFilterContactsTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterContacts *SupergroupMembersFilterContacts) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterContactsType
}

// SupergroupMembersFilterAdministrators Returns the owner and administrators
type SupergroupMembersFilterAdministrators struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterAdministrators
func (supergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators) MessageType() string {
	return "supergroupMembersFilterAdministrators"
}

// NewSupergroupMembersFilterAdministrators creates a new SupergroupMembersFilterAdministrators
//
func NewSupergroupMembersFilterAdministrators() *SupergroupMembersFilterAdministrators {
	supergroupMembersFilterAdministratorsTemp := SupergroupMembersFilterAdministrators{
		tdCommon: tdCommon{Type: "supergroupMembersFilterAdministrators"},
	}

	return &supergroupMembersFilterAdministratorsTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterAdministratorsType
}

// SupergroupMembersFilterSearch Used to search for supergroup or channel members via a (string) query
type SupergroupMembersFilterSearch struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterSearch
func (supergroupMembersFilterSearch *SupergroupMembersFilterSearch) MessageType() string {
	return "supergroupMembersFilterSearch"
}

// NewSupergroupMembersFilterSearch creates a new SupergroupMembersFilterSearch
//
// @param query Query to search for
func NewSupergroupMembersFilterSearch(query string) *SupergroupMembersFilterSearch {
	supergroupMembersFilterSearchTemp := SupergroupMembersFilterSearch{
		tdCommon: tdCommon{Type: "supergroupMembersFilterSearch"},
		Query:    query,
	}

	return &supergroupMembersFilterSearchTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterSearch *SupergroupMembersFilterSearch) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterSearchType
}

// SupergroupMembersFilterRestricted Returns restricted supergroup members; can be used only by administrators
type SupergroupMembersFilterRestricted struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterRestricted
func (supergroupMembersFilterRestricted *SupergroupMembersFilterRestricted) MessageType() string {
	return "supergroupMembersFilterRestricted"
}

// NewSupergroupMembersFilterRestricted creates a new SupergroupMembersFilterRestricted
//
// @param query Query to search for
func NewSupergroupMembersFilterRestricted(query string) *SupergroupMembersFilterRestricted {
	supergroupMembersFilterRestrictedTemp := SupergroupMembersFilterRestricted{
		tdCommon: tdCommon{Type: "supergroupMembersFilterRestricted"},
		Query:    query,
	}

	return &supergroupMembersFilterRestrictedTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterRestricted *SupergroupMembersFilterRestricted) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterRestrictedType
}

// SupergroupMembersFilterBanned Returns users banned from the supergroup or channel; can be used only by administrators
type SupergroupMembersFilterBanned struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterBanned
func (supergroupMembersFilterBanned *SupergroupMembersFilterBanned) MessageType() string {
	return "supergroupMembersFilterBanned"
}

// NewSupergroupMembersFilterBanned creates a new SupergroupMembersFilterBanned
//
// @param query Query to search for
func NewSupergroupMembersFilterBanned(query string) *SupergroupMembersFilterBanned {
	supergroupMembersFilterBannedTemp := SupergroupMembersFilterBanned{
		tdCommon: tdCommon{Type: "supergroupMembersFilterBanned"},
		Query:    query,
	}

	return &supergroupMembersFilterBannedTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterBanned *SupergroupMembersFilterBanned) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterBannedType
}

// SupergroupMembersFilterMention Returns users which can be mentioned in the supergroup
type SupergroupMembersFilterMention struct {
	tdCommon
	Query           string `json:"query"`             // Query to search for
	MessageThreadId int64  `json:"message_thread_id"` // If non-zero, the identifier of the current message thread
}

// MessageType return the string telegram-type of SupergroupMembersFilterMention
func (supergroupMembersFilterMention *SupergroupMembersFilterMention) MessageType() string {
	return "supergroupMembersFilterMention"
}

// NewSupergroupMembersFilterMention creates a new SupergroupMembersFilterMention
//
// @param query Query to search for
// @param messageThreadId If non-zero, the identifier of the current message thread
func NewSupergroupMembersFilterMention(query string, messageThreadId int64) *SupergroupMembersFilterMention {
	supergroupMembersFilterMentionTemp := SupergroupMembersFilterMention{
		tdCommon:        tdCommon{Type: "supergroupMembersFilterMention"},
		Query:           query,
		MessageThreadId: messageThreadId,
	}

	return &supergroupMembersFilterMentionTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterMention *SupergroupMembersFilterMention) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterMentionType
}

// SupergroupMembersFilterBots Returns bot members of the supergroup or channel
type SupergroupMembersFilterBots struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterBots
func (supergroupMembersFilterBots *SupergroupMembersFilterBots) MessageType() string {
	return "supergroupMembersFilterBots"
}

// NewSupergroupMembersFilterBots creates a new SupergroupMembersFilterBots
//
func NewSupergroupMembersFilterBots() *SupergroupMembersFilterBots {
	supergroupMembersFilterBotsTemp := SupergroupMembersFilterBots{
		tdCommon: tdCommon{Type: "supergroupMembersFilterBots"},
	}

	return &supergroupMembersFilterBotsTemp
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterBots *SupergroupMembersFilterBots) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterBotsType
}
