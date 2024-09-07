package main

import (
	"breakfastAndBedWebApp/pkg/config"
	"breakfastAndBedWebApp/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"            // library used for route handling
	"github.com/go-chi/chi/middleware" // library used for using middleware
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/awesome", handlers.Repo.Awesome)

	// setting up a file server to load images or other static files into the page
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
