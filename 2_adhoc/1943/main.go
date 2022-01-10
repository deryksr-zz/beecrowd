package main

import "fmt"

func getPosition(value int) string {
	switch {
	case value == 1:
		return "Top 1"
	case value > 1 && value < 4:
		return "Top 3"
	case value > 3 && value < 6:
		return "Top 5"
	case value > 5 && value < 11:
		return "Top 10"
	case value > 10 && value < 26:
		return "Top 25"
	case value > 25 && value < 51:
		return "Top 50"
	default:
		return "Top 100"
	}

}

func main() {
	var value int
	fmt.Scanf("%d", &value)
	fmt.Println(getPosition(value))
}
