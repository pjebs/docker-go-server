package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// TestHandler will check if the output of the server is "hello world"
func TestHandler(t *testing.T) {

	w := httptest.NewRecorder()

	handler(w, nil)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("could not interpret response")
	}

	desiredResult := "hello world"

	if desiredResult != string(body) {
		t.Fatalf("response not correct. response=%s; required=%s", body, desiredResult)
	}

}
