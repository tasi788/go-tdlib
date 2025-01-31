package tdlib

import (
	"encoding/json"
	"fmt"
)

// ImportedContacts Represents the result of an ImportContacts request
type ImportedContacts struct {
	tdCommon
	UserIds       []int64 `json:"user_ids"`       // User identifiers of the imported contacts in the same order as they were specified in the request; 0 if the contact is not yet a registered user
	ImporterCount []int32 `json:"importer_count"` // The number of users that imported the corresponding contact; 0 for already registered users or if unavailable
}

// MessageType return the string telegram-type of ImportedContacts
func (importedContacts *ImportedContacts) MessageType() string {
	return "importedContacts"
}

// NewImportedContacts creates a new ImportedContacts
//
// @param userIds User identifiers of the imported contacts in the same order as they were specified in the request; 0 if the contact is not yet a registered user
// @param importerCount The number of users that imported the corresponding contact; 0 for already registered users or if unavailable
func NewImportedContacts(userIds []int64, importerCount []int32) *ImportedContacts {
	importedContactsTemp := ImportedContacts{
		tdCommon:      tdCommon{Type: "importedContacts"},
		UserIds:       userIds,
		ImporterCount: importerCount,
	}

	return &importedContactsTemp
}

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
// @param contacts The list of contacts to import or edit; contacts' vCard are ignored and are not imported
func (client *Client) ImportContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "importContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err
}

// ChangeImportedContacts Changes imported contacts using the list of contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts. Query result depends on the result of the previous query, so only one query is possible at the same time
// @param contacts The new list of contacts, contact's vCard are ignored and are not imported
func (client *Client) ChangeImportedContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "changeImportedContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err
}
