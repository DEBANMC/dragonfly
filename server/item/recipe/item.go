package recipe

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"math"
)

// InputItem is a recipe item. It is an item stack inherently, but it also has an extra value for if it applies to all types.
type InputItem struct {
	item.Stack
	// Variants is true if the item applies to all of its possible variants. This is so it can be rendered in the recipe book
	// for alternative types.
	Variants bool
}

// inputItemData contains data for an item that is inputted to a crafting menu.
type inputItemData struct {
	// Name is the name of the item being inputted.
	Name string `nbt:"name"`
	// MetadataValue is the meta of the item. This can change the item almost completely, or act as durability.
	MetadataValue int32 `nbt:"meta"`
	// Count is the amount of the item.
	Count int32 `nbt:"count"`
}

// inputItem converts an inputItemData to an InputItem.
func (i inputItemData) inputItem() (InputItem, bool) {
	it, ok := world.ItemByName(i.Name, int16(i.MetadataValue))
	if !ok {
		return InputItem{}, false
	}

	return InputItem{Stack: item.NewStack(it, int(i.Count)), Variants: i.MetadataValue == math.MaxInt16}, true
}

// inputItems is an array of input items.
type inputItems []inputItemData

// Items converts inputItems into an array of input items.
func (i inputItems) Items() (s []InputItem, ok bool) {
	for _, it := range i {
		st, ok := it.inputItem()
		if !ok {
			return nil, false
		}
		s = append(s, st)
	}
	return s, true
}

// outputItem is an item that is outputted after crafting.
type outputItem struct {
	// Name is the name of the item being output.
	Name string `nbt:"name"`
	// MetadataValue is the meta of the item. This can change the item almost completely, or act as durability.
	MetadataValue int32 `nbt:"meta"`
	// Count is the amount of the item.
	Count int16 `nbt:"count"`
	// State is included if the output is a block. If it's not included, the meta can be discarded and the output item can be incorrect.
	State struct {
		Name       string                 `nbt:"name"`
		Properties map[string]interface{} `nbt:"states"`
		Version    int32                  `nbt:"version"`
	} `nbt:"block"`
	// NBTData contains extra NBTData which may modify the item in other, more discreet ways.
	NBTData map[string]interface{} `nbt:"data"`
}

// Stack converts an output item to an item stack.
func (o outputItem) Stack() (item.Stack, bool) {
	var it world.Item
	var ok bool
	if o.State.Version != 0 {
		if b, ok := world.BlockByName(o.State.Name, o.State.Properties); ok {
			if it, ok = b.(world.Item); !ok {
				return item.Stack{}, false
			}
		}
	} else {
		it, ok = world.ItemByName(o.Name, int16(o.MetadataValue))
		if !ok {
			return item.Stack{}, false
		}
	}

	return item.NewStack(it, int(o.Count)), true
}
