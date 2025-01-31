package tdlib

import (
	"encoding/json"
	"fmt"
)

// StorageStatisticsFast Contains approximate storage usage statistics, excluding files of unknown file type
type StorageStatisticsFast struct {
	tdCommon
	FilesSize                int64 `json:"files_size"`                  // Approximate total size of files, in bytes
	FileCount                int32 `json:"file_count"`                  // Approximate number of files
	DatabaseSize             int64 `json:"database_size"`               // Size of the database
	LanguagePackDatabaseSize int64 `json:"language_pack_database_size"` // Size of the language pack database
	LogSize                  int64 `json:"log_size"`                    // Size of the TDLib internal log
}

// MessageType return the string telegram-type of StorageStatisticsFast
func (storageStatisticsFast *StorageStatisticsFast) MessageType() string {
	return "storageStatisticsFast"
}

// NewStorageStatisticsFast creates a new StorageStatisticsFast
//
// @param filesSize Approximate total size of files, in bytes
// @param fileCount Approximate number of files
// @param databaseSize Size of the database
// @param languagePackDatabaseSize Size of the language pack database
// @param logSize Size of the TDLib internal log
func NewStorageStatisticsFast(filesSize int64, fileCount int32, databaseSize int64, languagePackDatabaseSize int64, logSize int64) *StorageStatisticsFast {
	storageStatisticsFastTemp := StorageStatisticsFast{
		tdCommon:                 tdCommon{Type: "storageStatisticsFast"},
		FilesSize:                filesSize,
		FileCount:                fileCount,
		DatabaseSize:             databaseSize,
		LanguagePackDatabaseSize: languagePackDatabaseSize,
		LogSize:                  logSize,
	}

	return &storageStatisticsFastTemp
}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics. Can be called before authorization
func (client *Client) GetStorageStatisticsFast() (*StorageStatisticsFast, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStorageStatisticsFast",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatisticsFast StorageStatisticsFast
	err = json.Unmarshal(result.Raw, &storageStatisticsFast)
	return &storageStatisticsFast, err
}
