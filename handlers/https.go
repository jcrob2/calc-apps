package handlers

import (
	"fmt"
	"github.com/jcrob2/calc-lib/calc"
	"net/http"
	"strconv"
)

type HttpsHandler struct {
	Calculator calc.Calculator
}

func (this HttpsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	a, err := strconv.Atoi(req.URL.Query().Get("a"))
	if err != nil {
		//res.Header()
		res.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(res, "Error: %s", err)
	}
	b, err := strconv.Atoi(req.URL.Query().Get("b"))
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(res, "Error: %d", err)
	}
	result := this.Calculator.Calculate(a, b)
	fmt.Fprintf(res, "Result: %d", result)
}

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/add", HttpsHandler{Calculator: calc.Addition{}})
	router.Handle("/sub", HttpsHandler{Calculator: calc.Subtraction{}})
	router.Handle("/mult", HttpsHandler{Calculator: calc.Multiplication{}})
	router.Handle("/div", HttpsHandler{Calculator: calc.Division{}})

	return router
}
