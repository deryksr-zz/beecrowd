package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var times int

	fmt.Scanf("%d", &times)
	regex := regexp.MustCompile(`\w+`)
	reader := bufio.NewReader(os.Stdin)

	for times > 0 {
		line, _ := reader.ReadString('\n')
		words := regex.FindAllString(line, -1)
		result := []rune{}
		for _, current := range words {
			result = append(result, rune(current[0]))
		}
		fmt.Println(string(result))
		times--
	}
}
