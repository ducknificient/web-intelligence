package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type AlodokterCrawlerService interface {
	GetNamaPenyakit() (dataList []entity.AlodokterPenyakit, err error)
	GetNamaObat() (dataList []entity.AlodokterObat, err error)
	CheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error)
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

func (a *AlodokterCrawler) GetNamaPenyakit() (dataList []entity.AlodokterPenyakit, err error) {

	// open json
	var (
		errMsg   string
		filename string
		// dataList []entity.AlodokterPenyakit
	)

	dataList = make([]entity.AlodokterPenyakit, 0)

	filename = `/home/spil/Projects/web-intelligence/temp/alodokter/penyakit.json`

	jsonFile, err := os.Open(filename)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to load config.json :" + err.Error())
		err := errors.New(errMsg)
		return dataList, err
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to read jsonFile :" + err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	err = json.Unmarshal(byteData, &dataList)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to unmarshall jsonFile -> byteData :" + err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	// for _, b := range dataList {

	// 	newurl := "https://www.alodokter.com/" + b.Permalink
	// 	fmt.Printf("%#v\n", newurl)
	// }

	return dataList, err
}

func (a *AlodokterCrawler) GetNamaObat() (dataList []entity.AlodokterObat, err error) {

	// open json
	var (
		errMsg   string
		filename string
		// dataList []entity.AlodokterPenyakit
	)

	dataList = make([]entity.AlodokterObat, 0)

	filename = `/home/spil/Projects/web-intelligence/temp/alodokter/obat.json`

	jsonFile, err := os.Open(filename)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to load config.json :" + err.Error())
		err := errors.New(errMsg)
		return dataList, err
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to read jsonFile :" + err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	err = json.Unmarshal(byteData, &dataList)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to unmarshall jsonFile -> byteData :" + err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	return dataList, err
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
