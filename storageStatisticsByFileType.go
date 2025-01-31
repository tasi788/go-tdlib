package tdlib

import (
	"encoding/json"
)

// StorageStatisticsByFileType Contains the storage usage statistics for a specific file type
type StorageStatisticsByFileType struct {
	tdCommon
	FileType FileType `json:"file_type"` // File type
	Size     int64    `json:"size"`      // Total size of the files, in bytes
	Count    int32    `json:"count"`     // Total number of files
}

// MessageType return the string telegram-type of StorageStatisticsByFileType
func (storageStatisticsByFileType *StorageStatisticsByFileType) MessageType() string {
	return "storageStatisticsByFileType"
}

// NewStorageStatisticsByFileType creates a new StorageStatisticsByFileType
//
// @param fileType File type
// @param size Total size of the files, in bytes
// @param count Total number of files
func NewStorageStatisticsByFileType(fileType FileType, size int64, count int32) *StorageStatisticsByFileType {
	storageStatisticsByFileTypeTemp := StorageStatisticsByFileType{
		tdCommon: tdCommon{Type: "storageStatisticsByFileType"},
		FileType: fileType,
		Size:     size,
		Count:    count,
	}

	return &storageStatisticsByFileTypeTemp
}

// UnmarshalJSON unmarshal to json
func (storageStatisticsByFileType *StorageStatisticsByFileType) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Size  int64 `json:"size"`  // Total size of the files, in bytes
		Count int32 `json:"count"` // Total number of files
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	storageStatisticsByFileType.tdCommon = tempObj.tdCommon
	storageStatisticsByFileType.Size = tempObj.Size
	storageStatisticsByFileType.Count = tempObj.Count

	fieldFileType, _ := unmarshalFileType(objMap["file_type"])
	storageStatisticsByFileType.FileType = fieldFileType

	return nil
}
