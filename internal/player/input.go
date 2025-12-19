package player

import (
	"gnake/internal/math"

	"github.com/hajimehoshi/ebiten"
)

const (
	MoveUpKey    = ebiten.KeyW
	MoveDownKey  = ebiten.KeyS
	MoveLeftKey  = ebiten.KeyA
	MoveRightKey = ebiten.KeyD
)

// IsKeyPressed returns true if the given key is pressed
func IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

// PlayerController handles player input
type PlayerController struct{}

func NewPlayerController() PlayerController {
	return PlayerController{}
}

func (p *PlayerController) Direction() math.Vector {
	dir := math.NewZeroVector()

	if IsKeyPressed(MoveLeftKey) {
		dir.X--
	}

	if IsKeyPressed(MoveRightKey) {
		dir.X++
	}

	if IsKeyPressed(MoveUpKey) {
		dir.Y--
	}

	if IsKeyPressed(MoveDownKey) {
		dir.Y++
	}

	return dir.Normalized()
}
