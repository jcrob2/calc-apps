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
	err := csvh.Handle(i)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if o.String() != want {

	}
}
