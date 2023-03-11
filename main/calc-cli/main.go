package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jcrob2/calc-apps/handlers"
	"github.com/jcrob2/calc-lib/calc"
	"log"
	"os"
)

func main() {
	var opFlag string
	flag.StringVar(&opFlag, "op", "+", "Sets which calculator we want to use")
	flag.Parse()

	var calculator calc.Calculator

	switch opFlag {
	case "+":
		calculator = calc.Addition{}
	case "-":
		calculator = calc.Subtraction{}
	case "*":
		calculator = calc.Multiplication{}
	case "/":
		calculator = calc.Division{}
	default:
		err := fmt.Errorf("%w: %s", unsupportedOperand, opFlag)
		log.Fatal(err)
	}

	handler := handlers.BridgeConstructor(os.Stdout, calculator)
	err := handler.Handle(flag.Args())

	if err != nil {
		log.Fatal(err)
	}
}

var unsupportedOperand = errors.New("unsupported operand")
