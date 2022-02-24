package item

// Totem protects the player from death
type Totem struct {
}

// EncodeItem ...
func (a Totem) EncodeItem() (name string, meta int16) {
	return "minecraft:totem_of_undying", 0
}

func (Totem) MaxCount() int {
	return 1
}
