package main

import (
	"fmt"
)

func main() {
	var total, people int
	var last, current, result string
	var flag bool
	fmt.Scanf("%d", &total)

	for total > 0 {
		fmt.Scanf("%d", &people)
		result = ""
		flag = false
		last = ""
		for people > 0 {
			fmt.Scanf("%s", &current)
			if flag && current != last && result != "ingles" {
				result = "ingles"
			} else if !flag {
				flag = true
				last = current
				result = current
			}
			people--
		}
		fmt.Println(result)
		total--
	}
}
