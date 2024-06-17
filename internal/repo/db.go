package repo

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

// db is the database connection pool
var db *sql.DB

func CreateDbPool(dsn string) (*sql.DB, error) {

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5 * time.Minute)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	db = conn

	return db, nil
}
