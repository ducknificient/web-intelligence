package controller

import (
	"net/http"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/logger"
	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Options(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	About(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	ServeFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	DatabasePing(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	StartCrawling(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	CrawlpageList(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type DefaultController struct {
	response       Response
	crawlerService CrawlerService
}

func NewDefaultController(logger logger.Logger) (default_controller *DefaultController) {
	return &DefaultController{
		response: &DefaultResponse{
			Logger: logger,
		},
	}
}

func (u *DefaultController) Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer u.response.Panic(w, r)

	u.response.DefaultText(w, http.StatusOK, true, *config.Conf.Version)
	return
}

func (u *DefaultController) About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer u.response.Panic(w, r)

	u.response.DefaultText(w, http.StatusOK, true, *config.Conf.Version)
	return
}

func (u *DefaultController) Options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer u.response.Panic(w, r)

	return
}

func (u *DefaultController) DatabasePing(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer u.response.Panic(w, r)

	return
}

func (u *DefaultController) ServeFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer u.response.Panic(w, r)

	return
}
