package datastore

import "github.com/ducknificient/web-intelligence/go/entity"

type Datastore interface {
	StoreD(pagesource string, link string, task string, documenttype string, mimetype string) (err error)
	StoreDocument(link string, task string, filename string, document []byte, documentcontenttype string) (err error)
	StoreE(link string, href string, task string) (err error)
	ContainsD(url string) (contains bool, err error)
	GetExistingQueue(task string) (queue []string, err error)
	GetLatestSeedUrl(task string, seedurl string) (res_seedurl string, err error)
	CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error)
	CrawlpageListParsed(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListParsedData, err error)
	CheckUrlIsExist(link string) (isexist bool, err error)
	GetListImageUrl() (dataList []entity.CookpadRecipeList, err error)
	GetAlodokterListParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error)
	GetCookpadListImage() (dataList []entity.CookpadImageList, err error)
}
