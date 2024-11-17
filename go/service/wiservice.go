package service

import (
	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type WIService struct {
	config   config.Configuration
	logger   logger.Logger
	database datastore.Datastore
	// SeedURL   string
	Task      string
	Datastore datastore.Datastore
	IsStop    bool
}

func NewWIService(config config.Configuration, logger logger.Logger, datastore datastore.Datastore) (c *WIService) {
	return &WIService{
		config:    config,
		logger:    logger,
		Datastore: datastore,
		IsStop:    false,
	}

}
