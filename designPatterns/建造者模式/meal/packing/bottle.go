package packing

type Bottle struct {
	Packing
}

func (w *Bottle) pack() string {
	return "Bottle"
}
