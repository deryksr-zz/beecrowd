package main

import "fmt"

func main() {
	var total, current int

	fmt.Scanf("%d", &total)

	for i := 1; i <= total; i++ {
		fmt.Scanf("%d", &current)
		fmt.Printf("resposta %d: %d\n", i, current)
	}
}
