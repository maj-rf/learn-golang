package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word exists in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("word doesn't exist in dictionary", func(t *testing.T) {
		_, got := dictionary.Search("testing")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q error want %q", got, want)
	}
}
