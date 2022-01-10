package main

import (
	"fmt"
	"strings"
)

func splitWord(word []string, value string) (left, right []string) {
	var idx int
	for idx = 0; idx < len(word); idx++ {
		if word[idx] == value {
			break
		}
	}
	return word[:idx], word[idx+1:]
}

func getNext(preOrder []string, substring []string) string {
	for _, current := range preOrder {
		for _, value := range substring {
			if current == value {
				return value
			}
		}
	}
	return ""
}

func remove(value string, word []string) []string {
	for idx := 0; idx < len(word); idx++ {
		if word[idx] == value {
			return append(word[:idx], word[idx+1:]...)
		}
	}
	return word
}

func printPostOrder(start string, preOrder []string, inOrder []string) {
	preOrder = remove(start, preOrder)
	if len(inOrder) == 0 {
		return
	}
	left, right := splitWord(inOrder, start)
	nextLeft := getNext(preOrder, left)
	nextRight := getNext(preOrder, right)
	printPostOrder(nextLeft, preOrder, left)
	printPostOrder(nextRight, preOrder, right)
	fmt.Printf(start)
}

func main() {
	var preOrder, inOrder string
	for {
		_, err := fmt.Scanf("%s %s", &preOrder, &inOrder)
		if err != nil {
			break
		}

		pre := strings.Split(preOrder, "")
		in := strings.Split(inOrder, "")

		printPostOrder(pre[0], pre, in)
		fmt.Println()

	}

}
