package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configpackage "github.com/ducknificient/web-intelligence/go/config"
	controllerpackage "github.com/ducknificient/web-intelligence/go/controller"
	"github.com/ducknificient/web-intelligence/go/datastore"
	loggerpackage "github.com/ducknificient/web-intelligence/go/logger"
	routerpackage "github.com/ducknificient/web-intelligence/go/router"
	"github.com/ducknificient/web-intelligence/go/service"

	serverpackage "github.com/ducknificient/web-intelligence/go/server"
)

var (
	configPath string
	config     *configpackage.AppConfiguration
	err        error
)

func init() {

	// reading from command line
	var args = os.Args
	configPath = args[1]

	// config file path
	config, err = configpackage.NewConfiguration(configPath)
	if err != nil {
		panic(err)
	}
}

func main() {

	/*
		SETUP CONTEXT
	*/
	// https://dasarpemrogramangolang.novalagung.com/A-pipeline-context-cancellation.html
	var (
		ctx = context.Background()
		err error
	)

	/*
		SETUP LOGGER
	*/

	// create logger
	logger, err := loggerpackage.NewLogger(config)
	if err != nil {
		// logger.Error(fmt.Sprintf("unable to init logger : %v", err.Error()))
		panic(err.Error())
	}

	err = logger.CheckEmptyLog()
	if err != nil {
		panic(err)
	}

	err = logger.SetupCrawlLogFile()
	if err != nil {
		panic(err)
	}

	/*
		SETUP DATASTORE IMPLEMENTATION

		// MODEL / REPOSITORY / DAO
	*/

	mapPgDB, err := datastore.NewPostgresModelList(ctx, config, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("unable to make new postgres model : %v", err.Error()))
		panic(err.Error())
	}

	postgresDB := mapPgDB[*config.SelectedPgDB]

	/*
		SETUP SERVICE
	*/

	wiservice := service.NewWIService(config, logger, postgresDB)

	/*
		SETUP CONTROLLER
	*/

	// init controller response
	httpresponse := controllerpackage.NewHTTPResponse(logger)

	// init default http controller
	handler := controllerpackage.NewHTTPController(config, logger, httpresponse)

	// inject service to handler
	handler.NewWIService(wiservice)

	/*
		SETUP ROUTER
	*/

	// init middleware
	middleware := routerpackage.NewMiddleware(config, logger, httpresponse)

	// init router
	httprouter := routerpackage.NewRouter(handler, middleware, config)

	/*
		SETUP SERVER
	*/

	// init http server
	appIp := *config.AppIP + ":" + *config.AppPort
	httpserver := serverpackage.NewHTTPServer(&http.Server{
		Addr:    appIp,
		Handler: httprouter.Router,
	}, logger)

	if *config.HTTPS == `true` {
		httpserver.SetCertificate(*config.Certificate, *config.CertificateKey)

		// run server
		httpserver.RunTLS()

	} else {
		httpserver.Run()
	}

	/*
		Graceful shutdown
	*/

	// Wait for kill signal of channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Shutdown server
	logger.Info("Shutting down server...")
	err = httpserver.Shutdown(ctx)
	if err != nil {
		// fmt.Errorf(fmt.Sprintf("Server forced to shutdown : %v\n", err.Error()))
		logger.Fatal(fmt.Sprintf("Server forced to shutdown : %v\n", err.Error()))
	}

}

/*

func TestExtractUrl(conn *pgxpool.Pool) {

	var (
		q         string
		inputurl  string
		inputhtml string

		resurl  sql.NullString
		reshtml sql.NullString
		err     error
	)

	q = `SELECT t.u,t.du FROM webintelligence.tabled t WHERE u = 'https://www.azlyrics.com/a.html' `
	fmt.Println(q)

	err = conn.QueryRow(context.Background(), q).Scan(&resurl, &reshtml)
	if err != nil {
		panic(err)
	}

	inputurl = resurl.String
	inputhtml = reshtml.String

	// fmt.Printf("\n%#v,%#v\n", inputurl, inputhtml)

	List := ExtractURL(inputurl, inputhtml)

	// file, err := os.Create("href_list.txt")

	for a, b := range List {
		fmt.Printf("%#v. href : '%#v' \n", a, b)
		// fmt.Fprintf(file, "%#v. href : '%#v' \n", a, b)

		if a == 10 {
			// break
		}
	}
}

*/
