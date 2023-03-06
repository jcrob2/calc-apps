package main

import (
	"github.com/jcrob2/calc-apps/handlers"
	"github.com/jcrob2/calc-lib/calc"
	"os"
)

func main() {
	var calculator calc.Calculator
	o := os.Stdout
	i := os.Args

	var csvh handlers.CsvHandler = handlers.CsvHandlerConstructor(o, calculator)

	err := csvh.Handle(i)

	if err != nil {

	}

}
