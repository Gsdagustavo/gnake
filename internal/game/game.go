package game

import (
	"errors"
	"gnake/internal/core"
	"gnake/internal/player"
	"gnake/internal/util"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Player         player.Player
	GameBackground ebiten.Image
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Move player
	g.Player.Entity.Move(g.Player.PlayerController.Direction())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game background
	err := screen.DrawImage(&g.GameBackground, &ebiten.DrawImageOptions{})
	if err != nil {
		panic(errors.Join(core.ErrDrawGameBackgroundImage, err))
	}

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

func NewGame() (*Game, error) {
	file, err := core.AssetsFS.Open("assets/background.png")
	if err != nil {
		return nil, errors.Join(core.ErrLoadGameImage, err)
	}
	defer util.DeferFileClose(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errors.Join(core.ErrLoadGameImage, err)
	}

	if img == nil {
		return nil, core.ErrInvalidGameImage
	}

	bg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, errors.Join(core.ErrConvertImage, err)
	}

	game := &Game{
		Player:         player.NewPlayer(),
		GameBackground: *bg,
	}

	return game, nil
}
