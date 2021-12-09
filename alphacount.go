package piscine

import (
	"fmt"
)

func AlphaCount(s string) int {
	count := 0
	for i := range s {
		if i >= 65 && i <= 90 || i >= 97 && i <= 122 {
			count = count + 1
		}
	}
	return count
}
