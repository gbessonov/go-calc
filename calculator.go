package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Evaluate the expression tree
func (n *Node) Evaluate() float64 {
	if n.Left == nil && n.Right == nil {
		val, _ := strconv.ParseFloat(n.Value, 64)
		return val
	}

	leftVal := n.Left.Evaluate()
	rightVal := n.Right.Evaluate()

	switch n.Value {
	case "+":
		return leftVal + rightVal
	case "-":
		return leftVal - rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		return leftVal / rightVal
	case "**":
		return math.Pow(leftVal, rightVal)
	default:
		panic("Unknown operator: " + n.Value)
	}
}

// Construct an expression tree from a postfix expression
func BuildExpressionTree(tokens []string) *Node {
	stack := []*Node{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "**":
			if len(stack) < 2 {
				panic("Invalid expression")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &Node{Value: token, Left: left, Right: right})
		default:
			stack = append(stack, &Node{Value: token})
		}
	}

	if len(stack) != 1 {
		panic("Invalid expression")
	}

	return stack[0]
}

// Convert infix expression to postfix notation
func InfixToPostfix(expression string) []string {
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "**": 3, "(": 0}
	var output []string
	var stack []string

	// Tokenize: Spaces around operators and brackets
	re := regexp.MustCompile(`(\*\*|[\+\-\*/\(\)])`)
	expression = re.ReplaceAllString(expression, " $1 ")
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // Remove '('
		case "+", "-", "*", "/", "**":
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		default:
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output
}

func main() {
	fmt.Println("Enter an arithmetic expression (e.g., (3 + 5) * 2, 2**3):")
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	postfix := InfixToPostfix(expression)
	exprTree := BuildExpressionTree(postfix)

	fmt.Println("Result:", exprTree.Evaluate())
}
