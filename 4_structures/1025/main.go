package main

import (
	"fmt"
)

func clean(array *[10001]int, values map[int]bool) {
	for k, _ := range values {
		array[k] = 0
	}
}

func main() {
	var queries, marables, current, value, maxValue int
	var storage map[int]bool
	var data [10001]int

	value = 1
	for {
		_, err := fmt.Scanf("%d %d", &marables, &queries)
		if err != nil || (marables == 0 && queries == 0) {
			break
		}

		maxValue = 0
		storage = make(map[int]bool)

		for marables > 0 {
			fmt.Scanf("%d", &current)
			if current > maxValue {
				maxValue = current
			}
			if _, ok := storage[current]; !ok {
				storage[current] = true
			}
			data[current]++
			marables--
		}

		fmt.Printf("CASE# %d:\n", value)

		for queries > 0 {
			fmt.Scanf("%d", &current)

			if data[current] != 0 {
				position := 1
				for i := 0; i < current; i++ {
					position += data[i]
				}
				fmt.Printf("%d found at %d\n", current, position)
			} else {
				fmt.Printf("%d not found\n", current)
			}
			queries--
		}

		clean(&data, storage)
		value++
	}
}
