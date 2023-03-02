package handlers

import (
	"errors"
	"fmt"
	"github.com/jcrob2/calc-lib/calc"
	"io"
	"strconv"
)

func BridgeConstructor(w io.Writer, c calc.Calculator) Bridge {
	return Bridge{W: w, C: c}
}

type Handler interface {
	Handle(r []string) error
}

type Bridge struct {
	W io.Writer
	C calc.Calculator
}

func (br Bridge) Handle(r []string) error {
	a, err := strconv.Atoi(r[3])
	if err != nil {
		return fmt.Errorf("%w:%s", invalidArgumentError, err.Error())
	}

	b, err := strconv.Atoi(r[4])
	if err != nil {
		return fmt.Errorf("%w:%s", invalidArgumentError, err.Error())
	}

	opFlag := r[2]

	result := br.C.Calculate(a, b)
	_, err = fmt.Fprintf(br.W, "%d %s %d = %d", a, opFlag, b, result)
	if err != nil {
		return err
	}

	return nil
}

var invalidArgumentError = errors.New("invalid argument")
