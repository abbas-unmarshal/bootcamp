package area

import (
	"fmt"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	t.Run("Should return 20 as area when width is 4 and height is 5", func(*testing.T) {
		rectangle := Rectangle{width: 4, height: 5}
		got := rectangle.Area()
		want := 20
		if got != want {
			t.Fail()
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
	t.Run("Should return 60 as area when width is 10 and height is 6", func(*testing.T) {
		rectangle := Rectangle{width: 10, height: 6}
		got := rectangle.Area()
		want := 60
		if got != want {
			t.Fail()
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
}

func TestRectanglePerimeter(t *testing.T) {
	t.Run("Should return 18 as perimeter when width is 4 and height is 5", func(*testing.T) {
		rectangle := Rectangle{width: 4, height: 5}
		got := rectangle.Perimeter()
		want := 18
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
	t.Run("Should return 30 as perimeter when width is 10 and height is 5", func(*testing.T) {
		rectangle := Rectangle{width: 10, height: 5}
		got := rectangle.Perimeter()
		want := 30
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
}

func TestSumPerimeterArea(t *testing.T) {
	t.Run("Expecting 38 for width 4 and height 5", func(*testing.T) {
		rectangle := Rectangle{width: 4, height: 5}
		got := rectangle.SumPerimeterArea()
		want := 38
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
	t.Run("Expecting 80 for width 10 and height 5", func(*testing.T) {
		rectangle := Rectangle{width: 10, height: 5}
		got := rectangle.SumPerimeterArea()
		want := 80
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ", want, got))
		}
	})
}
