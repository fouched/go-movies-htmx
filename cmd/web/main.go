package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-movies-htmx/internal/config"
	"github.com/fouched/go-movies-htmx/internal/handlers"
	"log"
	"net/http"
	"time"
)

const port = ":9080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	initApp()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func initApp() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)
}
