package main

import (
	"github.com/go-martini/martini"
)


func Greeting() string {
	return "Hello world!"
}


func MakeApp() *martini.ClassicMartini {
	m := martini.Classic()
	repo := new(DummyProjectRepository)
	repo.Data = make(map[string] *Project)

	m.MapTo(repo,
		(*ProjectRepository)(nil))
	m.Get("/", Greeting)
	m.Get("/projects", ProjectCollection)
	return m
}
func main() {
	m := MakeApp()
	m.Run()
}
