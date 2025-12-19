package entity

type Enemy struct {
	Entity Entity
	Health int
	Damage int
}

func (e *Enemy) MoveTowardsPlayer(plr Player) {
	sub := plr.Entity.Position.Sub(e.Entity.Position)
	direction := sub.Normalized()
	e.Entity.Move(direction)
}

func NewEnemy(entity Entity, health, damage int) Enemy {
	return Enemy{
		Entity: entity,
		Health: health,
		Damage: damage,
	}
}
