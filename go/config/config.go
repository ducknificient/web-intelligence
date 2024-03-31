package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Configuration interface {
	GetConfiguration() (config *AppConfiguration)
}

type PgInfo struct {
	Host        *string `json:"host"` // server ip
	DbName      *string `json:"dbname"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Port        *string `json:"port"`
	SSLMode     *string `json:"sslmode"`
	PoolMaxConn *string `json:"poolmaxconn"`
}

type AppConfiguration struct {
	// app configuration
	Appname    *string `json:"appname"`
	AppDomain  *string `json:"appdomain"`
	AppIP      *string `json:"appip"`
	AppPort    *string `json:"appport"`
	Production *string `json:"production"`
	Version    *string `json:"version"`

	// misc
	FileSep *string `json:"filesep"`

	// DB         *DbInfo   `json:"dbinfo"`
	PgList       *[]PgInfo `json:"pglist"`
	SelectedPgDB *string   `json:"selectedpgdb"`

	// path
	PathLog   *string `json:"pathlog"`
	PathTemp  *string `json:"pathtemp"`
	PathPdf   *string `json:"pathpdf"`
	PathXlsx  *string `json:"pathxlsx"`
	PathCsv   *string `json:"pathcsv"`
	PathGraph *string `json:"pathgraph"`

	PathCrawlLog   *string `json:"pathcrawllog"`
	PathCrawlError *string `json:"pathcrawlerror"`
	PathCrawlHref  *string `json:"pathcrawlhref"`
	PathCrawlPdf   *string `json:"pathcrawlpdf"`

	// authentication
	AesKey      *string `json:"aeskey"`
	CookieName  *string `json:"cookiename"`
	ServiceName *string `json:"servicename"`
	UserHeader  *string `json:"userheader"`
}

func NewConfiguration(filename string) (config *AppConfiguration, err error) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to load config.json :" + err.Error())
		return config, err
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		err = fmt.Errorf("Unable to read jsonFile :" + err.Error())
		return config, err
	}

	err = json.Unmarshal(byteData, &config)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshall jsonFile -> byteData :" + err.Error())
		return config, err
	}

	return config, err
}

func (c *AppConfiguration) GetConfiguration() (config *AppConfiguration) {
	return c
}
