package main

import "fmt"

func hashFunction(value int, base int) int {
	return value % base
}

func main() {
	var storage map[int][]int
	var testCases, keys, base, current int

	fmt.Scanf("%d", &testCases)
	keys = 0
	for testCases > 0 {
		storage = make(map[int][]int, 0)

		if keys != 0 {
			fmt.Println()
		}

		fmt.Scanf("%d %d", &base, &keys)
		for i := 0; i < keys; i++ {
			fmt.Scanf("%d", &current)
			pos := hashFunction(current, base)
			list, _ := storage[pos]
			storage[pos] = append(list, current)
		}

		for aux := 0; aux < base; aux++ {
			if list, ok := storage[aux]; !ok {
				fmt.Printf("%d -> \\", aux)
			} else {
				fmt.Printf("%d", aux)
				for _, elem := range list {
					fmt.Printf(" -> %d", elem)
				}
				fmt.Print(" -> \\")
			}
			fmt.Println()
		}
		testCases--
	}
}
