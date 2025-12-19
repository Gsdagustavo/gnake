package entity

import (
	"gnake/internal/graphics"
	"gnake/internal/math"
)

// Entity represents an object in the game, with a position, size, speed and color
type Entity struct {
	Position      math.Vector
	Size          math.Vector
	Speed         math.Vector
	MovementSpeed float64
	Color         graphics.Color
}

func (e *Entity) Move(direction math.Vector) {
	direction.Scale(e.MovementSpeed)
	e.Position.Add(direction)
}

// CheckCollision checks if two entities collide
func CheckCollision(this Entity, other Entity) bool {
	thisPos := this.Position
	thisSize := this.Size

	otherPos := other.Position
	otherSize := other.Size

	return thisPos.X < otherPos.X+otherSize.X && thisPos.X+thisSize.X > otherPos.X && thisPos.Y < otherPos.Y+otherSize.Y && thisPos.Y+thisSize.Y > otherPos.Y
}
