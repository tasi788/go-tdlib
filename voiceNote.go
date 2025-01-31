package tdlib

// VoiceNote Describes a voice note. The voice note must be encoded with the Opus codec, and stored inside an OGG container. Voice notes can have only a single audio channel
type VoiceNote struct {
	tdCommon
	Duration int32  `json:"duration"`  // Duration of the voice note, in seconds; as defined by the sender
	Waveform []byte `json:"waveform"`  // A waveform representation of the voice note in 5-bit format
	MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender
	Voice    *File  `json:"voice"`     // File containing the voice note
}

// MessageType return the string telegram-type of VoiceNote
func (voiceNote *VoiceNote) MessageType() string {
	return "voiceNote"
}

// NewVoiceNote creates a new VoiceNote
//
// @param duration Duration of the voice note, in seconds; as defined by the sender
// @param waveform A waveform representation of the voice note in 5-bit format
// @param mimeType MIME type of the file; as defined by the sender
// @param voice File containing the voice note
func NewVoiceNote(duration int32, waveform []byte, mimeType string, voice *File) *VoiceNote {
	voiceNoteTemp := VoiceNote{
		tdCommon: tdCommon{Type: "voiceNote"},
		Duration: duration,
		Waveform: waveform,
		MimeType: mimeType,
		Voice:    voice,
	}

	return &voiceNoteTemp
}
