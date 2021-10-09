package meal

import "builder/meal/item"

type Meal struct {
	items []string
	total float64
}

func (m *Meal) AddItem(item item.Item) {
	m.items = append(m.items, item.GetName())
	m.total += item.GetPrice()
}

func (m *Meal) GetTotal() float64 {
	return m.total
}

func (m *Meal) GetItems() []string {
	return m.items
}
