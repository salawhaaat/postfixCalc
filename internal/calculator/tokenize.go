package calculator

import "strings"

// Tokenize breaks the expression into a slice of tokens.
func (c *Calculator) Tokenize(expression string) []string {
	var infixTokens []string
	currentToken := ""
	for _, char := range expression {
		if char == ' ' || char == '\n' {
			continue
		} else if strings.ContainsRune(c.OperatorTokens, char) {
			if currentToken != "" {
				infixTokens = append(infixTokens, currentToken)
			}
			currentToken = ""
			infixTokens = append(infixTokens, string(char))
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		infixTokens = append(infixTokens, currentToken)
	}
	// fmt.Println(infixTokens) // Debug print
	return infixTokens
}
