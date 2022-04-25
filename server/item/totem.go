package item

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
)

// Totem protects the player from death
type Totem struct {
}

// EncodeItem ...
func (Totem) EncodeItem() (name string, meta int16) {
	return "minecraft:totem_of_undying", 0
}

func (Totem) MaxCount() int {
	return 1
}

func (Totem) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration is the duration consuming the item takes. If the player is using the item for at least
// this duration, the item will be consumed and have its Consume method called.
func (Totem) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (Totem) Consume(w *world.World, c ConsumerTotem) Stack {
	c.ClearEffects()
	c.AddHealth(2)
	c.AddEffect(effect.New(effect.Absorption{}, 2, 5*time.Second))
	c.AddEffect(effect.New(effect.Regeneration{}, 2, 45*time.Second))
	c.AddEffect(effect.New(effect.FireResistance{}, 1, 40*time.Second))

	w.AddParticle(c.Position(), particle.EndermanTeleportParticle{})
	w.AddParticle(c.Position(), particle.HugeExplosion{})
	w.PlaySound(c.Position(), sound.TotemUsed{})
	return Stack{}
}
