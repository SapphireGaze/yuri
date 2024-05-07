package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB {
	cfg := &pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: "yuri",
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			fmt.Printf("Successfully connected to database as %s\n", os.Getenv("DB_USER"))
			return nil
		},
	}

	db := pg.Connect(cfg)

	return db
}
