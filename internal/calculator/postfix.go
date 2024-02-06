package calculator

// ConvertToPostfix converts the infix expression to postfix notation.
func (c *Calculator) ConvertToPostfix(infixTokens []string) []string {
	var postfixTokens []string
	operatorStack := []string{"("}
	infixTokens = append(infixTokens, ")")
	for _, token := range infixTokens {
		if _, isOperator := c.PrecedenceMap[token]; !isOperator {
			postfixTokens = append(postfixTokens, token)
		} else if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for operatorStack[len(operatorStack)-1] != "(" {
				postfixTokens = append(postfixTokens, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else {
			for len(operatorStack) > 0 && c.PrecedenceMap[token] <= c.PrecedenceMap[operatorStack[len(operatorStack)-1]] {
				postfixTokens = append(postfixTokens, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		}
	}
	//fmt.Println(postfixTokens) // Debug print
	return postfixTokens
}
