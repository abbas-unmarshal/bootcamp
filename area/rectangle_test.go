package area

import (
	"fmt"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	t.Run("Should return area", func( *testing.T) {
		got := RectangleArea(4,5)
		want := 20
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
	t.Run("Should return area", func( *testing.T) {
		got := RectangleArea(5,6)
		want := 30
		if got != want {
			t.Error(fmt.Sprintf("Expected %v got %v ",want,got))
		}
	})
}