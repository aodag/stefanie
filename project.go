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

func ProjectCollection(req *http.Request, repo ProjectRepository) string {
	name := req.PostFormValue("name")
	project := NewProject(name)
	repo.Add(project)
	return project.Name
}

type ProjectRepository interface {
	Add(*Project)
}

type DummyProjectRepository struct {
	Data map[string] *Project
}

func (this *DummyProjectRepository) Add(project *Project) {
	this.Data[project.Name] = project
}
