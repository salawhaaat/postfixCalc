package calculator

import (
	"fmt"
	"strconv"
)

// EvaluatePostfix evaluates the postfix expression and returns the result.
func (c *Calculator) EvaluatePostfix(postfixTokens []string) (float64, error) {
	operandStack := []float64{}
	for _, token := range postfixTokens {
		if c.IsOperator(token) {
			if len(operandStack) < 2 {
				return 0, fmt.Errorf("invalid expression: insufficient operands for operator %s", token)
			}
			secondOperand := operandStack[len(operandStack)-1]
			firstOperand := operandStack[len(operandStack)-2]
			result := c.applyOperator(firstOperand, secondOperand, token)
			operandStack = operandStack[:len(operandStack)-2]
			operandStack = append(operandStack, result)
		} else {
			if val, err := strconv.ParseFloat(token, 64); err == nil {
				operandStack = append(operandStack, val)
			}
		}
	}
	if len(operandStack) != 1 {
		return 0, fmt.Errorf("invalid expression: insufficient operators or too many operators")
	}
	return operandStack[0], nil
}
