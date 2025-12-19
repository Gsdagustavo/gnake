package main

import (
	"gnake/entities"

	"github.com/hajimehoshi/ebiten"
)

const (
	WindowTitle = "GoSnake"
)

func main() {
	game := entities.NewGame()
	ebiten.SetWindowSize(entities.WindowWidth, entities.WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
