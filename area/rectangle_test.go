package area

import (
	"fmt"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	t.Run("Should return 20 as area when width is 4 and height is 5", func( *testing.T) {
		got := RectangleArea(4,5)
		want := 20
		if got != want {
			t.Fail()
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
	t.Run("Should return 60 as area when width is 10 and height is 6", func( *testing.T) {
		got := RectangleArea(10,6)
		want := 60
		if got != want {
			t.Fail()
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
}

func TestRectanglePerimeter(t *testing.T) {
	t.Run("Should return 18 as perimeter when width is 4 and height is 5", func( *testing.T) {
		got := RectanglePerimeter(4,5)
		want := 18
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
	t.Run("Should return 18 as perimeter when width is 4 and height is 5", func( *testing.T) {
		got := RectanglePerimeter(10,5)
		want := 30
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
}

func TestSumPerimeterArea(t *testing.T) {
	t.Run("Expecting 38 for width 4 and height 5", func( *testing.T) {
		got := SumPerimeterArea(4,5)
		want := 38
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
	t.Run("Expecting 80 for width 10 and height 5", func( *testing.T) {
		got := SumPerimeterArea(10,5)
		want := 80
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
}

