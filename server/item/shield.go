package item

// Shield protects the player from damage
type Shield struct {
}

// EncodeItem ...
func (Shield) EncodeItem() (name string, meta int16) {
	return "minecraft:shield", 0
}

func (Shield) MaxCount() int {
	return 1
}
