package main

import (
	"AT3/network"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userJSON = `{"name": "bobby","description": "a dude", "id": 3}`
var expected = `"name":"Taco","description":"A food item"`
var expected2 = `"name":"Paper Plane","description":"A folded piece of paper, in the shape of a plane"`
var expected3 = `"name":"bobby","description":"a dude"`

func TestGetsamples(t *testing.T) {
	samphandler, e := network.NewSampleHandle()

	url := "/sample"
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// c.SetPath("/sample")
	// c.JSON(http.StatusOK, Devices{"Jhon Doe", "Middle Way"})

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, samphandler.GetSamples(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), expected)
		assert.Contains(t, rec.Body.String(), expected2)
	}
}

func TestGetsample(t *testing.T) {
	samphandler, e := network.NewSampleHandle()

	url := "/sample/2"
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, samphandler.GetSample(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), expected)
	}
}

func TestDeleteSample(t *testing.T) {
	samphandler, e := network.NewSampleHandle()

	url := "/sample/2"
	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, samphandler.DeleteSample(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotContains(t, rec.Body.String(), expected2)
	}
}

func TestCreateSample(t *testing.T) {
	samphandler, e := network.NewSampleHandle()

	url := "/sample"

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, samphandler.CreateSample(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), expected3)
	}
}

func TestUpdateSample(t *testing.T) {
	samphandler, e := network.NewSampleHandle()

	url := "/sample"

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	res := rec.Result()
	defer res.Body.Close()

	//Assume the second string is replace with new string
	if assert.NoError(t, samphandler.UpdateSample(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), expected3)
		assert.NotContains(t, rec.Body.String(), expected2)
	}
}
