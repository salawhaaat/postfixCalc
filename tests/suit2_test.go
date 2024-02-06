package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Parametrized test cases
var TestCases = []struct {
	Name           string
	Expression     string
	ExpectedResult interface{}
	ExpectedError  bool
}{
	{
		Name:           "8.Invalid Calculation",
		Expression:     "5 + @ 3",
		ExpectedResult: 0,
		ExpectedError:  true,
	},
	{
		Name:           "9.Valid Operator Check",
		Expression:     "+",
		ExpectedResult: 0,
		ExpectedError:  true,
	},
	{
		Name:           "10.Invalid Operator Check",
		Expression:     "5",
		ExpectedResult: false,
		ExpectedError:  true,
	},
	{
		Name:           "11.Valid Multiplication",
		Expression:     "5 * 3",
		ExpectedResult: 15.0,
		ExpectedError:  false,
	},
	{
		Name:           "12.Invalid Division by Zero",
		Expression:     "6 / 0",
		ExpectedResult: 0,
		ExpectedError:  true,
	},
}

func (suite *CalculatorTestSuite) TestParametrizedCalc() {
	for _, tc := range TestCases {
		suite.Run(tc.Name, func() {
			switch tc.ExpectedResult.(type) {
			case float64:
				// Test Calculate
				result, err := suite.Calc.Calculate(tc.Expression)
				if tc.ExpectedError {
					assert.Error(suite.T(), err)
				} else {
					assert.NoError(suite.T(), err)
					assert.Equal(suite.T(), tc.ExpectedResult, result)
				}
			case bool:
				// Test isOperator
				result := suite.Calc.IsOperator(tc.Expression)
				assert.Equal(suite.T(), tc.ExpectedResult, result)
			}
		})
	}
}

func TestRunSuite2(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
