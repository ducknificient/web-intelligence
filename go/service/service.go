package service

import "github.com/ducknificient/web-intelligence/go/entity"

type DefaultService interface {
	// crawl
	Crawling(seedurl string, task string) (err error)
	TestCrawling()
	StartCrawling() (err error)
	StopCrawling() (err error)
	CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error)
	CrawlpageListParsed(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListParsedData, err error)

	// alodokter
	AlodokterGetNamaPenyakit() (dataList []entity.AlodokterPenyakit, err error)
	AlodokterGetNamaObat() (dataList []entity.AlodokterObat, err error)
	AlodokterCheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error)
	AlodokterGetListDataParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error)

	// cookpad
	CookpadGetListImageUrl() (dataList []entity.CookpadValidation, err error)
	CookpadGetImageList() (dataList []entity.CookpadImageList, err error)
	CookpadSaveImageToLocal(param entity.CookpadSaveImageParam) (err error)

	// halodoc
	HalodocGetListPenyakit() (dataList []entity.HalodocListPenyakit, err error)
}
