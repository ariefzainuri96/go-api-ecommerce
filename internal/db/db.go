package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func New(addr string, maxOpenCons, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)

	if err != nil {
		log.Println("openError:", err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenCons)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)

	if err != nil {
		log.Println("parseDurationError:", err.Error())
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		log.Println("pingError:", err.Error())
		return nil, err
	}

	return db, nil
}
