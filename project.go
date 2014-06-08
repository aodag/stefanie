package main

import (
	"net/http"
)

type Project struct {
	Name string
}

func NewProject(name string) *Project {
	project := new(Project)
	project.Name = name
	return project
}

func ProjectCollection(req *http.Request) string {
	return ""
}
