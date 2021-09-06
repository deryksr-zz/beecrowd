package main

import (
	"fmt"
	"sort"
	"strings"
)

type OperatorFunc func(first, second element) []byte

type element struct {
	operator   bool
	bracket    bool
	precedence int
	function   *OperatorFunc
	names      []byte
	debug      string
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
				if operators[len(operators)-1].bracket {
					break
				}
				value, operators = pop(operators)
				postfix = push(postfix, value)
			}
			_, operators = pop(operators)
		} else if input[index] == '{' {
			index++
			value = element{
				operator: false,
				bracket:  false,
				function: nil,
				names:    []byte{},
			}
			for {
				if input[index] == '}' {
					break
				}
				value.names = append(value.names, input[index])
				index++
			}
			postfix = push(postfix, value)
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

func calculate(postfix []element) element {
	var stack []element
	var left, right element

	for len(postfix) > 0 {
		value := postfix[0]
		postfix = postfix[1:]

		if value.operator {
			left, stack = pop(stack)
			right, stack = pop(stack)
			stack = append(stack, element{
				operator:   false,
				bracket:    false,
				precedence: 0,
				names:      (*value.function)(right, left),
			})
		} else {
			stack = append(stack, value)
		}
	}
	return stack[0]
}

func getAnswer(text string) string {
	postfix := infixToPostfix([]byte(text))
	answer := calculate(postfix)

	result := strings.Split(string(answer.names), "")
	sort.Strings(result)

	return "{" + strings.Join(result, "") + "}"
}

func main() {
	var input string
	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		fmt.Println(getAnswer(input))
	}
}
