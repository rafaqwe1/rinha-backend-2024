package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaqwe1/rinha-backend-2024/cmd"
	"github.com/rafaqwe1/rinha-backend-2024/infrastructure/database"
)

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := "postgresql://admin:123@" + host + ":5432/rinha?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute(database.NewClientDbPostgres(pool), database.NewTransactionDbPostgres(pool))
}
