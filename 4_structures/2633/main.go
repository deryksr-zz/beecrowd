package main

import (
	"fmt"
	"sort"
	"strings"
)

type pair struct {
	name  string
	value int
}

func getAnswer(group []pair) string {
	var aux []string
	for _, value := range group {
		aux = append(aux, value.name)
	}
	return strings.Join(aux, " ")
}

func main() {
	var storage []pair
	var size, value int
	var name string
	for {
		_, err := fmt.Scanf("%d", &size)
		if err != nil {
			break
		}

		for i := 0; i < size; i++ {
			fmt.Scanf("%s %d", &name, &value)
			storage = append(storage, pair{
				name:  name,
				value: value,
			})
		}

		sort.Slice(storage, func(i, j int) bool {
			return storage[i].value < storage[j].value
		})
		fmt.Println(getAnswer(storage))
		storage = nil
	}

}
