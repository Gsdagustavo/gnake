package entities

import (
	"gnake/entities/core"
	"image/color"
)

type Player struct {
	Entity core.Entity
	Score  int
}

func NewPlayer() *Player {
	return &Player{
		Entity: core.Entity{
			Position: core.NewVector(WindowWidth/2, WindowHeight/2),
			Size:     core.NewVector(20, 20),
			Color:    &color.RGBA{R: 255, A: 255},
		},
	}
}

func (p *Player) Move(direction core.Vector) {
	p.Entity.Position.Add(direction)
}
