package game

import (
	"gnake/internal/core"
	"gnake/internal/player"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Player player.Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Move player
	g.Player.Entity.Move(g.Player.PlayerController.Direction())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw player
	plr := g.Player
	plrPos := plr.Entity.Position
	plrSize := plr.Entity.Size
	plrColor := plr.Entity.Color
	ebitenutil.DrawRect(screen, plrPos.X, plrPos.Y, plrSize.X, plrSize.Y, plrColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return core.WindowWidth, core.WindowHeight
}

func NewGame() *Game {
	return &Game{
		Player: player.NewPlayer(),
	}
}
