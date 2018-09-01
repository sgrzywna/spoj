package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
// Stack implementation
//

// Stack top node.
type Stack struct {
	top    *node
	length int
}

type node struct {
	value rune
	prev  *node
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Clear removes all elements from the stack
func (stack *Stack) Clear() {
	stack.top = nil
	stack.length = 0
}

// Len returns the number of items in the stack
func (stack *Stack) Len() int {
	return stack.length
}

// Peek views the top item on the stack
func (stack *Stack) Peek() rune {
	if stack.length == 0 {
		return 0
	}
	return stack.top.value
}

// Pop the top item of the stack and return it
func (stack *Stack) Pop() rune {
	if stack.length == 0 {
		return 0
	}

	n := stack.top
	stack.top = n.prev
	stack.length--
	return n.value
}

// Push a value onto the top of the stack
func (stack *Stack) Push(value rune) {
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}

//
// Parser implementation
//

// Parser implements shunting-yard algorithm
type Parser struct {
	stack  *Stack
	output string
}

// NewParser creates parser
func NewParser() *Parser {
	return &Parser{NewStack(), ""}
}

// Reset resets parser state.
func (parser *Parser) Reset() {
	parser.stack.Clear()
	parser.output = ""
}

func (parser *Parser) isOperand(token rune) bool {
	return token >= 'a' && token <= 'z'
}

func (parser *Parser) getOperatorPriority(operator rune) int {
	// +, -, *, /, ^
	priority := 0
	switch operator {
	case '+':
		priority = 0
	case '-':
		priority = 1
	case '*':
		priority = 2
	case '/':
		priority = 3
	case '^':
		priority = 4
	}
	return priority
}

func (parser *Parser) parse(input string) string {
	parser.Reset()
	for _, token := range input {
		if parser.isOperand(token) {
			parser.output = parser.output + string(token)
		} else if token == '(' {
			parser.stack.Push(token)
		} else if token == ')' {
			for parser.stack.Peek() != '(' {
				parser.output = parser.output + string(parser.stack.Pop())
			}
			parser.stack.Pop()
		} else {
			precedence := parser.getOperatorPriority(token)
			for precedence < parser.getOperatorPriority(parser.stack.Peek()) {
				parser.output = parser.output + string(parser.stack.Pop())
			}
			parser.stack.Push(token)
		}
	}
	for parser.stack.Len() > 0 {
		parser.output = parser.output + string(parser.stack.Pop())
	}
	return parser.output
}

func main() {
	parser := NewParser()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	for i := 0; i < t; i++ {
		scanner.Scan()
		expression := strings.TrimSpace(scanner.Text())
		res := parser.parse(expression)
		fmt.Fprintln(os.Stdout, res)
	}
}
