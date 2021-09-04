package main

import "fmt"

func fill(matrix [][]bool, segments int) {
	var first, second int
	for index := 0; index < segments; index++ {
		fmt.Scanf("%d %d", &first, &second)
		matrix[first][second] = true
		matrix[second][first] = true
	}
}

func getAnswer(numberLeds int, leds int) string {
	if numberLeds == leds {
		return "COMPLETO"
	} else {
		return "INCOMPLETO"
	}
}

func shouldVisit(value int, visited []int) bool {
	for _, current := range visited {
		if value == current {
			return false
		}
	}
	return true
}

func run(line []bool, matrix [][]bool, visited *[]int) {
	for index := 1; index < len(line); index++ {
		if line[index] && shouldVisit(index, *visited) {
			*visited = append(*visited, index)
			run(matrix[index], matrix, visited)
		}
	}
}

func main() {
	var leds, segments int
	fmt.Scanf("%d %d", &leds, &segments)

	matrix := make([][]bool, leds+1)
	for index := range matrix {
		matrix[index] = make([]bool, leds+1)
	}

	fill(matrix, segments)
	visited := make([]int, 0)

	run(matrix[1], matrix, &visited)
	fmt.Println(getAnswer(len(visited), leds))
}
