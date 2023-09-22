package main

import (
	"context"
	"market/api"
	"market/api/handler"
	"market/config"
	"market/pkg/logger"
	"market/storage/db"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("Project Name: Market", logger.LevelInfo)
	strg, err := db.NewStorage(context.Background(), *cfg)
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(strg, *cfg, log)

	r := api.NewServer(h)
	r.Run(":3000")
}
