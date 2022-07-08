package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSampleAPI(t *testing.T) {
	app := NewServer()
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req);
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t,  string(body), "Hello, World 👋!")
}