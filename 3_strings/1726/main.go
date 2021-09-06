package main

import "fmt"

type OperatorFunc func(first, second element) []byte

type element struct {
	operator   bool
	bracket    bool
	precedence int
	function   *OperatorFunc
	names      []byte
}

type node struct {
	data  element
	left  *node
	right *node
}

func contains(value byte, array []byte) bool {
	for index := range array {
		if array[index] == value {
			return true
		}
	}
	return false
}

func union(first, second element) []byte {
	result := first.names
	for _, value := range second.names {
		if !contains(value, result) {
			result = append(result, value)
		}
	}
	return result
}

func subtr(first, second element) []byte {
	var result []byte
	for _, value := range first.names {
		if !contains(value, second.names) {
			result = append(result, value)
		}
	}
	return result
}

func inner(first, second element) []byte {
	var result []byte
	for _, value := range first.names {
		if contains(value, second.names) {
			result = append(result, value)
		}
	}
	return result
}

func getPrecedence(value byte) int {
	if value == '*' {
		return 2
	}
	return 1
}

func getFunction(value byte) *OperatorFunc {
	var result OperatorFunc
	switch value {
	case '*':
		result = inner
	case '-':
		result = subtr
	case '+':
		result = union
	}
	return &result
}

func push(stack []element, value element) []element {
	return append(stack, value)
}

func pop(stack []element) (element, []element) {
	if len(stack) != 0 {
		size := len(stack) - 1
		result := stack[size]
		stack = stack[:size]
		return result, stack
	}
	return element{}, []element{}
}

func infixToPostfix(input []byte) []element {
	var postfix []element
	var operators []element
	var value element

	for index := 0; index < len(input); index++ {
		if input[index] == '(' {
			operators = push(operators, element{
				function: nil,
				operator: true,
				bracket:  true,
			})
		} else if input[index] == ')' {
			for {
				if input[index] != ')' {
					break
				}
				value, operators = pop(operators)
				postfix = push(postfix, value)
			}
		} else if input[index] == '{' {
			index++
			value = element{
				operator: false,
				bracket:  false,
				function: nil,
			}
			for {
				if input[index] == '}' {
					break
				}
				value.names = append(value.names, input[index])
				index++
			}
			if len(value.names) > 0 {
				postfix = push(postfix, value)
			}
		} else {
			for len(operators) > 0 && (getPrecedence(input[index]) <= operators[len(operators)-1].precedence) {
				value, operators = pop(operators)
				postfix = push(postfix, value)
			}
			operators = push(operators, element{
				operator:   true,
				bracket:    false,
				function:   getFunction(input[index]),
				precedence: getPrecedence(input[index]),
			})
		}
	}

	for len(operators) > 0 {
		value, operators = pop(operators)
		postfix = push(postfix, value)
	}

	return postfix
}

func createBET(postfix []element) *node {

	return nil
}

func calculate(root node) string {

	return ""
}

func main() {
	text := "{ABCD}-{CZ}"
	postfix := infixToPostfix([]byte(text))
	fmt.Println(len(postfix))
}
