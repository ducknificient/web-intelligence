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

func (a *WIService) AlodokterGetNamaPenyakit() (dataList []entity.AlodokterPenyakit, err error) {

	// open json
	var (
		errMsg   string
		filename string
		// dataList []entity.AlodokterPenyakit
	)

	dataList = make([]entity.AlodokterPenyakit, 0)

	filename = `/home/spil/jeremy/Projects/s2/web-intelligence/temp/alodokter/penyakit.json`

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

func (a *WIService) AlodokterGetNamaObat() (dataList []entity.AlodokterObat, err error) {

	// open json
	var (
		errMsg   string
		filename string
		// dataList []entity.AlodokterPenyakit
	)

	dataList = make([]entity.AlodokterObat, 0)

	filename = `/home/spil/jeremy/Projects/s2/web-intelligence/temp/alodokter/obat.json`

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

func (a *WIService) AlodokterCheckUrlIsExist(dataList []entity.AlodokterPenyakit) (listUrl []entity.AlodokterValidation, err error) {

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

func (a *WIService) AlodokterGetListDataParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error) {

	dataList, err = a.database.GetAlodokterListParsed(param)
	if err != nil {
		return dataList, err
	}

	return dataList, err

}
