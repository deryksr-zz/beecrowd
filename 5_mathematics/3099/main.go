package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var leitura int
	fmt.Scanf("%d", &leitura)
	var result int
	var analyser string
	for index := 1; index <= leitura; index++ {
		analyser = strconv.Itoa(index)
		result += strings.Count(analyser, "1") + strings.Count(analyser, "7")
	}
	fmt.Printf("%d\n", result)
}
