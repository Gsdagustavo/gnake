package player

import (
	"gnake/internal/core"
	"gnake/internal/entity"
	"gnake/internal/graphics"
	"gnake/internal/math"
)

const (
	DefaultPlayerWidth         = 20
	DefaultPlayerHeight        = 20
	DefaultPlayerMovementSpeed = 8
)

var (
	DefaultPlayerColor = graphics.Blue
)

type Player struct {
	Entity           entity.Entity
	PlayerController PlayerController
	Score            int
}

func NewPlayer() Player {
	return Player{
		Entity: entity.Entity{
			Position:      math.NewVector(core.WindowWidth/2, core.WindowHeight/2),
			Size:          math.NewVector(DefaultPlayerWidth, DefaultPlayerHeight),
			Speed:         math.Vector{},
			MovementSpeed: DefaultPlayerMovementSpeed,
			Color:         DefaultPlayerColor,
		},
		PlayerController: NewPlayerController(),
	}
}
