package item

import (
	"builder/meal/packing"
)

type Drink struct {
	Item
}

func (b *Drink) GetName() string {
	return "饮料"
}

func (b *Drink) GetPrice() float64 {
	return 5.5
}

func (b *Drink) GetPacking() packing.Packing {
	var p packing.Bottle
	return p.Packing
}
