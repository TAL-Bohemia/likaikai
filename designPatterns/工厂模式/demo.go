package main

import (
	"factory/shape"
	"factory/shape2"
)

func main() {
	// 方式1
	s1 := new(shape.Square).GetName()
	println(s1)

	// 方式2
	var square shape2.Shape
	s2 := square.New(1).GetName()
	println(s2)
}
