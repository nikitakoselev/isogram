package isogram

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	m := map[int32]bool{}

	fmt.Printf("Word: %s\n", word)

	for _, c := range word {
		if c == '-' || c == ' ' {
			continue
		}
		fmt.Printf("Character %c ", c)
		fmt.Printf("m[%d] == %t\n", c, m[c])
		if m[c] == true {
			return false
		} else {
			m[c] = true
		}
	}
	return true
}
