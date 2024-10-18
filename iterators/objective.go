package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func printAllUsers(ctx context.Context, conn *pgx.Conn) error {
	rows, err := conn.Query(ctx, "SELECT id, name FROM users")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return err
		}
		fmt.Println(id, name)
	}

	// Должен отрабатывать так.
	// for id, name := range users(conn) {
	// 	fmt.Println(id, name)
	// }

	return nil
}
