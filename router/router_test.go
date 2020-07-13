package router

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("filename", "resources/input_routes.csv")
}

func makeRequest(method, url string, body io.Reader) (error, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return err, nil
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)
	handler.ServeHTTP(rr, req)

	return nil, rr
}

func TestGetAll(t *testing.T) {
	err, rr := makeRequest("GET", "/api/routes", nil)
	if err != nil {
		t.Fatal(err)
	}

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
	err, rr := makeRequest("POST", "/api/routes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFindBestRoute(t *testing.T) {
	values := []byte(`"origin": {"GRU"}, "destiny": {"CDG"}}`)
	err, rr := makeRequest("GET", "/api/bestroutes", bytes.NewBuffer(values))
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
