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

func hasConnectionToVisit(neighbors []int, tarjan []Node) bool {
	for _, value := range neighbors {
		if tarjan[value].Visited == false {
			return true
		}
	}
	return false
}

func continueFill(tarjan []Node, graph map[int][]int, start int) {
	tarjan[start].LowLink = start
	tarjan[start].Visited = true

	paths := graph[start]
	for _, value := range paths {
		if tarjan[value].Visited == false {
			continueFill(tarjan, graph, value)
		}
		tarjan[start].LowLink = min(tarjan[start].LowLink, tarjan[value].LowLink)
	}
}

func getSCCFromStart(tarjan []Node, graph map[int][]int, start int) {
	tarjan[start].LowLink = start
	tarjan[start].Visited = true

	paths := graph[start]
	for _, value := range paths {
		if tarjan[value].Visited == false {
			getSCCFromStart(tarjan, graph, value)
		} else if hasConnectionToVisit(graph[value], tarjan) {
			continueFill(tarjan, graph, value)
		}
		tarjan[start].LowLink = min(tarjan[start].LowLink, tarjan[value].LowLink)
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

		getSCCFromStart(tarjan, graph, 0)
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
