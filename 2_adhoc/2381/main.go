package main

import (
	"fmt"
	"sort"
)

func main() {
	var total, pos int
	var current string
	var alunos []string

	fmt.Scanf("%d %d", &total, &pos)
	for total > 0 {
		fmt.Scanf("%s", &current)
		alunos = append(alunos, current)
		total--
	}

	sort.Strings(alunos)
	fmt.Println(alunos[pos-1])
}
