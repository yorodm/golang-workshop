package math

import "testing"

const in int = 12345
const want int = 280

func TestSquaresPlusCubes(t *testing.T) {
	got := SquaresPlusCubes(in)
	if got != want {
		t.Errorf("SquaresPlusCubes(%d) == %d, want %d", in, got, want)
	}
}

func BenchmarkSquaresPlusCubes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SquaresPlusCubes(in)
	}
}

func TestSquaresPlusCubesMT(t *testing.T) {
	c := make(chan int)
	go SquaresPlusCubesMT(in, c)
	got := <-c

	if got != want {
		t.Errorf("SquaresPlusCubesMT(%d) == %d, want %d", in, got, want)
	}
}

func BenchmarkSquaresPlusCubesMT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := make(chan int)
		go SquaresPlusCubesMT(in, c)
		got := <-c

		if got != want {
			b.Errorf("SquaresPlusCubesMT(%d) == %d, want %d", in, got, want)
		}
	}
}
