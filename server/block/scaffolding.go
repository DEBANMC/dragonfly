package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// Scaffolding is a block affected by gravity. It can come in a red variant.
type Scaffolding struct {
	gravityAffected
	transparent
	solid
}

// SoilFor ...
func (s Scaffolding) SoilFor(block world.Block) bool {
	_, ok := block.(DeadBush)
	return ok
}

// NeighbourUpdateTick ...
func (s Scaffolding) NeighbourUpdateTick(pos, _ cube.Pos, w *world.World) {
	//s.fall(s, pos, w)
}

// BreakInfo ...
func (s Scaffolding) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, alwaysHarvestable, shovelEffective, oneOf(s))
}

// EncodeItem ...
func (s Scaffolding) EncodeItem() (name string, meta int16) {
	return "minecraft:scaffolding", 0
}

// EncodeBlock ...
func (s Scaffolding) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:scaffolding", nil
}
