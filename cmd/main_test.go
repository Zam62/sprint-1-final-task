package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandlerSuccessCase(t *testing.T) {
	mcPostBody := map[string]interface{}{
		"exspression": "2*s2",
	}
	body, _ := json.Marshal(mcPostBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calculate/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}

	if string(data) != "" {
		t.Errorf("Expression is empty. You must supply an exspression but got %v", string(data))
	}
}

// func TestRequestHandlerEmptyExpressionCase(t *testing.T) {
// 	expected := "Expression is empty. You must supply an exspression"
// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate?expression=", nil)
// 	w := httptest.NewRecorder()
// 	RequestHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	if res.StatusCode != http.StatusBadRequest {
// 		t.Errorf("wrong status code")
// 	}

// 	if string(data) != expected {
// 		t.Errorf("Expression is empty. You must supply an exspression but got %v", string(data))
// 	}
// }

// func TestRequestHandlerBadRequestCase(t *testing.T) {
// 	expected := "Bad request"
// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate?expression=/%", nil)
// 	w := httptest.NewRecorder()
// 	RequestHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	if res.StatusCode != http.StatusBadRequest {
// 		t.Errorf("wrong status code")
// 	}

// 	if string(data) != expected {
// 		t.Errorf("Expression is empty. You must supply an exspression but got %v", string(data))
// 	}
// }

// func TestRequestHandlerWrongPortCase(t *testing.T) {
// 	expected := "Bad request"
// 	req := httptest.NewRequest(http.MethodGet, "localhost://1111", nil)
// 	w := httptest.NewRecorder()
// 	RequestHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	if res.StatusCode != http.StatusBadRequest {
// 		t.Errorf("wrong status code")
// 	}

// 	if string(data) != expected {
// 		t.Errorf("Expression is empty. You must supply an exspression but got %v", string(data))
// 	}
// }

// // Test middleware
// // func TestCalculateMiddlewareSuccessCase(t *testing.T) {
// // 	calculateHandler := func(w http.ResponseWriter, r *http.Request) {
// // 		r.WithContext(context.WithValue(r.Context(), "something", "hello"))
// // 	}

// // 	req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
// // 	res := httptest.NewRecorder()
// // 	calculateHandler(res, req)

// // 	tim := calculateMiddleware(calculateHandler)
// // 	tim.ServeHTTP(res, req)

// // 	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
// // }
