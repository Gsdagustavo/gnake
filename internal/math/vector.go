package math

import "math"

// Vector represents a 2D vector
type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

func NewZeroVector() Vector {
	return Vector{0, 0}
}

// Add adds the other vector to the current one
func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

// Sub subrtracts the other vector from the current one
func (v *Vector) Sub(other Vector) Vector {
	return NewVector(v.X-other.X, v.Y-other.Y)
}

// Scale multiplies the vector factors by a factor
func (v *Vector) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}

// Normalize normalizes the vector factors
func (v *Vector) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X = v.X / magnitude
	v.Y = v.Y / magnitude
}

// Normalized returns a new vector with the normalized factors
func (v *Vector) Normalized() Vector {
	if v.X == 0 && v.Y == 0 {
		return *v
	}

	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	x := v.X / magnitude
	y := v.Y / magnitude
	return Vector{x, y}
}
