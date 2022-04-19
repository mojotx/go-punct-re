# Go Punctuation Demo

This code will print all of the printable "punctuation" ASCII characters in the range of 0-127. 

Note that it uses the `[:punct:]` character class, as described in the section, "ASCII character classes,"
in the [Google RE2 Syntax Documentation](https://github.com/google/re2/wiki/Syntax).

```go
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
```