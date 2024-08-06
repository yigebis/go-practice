package main

import (
	"fmt"
	"strings"
)

func freqCounter(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Split(s, " ")

	for _, word := range words {
		lower := strings.ToLower(word)
		_, exists := count[lower]
		if exists {
			count[lower] += 1
		} else {
			count[lower] = 1
		}
	}

	return count
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	// fmt.Print(freqCounter("lorem ipsum lorem lorem ipsum i IPsum Lorem I"))

	fmt.Println(isPalindrome("abeba"))
	fmt.Println(isPalindrome("alemu"))
}
