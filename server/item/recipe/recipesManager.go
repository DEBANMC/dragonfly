package recipe

//g:embed crafting_data.nbt
//var vanillaCraftingData []byte

type vanillaRecipes struct {
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
