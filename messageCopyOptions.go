package tdlib

// MessageCopyOptions Options to be used when a message content is copied without reference to the original sender. Service messages and messageInvoice can't be copied
type MessageCopyOptions struct {
	tdCommon
	SendCopy       bool           `json:"send_copy"`       // True, if content of the message needs to be copied without reference to the original sender. Always true if the message is forwarded to a secret chat or is local
	ReplaceCaption bool           `json:"replace_caption"` // True, if media caption of the message copy needs to be replaced. Ignored if send_copy is false
	NewCaption     *FormattedText `json:"new_caption"`     // New message caption; pass null to copy message without caption. Ignored if replace_caption is false
}

// MessageType return the string telegram-type of MessageCopyOptions
func (messageCopyOptions *MessageCopyOptions) MessageType() string {
	return "messageCopyOptions"
}

// NewMessageCopyOptions creates a new MessageCopyOptions
//
// @param sendCopy True, if content of the message needs to be copied without reference to the original sender. Always true if the message is forwarded to a secret chat or is local
// @param replaceCaption True, if media caption of the message copy needs to be replaced. Ignored if send_copy is false
// @param newCaption New message caption; pass null to copy message without caption. Ignored if replace_caption is false
func NewMessageCopyOptions(sendCopy bool, replaceCaption bool, newCaption *FormattedText) *MessageCopyOptions {
	messageCopyOptionsTemp := MessageCopyOptions{
		tdCommon:       tdCommon{Type: "messageCopyOptions"},
		SendCopy:       sendCopy,
		ReplaceCaption: replaceCaption,
		NewCaption:     newCaption,
	}

	return &messageCopyOptionsTemp
}
