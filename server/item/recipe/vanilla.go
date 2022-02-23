package recipe

import (
	// Insure all blocks and items are registered before trying to load vanilla recipes.
	_ "github.com/df-mc/dragonfly/server/block"
	_ "github.com/df-mc/dragonfly/server/item"

	_ "embed"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

//go:embed crafting_data.nbt
var vanillaCraftingData []byte

func init() {
	var vanillaRecipes struct {
		Shaped []struct {
			Input    inputItems `nbt:"input"`
			Output   outputItem `nbt:"output"`
			Width    int32      `nbt:"width"`
			Height   int32      `nbt:"height"`
			Block    string     `nbt:"block"`
			Priority int32      `nbt:"priority"`
		} `nbt:"shaped"`
		Shapeless []struct {
			Input    inputItems `nbt:"input"`
			Output   outputItem `nbt:"output"`
			Block    string     `nbt:"block"`
			Priority int32      `nbt:"priority"`
		} `nbt:"shapeless"`
	}

	if err := nbt.Unmarshal(vanillaCraftingData, &vanillaRecipes); err != nil {
		panic(err)
	}

	for _, s := range vanillaRecipes.Shapeless {
		input, inputOK := s.Input.Items()
		output, outputOk := s.Output.Stack()
		if !inputOK || !outputOk {
			continue
		}
		Register(&ShapelessRecipe{recipe{
			input:  input,
			output: output,
			block:  s.Block,
		}})
	}

	for _, s := range vanillaRecipes.Shaped {
		input, inputOK := s.Input.Items()
		output, outputOK := s.Output.Stack()
		if !inputOK || !outputOK {
			continue
		}
		Register(&ShapedRecipe{
			Dimensions: Dimensions{int(s.Width), int(s.Height)},
			recipe: recipe{
				input:  input,
				output: output,
				block:  s.Block,
			},
		})
	}
}
