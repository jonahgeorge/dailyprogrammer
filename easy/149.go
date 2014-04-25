// Daily Programmer #149 (Easy)
// April 24, 2014
package main

import (
	"fmt"
	"regexp"
)

func main() {
	inputs := []string{
		"all those who believe in psychokinesis raise my hand",
		"did you hear about the excellent farmer who was outstanding in his field",
	}

	for _, str := range inputs {
		fmt.Printf("%s\n", str)

		r, _ := regexp.Compile("[aeiou\\s]")
		fmt.Printf("%s\n", r.ReplaceAllString(str, ""))

		r, _ = regexp.Compile("[^aeiou]")
		fmt.Printf("%s\n\n", r.ReplaceAllString(str, ""))
	}
}
