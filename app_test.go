package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
)


func TestApp(t *testing.T) {
	m := MakeApp()
	ts := httptest.NewServer(m)
	defer ts.Close()
	http.Get(ts.URL)
	_, err := http.PostForm(ts.URL + "/projects",
		url.Values{"name": {"testing-project"}})
	if err != nil {
		t.Errorf("error with posting project: %s", err.Error())
	}
}
