package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	configpackage "go-cloud/config"
	"go-cloud/datastore"
	loggerpackage "go-cloud/logger"
	"go-cloud/service"
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
	// configPath = "/home/ducknificient/crawlmachine/config-cloud.json"

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

	crawler := service.NewCrawler(config, logger, postgresDB)
	// alodokter := service.NewAlodokterService(config, logger, postgresDB)

	// https://www.alodokter.com/komunitas/diskusi/penyakit/page/2

	dataList := make([]string, 50000)
	startidx := 11500
	for _, b := range dataList {
		startidx++

		newtask := `ALODOKTER-CHAT`
		newurl := fmt.Sprintf("https://www.alodokter.com/komunitas/diskusi/penyakit/page/%v", startidx)

		// fmt.Printf("%#v\n", newurl)
		err = crawler.Crawling(newurl, newtask)
		if err != nil {
			fmt.Printf("crawl error, %#v\n", err)
			return
		}

		fmt.Printf("", b)
		fmt.Printf("%v\n", newurl)

		time.Sleep(300 * time.Millisecond)
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
