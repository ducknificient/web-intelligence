package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/general"
	"github.com/julienschmidt/httprouter"
)

type CrawlerService interface {
	Crawling(seedurl string, task string) (err error)
	CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error)
}

func (c *DefaultController) NewCrawlerService(service CrawlerService) {
	c.crawlerService = service
}

func (u *DefaultController) StartCrawling(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	err = u.crawlerService.Crawling(request.Task, request.SeedURL)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
		return
	}

	return
}

func (u *DefaultController) CrawlpageList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	dataList, err = u.crawlerService.CrawlpageList(param)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "Unable to get crawlpage list.")
		return
	}

	dataListWrap := entity.CrawlpageListDataWrap{Total: len(dataList), Data: dataList}
	u.response.CustomResponse(w, r, prefixLog, "ok", dataListWrap)
	return
}
