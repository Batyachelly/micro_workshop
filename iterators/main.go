package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	dsn := "host=localhost port=5432 password=example user=postgres database=mydatabase"
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	if err := migrate(ctx, conn); err != nil {
		log.Fatal(err)
	}

	if err := printAllUsers(ctx, conn); err != nil {
		log.Fatal(err)
	}
}

func migrate(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
		DROP TABLE IF EXISTS users;

		CREATE TABLE IF NOT EXISTS users
		(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);

		INSERT INTO users (name) SELECT 'User ' || generate_series(1, 10);
	`)
	if err != nil {
		return err
	}

	return nil
}
