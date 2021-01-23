package math

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var in int = 10
	var want int = 55

	got := Fibonacci(in)
	if got != want {
		t.Errorf("Fibonacci(%d) == %d, want %d", in, got, want)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(10)
	}
}

func ExampleFibonacci() {
	fmt.Println(Fibonacci(10))
	// Output: 55
}
