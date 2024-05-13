package main

import (
	"fmt"
	"net/http"

	"github.com/sudhir512kj/go-hello-web/pkg/config"
	"github.com/sudhir512kj/go-hello-web/pkg/handlers"
	"github.com/sudhir512kj/go-hello-web/pkg/render"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
