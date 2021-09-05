package main

type node struct {
	operator bool
	function *(func(first, second node) []byte)
	names    []byte
}

func contains(value byte, array []byte) bool {
	for index := range array {
		if array[index] == value {
			return true
		}
	}
	return false
}

func union(first, second node) []byte {
	result := first.names
	for _, value := range second.names {
		if !contains(value, result) {
			result = append(result, value)
		}
	}
	return result
}

func subtr(first, second node) []byte {
	var result []byte
	for _, value := range first.names {
		if !contains(value, second.names) {
			result = append(result, value)
		}
	}
	return result
}

func inner(first, second node) []byte {
	var result []byte
	for _, value := range first.names {
		if contains(value, second.names) {
			result = append(result, value)
		}
	}
	return result
}

func push(stack []byte, value byte) []byte {
	return append(stack, value)
}

func pop(stack []byte) (byte, []byte) {
	if len(stack) != 0 {
		size := len(stack) - 1
		result := stack[size]
		stack = stack[:size]
		return result, stack
	}
	return 0, []byte{}
}

func infixToPostfix(input []byte) []node {
	var postfix []node

	return postfix
}

func main() {

}
