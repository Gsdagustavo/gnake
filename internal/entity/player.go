package entity

import (
	"errors"
	"gnake/internal/core"
	"gnake/internal/math"
	"gnake/internal/player"
)

const (
	DefaultPlayerWidth         = 20
	DefaultPlayerHeight        = 20
	DefaultPlayerMovementSpeed = 8
)

type Player struct {
	Entity           Entity
	PlayerController player.PlayerController
	Score            int
}

func NewPlayer() (*Player, error) {
	img, err := core.LoadAsset("player/player.png")
	if err != nil {
		return nil, errors.Join(errors.New("failed to load player image"), err)
	}

	if img == nil {
		return nil, errors.New("invalid player image")
	}

	plr := &Player{
		Entity: Entity{
			Position:      math.NewVector(core.WindowWidth/2, core.WindowHeight/2),
			Size:          math.NewVector(DefaultPlayerWidth, DefaultPlayerHeight),
			Speed:         math.Vector{},
			MovementSpeed: DefaultPlayerMovementSpeed,
			Sprite:        img,
		},
		PlayerController: player.NewPlayerController(),
	}

	return plr, nil
}
