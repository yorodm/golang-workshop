package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Run only one of the below functions
	cmdArgs()
	cmdFlags()
	cmdEnv()
}

func cmdArgs() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println("With program name: ", argsWithProg)
	fmt.Println("Just args:", argsWithoutProg)
	fmt.Println("3rd Argument:", arg)
}

func cmdFlags() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("tail:", flag.Args())
}

func cmdEnv() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "->", pair[1])
	}
}
