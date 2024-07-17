package main

import (
	"github.com/alkosmas92/xm-golang/internal/database"
	"github.com/alkosmas92/xm-golang/internal/logs"
	"github.com/alkosmas92/xm-golang/internal/server"
	"log"
)

func main() {
	logger, err := logs.Initialize()
	if err != nil {
		panic(err)
	}

	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := server.Run(logger, db); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
