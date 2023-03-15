package main

import (
	"github.com/jcrob2/calc-apps/handlers"
	"github.com/jcrob2/calc-lib/calc"
	"os"
)

func main() {

	println("Starting...")
	var calculator calc.Calculator

	//TODO: this is the only way I got this to actually read form the args. There has to be a more generic solution
	//in := os.Args[1]
	//i := strings.NewReader(in)

	csvh := handlers.CsvHandlerConstructor(os.Stdin, os.Stdout, calculator)

	println("\nOutput")
	err := csvh.Handle()

	if err != nil {

	}

}
