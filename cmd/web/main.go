package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/supWRLD/bookings/pckg/config"
	"github.com/supWRLD/bookings/pckg/handlers"
	"github.com/supWRLD/bookings/pckg/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

const port = ":8080"

func main() {
	// is it in production
	app.InPoduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InPoduction

	app.Session = session

	app.TemplateCache, _ = render.CreateTemplateCache()
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Routes
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	serve := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("starting application on port", port)

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	// getting requests from web
	// _ = http.ListenAndServe(port, nil)
}
