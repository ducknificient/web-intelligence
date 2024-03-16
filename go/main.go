package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/controller"
	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/logger"
	"github.com/ducknificient/web-intelligence/go/router"
	"github.com/julienschmidt/httprouter"
)

/*initialization config file*/
func init() {
	config.Conf = config.NewConfiguration("dev")
}

func main() {

	// setup context
	var (
		ctx = context.Background()
		err error
	)

	// setup logger
	default_logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	default_logger.PathCrawlLog = *config.Conf.PathCrawlLog
	default_logger.PathCrawlError = *config.Conf.PathCrawlError
	default_logger.PathCrawlPdf = *config.Conf.PathCrawlPdf
	err = default_logger.SetupCrawlLogFile()
	if err != nil {
		panic(err)
	}

	err = default_logger.CheckEmptyLog()
	if err != nil {
		// logger.DefaultLogger.Logger.Error("")
		panic(err)
	}

	// postgresDB := datastore.NewPostgreSQLDB(ctx, default_logger)
	// err = postgresDB.Connect()
	// if err != nil {
	// 	default_logger.Error(fmt.Sprintf("unable to connect : %v", err.Error()))
	// 	panic(err.Error())
	// }

	listPgDatabase := datastore.GetListPgDatabase(default_logger)
	mapDB := make(map[string]datastore.Datastore)

	for _, postgresDB := range listPgDatabase {

		postgresDB.Ctx = ctx
		err := postgresDB.Connect()
		if err != nil {
			default_logger.Error(fmt.Sprintf("unable to connect : %v", err.Error()))
			panic(err.Error())
		}

		err = postgresDB.Conn.Ping(ctx)
		if err != nil {
			panic(err.Error())
		}

		mapDB[*postgresDB.PgInfo.DbName] = &postgresDB

	}

	postgresDB := mapDB[*config.Conf.SelectedPgDB]

	// 6. setup model (repository)
	// postgresDB := &datastore.PostgresDB{
	// 	Ctx:        ctx,
	// 	Pool:       postgrespool.Conn,
	// 	BackupPool: postgrespool.Conn,
	// 	VectorPool: pgvectorpool.Conn,
	// }

	crawler := NewCrawler(postgresDB)

	// init controller

	default_controller := controller.NewDefaultController(default_logger)
	default_controller.NewCrawlerService(crawler)

	// init schmith router

	default_router := httprouter.New()

	handler_router := router.NewDefaultRouter(default_controller)
	handler_router.SetRouter(default_router)

	// ipPort := ":" + "8090"
	// fmt.Println("App listening on " + ipPort)
	appIp := *config.Conf.AppIP + ":" + *config.Conf.AppPort
	server := &http.Server{Addr: appIp, Handler: handler_router.GetHandler()}

	go func() {
		default_logger.Info("App listening on " + server.Addr)
		if err := server.ListenAndServe(); err != nil {
			panic(err)
			return
		}
	}()

	// 11. Graceful shutdownra

	// Wait for kill signal of channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// fmt.Println("before quit")
	// This blocks until a signal is passed into the quit channel
	<-quit
	fmt.Println("after quit")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	// log.Println("Shutting down server...")
	default_logger.Info("Shutting down server...")
	err = server.Close()
	if err != nil {
		//log.Fatalf("Server forced to shutdown: %v\n", err)
		default_logger.Error(fmt.Sprintf("Server forced to shutdown : %v\n", err.Error()))
		// logger.Fatal(fmt.Sprintf("Server forced to shutdown : %v\n", err.Error()))
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
