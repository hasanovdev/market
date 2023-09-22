package handler

import (
	"market/config"
	"market/pkg/logger"
	"market/storage"
)

type Handler struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
}

func NewHandler(strg storage.StorageI, cfg config.Config, log logger.LoggerI) *Handler {
	return &Handler{strg: strg, cfg: cfg, log: log}
}
