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
	body := bytes.NewBufferString("name=testing-project")
	req, err := http.NewRequest("POST", "/", body)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	repo := new(DummyProjectRepository)
	repo.Data = make(map[string] *Project)
	if err != nil {
		panic(err.Error())
	}
	res := ProjectCollection(req, repo)
	if res != "testing-project" {
		t.Errorf("res body: %s", res)
	}
	result := repo.Data["testing-project"]

	if result == nil {
		t.Errorf("project is not registered")
	}

}
