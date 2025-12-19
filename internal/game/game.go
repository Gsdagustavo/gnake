package game

import (
	"gnake/internal/core"
	"gnake/internal/entity"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Player *entity.Player
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
	return core.WindowWidth, core.WindowHeight
}

func NewGame() *Game {
	return &Game{
		Player: entity.NewPlayer(),
	}
}
