package main

import (
	"breakfastAndBedWebApp/pkg/config"
	"breakfastAndBedWebApp/pkg/handlers"
	"breakfastAndBedWebApp/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2" // library used for session management
)

const hostPort = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// creating template cache once to be used by the whole application
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// setting the template cache to be used by the application globally
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	// setting the config to be used inside render package
	render.NewTemplates(&app)

	// Starting the application
	fmt.Printf("Starting application on: %s \n", hostPort)

	srv := &http.Server{
		Addr:    hostPort,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
