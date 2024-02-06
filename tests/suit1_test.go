package tests

import (
	"stackCalc/internal/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	Calc *calculator.Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.Calc = calculator.New()
}

func (suite *CalculatorTestSuite) TearDownTest() {
	// Clean up any test-specific resources here
}

// Parametrized test cases
var conversionTestCases = []struct {
	Name            string
	InfixExpression string
	PostfixResult   []string
	Tokenize        []string
	EvaluateResult  float64
	ExpectedError   bool
}{
	{
		Name:            "1.Valid Infix to Postfix Conversion",
		InfixExpression: "5 + 3 * (7 - 2)",
		PostfixResult:   []string{"5", "3", "7", "2", "-", "*", "+"},
		Tokenize:        []string{"5", "+", "3", "*", "(", "7", "-", "2", ")"},
		EvaluateResult:  20.0,
		ExpectedError:   false,
	},
	{
		Name:            "2.Empty Infix Expression",
		InfixExpression: "",
		PostfixResult:   nil,
		Tokenize:        nil,
		EvaluateResult:  0.0,
		ExpectedError:   true,
	},
	{
		Name:            "3.Infix Expression with Invalid Characters",
		InfixExpression: "3 + @ 7",
		PostfixResult:   []string{"3", "@7", "+"},
		Tokenize:        []string{"3", "+", "@7"},
		EvaluateResult:  0.0,
		ExpectedError:   true,
	},
	{
		Name:            "4.Valid Postfix Evaluation",
		InfixExpression: "5 + 3 * (7 - 2)",
		PostfixResult:   []string{"5", "3", "7", "2", "-", "*", "+"},
		Tokenize:        []string{"5", "+", "3", "*", "(", "7", "-", "2", ")"},
		EvaluateResult:  20.0,
		ExpectedError:   false,
	},
	{
		Name:            "5.Invalid Postfix Expression",
		InfixExpression: "5 + @ 3",
		PostfixResult:   []string{"5", "@3", "+"},
		Tokenize:        []string{"5", "+", "@3"},
		EvaluateResult:  0.0,
		ExpectedError:   true,
	},
	{
		Name:            "6.Tokenize Empty Expression",
		InfixExpression: "",
		PostfixResult:   nil,
		Tokenize:        nil,
		EvaluateResult:  0.0,
		ExpectedError:   true,
	},
	{
		Name:            "7.Tokenize Expression with Spaces",
		InfixExpression: "   5   +   3  * 2 ",
		PostfixResult:   []string{"5", "3", "2", "*", "+"},
		Tokenize:        []string{"5", "+", "3", "*", "2"},
		EvaluateResult:  11.0,
		ExpectedError:   false,
	},
}

func (suite *CalculatorTestSuite) TestConversionCases() {
	for _, tc := range conversionTestCases {
		suite.Run(tc.Name, func() {
			// Test Infix to Postfix Conversion
			actualPostfix := suite.Calc.ConvertToPostfix(suite.Calc.Tokenize(tc.InfixExpression))

			assert.Equal(suite.T(), tc.PostfixResult, actualPostfix)

			// Test Postfix Evaluation
			actualResult, err := suite.Calc.EvaluatePostfix(actualPostfix)
			if tc.ExpectedError {
				assert.Error(suite.T(), err)
			} else {
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), tc.EvaluateResult, actualResult)
			}

			// Test Tokenize
			actualTokens := suite.Calc.Tokenize(tc.InfixExpression)
			assert.Equal(suite.T(), tc.Tokenize, actualTokens)
		})
	}
}

func TestRunSuite1(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
