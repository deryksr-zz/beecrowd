package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var lines, maxSpaces int
	var current string
	var result []string

	reader := bufio.NewReader(os.Stdin)

	for {
		value, _ := reader.ReadString('\n')
		lines, _ = strconv.Atoi(value[:len(value)-1])
		if lines == 0 {
			break
		}

		maxSpaces = 0
		if result != nil {
			fmt.Println()
		}

		result = nil

		for lines > 0 {
			current, _ = reader.ReadString('\n')
			words := regexp.MustCompile(`\w+`).FindAllString(current, -1)
			current = strings.Join(words, " ")

			if len(current) > maxSpaces {
				maxSpaces = len(current)
			}

			result = append(result, current)
			lines--
		}

		for i := 0; i < len(result); i++ {
			spaces := strings.Repeat(" ", maxSpaces-len(result[i]))
			fmt.Println(spaces + result[i])
		}
	}

}
