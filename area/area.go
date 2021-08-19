package area

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

func (r Rectangle) SumPerimeterArea() int {
	return r.Area() + r.Perimeter()
}
