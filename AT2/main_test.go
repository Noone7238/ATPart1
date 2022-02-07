package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {

	sh := NewSampleHandle()
	req, err := http.NewRequest("GET", "/sample/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	var xi []Item
	json.Unmarshal(bodyBytes, &xi)
	var failed bool = false
	for _, item := range xi {
		_, ok := sh.items[item.ID]
		if !ok {
			failed = true
		}
	}
	if failed {
		t.Errorf("handler returned unexpected body: got %v want %v",
			xi, sh.items)
	}
}

func TestGetAtID(t *testing.T) {

	sh := NewSampleHandle()
	req, err := http.NewRequest("GET", "/sample/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"name":"Paper Plane","description":"A folded piece of paper, in the shape of a plane","id":"2"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAtBadID(t *testing.T) {

	sh := NewSampleHandle()
	req, err := http.NewRequest("GET", "/sample/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnsupportedMediaType {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnsupportedMediaType)
	}

	// Check the response body is what we expect.
	expected := "ID not found"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestADD(t *testing.T) {

	sh := NewSampleHandle()
	var jsonStr = []byte(`{"name":"sword","description":"weapon used in melee","id":"3"}`)

	req, err := http.NewRequest("POST", "/sample/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ID string = "3"
	_, ok := sh.items[ID]
	if !ok {
		t.Errorf("could not find item in stored items: expected %v stored %v",
			jsonStr, sh.items)
	}
}

func TestEditAtID(t *testing.T) {

	sh := NewSampleHandle()
	var jsonStr = []byte(`{"name":"sword","description":"weapon used in melee","id":"2"}`)

	req, err := http.NewRequest("POST", "/sample/2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ID string = "2"
	_, ok := sh.items[ID]
	if sh.items[ID].Name != "sword" || sh.items[ID].Description != "weapon used in melee" {
		ok = false
	}
	if !ok {
		t.Errorf("could not find item in stored items: expected %v stored %v",
			jsonStr, sh.items)
	}
}

func TestRemoveAtID(t *testing.T) {

	sh := NewSampleHandle()
	var jsonStr = []byte(`{"name":"Paper Plane","description":"A folded piece of paper, in the shape of a plane","id":"2"}`)

	req, err := http.NewRequest("DELETE", "/sample/2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ID string = "2"
	_, ok := sh.items[ID]

	if ok {
		t.Errorf("Item was still found in map, should have been removed")
	}
}

func TestBadRemoveAtID(t *testing.T) {

	sh := NewSampleHandle()
	var jsonStr = []byte(`{"name":"Paper Plane","description":"A folded piece of paper, in the shape of a plane","id":"5"}`)

	req, err := http.NewRequest("DELETE", "/sample/5", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.selector)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnsupportedMediaType {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnsupportedMediaType)
	}

	var ID string = "5"
	_, ok := sh.items[ID]

	if ok {
		t.Errorf("Somehow you found something not supposed to be in the map.")
	}
}
