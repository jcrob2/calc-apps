package main

import (
	"errors"
	"flag"
	"fmt"
	handler2 "github.com/jcrob2/calc-apps/handlers"
	"github.com/jcrob2/calc-lib/calc"
	"log"
	"os"
)

func main() {

	opFlag := flag.String("op", "+", "Sets which calculator we want to use")
	flag.Parse()

	var calculator calc.Calculator

	switch *opFlag {
	case "+":
		calculator = calc.Addition{}
	case "-":
		calculator = calc.Subtraction{}
	case "*":
		calculator = calc.Multiplication{}
	case "/":
		calculator = calc.Division{}
	default:
		err := fmt.Errorf("%w: %s", unsupportedOperand, *opFlag)
		log.Fatal(err)
	}

	o := os.Stdout
	i := os.Args

	var handler handler2.Handler = handler2.BridgeConstructor(o, calculator)
	err := handler.Handle(i)

	if err != nil {
		log.Fatal(err)
	}
}

var unsupportedOperand = errors.New("unsupported operand")
