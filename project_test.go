package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestNewProject(t *testing.T) {
	result := NewProject("testing project")

	if result.Name != "testing project" {
		t.Errorf("project name is %s", result.Name)
	}
}

func TestProjectCollection(t *testing.T) {
	body := bytes.NewBufferString("")
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		panic(err.Error())
	}
	ProjectCollection(req)

}
