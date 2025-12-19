package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	WindowWidth  = 800
	WindowHeight = 600
)

const (
	TargetFPS = 60
)

type Game struct {
	Player *Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw player
	plrPos := g.Player.Entity.Position
	plrSize := g.Player.Entity.Size
	plrColor := g.Player.Entity.Color
	ebitenutil.DrawRect(screen, plrPos.X, plrPos.Y, plrSize.X, plrSize.Y, *plrColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func NewGame() *Game {
	return &Game{
		Player: NewPlayer(),
	}
}
