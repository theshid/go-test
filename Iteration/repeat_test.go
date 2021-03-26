package Iteration

import (

	"testing"
)

func TestRepeat(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("iteration 6 times", func(t *testing.T) {
		repeated := Repeat("a",6)
		expected := "aaaaaa"
		/*if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}*/
		assertCorrectMessage(t,repeated,expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}