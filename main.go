package main

import (
	"errors"
	"gnake/internal/core"
	"gnake/internal/game"

	"github.com/hajimehoshi/ebiten"
)

const (
	WindowTitle = "GoSnake"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		panic(errors.Join(errors.New("failed to create game"), err))
	}

	ebiten.SetWindowSize(core.WindowWidth, core.WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)

	err = ebiten.RunGame(g)
	if err != nil {
		panic(errors.Join(errors.New("failed to run game"), err))
	}
}
