package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	re := regexp.MustCompile(`[[:punct:]]`)

	for i := rune(0); i <= 127; i++ {
		s := string(i)
		if strconv.IsPrint(i) && re.MatchString(s) {
			fmt.Printf("%d: '%s'\n", i, string(i))
		}
	}
}
