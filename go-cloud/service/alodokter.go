package service

import (
	"fmt"
	"go-cloud/entity"

	"go-cloud/config"
	"go-cloud/datastore"
	"go-cloud/logger"
)

type AlodokterCrawlerService interface {
	GetNamaPenyakit() (dataList []entity.AlodokterPenyakit, err error)
	GetNamaObat() (dataList []entity.AlodokterObat, err error)
	CheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error)
	GetListDataParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error)
}

type AlodokterCrawler struct {
	config   config.Configuration
	logger   logger.Logger
	database datastore.Datastore
}

func NewAlodokterService(config config.Configuration, logger logger.Logger, database datastore.Datastore) *AlodokterCrawler {

	return &AlodokterCrawler{
		config:   config,
		logger:   logger,
		database: database,
	}
}

func (a *AlodokterCrawler) CheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error) {

	var isexist bool = false

	for _, b := range dataList {
		newurl := "https://www.alodokter.com/" + b.Permalink
		isexist, err = a.database.CheckUrlIsExist(newurl)
		if err != nil {
			fmt.Printf("error: %#v \n", err.Error())
			continue
		}

		data := entity.AlodokterValidation{
			Url:     newurl,
			IsExist: isexist,
		}

		listUrl = append(listUrl, data)

	}

	return listUrl, err

}

func (a *AlodokterCrawler) GetListDataParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error) {

	dataList, err = a.database.GetAlodokterListParsed(param)
	if err != nil {
		return dataList, err
	}

	return dataList, err

}
