package item

import (
	"builder/meal/packing"
)

type Burger struct {
	Item
}

func (b *Burger) GetName() string {
	return "汉堡"
}

func (b *Burger) GetPrice() float64 {
	return 12
}

func (b *Burger) GetPacking() packing.Packing {
	var p packing.Wrapper
	return p.Packing
}
