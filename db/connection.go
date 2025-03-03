package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type DBConnection struct {
	Conn *pgx.Conn
	Ctx  context.Context
}

func ConnectDB(databaseURL string) *DBConnection {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to database")
	return &DBConnection{Conn: conn, Ctx: ctx}
}

func (db *DBConnection) Close(ctx context.Context) {
	db.Conn.Close(ctx)
}
