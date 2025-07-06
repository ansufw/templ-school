package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type config struct {
	port int
	env  string
	api  string

	db struct {
		dsn string
	}

	stripe struct {
		secret string
		key    string
	}

	secretkey string
	frontend  string
}

func appConfig() *application {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001/api", "URL to API")
	flag.StringVar(&cfg.db.dsn, "dsn", "ananto:secret@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "(dsn) database string")
	flag.StringVar(&cfg.secretkey, "secret", "jB7Mig0ZWNoqGZGqxubinIQXYxzpKt9b", "secret key")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to front end")

	flag.Parse()

	var l zerolog.Logger
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	l = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})

	app := &application{
		config: cfg,
		log:    &l,
	}

	return app

}
