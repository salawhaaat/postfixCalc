package main

import (
	"bufio"
	"fmt"
	"os"
	"stackCalc/internal/calculator"
	"strings"
)

func main() {
	for {
		fmt.Print("Enter expression (or type 'exit' to quit): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		trimmedInput := strings.TrimSpace(input)
		if strings.ToLower(trimmedInput) == "exit" {
			break
		}

		calc := calculator.New()
		result, err := calc.Calculate(trimmedInput)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Result: %.2f\n", result)
		}
	}
}
