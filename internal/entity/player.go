package entity

import (
	"gnake/internal/core"
	"gnake/internal/graphics"
	"gnake/internal/math"
)

const (
	DefaultPlayerWidth  = 20
	DefaultPlayerHeight = 20
)

type Player struct {
	Entity Entity
	Score  int
}

func NewPlayer() Player {
	return Player{
		Entity: Entity{
			Position: math.NewVector(core.WindowWidth/2, core.WindowHeight/2),
			Size:     math.NewVector(DefaultPlayerWidth, DefaultPlayerHeight),
			Color:    graphics.Blue,
		},
	}
}

func (p *Player) Move(direction math.Vector) {
	p.Entity.Position.Add(direction)
}
