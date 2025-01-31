package tdlib

// Point A point on a Cartesian plane
type Point struct {
	tdCommon
	X float64 `json:"x"` // The point's first coordinate
	Y float64 `json:"y"` // The point's second coordinate
}

// MessageType return the string telegram-type of Point
func (point *Point) MessageType() string {
	return "point"
}

// NewPoint creates a new Point
//
// @param x The point's first coordinate
// @param y The point's second coordinate
func NewPoint(x float64, y float64) *Point {
	pointTemp := Point{
		tdCommon: tdCommon{Type: "point"},
		X:        x,
		Y:        y,
	}

	return &pointTemp
}
