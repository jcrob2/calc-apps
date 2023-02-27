package main

import (
	"github.com/jcrob2/calc-lib/calc"
	handler2 "github.com/jcrob2/calc-lib/calc_handler"
	"log"
	"os"
)

func main() {

	var calculator calc.Calculator = calc.Addition{}

	o := os.Stdout
	i := os.Args

	var handler handler2.Handler = handler2.BridgeConstructor(o, calculator)
	err := handler.Handle(i)

	if err != nil {
		log.Fatal(err)
	}
}
