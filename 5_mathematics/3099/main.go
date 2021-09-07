package main

import (
	"fmt"
	"strings"
)

func getNewTight(value int, current int) int {
	if value == current {
		return 1
	}
	return 0
}

func getAnswer(index int, tight int, digits []int) int {
	var limit int
	result := 0

	if index < 0 {
		return 0
	}

	if tight == 0 {
		limit = digits[index]
	} else {
		limit = 9
	}

	for i := 0; i < limit; i++ {

		result += getAnswer(index-1, i, digits)
	}

	return result
}

func strToInt(text string) (result []int) {
	array := strings.Split(text, "")
	for _, value := range array {
		conv := int(value[0] - '0')
		result = append(result, conv)
	}
	return result
}

func main() {
	var input string
	fmt.Scanf("%s", &input)

	digits := strToInt(input)
	answer := getAnswer(len(digits)-1, 0, digits)

	fmt.Println(answer)
}
