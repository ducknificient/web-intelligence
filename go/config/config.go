package config

import (
	"encoding/json"
	"io"
	"os"
)

var (
	Conf Configuration
)

// type DbInfo struct {
// 	Host        *string `json:"host"` // server ip
// 	DbName      *string `json:"dbname"`
// 	Username    *string `json:"username"`
// 	Password    *string `json:"password"`
// 	Port        *string `json:"port"`
// 	SSLMode     *string `json:"sslmode"`
// 	PoolMaxConn *string `json:"poolmaxconn"`
// }

type PgInfo struct {
	Host        *string `json:"host"` // server ip
	DbName      *string `json:"dbname"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Port        *string `json:"port"`
	SSLMode     *string `json:"sslmode"`
	PoolMaxConn *string `json:"poolmaxconn"`
}

type Configuration struct {
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

func NewConfiguration(environment string) (config Configuration) {

	var (
		filename      string
		environmentOK bool = true
	)

	switch environment {
	case "dev":
		filename = "config/config-dev.json"
	case "prod-test":
		filename = "config/config-prod-test.json"
	case "prod-api6":
		filename = "config/config-prod-api6.json"
	case "prod-nanika":
		filename = "config/config-prod-nanika.json"
	default:
		filename = "config/config.json"
		environmentOK = false
	}

	if !environmentOK {
		panic("Unable to find suitable environment")
	}

	jsonFile, err := os.Open(filename)
	if err != nil {
		panic("Unable to load config.json :" + err.Error())
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		panic("Unable to read jsonFile :" + err.Error())
	}

	err = json.Unmarshal(byteData, &config)
	if err != nil {
		panic("Unable to unmarshall jsonFile -> byteData :" + err.Error())
	}

	return config
}
