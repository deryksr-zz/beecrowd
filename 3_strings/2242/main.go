package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(value []string) bool {
	start := 0
	end := len(value) - 1
	for start < end {
		if value[start] != value[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	var input string
	fmt.Scanf("%s", &input)

	words := regexp.MustCompile(`[^aeiou]+`).Split(input, -1)
	result := []string{}
	for _, value := range words {
		result = append(result, strings.Split(value, "")...)
	}
	if isPalindrome(result) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
