package entity

type Enemy struct {
	Entity Entity
	Health int
	Damage int
}

func (e *Enemy) MoveTowardsPlayer(plr Player) {
	e.Entity.Move(plr.Entity.Position)
}

func NewEnemy(entity Entity, health, damage int) Enemy {
	return Enemy{
		Entity: entity,
		Health: health,
		Damage: damage,
	}
}
