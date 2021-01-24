package stringutil

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Reverse("Hello, world")
        }
}

func ExampleReverse() {
    fmt.Println(Reverse("Hello, world"))
    // Output: dlrow ,olleH
}
