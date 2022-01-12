package main

import (
	"fmt"
	"strings"
)

func getAnswer(value string) string {
	if len(value) != 3 {
		return value
	}
	switch {
	case value[0] == 'U' && value[1] == 'R':
		return "URI"
	case value[0] == 'O' && value[1] == 'B':
		return "OBI"
	default:
		return value
	}
}

func main() {
	var qtd int
	var word string
	fmt.Scanf("%d", &qtd)
	result := []string{}

	for qtd > 0 {
		fmt.Scanf("%s", &word)
		result = append(result, getAnswer(word))
		qtd--
	}
	fmt.Println(strings.Join(result, " "))
}
