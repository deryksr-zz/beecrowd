package main

import (
	"fmt"
	"sort"
)

type Road struct {
	Src  int
	Dest int
	Cost int
}

type Graph []Road

func (g Graph) Len() int {
	return len(g)
}

func (g Graph) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Graph) Less(i, j int) bool {
	return g[i].Cost < g[j].Cost
}

func FindParentRoot(dsu []int, value int) int {
	for dsu[value] != -1 {
		value = dsu[value]
	}
	return value
}

func hasConnection(dsu []int, first, second int) bool {
	firstParent := FindParentRoot(dsu, first)
	secondParent := FindParentRoot(dsu, second)
	return firstParent == secondParent
}

func Merge(dsu []int, first, second int) {
	firstParent := FindParentRoot(dsu, first)
	secondParent := FindParentRoot(dsu, second)
	dsu[firstParent] = secondParent
	discovered++
}

var discovered int

func main() {
	var junctions, roads int

	for {
		fmt.Scanf("%d %d", &junctions, &roads)
		if junctions == 0 && roads == 0 {
			break
		}

		graph := make(Graph, 0)
		total := 0
		for roads > 0 {
			temp := Road{}
			fmt.Scanf("%d %d %d", &temp.Src, &temp.Dest, &temp.Cost)
			graph = append(graph, temp)
			total += temp.Cost
			roads--
		}

		sort.Sort(graph)
		dsu := make([]int, junctions)
		for key := range dsu {
			dsu[key] = -1
		}
		saving := 0
		discovered = 0
		for _, value := range graph {
			if !hasConnection(dsu, value.Src, value.Dest) {
				Merge(dsu, value.Src, value.Dest)
				saving += value.Cost
				if discovered == junctions-1 {
					break
				}
			}
		}
		fmt.Println(total - saving)
	}

}
