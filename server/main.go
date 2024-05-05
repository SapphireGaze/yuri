package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sapphiregaze/yuri-server/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.ConnectDB()
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
}
