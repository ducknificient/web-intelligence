package service

import (
	"fmt"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type CookpadCrawler struct {
	config   config.Configuration
	logger   logger.Logger
	database datastore.Datastore
}

func NewCookpadService(config config.Configuration, logger logger.Logger, database datastore.Datastore) *CookpadCrawler {

	return &CookpadCrawler{
		config:   config,
		logger:   logger,
		database: database,
	}
}

func (a *WIService) CookpadGetListImageUrl() (dataList []entity.CookpadValidation, err error) {

	var (
		allList []entity.CookpadRecipeList = make([]entity.CookpadRecipeList, 0)
	)

	allList, err = a.database.GetListImageUrl()
	if err != nil {
		return dataList, err
	}

	for _, b := range allList {

		var isexist bool = false
		// newurl := "https://www.alodokter.com/" + b.Permalink
		isexist, err = a.database.CheckUrlIsExist(b.Url)
		if err != nil {
			fmt.Printf("error: %#v \n", err.Error())
			continue
		}

		data := entity.CookpadValidation{
			Url:     b.Url,
			IsExist: isexist,
		}

		fmt.Printf("exist:%#v \n", isexist)

		dataList = append(dataList, data)

	}

	return dataList, err
}

func (a *WIService) CookpadCheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error) {

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

func (a *WIService) CookpadGetImageList() (dataList []entity.CookpadImageList, err error) {

	dataList, err = a.database.GetCookpadListImage()
	if err != nil {
		fmt.Printf("error: %#v \n", err.Error())
		return dataList, err
	}

	return dataList, err
}

func (a *WIService) CookpadSaveImageToLocal(param entity.CookpadSaveImageParam) (err error) {

	return err
}
