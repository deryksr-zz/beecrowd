package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var total int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &total)
	for total > 0 {
		line, _ := reader.ReadString('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		result := make([]string, 0)

		for _, value := range strings.Split(line, " ") {
			result = append(result, value)
		}

		sort.SliceStable(result, func(i, j int) bool {
			return len(result[i]) > len(result[j])
		})

		fmt.Println(strings.Join(result, " "))
		total--
	}

}
