package main

import "fmt"

func main() {
	var cases, current, previous, answer int
	var results []int

	for {
		fmt.Scanf("%d", &cases)

		if cases == 0 {
			break
		}

		answer = 0
		previous = 0
		for cases > 0 {
			fmt.Scanf("%d", &current)

			if previous == 0 || current-previous > 10 {
				answer += 10
			} else {
				answer += current - previous
			}

			previous = current
			cases--
		}
		results = append(results, answer)
	}

	for _, value := range results {
		fmt.Println(value)
	}
}
