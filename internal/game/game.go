package game

import (
	"errors"
	"gnake/internal/entity"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	MinimumEnemiesOnScreen = 1
	MaximumEnemiesOnScreen = 50

	ChangeToSpawnEnemiesEverySecond = 0.25
)

type Game struct {
	Player         entity.Player
	Enemies        []entity.Enemy
	GameBackground ebiten.Image
	SpawnTimer     float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Move player
	g.Player.Entity.Move(g.Player.PlayerController.Direction())

	// Move enemies
	for _, enemy := range g.Enemies {
		enemy.MoveTowardsPlayer(g.Player)
	}

	// Accumulate time for en enemy spawner
	g.SpawnTimer += 1.0 / ebiten.CurrentTPS()

	if g.SpawnTimer >= 1.0 {
		g.SpawnTimer = 0

		if rand.Float64() < ChangeToSpawnEnemiesEverySecond {
			quantEnemies := len(g.Enemies)
			if quantEnemies < MaximumEnemiesOnScreen {
				enemies := GenerateEnemies(MaximumEnemiesOnScreen - quantEnemies)
				g.Enemies = append(g.Enemies, enemies...)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw player sprite
	plr := g.Player
	plrPos := plr.Entity.Position
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(plrPos.X, plrPos.Y)
	err := screen.DrawImage(plr.Entity.Sprite, opts)
	if err != nil {
		panic(errors.Join(errors.New("failed to draw player image"), err))
	}

	// Draw enemies sprites
	for _, enemy := range g.Enemies {
		opts = &ebiten.DrawImageOptions{}

		e := enemy.Entity
		opts.GeoM.Translate(e.Position.X, e.Position.Y)
		err = screen.DrawImage(e.Sprite, opts)
		if err != nil {
			panic(errors.Join(errors.New("failed to draw enemy image"), err))
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() (*Game, error) {
	plr, err := entity.NewPlayer()
	if err != nil {
		return nil, errors.Join(errors.New("failed to create player"), err)
	}

	if plr == nil {
		return nil, errors.Join(errors.New("invalid player instance"), err)
	}

	game := &Game{
		Player: *plr,
	}

	return game, nil
}
