package tools

import (
	"io/ioutil"
	"os"
)

// Calculate a Fibonacci number
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func BytesInFile(filename string) (int, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	s, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, err
	}
	return len(s), nil
}
