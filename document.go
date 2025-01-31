package tdlib

// Document Describes a document of any type
type Document struct {
	tdCommon
	FileName      string         `json:"file_name"`     // Original name of the file; as defined by the sender
	MimeType      string         `json:"mime_type"`     // MIME type of the file; as defined by the sender
	Minithumbnail *Minithumbnail `json:"minithumbnail"` // Document minithumbnail; may be null
	Thumbnail     *Thumbnail     `json:"thumbnail"`     // Document thumbnail in JPEG or PNG format (PNG will be used only for background patterns); as defined by the sender; may be null
	Document      *File          `json:"document"`      // File containing the document
}

// MessageType return the string telegram-type of Document
func (document *Document) MessageType() string {
	return "document"
}

// NewDocument creates a new Document
//
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file; as defined by the sender
// @param minithumbnail Document minithumbnail; may be null
// @param thumbnail Document thumbnail in JPEG or PNG format (PNG will be used only for background patterns); as defined by the sender; may be null
// @param document File containing the document
func NewDocument(fileName string, mimeType string, minithumbnail *Minithumbnail, thumbnail *Thumbnail, document *File) *Document {
	documentTemp := Document{
		tdCommon:      tdCommon{Type: "document"},
		FileName:      fileName,
		MimeType:      mimeType,
		Minithumbnail: minithumbnail,
		Thumbnail:     thumbnail,
		Document:      document,
	}

	return &documentTemp
}
