package main

import "fmt"

func getResult(start string, final string) int {
	result := 0
	for index := range start {
		letter := start[index]
		for {
			if letter == final[index] {
				break
			} else if letter == 'z' {
				letter = byte('a')
			} else {
				letter++
			}
			result++
		}
	}
	return result
}

func main() {
	var start, final string
	var totalCases int

	fmt.Scanf("%d", &totalCases)
	for {
		if totalCases == 0 {
			break
		}
		fmt.Scanf("%s %s", &start, &final)
		fmt.Println(getResult(start, final))
		totalCases--
	}
}
