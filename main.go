package main

import (
	"gnake/internal/core"
	"gnake/internal/game"

	"github.com/hajimehoshi/ebiten"
)

const (
	WindowTitle = "GoSnake"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(core.WindowWidth, core.WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
