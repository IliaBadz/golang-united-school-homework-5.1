package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	end := Point{s.start.x + int(s.a), s.start.y + int(s.a)}

	return end
}

func (s Square) Area() uint {
	area := s.a * s.a
	return area
}

func (s Square) Perimeter() uint {
	perimeter := s.a * 4
	return perimeter
}
