package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ducknificient/web-intelligence/go/config"
	"github.com/ducknificient/web-intelligence/go/datastore"
	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/logger"
)

type HalodocCrawler struct {
	config   config.Configuration
	logger   logger.Logger
	database datastore.Datastore
}

func NewHalodocService(config config.Configuration, logger logger.Logger, database datastore.Datastore) *HalodocCrawler {

	return &HalodocCrawler{
		config:   config,
		logger:   logger,
		database: database,
	}
}

func (a *WIService) HalodocGetListPenyakit() (dataList []entity.HalodocListPenyakit, err error) {

	// open json
	var (
		errMsg   string
		filename string
		// dataList []entity.AlodokterPenyakit
	)

	dataList = make([]entity.HalodocListPenyakit, 0)

	var datalist2 []entity.HalodocResultData

	var Alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "X", "Y", "Z"}

	for _, alphabet := range Alphabet {
		filename = `/home/spil/jeremy/Projects/s2/web-intelligence/datastore/halodoc2/` + strings.ToLower(alphabet) + `.json`
		var data entity.HalodocListPenyakit

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

		err = json.Unmarshal(byteData, &data)
		if err != nil {
			errMsg = fmt.Sprintf("Unable to unmarshall jsonFile -> byteData :" + err.Error())
			err = errors.New(errMsg)
			return dataList, err
		}

		// fmt.Printf("len: %#v", len(data.Result))
		// fmt.Printf("len: %#v\n", data.Total)

		// for _, data := range data.Result {

		// }

		datalist2 = append(datalist2, data.Result...)

		fmt.Printf("%#v. \n", len(data.Result))

	}

	fmt.Printf("total: %#v.\n", len(datalist2))

	return dataList, err
}
