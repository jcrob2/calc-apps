package handlers

import (
	"errors"
	"github.com/jcrob2/calc-lib/calc"
	"testing"
)

func TestInvalidFirstParameter(t *testing.T) {
	calculator := calc.Addition{}
	var handler Handler = BridgeConstructor(nil, calculator)
	testString := []string{"", "a", "1"}
	err := handler.Handle(testString)

	if !errors.Is(err, invalidArgumentError) {
		t.Errorf("Expected non-nil error, but got nil")
	}

	t.Log(err)
}

func TestInvalidSecondParameter(t *testing.T) {
	calculator := calc.Addition{}
	var handler Handler = BridgeConstructor(nil, calculator)
	testString := []string{"", "1", "a"}
	err := handler.Handle(testString)

	if !errors.Is(err, invalidArgumentError) {
		t.Errorf("Expected non-nil error, but got nil")
	}

	t.Log(err)
}

func TestCalculate(t *testing.T) {

	out := FakeWriter{}

	calculator := FakeCalculator{}
	var handler Handler = BridgeConstructor(&out, calculator)

	testString := []string{"", "5", "7"}
	err := handler.Handle(testString)

	test := string(out.output)

	if test != "5 + 7 = 42" {
		t.Error("Calculation incorrect:", test)
	}

	if err != nil {
		t.Error("Unexpected Error", err)
	}

}

func TestWriteError(t *testing.T) {
	calculator := FakeCalculator{}
	writerErr := errors.New("writer error")
	out := FakeWriter{err: writerErr}
	var handler Handler = BridgeConstructor(&out, calculator)

	args := []string{"", "5", "7"}
	err := handler.Handle(args)

	if err != writerErr {
		t.Error("Unexpected Error:", err)
	}
}

type FakeCalculator struct {
	num1 int
	num2 int
}

func (comp FakeCalculator) Calculate(a, b int) int {

	return 42
}

type FakeWriter struct {
	output []byte
	err    error
}

func (w *FakeWriter) Write(p []byte) (int, error) {
	w.output = p
	return len(p), w.err
}
