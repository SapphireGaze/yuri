package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sapphiregaze/yuri-server/internal/database"
	"github.com/sapphiregaze/yuri-server/internal/redirect"
	"github.com/sapphiregaze/yuri-server/internal/route"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load enviornment variables: %v", err)
	}

	db := database.ConnectDB()
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Failed to start database: %v", err)
	}
	defer db.Close()

	if err := redirect.CreateSchema(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	switch strings.ToLower(os.Getenv("GIN_MODE")) {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	if err := route.InitRoutes().Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
