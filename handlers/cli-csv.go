package handlers

import (
	"encoding/csv"
	"github.com/jcrob2/calc-lib/calc"
	"io"
	"strconv"
)

type Equation struct {
	A  int
	B  int
	Op string
}

func EquationConstructor(a, b int, op string) Equation { return Equation{A: a, B: b, Op: op} }

type CsvHandler struct {
	R *csv.Reader
	W *csv.Writer
	C calc.Calculator
}

func CsvHandlerConstructor(r io.Reader, w io.Writer, c calc.Calculator) CsvHandler {
	return CsvHandler{R: csv.NewReader(r), W: csv.NewWriter(w), C: c}
}

func (ch CsvHandler) Handle(r io.Reader) error {
	defer ch.W.Flush()
	var err error

	for {
		record, err := ch.R.Read()
		if err == io.EOF {
			break
		}
		arg1, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		arg2, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}
		op := record[1]

		var calculator calc.Calculator

		addCalc := calc.Addition{}
		subCalc := calc.Subtraction{}
		multCalc := calc.Multiplication{}
		divCalc := calc.Division{}

		switch op {
		case "+":
			calculator = addCalc
		case "-":
			calculator = subCalc
		case "*":
			calculator = multCalc
		case "/":
			calculator = divCalc
		default:
			continue
		}

		result := calculator.Calculate(arg1, arg2)

		record = append(record, strconv.Itoa(result))

		ch.W.Write(record)
	}
	return err
}
