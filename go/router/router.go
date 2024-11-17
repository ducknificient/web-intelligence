package router

import (
	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/controller"
	"github.com/julienschmidt/httprouter"
)

type DefaultRouter struct {
	Router *httprouter.Router
}

func NewRouter(controller controller.HTTPController, middleware Middleware, config config.Configuration) *DefaultRouter {

	router := httprouter.New()

	router.HandlerFunc("POST", "/path", controller.Options)

	/* ALL */
	router.GET("/", httprouter.WrapF(controller.Root))
	// router.Use(middleware.EnableCors)

	router.OPTIONS("/*path", httprouter.WrapF(controller.Options))
	router.GET("/serve/:path/:identifier/:file", httprouter.WrapF(controller.ServeFile))

	/* DEFAULT */
	router.GET("/about/version", httprouter.WrapF(controller.About))
	router.POST("/database/ping", httprouter.WrapF(controller.DatabasePing))

	/* BUSINESS LOGIC */
	router.POST("/crawl/start", httprouter.WrapF(controller.StartCrawling))
	router.POST("/crawl/start/multiple", httprouter.WrapF(controller.StartMultipleCrawling))
	router.POST("/crawl/stop", httprouter.WrapF(controller.StopCrawling))
	router.POST("/crawl/list", httprouter.WrapF(controller.CrawlpageList))
	router.POST("/crawl/list/parsed", httprouter.WrapF(controller.CrawlpageListParsed))

	router.POST("/alodokter/crawler", httprouter.WrapF(controller.AlodokterCrawler))
	router.POST("/alodokter/crawler/CHAT", httprouter.WrapF(controller.AlodokterCrawlerChat))

	router.POST("/alodokter/checkurl", httprouter.WrapF(controller.AlodokterCheckUrl))
	router.POST("/alodokter/list/export", httprouter.WrapF(controller.AlodokterListExport))

	router.POST("/halodoc/checkurl", httprouter.WrapF(controller.HalodocListPenyakit))

	router.POST("/cookpad/crawler", httprouter.WrapF(controller.CookpadCrawler))
	router.POST("/cookpad/crawler/image", httprouter.WrapF(controller.CookpadImageCrawler))
	router.POST("/cookpad/getimage", httprouter.WrapF(controller.CookpadGetImage))

	return &DefaultRouter{
		Router: router,
	}
}
