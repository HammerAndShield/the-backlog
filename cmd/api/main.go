package main

import (
	"fmt"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config config
	logger *httplog.Logger
}

func main() {
	app := &application{
		config: config{
			port: 8080,
			env:  "dev",
		},
	}

	logger := httplog.NewLogger("http-logger", httplog.Options{
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		Tags: map[string]string{
			"env": app.config.env,
		},
		QuietDownRoutes: []string{
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
	})

	app.logger = logger

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	logger.Debug("starting server")

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
