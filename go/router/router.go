package router

import (
	"net/http"

	"github.com/ducknificient/web-intelligence/go/controller"
	"github.com/julienschmidt/httprouter"
)

type Router interface {
	GetHandler() (handler http.Handler)
}

type DefaultRouter struct {
	Router     *httprouter.Router
	Controller controller.Controller
}

func NewDefaultRouter(controller controller.Controller) *DefaultRouter {
	return &DefaultRouter{
		Controller: controller,
	}
}

func (s *DefaultRouter) SetRouter(router *httprouter.Router) {
	s.Router = router
}

func (s *DefaultRouter) SetResponse(router *httprouter.Router) {

}

func (s *DefaultRouter) GetHandler() http.Handler {

	router := s.Router
	defaultController := s.Controller

	/* ALL */
	router.GET("/", defaultController.Root)
	// router.Use(middleware.EnableCors)

	router.OPTIONS("/*path", defaultController.Options)
	router.GET("/serve/:path/:identifier/:file", defaultController.ServeFile)

	/* DEFAULT */
	router.GET("/about/version", defaultController.About)
	router.POST("/database/ping", defaultController.DatabasePing)

	/* BUSINESS LOGIC */
	router.POST("/crawl/start", defaultController.StartCrawling)
	router.POST("/crawl/start/multiple", defaultController.StartMultipleCrawling)
	router.POST("/crawl/stop", defaultController.StopCrawling)
	router.POST("/crawl/list", defaultController.CrawlpageList)
	router.POST("/crawl/list/parsed", defaultController.CrawlpageListParsed)

	return router
}
