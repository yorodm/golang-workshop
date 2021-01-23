package main

import (
	"flag"
	"fmt"
	"stringutil"
)

func main() {
	in := flag.String("text", "!oG ,olleH", "a string")
	flag.Parse()

	fmt.Printf(stringutil.Reverse(*in))
}
