package tdlib

// ColorReplacement Describes a color replacement for animated emoji
type ColorReplacement struct {
	tdCommon
	OldColor int32 `json:"old_color"` // Original animated emoji color in the RGB24 format
	NewColor int32 `json:"new_color"` // Replacement animated emoji color in the RGB24 format
}

// MessageType return the string telegram-type of ColorReplacement
func (colorReplacement *ColorReplacement) MessageType() string {
	return "colorReplacement"
}

// NewColorReplacement creates a new ColorReplacement
//
// @param oldColor Original animated emoji color in the RGB24 format
// @param newColor Replacement animated emoji color in the RGB24 format
func NewColorReplacement(oldColor int32, newColor int32) *ColorReplacement {
	colorReplacementTemp := ColorReplacement{
		tdCommon: tdCommon{Type: "colorReplacement"},
		OldColor: oldColor,
		NewColor: newColor,
	}

	return &colorReplacementTemp
}
