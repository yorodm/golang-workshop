package main

import "fmt"

func main() {
	fmt.Printf(Reverse("!oG ,olleH"))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = ChangeVowelCase(r[j]), ChangeVowelCase(r[i])
	}
	return string(r)
}

func ChangeVowelCase(s rune) rune {
	var ret string = string(s)
	
	switch ret {
		case "A":
			ret = "a"
		case "E":
			ret = "e"
		case "I":
			ret = "i"
		case "O":
			ret = "o"
		case "U":
			ret = "u"
		case "a":
			ret = "A"
		case "e":
			ret = "E"
		case "i":
			ret = "I"
		case "o":
			ret = "O"
		case "u":
			ret = "U"
	}
		
	return []rune(ret)[0]
}
