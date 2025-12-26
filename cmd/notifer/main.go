package main

import (
	"delayed-notifier/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	zlog.Init()

	cfg := config.LoadConfig()

	zlog.Logger.Info().
		Str("host", cfg.Postgres.Host).
		Int("port", cfg.Postgres.Port).
		Str("user", cfg.Postgres.User).
		Str("database", cfg.Postgres.Name).
		Msg("PostgreSQL configuration loaded")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=10",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer db.Close()

	zlog.Logger.Info().Msg("Successfully connected to PostgreSQL!")
}
