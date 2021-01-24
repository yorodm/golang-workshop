package tools

import (
	"fmt"
	"path/filepath"
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

func TestBytesInFile(t *testing.T) {
	l, err := BytesInFile(filepath.Join("testdata", "hello_world"))
	if err != nil {
		t.Errorf("Error %s", err)
		t.FailNow()
	}
	if l != 12 { // New line at end of file
		t.Errorf("Got wrong number of bytes %d", l)
	}
}
