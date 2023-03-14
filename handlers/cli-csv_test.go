package handlers

import (
	"bytes"
	"github.com/jcrob2/calc-lib/calc"
	"strings"
	"testing"
)

func TestCsv(t *testing.T) {
	input := "1,+,2\n" +
		"2,-,1\n" +
		"NaN,+,2\n" +
		"1,+,NaN\n" +
		"1,nop,2\n" +
		"3,*,4\n" +
		"20,/,10\n"

	want := "1,+,2,3\n" +
		"2,-,1,1\n" +
		"3,*,4,12\n" +
		"20,/,10,2\n"

	var calculator calc.Calculator
	i := strings.NewReader(input)
	o := &bytes.Buffer{}

	csvh := CsvHandlerConstructor(i, o, calculator)
	err := csvh.Handle()

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if o.String() != want {
		t.Errorf("Error: ")
	}
}

func TestInvalidFirstParam(t *testing.T) {
	input := "a,+,1"
	var calculator calc.Calculator
	out := CustomWriter{}
	csvh := CsvHandlerConstructor(strings.NewReader(input), out, calculator)
	csvh.Handle()
	if out.output != nil {
		t.Errorf("Invalid first parameter was accepted as valid")
	}
}

func TestInvalidSecondParam(t *testing.T) {
	input := "1,+,a"
	var calculator calc.Calculator
	out := CustomWriter{}
	csvh := CsvHandlerConstructor(strings.NewReader(input), out, calculator)
	csvh.Handle()

	if out.output != nil {
		t.Errorf("Invalid second parameter was accepted as valid")
	}
}

func TestInvalidOp(t *testing.T) {
	input := "1,nop,1"
	var calculator calc.Calculator
	out := CustomWriter{}
	csvh := CsvHandlerConstructor(strings.NewReader(input), out, calculator)
	csvh.Handle()

	if out.output != nil {
		t.Errorf("Invalid operation was accepted as valid")
	}
}

func TestWriteErr(t *testing.T) {

}

func TestReadErr(t *testing.T) {

}

type CustomWriter struct {
	output []byte
	err    error
}

func (this CustomWriter) Write(p []byte) (int, error) {
	this.output = p
	return len(p), this.err
}
