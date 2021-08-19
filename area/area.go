package area

func RectangleArea(w int, h int) int{
	return w*h
}

func RectanglePerimeter(w int, h int) int{
	return 2*(w+h)
}

func SumPerimeterArea(w,h int) int{
	return RectangleArea(w,h)+RectanglePerimeter(w,h)
}