package main

import "prototype/shape"

func main() {
	var s shape.Shape
	// s:= &shape.Shape{}

	s.SetType(1)
	s.SetName("圆")

	s2 := s.Clone()
	s2.SetType(2)

	s3 := s.Clone()
	s3.SetName("长方形")

	s.Show()
	s2.Show()
	s3.Show()

}
