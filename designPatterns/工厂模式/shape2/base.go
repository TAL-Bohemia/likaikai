package shape2

type ShapeI interface {
	GetName() string
}

type Shape struct {
}

func (s *Shape) New(shapeType int) ShapeI {
	switch shapeType {
	case 1:
		return &Square{}
	case 2:
		return &Rectangle{}
	}

	return nil
}
