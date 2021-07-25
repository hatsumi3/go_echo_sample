package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON       = `{"name":"Jon Snow","email":"jon@labstack.com"}`
	userJSONError  = `{"name":12345,"email":"jon@labstack.com"}`
	invalidRequest = `{"Code":100,"Message":"invalid request"}`
)

func Test_hello(t *testing.T) {
	// Setup
	e := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

func Test_show(t *testing.T) {
	// Setup
	e := NewRouter()

	q := make(url.Values)
	q.Set("name", "Jon Snow")
	q.Set("email", "jon@labstack.com")

	req := httptest.NewRequest(http.MethodGet, "/user?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, show(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON+"\n", rec.Body.String())
	}

}

func Test_show_error(t *testing.T) {
	// Setup
	e := NewRouter()

	// binding errpr
	q := make(url.Values)
	q.Set("email", "Jon Snow")
	q.Set("email", "jon@labstack.com")

	req := httptest.NewRequest(http.MethodGet, "/user?"+q.Encode(), strings.NewReader(userJSONError))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.Error(t, show(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, invalidRequest+"\n", rec.Body.String())
	}

}

func Test_display(t *testing.T) {
	// Setup
	e := NewRouter()

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, display(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON+"\n", rec.Body.String())
	}
}

func Test_display_error(t *testing.T) {
	// Setup
	e := NewRouter()

	// binding errpr
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJSONError))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.Error(t, display(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, invalidRequest+"\n", rec.Body.String())
	}
}
