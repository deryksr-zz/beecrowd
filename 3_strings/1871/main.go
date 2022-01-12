package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var first, second int
	rxp := regexp.MustCompile(`[^0]+`)

	for {
		fmt.Scanf("%d %d", &first, &second)
		if first == 0 && second == 0 {
			break
		}
		sum := first + second
		numbers := strings.Join(rxp.FindAllString(strconv.Itoa(sum), -1), "")
		fmt.Println(numbers)
	}

}
