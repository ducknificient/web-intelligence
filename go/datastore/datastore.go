package datastore

import "github.com/ducknificient/web-intelligence/go/entity"

type Datastore interface {
	StoreD(pagesource string, link string, task string) (err error)
	StorePdf(filename string, link string, task string, pdf []byte) (err error)
	StoreE(link string, href string, task string) (err error)
	ContainsD(url string) (contains bool, err error)
	CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error)
}
