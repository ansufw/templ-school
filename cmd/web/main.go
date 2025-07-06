package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rs/zerolog"
)

type application struct {
	config  config
	log     *zerolog.Logger
	version string
	Session *scs.SessionManager
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := appConfig()

	r := newRouter(app)

	err := app.serve(r)
	if err != nil {
		app.log.Fatal().Err(err).Msg("error servee")
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func (app *application) serve(r http.Handler) error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           r,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.log.Info().Str("env", app.config.env).Int("port", app.config.port).Msg("starting HTTP server")

	return srv.ListenAndServe()
}
