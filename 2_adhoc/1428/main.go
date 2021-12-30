package main

import "fmt"

func main() {
	var total, column, row int
	var first, second int
	var result []int

	fmt.Scanf("%d", &total)

	for total > 0 {
		fmt.Scanf("%d %d", &column, &row)

		first = (column / 3)
		second = (row / 3)

		if row-second*3 > 2 {
			second++
		}
		if column-first*3 > 2 {
			first++
		}

		result = append(result, first*second)
		total--
	}

	for _, value := range result {
		fmt.Println(value)
	}

}
