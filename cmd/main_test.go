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
		"exspression": "2*2",
	}
	body, _ := json.Marshal(mcPostBody)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/calculate/", bytes.NewReader(body))

	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	RequestHandler(rr, req)
	res := rr.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code, expected: %d", res.StatusCode)
	}

	if string(data) != "" {
		t.Errorf("Expression is empty. You must supply an exspression but got %v", string(data))
	}
}
