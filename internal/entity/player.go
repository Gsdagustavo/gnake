package entity

import (
	"gnake/internal/core"
	"image/color"
)

type Player struct {
	Entity Entity
	Score  int
}

func NewPlayer() *Player {
	return &Player{
		Entity: Entity{
			Position: NewVector(core.WindowWidth/2, core.WindowHeight/2),
			Size:     NewVector(20, 20),
			Color:    &color.RGBA{R: 255, A: 255},
		},
	}
}

func (p *Player) Move(direction Vector) {
	p.Entity.Position.Add(direction)
}
