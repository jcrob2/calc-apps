package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
import "github.com/smartystreets/gunit"
import "github.com/smartystreets/assertions/should"

func TestHttpHandler(t *testing.T) {
	gunit.Run(new(HttpHandler), t)
}

type HttpHandler struct {
	*gunit.Fixture
	handler  http.Handler
	response *httptest.ResponseRecorder
	request  *http.Request
	buffer   string
	code     int
}

func (this *HttpHandler) Setup() {
	this.handler = SetupRouter()
	this.response = httptest.NewRecorder()
}

func (this *HttpHandler) TestValid() {
	this.httpGet("/add?a=1&b=3")

	this.So(this.buffer, should.Equal, "Result: 4")
	this.So(this.code, should.Equal, http.StatusOK)
}

func (this *HttpHandler) TestFirstParameter() {
	this.httpGet("/add?a=a&b=3")

	this.So(this.code, should.Equal, http.StatusUnprocessableEntity)
}

func (this *HttpHandler) TestSecondParameter() {
	this.httpGet("/add?a=1&b=b")

	this.So(this.code, should.Equal, http.StatusUnprocessableEntity)
}

func (this *HttpHandler) TestValidOperator() {
	this.httpGet("/ad?a=1&b=3")

	this.So(this.code, should.Equal, http.StatusNotFound)
}

func (this *HttpHandler) httpGet(url string) {
	this.request = httptest.NewRequest(http.MethodGet, url, nil)
	this.handler.ServeHTTP(this.response, this.request)
	this.buffer = this.response.Body.String()
	this.code = this.response.Code
}
