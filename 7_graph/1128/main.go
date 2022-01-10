package main

import (
	"fmt"
)

func addValue(graph map[int][]int, src, dest int) {
	if value, exists := graph[src]; !exists {
		graph[src] = []int{dest}
	} else {
		graph[src] = append(value, dest)
	}
}

func min(first, second int) int {
	if first > second {
		return second
	}
	return first
}

type Node struct {
	LowLink int
	Visited bool
}

func getSCCFromStart(tarjan []Node, graph map[int][]int, start int) (result []int) {
	tarjan[start].LowLink = start
	tarjan[start].Visited = true

	paths := graph[start]
	for _, value := range paths {
		if tarjan[value].Visited == false {
			result = append(result, getSCCFromStart(tarjan, graph, value)...)
		}
		tarjan[start].LowLink = min(tarjan[start].LowLink, tarjan[value].LowLink)
	}
	return append(result, start)
}

func createSCC(tarjan []Node, graph map[int][]int, start int) {
	stack := getSCCFromStart(tarjan, graph, start)
	index := len(stack) - 1
	for index >= 0 {
		current := stack[index]
		paths := graph[current]
		for _, value := range paths {
			tarjan[current].LowLink = min(tarjan[current].LowLink, tarjan[value].LowLink)
		}
		index--
	}
}

func main() {
	var cities, streets int
	for {
		_, err := fmt.Scanf("%d %d", &cities, &streets)
		if cities == 0 && streets == 0 || err != nil {
			break
		}

		graph := make(map[int][]int)
		tarjan := make([]Node, cities)
		for key := range tarjan {
			tarjan[key] = Node{LowLink: -1, Visited: false}
		}

		var src, dest, direction int
		for streets > 0 {
			fmt.Scanf("%d %d %d", &src, &dest, &direction)
			addValue(graph, src-1, dest-1)
			if direction == 2 {
				addValue(graph, dest-1, src-1)
			}
			streets--
		}

		createSCC(tarjan, graph, 0)
		result := 1
		for _, value := range tarjan {
			if value.LowLink == -1 || value.LowLink != 0 {
				result = 0
				break
			}
		}
		fmt.Println(result)
	}
}
