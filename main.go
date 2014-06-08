package main

import (
	"github.com/go-martini/martini"
	"os"
)


type Application struct {
	DatabaseURL string
}

func Greeting() string {
	return "Hello world!"
}


func (app *Application) MakeWebApp() *martini.ClassicMartini {
	dbmap := InitDB(app.DatabaseURL)
	repo := NewGorpProjectRepository(dbmap)

	m := martini.Classic()
	m.MapTo(repo,
		(*ProjectRepository)(nil))
	m.Get("/", Greeting)
	m.Get("/projects", ProjectCollection)
	m.Post("/projects", ProjectCollection)
	return m
}

func NewApplication(databaseURL string) *Application {
	app := new(Application)
	app.DatabaseURL = databaseURL
	return app
}

func main() {
	app := NewApplication(os.Getenv("DATABASE_URL"))
	m := app.MakeWebApp()
	m.Run()
}
