package tdlib

// DatedFile File with the date it was uploaded
type DatedFile struct {
	tdCommon
	File *File `json:"file"` // The file
	Date int32 `json:"date"` // Point in time (Unix timestamp) when the file was uploaded
}

// MessageType return the string telegram-type of DatedFile
func (datedFile *DatedFile) MessageType() string {
	return "datedFile"
}

// NewDatedFile creates a new DatedFile
//
// @param file The file
// @param date Point in time (Unix timestamp) when the file was uploaded
func NewDatedFile(file *File, date int32) *DatedFile {
	datedFileTemp := DatedFile{
		tdCommon: tdCommon{Type: "datedFile"},
		File:     file,
		Date:     date,
	}

	return &datedFileTemp
}
