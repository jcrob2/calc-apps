package handlers

import (
	"fmt"
	"github.com/jcrob2/calc-lib/calc"
	"io"
	"strconv"
	"strings"
)

type Equation struct {
	A  int
	B  int
	Op string
}

func EquationConstructor(a, b int, op string) Equation { return Equation{A: a, B: b, Op: op} }

type Equations struct {
	Eqns []Equation
}

func EquationsConstructor(s []Equation) Equations { return Equations{Eqns: s} }

type CsvHandler struct {
	W io.Writer
	C calc.Calculator
}

func CsvHandlerConstructor(w io.Writer, c calc.Calculator) CsvHandler {
	return CsvHandler{W: w, C: c}
}

func (ch CsvHandler) Handle(r []string) error {
	equations, err := ch.parseCsv(r)
	if err != nil {

	}

	for _, x := range equations.Eqns {
		switch x.Op {
		case "+":
			ch.C = calc.Addition{}
		case "-":
			ch.C = calc.Subtraction{}
		case "*":
			ch.C = calc.Multiplication{}
		case "/":
			ch.C = calc.Division{}
		default:
			continue
		}

		result := ch.C.Calculate(x.A, x.B)
		_, err = fmt.Fprintf(ch.W, "%d,%s,%d,%d\n", x.A, x.Op, x.B, result)

	}
	return nil
}

func (ch CsvHandler) parseCsv(r []string) (Equations, error) {
	var equations []Equation

	temp := strings.Split(r[1], "\n")

	for _, nl := range temp {
		parse := strings.Split(nl, ",")

		a, err := strconv.Atoi(parse[0])
		if err != nil {
			continue
		}

		op := parse[1]

		b, err := strconv.Atoi(parse[2])
		if err != nil {
			continue
		}

		equation := EquationConstructor(a, b, op)
		equations = append(equations, equation)
	}

	eqns := EquationsConstructor(equations)
	return eqns, nil
}
