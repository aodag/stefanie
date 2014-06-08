package main

import (
	"github.com/go-martini/martini"
)


func Greeting() string {
	return "Hello world!"
}

func main() {
	m := martini.Classic()
	m.MapTo(DummyProjectRepository{}, (*ProjectRepository)(nil))
	m.Get("/", Greeting)
	m.Get("/projects", ProjectCollection)
	m.Run()
}
