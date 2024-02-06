### Stack Calculator

This project implements a simple stack-based calculator that supports basic arithmetic operations. It consists of packages for the calculator logic and test suites for ensuring the correctness of the implementation.

### Project Structure

```golang
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   └── calculator
│       ├── calculator.go
│       ├── evaluate.go
│       ├── postfix.go
│       └── tokenize.go
└── tests
    ├── suit1_test.go
    └── suit2_test.go
```

- `cmd`: Contains the main application entry point.
- `internal/calculator`: Implements the calculator logic, including tokenization, conversion to postfix notation, and evaluation of postfix expressions.
- `tests`: Contains test suites for ensuring the correctness of the calculator implementation.

### Usage

To run the calculator application, execute the `main.go` file located in the `cmd` directory. Follow the on-screen instructions to input expressions and get the calculated results.

### How to Run Tests

The project includes two test suites:
1. `suit1_test.go`: Contains test cases for conversion, tokenization, and evaluation of infix expressions.
2. `suit2_test.go`: Contains parametrized test cases for various arithmetic operations and error cases.

To run the tests, use the following command:

```bash
go test ./...
```

### Dependencies

The project uses the following external dependencies:
- `github.com/stretchr/testify/assert`: Provides assertion functions for testing.
- `github.com/stretchr/testify/suite`: Provides support for test suites in Go testing.

