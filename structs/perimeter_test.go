package structs

import "testing"

/*
Type float64 for floating point nums (0.2345)
%.2f => format string for float64 with 2 decimal places
%g => format string for more precise decimal number
*/

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%g' '%.g' want", got, want)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
