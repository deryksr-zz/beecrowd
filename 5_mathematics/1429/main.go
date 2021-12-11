package main

import (
	"fmt"
)

func main() {
	factorial := []int{
		1,
		2,
		6,
		24,
		120,
	}
	var value string

	for {
		fmt.Scanf("%s", &value)
		result := 0
		if value == "0" {
			break
		}

		index := len(value) - 1
		aux := 0
		for index >= 0 {
			result += (int(value[aux]) - 48) * factorial[index]
			index--
			aux++
		}
		fmt.Println(result)
	}

}
