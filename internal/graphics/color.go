package graphics

import "math/rand"

const (
	// RGBAFactor is the factor to multiply the color values by
	// This is needed to convert the color values to the internal format
	RGBAFactor = 0x101
)

// Color is a wrapper for RGBA colors
type Color struct {
	R, G, B, A uint8
}

// NewColor creates a new color
func NewColor(R, G, B, A uint8) Color {
	return Color{R, G, B, A}
}

// NewRandomColor creates a new random color
func NewRandomColor() Color {
	return NewColor(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 1)
}

// RGBA returns the color as RGBA values
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) * RGBAFactor
	g = uint32(c.G) * RGBAFactor
	b = uint32(c.B) * RGBAFactor
	a = uint32(c.A) * RGBAFactor
	return r, g, b, a
}

// Predefined colors
var (
	White  = Color{255, 255, 255, 255}
	Black  = Color{0, 0, 0, 255}
	Red    = Color{255, 0, 0, 255}
	Green  = Color{0, 255, 0, 255}
	Blue   = Color{0, 0, 255, 255}
	Yellow = Color{255, 255, 0, 255}
)
