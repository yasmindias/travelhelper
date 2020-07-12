package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/routes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"Origin":"GRU","Destiny":"BRC","Cost":10},` +
		`{"Origin":"BRC","Destiny":"SCL","Cost":5},` +
		`{"Origin":"GRU","Destiny":"CDG","Cost":75},` +
		`{"Origin":"GRU","Destiny":"SCL","Cost":20},` +
		`{"Origin":"GRU","Destiny":"ORL","Cost":56},` +
		`{"Origin":"ORL","Destiny":"CDG","Cost":5},` +
		`{"Origin":"SCL","Destiny":"ORL","Cost":20}]`

	if rr.Body.String() != expected {
		t.Errorf("request returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreate(t *testing.T) {
	body := []byte(`{"Origin":"SCL","Destiny":"BRC","Cost":"15"}`)
	req, err := http.NewRequest("POST", "/api/routes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
