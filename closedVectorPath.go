package tdlib

// ClosedVectorPath Represents a closed vector path. The path begins at the end point of the last command
type ClosedVectorPath struct {
	tdCommon
	Commands []VectorPathCommand `json:"commands"` // List of vector path commands
}

// MessageType return the string telegram-type of ClosedVectorPath
func (closedVectorPath *ClosedVectorPath) MessageType() string {
	return "closedVectorPath"
}

// NewClosedVectorPath creates a new ClosedVectorPath
//
// @param commands List of vector path commands
func NewClosedVectorPath(commands []VectorPathCommand) *ClosedVectorPath {
	closedVectorPathTemp := ClosedVectorPath{
		tdCommon: tdCommon{Type: "closedVectorPath"},
		Commands: commands,
	}

	return &closedVectorPathTemp
}
