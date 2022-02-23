package block_internal

//lint:file-ignore ST1022 Exported variables in this package have compiler directives. These variables are not otherwise exposed to users.

import (
	_ "unsafe" // Imported for compiler directives.

	"github.com/DEBANMC/dragonfly/server/world"
	"github.com/DEBANMC/dragonfly/server/world/particle"
)

//go:linkname world_breakParticle github.com/DEBANMC/dragonfly/server/world.breakParticle
//noinspection ALL
var world_breakParticle func(b world.Block) world.Particle

func init() {
	world_breakParticle = func(b world.Block) world.Particle {
		return particle.BlockBreak{Block: b}
	}
}
