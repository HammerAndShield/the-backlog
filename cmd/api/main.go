package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
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
	wg     sync.WaitGroup
}

func main() {
	cfg := getConfigurationFlags()

	logger := httplog.NewLogger("http-logger", httplog.Options{
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		Tags: map[string]string{
			"env": cfg.env,
		},
		QuietDownRoutes: []string{
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
	})

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Debug("connection to DB established")
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Info("starting server")

	err = app.serve()
	if err != nil {
		logger.Error("error starting server", "error", err)
		return
	}
}

func getConfigurationFlags() config {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev | staging | prod")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("THEBACKLOG_DB_DSN"), "PostgreSQL connection string")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle conns")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max idle time")

	return cfg
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
