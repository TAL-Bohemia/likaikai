package main

import (
	"fmt"

	"builder/meal"
	"builder/meal/item"
)

func main() {
	var m meal.Meal

	m.AddItem(new(item.Burger))
	m.AddItem(new(item.Drink))
	m.AddItem(new(item.Burger))

	total := m.GetTotal()

	items := m.GetItems()

	fmt.Printf("总价：%f\n", total)
	fmt.Printf("所有食物：%s", items)
}
