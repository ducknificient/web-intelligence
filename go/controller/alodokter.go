package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/general"
	"github.com/ducknificient/web-intelligence/go/service"
)

func (c *DefaultController) NewAlodokterService(service service.AlodokterCrawlerService) {
	c.alodokterService = service
}

func (u *DefaultController) AlodokterCrawler(w http.ResponseWriter, r *http.Request) {
	prefixLog := `AlodokterCrawler`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.AlodokterCrawlerReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
		dataList []entity.AlodokterPenyakit
	)
	dataList, err = u.alodokterService.GetNamaPenyakit()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "gagal mendapatkan nama penyakit")
		return
	}

	fmt.Printf("total: %#v\n", len(dataList))

	for _, b := range dataList {

		newtask := `ALODOKTER-OBAT`
		newurl := "https://www.alodokter.com/" + b.Permalink
		// fmt.Printf("%#v\n", newurl)

		err = u.crawlerService.Crawling(newurl, newtask)
		if err != nil {
			u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
			return
		}

		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(1000000000 * time.Second)

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) AlodokterCheckUrl(w http.ResponseWriter, r *http.Request) {
	prefixLog := `AlodokterCheckUrl`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.AlodokterCrawlerReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
		dataList []entity.AlodokterPenyakit
	)
	dataList, err = u.alodokterService.GetNamaPenyakit()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "gagal mendapatkan nama penyakit")
		return
	}

	var listUrl []entity.AlodokterValidation
	fmt.Printf("total: %#v\n", len(dataList))
	listUrl, err = u.alodokterService.CheckUrlIsExist(dataList)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "gagal check url")
		return
	}

	var listNotExist []entity.AlodokterValidation

	listNotExist = make([]entity.AlodokterValidation, 0)

	for _, b := range listUrl {

		if !b.IsExist {
			fmt.Printf("url: %#v, %#v\n", b.IsExist, b.Url)
			listNotExist = append(listNotExist, b)
		}
	}

	fmt.Printf("not exist : %#v\n", len(listNotExist))

	// Create or open a JSON file for writing
	jsonData, err := json.Marshal(listNotExist)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	currentime := time.Now().Format("2006-01-02_15:04:05")

	file, err := os.Create(currentime + "_notexists_penyakit.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	u.response.Default(w, http.StatusOK, true, "ok")
	return

}
