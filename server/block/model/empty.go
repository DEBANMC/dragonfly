package model

import (
	"github.com/DEBANMC/dragonfly/server/block/cube"
	"github.com/DEBANMC/dragonfly/server/entity/physics"
	"github.com/DEBANMC/dragonfly/server/world"
)

// Empty is a model that is completely empty. It has no collision boxes or solid faces.
type Empty struct{}

// AABB returns an empty slice.
func (Empty) AABB(cube.Pos, *world.World) []physics.AABB {
	return nil
}

// FaceSolid always returns false.
func (Empty) FaceSolid(cube.Pos, cube.Face, *world.World) bool {
	return false
}
