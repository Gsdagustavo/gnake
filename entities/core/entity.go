package core

import "image/color"

// Entity represents an object in the game, with a position, size and speed
type Entity struct {
	// Position, Size and Speed are in pixels
	Position *Vector
	Size     *Vector
	Speed    *Vector
	Color    *color.RGBA
}

// Movable represents an object that can move
type Movable interface {
	// Move moves the entity
	Move()
}

// Collidable represents an object that can collide with other objects
type Collidable interface {
	// Collides returns true if the entity collides with the other object
	Collides(other Collidable) bool
}

// Drawable represents an object that can be drawn
type Drawable interface {
	// Draw draws the entity
	Draw()
}

// CheckCollision checks if two entities collide
func CheckCollision(this Entity, other Entity) bool {
	thisPos := this.Position
	thisSize := this.Size

	otherPos := other.Position
	otherSize := other.Size

	return thisPos.X < otherPos.X+otherSize.X && thisPos.X+thisSize.X > otherPos.X && thisPos.Y < otherPos.Y+otherSize.Y && thisPos.Y+thisSize.Y > otherPos.Y
}
