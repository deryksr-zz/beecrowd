package main

import "fmt"

func main() {
	var c, n int

	fmt.Scanf("%d %d", &c, &n)
	turns := c / n
	fmt.Println(c - turns*n)
}
