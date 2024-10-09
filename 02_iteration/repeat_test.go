package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat given string multiple times with the given numner", func(t *testing.T) {
		repeated := "hahahaha"
		expected := Repeat("ha", 4)

		if repeated != expected && strings.Count(expected, "ha") == 4 {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
	t.Run("works with strings that have spaces", func(t *testing.T) {
		repeated := "ha ha ha ha ha ha "
		expected := Repeat("ha ", 6)

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	repeated := Repeat("1", 6)
	fmt.Println(repeated)
	// Output: 111111

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("ha", 4)
	}
}
