package main

import (
	"github.com/go-martini/martini"
	"os"
)


func Greeting() string {
	return "Hello world!"
}


func MakeApp() *martini.ClassicMartini {
	dbmap := InitDB(os.Getenv("DATABASE_URL"))
	m := martini.Classic()
	repo := NewGorpProjectRepository(dbmap)
	//repo.Data = make(map[string] *Project)

	m.MapTo(repo,
		(*ProjectRepository)(nil))
	m.Get("/", Greeting)
	m.Get("/projects", ProjectCollection)
	m.Post("/projects", ProjectCollection)
	return m
}
func main() {
	m := MakeApp()
	m.Run()
}
