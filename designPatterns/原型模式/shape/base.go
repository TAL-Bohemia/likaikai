package shape

import "strconv"

type Shape struct {
	shapeType int
	name      string
}

func (s *Shape) SetName(name string) {
	s.name = name
}

func (s *Shape) getName() string {
	return s.name
}

func (s *Shape) SetType(shapeType int) {
	s.shapeType = shapeType
}

func (s *Shape) getType() int {
	return s.shapeType
}

func (s *Shape) Show() {
	println("name:" + s.name + "\nshapeType:" + strconv.Itoa(s.shapeType) + "\n")
}

func (s *Shape) Clone() *Shape {
	shape := *s
	return &shape
}
