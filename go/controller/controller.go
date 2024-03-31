package controller

import (
	"net/http"

	"github.com/ducknificient/web-intelligence/go/config"
	configpackage "github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/logger"
	loggerpackage "github.com/ducknificient/web-intelligence/go/logger"
	"github.com/ducknificient/web-intelligence/go/service"
)

type HTTPController interface {
	Root(w http.ResponseWriter, r *http.Request)
	Options(w http.ResponseWriter, r *http.Request)
	About(w http.ResponseWriter, r *http.Request)
	ServeFile(w http.ResponseWriter, r *http.Request)
	DatabasePing(w http.ResponseWriter, r *http.Request)
	StartCrawling(w http.ResponseWriter, r *http.Request)
	StartMultipleCrawling(w http.ResponseWriter, r *http.Request)
	StopCrawling(w http.ResponseWriter, r *http.Request)
	CrawlpageList(w http.ResponseWriter, r *http.Request)
	CrawlpageListParsed(w http.ResponseWriter, r *http.Request)
}

type DefaultController struct {
	config         config.Configuration
	logger         logger.Logger
	response       HTTPResponse
	crawlerService service.CrawlerService
	crawlStop      bool
}

// func NewDefaultController(logger logger.Logger) (default_controller *DefaultController) {
// 	return &DefaultController{
// 		logger: logger,
// 		response: &Response{
// 			Logger: logger,
// 		},
// 		crawlStop: false,
// 	}
// }

func NewHTTPController(config configpackage.Configuration, logger loggerpackage.Logger, res HTTPResponse) (defaultController *DefaultController) {

	defaultController = &DefaultController{
		config:    config,
		logger:    logger,
		response:  res,
		crawlStop: false,
	}

	return defaultController
}

func (u *DefaultController) Root(w http.ResponseWriter, r *http.Request) {
	defer u.response.Panic(w, r)

	u.response.DefaultText(w, http.StatusOK, true, *u.config.GetConfiguration().Version)
	return
}

func (u *DefaultController) About(w http.ResponseWriter, r *http.Request) {
	defer u.response.Panic(w, r)

	u.response.DefaultText(w, http.StatusOK, true, *u.config.GetConfiguration().Version)
	return
}

func (u *DefaultController) Options(w http.ResponseWriter, r *http.Request) {
	defer u.response.Panic(w, r)

	return
}

func (u *DefaultController) DatabasePing(w http.ResponseWriter, r *http.Request) {
	defer u.response.Panic(w, r)

	return
}

func (u *DefaultController) ServeFile(w http.ResponseWriter, r *http.Request) {
	defer u.response.Panic(w, r)

	// c.ByName("")

	return
}
