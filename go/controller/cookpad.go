package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/general"
)

// func (c *DefaultController) NewCookpadService(service service.DefaultService) {
// 	c.cookpadService = service
// }

func (u *DefaultController) CookpadCrawler(w http.ResponseWriter, r *http.Request) {
	prefixLog := `CookpadCrawler`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CookpadCrawlerReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var menu string

	if len(menu) == 0 {
		menu = `jawa%20timur`
	} else {
		menu = request.Menu
	}

	dataList := make([]string, 50000)
	startidx := u.config.GetConfiguration().StartIndex
	for _, b := range dataList {
		startidx++

		newtask := `COOKPAD-JAWATIMUR`
		newurl := fmt.Sprintf("https://cookpad.com/id/cari/%v?event=search.suggestion&order=recent&page=%v", menu, startidx)

		fmt.Printf("%#v\n", newurl)

		err = u.defaultService.Crawling(newurl, newtask)
		if err != nil {
			u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
			return
		}

		fmt.Printf("", b)
		fmt.Printf("%v\n", newurl)

		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(1000000000 * time.Second)

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) CookpadImageCrawler(w http.ResponseWriter, r *http.Request) {
	prefixLog := `CookpadImageCrawler`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CookpadCrawlerReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var menu string

	if len(menu) == 0 {
		menu = `jawa%20timur`
	} else {
		menu = request.Menu
	}

	var (
		dataList []entity.CookpadValidation
	)

	dataList, err = u.defaultService.CookpadGetListImageUrl()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "error get list image url")
		return
	}

	for _, b := range dataList {
		newtask := `COOKPAD-JAWATIMUR-IMAGE`
		// newurl := fmt.Sprintf("https://cookpad.com/id/cari/%v?event=search.suggestion&order=recent&page=%v", menu, startidx)
		newurl := b.Url

		fmt.Printf("%#v\n", newurl)

		err = u.defaultService.Crawling(newurl, newtask)
		if err != nil {
			u.response.Error(w, r, err, prefixLog, fmt.Sprintf("Unable to crawl."))
			return
		}

		fmt.Printf("", b)
		fmt.Printf("%v\n", newurl)

		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(1000000000 * time.Second)

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}

func (u *DefaultController) CookpadGetImage(w http.ResponseWriter, r *http.Request) {
	prefixLog := `CookpadGetImage`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.CookpadCrawlerReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var menu string

	if len(menu) == 0 {
		menu = `jawa%20timur`
	} else {
		menu = request.Menu
	}

	var (
		dataList []entity.CookpadImageList
	)

	dataList, err = u.defaultService.CookpadGetImageList()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "error get list image url")
		return
	}

	fmt.Printf("%#v \n", dataList[0])

	// param := entity.CookpadSaveImageParam{
	// 	List: dataList,
	// }

	// err = u.cookpadService.SaveImageToLocal(param)
	// if err != nil {
	// 	u.response.Error(w, r, err, prefixLog, "error get list image url")
	// 	return
	// }

	u.response.Default(w, http.StatusOK, true, "ok")
	return
}
