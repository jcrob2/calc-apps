package main

import (
	"github.com/jcrob2/calc-apps/handlers"
	"github.com/jcrob2/calc-lib/calc"
	"os"
	"strings"
)

func main() {
	var calculator calc.Calculator
	in := os.Args[1]

	//TODO: this is the only way I got this to actually read form the args. There has to be a more generic solution
	i := strings.NewReader(in)

	csvh := handlers.CsvHandlerConstructor(i, os.Stdout, calculator)

	println("\nOutput")
	err := csvh.Handle(os.Stdin)

	if err != nil {

	}

}
