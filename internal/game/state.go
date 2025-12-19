package game

import (
	"errors"
	"gnake/internal/core"
	"gnake/internal/entity"
	"gnake/internal/math"
	"math/rand"
)

const (
	DefaultEnemyWidth  = 20
	DefaultEnemyHeight = 20

	DefaultEnemyMoveSpeed = 5

	DefaultEnemyHealth = 100
	DefaultEnemyDamage = 1
)

func GenerateEnemies(quantity int) []entity.Enemy {
	enemies := make([]entity.Enemy, 0)

	for range quantity {
		x := rand.Intn(core.WindowWidth + DefaultEnemyWidth)
		y := rand.Intn(core.WindowHeight + DefaultEnemyHeight)

		initialPosition := math.NewVector(float64(x), float64(y))
		initialSize := math.NewVector(DefaultEnemyWidth, DefaultEnemyHeight)

		img, err := core.LoadAsset("enemy/skull_enemy.png")
		if err != nil {
			panic(err)
		}

		if img == nil {
			panic(errors.New("invalid enemy image"))
		}

		e := entity.NewEntity(initialPosition, initialSize, DefaultEnemyMoveSpeed, img)

		enemy := entity.NewEnemy(e, DefaultEnemyHealth, DefaultEnemyDamage)
		enemies = append(enemies, enemy)
	}

	return enemies
}
