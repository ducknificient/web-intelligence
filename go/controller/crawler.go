package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/general"
)

// func (c *DefaultController) NewCrawlerService(service service.DefaultService) {
// 	c.crawlerService = service
// }

func (u *DefaultController) StartCrawling(w http.ResponseWriter, r *http.Request) {
	prefixLog := `StartCrawling`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CrawlingReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
	// seedurl string
	// task    string
	)

	// task = `KEMENDAG`
	// seedurl = "https://www.kemendag.go.id/berita/perdagangan?page=8"

	// go func() (err error) {

	// 	return err
	// }()

	// if err != nil {
	// 	u.response.Error(w, r, err, prefixLog, fmt.Sprintf("go routines crawl error."))
	// 	u.logger.Error(fmt.Sprintf("go routines crawl error."))
	// 	return
	// }

	err = u.defaultService.Crawling(request.SeedURL, request.Task)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
		return
	}

	// err = u.crawlerService.StartCrawling()
	// u.crawlerService.TestCrawling()

	fmt.Println("crawl ended")

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) StartMultipleCrawling(w http.ResponseWriter, r *http.Request) {
	prefixLog := `StartMultipleCrawling`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CrawlingMultipleReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
	// seedurl string
	// task    string
	)

	// task = `KEMENDAG`
	// seedurl = "https://www.kemendag.go.id/berita/perdagangan?page=8"

	// go func() (err error) {

	// 	return err
	// }()

	// if err != nil {
	// 	u.response.Error(w, r, err, prefixLog, fmt.Sprintf("go routines crawl error."))
	// 	u.logger.Error(fmt.Sprintf("go routines crawl error."))
	// 	return
	// }

	for _, seedurl := range request.SeedURLList {

		// go func() (err error) {

		// 	return err
		// }()

		fmt.Printf("starting crawl for seedurl: %#v\n", seedurl)
		err = u.defaultService.Crawling(seedurl, request.Task)
		if err != nil {
			u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
			return
		}

		fmt.Println("crawl ended")
	}

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) StopCrawling(w http.ResponseWriter, r *http.Request) {
	prefixLog := `StopCrawling`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CrawlingReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	u.defaultService.StopCrawling()

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) CrawlpageList(w http.ResponseWriter, r *http.Request) {
	prefixLog := `CrawlpageList`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CrawlpageListReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
	// seedurl string
	// task    string
	)

	// task = `KEMENDAG`
	// seedurl = "https://www.kemendag.go.id/berita/perdagangan?page=8"

	var (
		param    entity.CrawlpageListParam
		dataList []entity.CrawlpageListData
	)

	param.Page = request.Page
	param.Count = request.Count
	param.Search = request.Search

	fmt.Printf("param: %#v\n", param)

	dataList, err = u.defaultService.CrawlpageList(param)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "Unable to get crawlpage list.")
		return
	}

	dataListWrap := entity.CrawlpageListDataWrap{Total: len(dataList), Data: dataList}
	u.response.CustomResponse(w, r, prefixLog, "ok", dataListWrap)
	return
}

func (u *DefaultController) CrawlpageListParsed(w http.ResponseWriter, r *http.Request) {
	prefixLog := `CrawlpageListParsed`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CrawlpageListReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
	// seedurl string
	// task    string
	)

	// task = `KEMENDAG`
	// seedurl = "https://www.kemendag.go.id/berita/perdagangan?page=8"

	var (
		param    entity.CrawlpageListParam
		dataList []entity.CrawlpageListParsedData
	)

	param.Page = request.Page
	param.Count = request.Count
	param.Search = request.Search

	fmt.Printf("param: %#v\n", param)

	dataList, err = u.defaultService.CrawlpageListParsed(param)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "Unable to get crawlpage list.")
		return
	}

	dataListWrap := entity.CrawlpageListParsedDataWrap{Total: len(dataList), Data: dataList}
	u.response.CustomResponse(w, r, prefixLog, "ok", dataListWrap)
	return
}
