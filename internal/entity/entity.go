package entity

import (
	"gnake/internal/math"

	"github.com/hajimehoshi/ebiten"
)

// Entity represents an object in the game, with a position, size, speed and color
type Entity struct {
	Position      math.Vector
	Size          math.Vector
	Speed         math.Vector
	MovementSpeed float64
	Sprite        *ebiten.Image
}

func NewEntity(position math.Vector, size math.Vector, movementSpeed float64, sprite *ebiten.Image) Entity {
	return Entity{position, size, math.NewZeroVector(), movementSpeed, sprite}
}

func (e *Entity) Move(direction math.Vector) {
	direction.Scale(e.MovementSpeed)
	e.Position.Add(direction)
}

func (e *Entity) GetCenter() math.Vector {
	centerX := e.Position.X + e.Size.X/2
	centerY := e.Position.Y + e.Size.Y/2
	return math.NewVector(centerX, centerY)
}

// CheckCollision checks if two entities collide
func CheckCollision(this Entity, other Entity) bool {
	thisPos := this.Position
	thisSize := this.Size

	otherPos := other.Position
	otherSize := other.Size

	return thisPos.X < otherPos.X+otherSize.X && thisPos.X+thisSize.X > otherPos.X && thisPos.Y < otherPos.Y+otherSize.Y && thisPos.Y+thisSize.Y > otherPos.Y
}
