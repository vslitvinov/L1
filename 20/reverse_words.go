package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

// reverseWords reverse words
func reverseWords(str string) string {
	words := strings.Fields(str)

	left, right := 0, len(words)-1

	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}


func main() {
	s := "snow dog sun"

	fmt.Println(s)
	fmt.Println(reverseWords(s))
}
