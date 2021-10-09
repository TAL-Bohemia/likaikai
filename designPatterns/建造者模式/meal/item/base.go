package item

import (
	"builder/meal/packing"
)

type Item interface {
	GetName() string
	GetPrice() float64
	GetPacking() packing.Packing
}
