package main

import (
	"fmt"
)

func main() {
	var first, second int
	fmt.Scanf("%d", &first)
	fmt.Scanf("%d", &second)
	fmt.Println(first % second)
}
