package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValid(t *testing.T) {

	//creates fake request to pass to router
	request := httptest.NewRequest(http.MethodGet, "/add?a=1&b=3", nil)
	response := httptest.NewRecorder()
	handler := SetupRouter()
	handler.ServeHTTP(response, request)
	code := response.Code

	if code != 200 {
		t.Errorf("Unexpected status code. Expected 200, got %d", code)
	}
}

func TestFirstParamFail(t *testing.T) {

	//creates fake request to pass to router
	request := httptest.NewRequest(http.MethodGet, "/add?a=a&b=3", nil)
	response := httptest.NewRecorder()
	handler := SetupRouter()
	handler.ServeHTTP(response, request)
	code := response.Code

	if code != 422 {
		t.Errorf("Unexpected status code. Expected 422, got %d", code)
	}
}

func TestSecondParamFail(t *testing.T) {

	//creates fake request to pass to router
	request := httptest.NewRequest(http.MethodGet, "/add?a=1&b=b", nil)
	response := httptest.NewRecorder()
	handler := SetupRouter()
	handler.ServeHTTP(response, request)
	code := response.Code

	if code != 422 {
		t.Errorf("Unexpected status code. Expected 422, got %d", code)
	}
}

func TestBadOp(t *testing.T) {

	//creates fake request to pass to router
	request := httptest.NewRequest(http.MethodGet, "/ad?a=1&b=6", nil)
	response := httptest.NewRecorder()
	handler := SetupRouter()
	handler.ServeHTTP(response, request)
	code := response.Code

	if code != 404 {
		t.Errorf("Unexpected status code. Expected 404, got %d", code)
	}
}
