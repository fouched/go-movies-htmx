package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-movies-htmx/internal/config"
	"github.com/fouched/go-movies-htmx/internal/handlers"
	"github.com/fouched/go-movies-htmx/internal/helpers"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"log"
	"net/http"
	"time"
)

const port = ":9080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	dbPool, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	// we have database connectivity, close it after app stops
	defer dbPool.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Printf("Starting application on %s\n", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func initApp() (*sql.DB, error) {
	// read from command line, the flag, default value and some help text
	flag.StringVar(&app.DSN,
		"dsn",
		"host=localhost port=5432 user=fouche password=javac dbname=movies sslmode=disable timezone=UTC connect_timeout=5",
		"Database connection string")
	flag.StringVar(&app.APIKey, "api-key", "2824634c24070c272d34b330e82cba7c", "api key")
	flag.Parse()

	dbPool, err := repo.CreateDbPool(app.DSN)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	} else {
		log.Println("Connected to database!")
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return dbPool, nil
}
