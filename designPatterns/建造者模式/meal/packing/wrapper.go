package packing

type Wrapper struct {
	Packing
}

func (w *Wrapper) pack() string {
	return "Wrapper"
}