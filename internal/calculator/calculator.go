package calculator

import (
	"fmt"
	"strings"
)

// Calculator provides basic arithmetic operations.
type Calculator struct {
	AllowedChars   string
	OperatorTokens string
	PrecedenceMap  map[string]int
}

// NewCalculator creates a new instance of the Calculator.
func New() *Calculator {
	return &Calculator{
		AllowedChars:   "0123456789+-*/() ",
		OperatorTokens: "+-*/()",
		PrecedenceMap: map[string]int{
			"*": 2,
			"/": 2,
			"+": 1,
			"-": 1,
			"(": -1,
			")": -1,
		},
	}
}

// Calculate evaluates the given mathematical expression and returns the result.
func (c *Calculator) Calculate(expression string) (float64, error) {
	for _, char := range expression {
		if !strings.ContainsRune(c.AllowedChars, char) {
			return 0, fmt.Errorf("invalid character '%c' in expression", char)
		}
	}
	infixTokens := c.Tokenize(expression)
	postfixTokens := c.ConvertToPostfix(infixTokens)
	result, err := c.EvaluatePostfix(postfixTokens)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// isOperator checks if a token is an operator.
func (c *Calculator) IsOperator(token string) bool {
	operators := map[string]struct{}{"+": {}, "-": {}, "*": {}, "/": {}}
	_, isOperator := operators[token]
	return isOperator
}

// applyOperator applies the given operator to two operands.
func (c *Calculator) applyOperator(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	}
	return 0
}
