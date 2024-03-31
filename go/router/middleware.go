package router

import (
	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/controller"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type Middleware interface {
}

type DefaultMiddleware struct {
	config   config.Configuration
	logger   logger.Logger
	response controller.HTTPResponse
}

func NewMiddleware(config config.Configuration, logger logger.Logger, response controller.HTTPResponse) (middleware *DefaultMiddleware) {

	return &DefaultMiddleware{
		config:   config,
		logger:   logger,
		response: response,
	}
}
