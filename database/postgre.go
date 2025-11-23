package database

import (
    "context"
    "log"

    "github.com/jackc/pgx/v5/pgxpool"
    "sppm/config" // IMPORT PACKAGE CONFIG
)

var Pg *pgxpool.Pool

func ConnectPostgres() {
    url := config.Env("POSTGRES_URL") 

    pool, err := pgxpool.New(context.Background(), url)
    if err != nil {
        log.Fatalf("❌ Failed create pool: %v", err)
    }

    if err := pool.Ping(context.Background()); err != nil {
        log.Fatalf("❌ PostgreSQL not responding: %v", err)
    }

    Pg = pool
    log.Println("✅ Connected to PostgreSQL!")
}
