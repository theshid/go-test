package Recursion

import "testing"

func TestRecursion(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("reversing a string", func(t *testing.T) {
		got := Reverse("rachid")
		want := "dihcar"
		assertCorrectMessage(t, got, want)
	})

}
