package model

import (
	"github.com/DEBANMC/dragonfly/server/block/cube"
	"github.com/DEBANMC/dragonfly/server/entity/physics"
	"github.com/DEBANMC/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Chain is a model used by chain blocks.
type Chain struct {
	// Axis is the axis which the chain faces.
	Axis cube.Axis
}

// AABB ...
func (c Chain) AABB(cube.Pos, *world.World) []physics.AABB {
	return []physics.AABB{physics.NewAABB(mgl64.Vec3{0.40625, 0.40625, 0.40625}, mgl64.Vec3{0.59375, 0.59375, 0.59375}).Stretch(c.Axis, 0.40625)}
}

// FaceSolid ...
func (Chain) FaceSolid(cube.Pos, cube.Face, *world.World) bool {
	return false
}
