package main

import (
	"fmt"
)

type Edge struct {
	Dest int
	Cost int
}

func exists(current int, visited map[int]bool) bool {
	_, ok := visited[current]
	return ok
}

func getMinimalRoads(roads []Edge, founded map[int]bool, discovered map[int]int) {
	for _, edge := range roads {
		if edge.Cost != 0 && !exists(edge.Dest, founded) {
			if value, ok := discovered[edge.Dest]; !ok {
				discovered[edge.Dest] = edge.Cost
			} else if value > edge.Cost {
				discovered[edge.Dest] = edge.Cost
			}
		}
	}
}

func main() {
	var junctions, roads, newCost, newPlace int
	var src, dest, cost, min int
	var startSrc, startDest int

	for {
		fmt.Scanf("%d %d", &junctions, &roads)
		if junctions == 0 && roads == 0 {
			break
		}

		graph := make(map[int][]Edge)
		total := 0

		for roads > 0 {
			fmt.Scanf("%d %d %d", &src, &dest, &cost)
			graph[src] = append(graph[src], Edge{Dest: dest, Cost: cost})
			graph[dest] = append(graph[dest], Edge{Dest: src, Cost: cost})
			if total == 0 || min > cost {
				min = cost
				startSrc = src
				startDest = dest
			}
			total += cost
			roads--

		}

		founded := map[int]bool{
			startSrc:  true,
			startDest: true,
		}

		discovered := make(map[int]int)
		getMinimalRoads(graph[startDest], founded, discovered)
		getMinimalRoads(graph[startSrc], founded, discovered)

		for len(founded) < junctions {
			newPlace = -1
			newCost = 0

			for key, value := range discovered {
				if newPlace == -1 || newCost > value {
					newPlace = key
					newCost = value
				}
			}

			delete(discovered, newPlace)
			founded[newPlace] = true
			min += newCost
			getMinimalRoads(graph[newPlace], founded, discovered)
		}

		fmt.Println(total - min)
	}

}
