package main

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word1 := strings.ToLower(word)
	j := len(word1) - 1
	for i := range word1 {
		if !unicode.IsLetter(rune(word1[i])) {
			continue
		}
		for j = range word1 {
			if i == j {
				break
			}
			if word1[i] == word1[j] {
				return false
			}
		}
		j--
	}
	return true
}
