package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ducknificient/web-intelligence/go/entity"
	"github.com/ducknificient/web-intelligence/go/general"
	"github.com/ducknificient/web-intelligence/go/service"
)

func (c *DefaultController) NewHalodocService(service service.HalodocCrawlerService) {
	c.halodocService = service
}

func (u *DefaultController) HalodocListPenyakit(w http.ResponseWriter, r *http.Request) {
	prefixLog := `HalodocListPenyakit`
	defer u.response.Panic(w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableMarshal)
		return
	}

	var (
		request entity.HalodocListPenyakitReq
	)

	err = json.Unmarshal(b, &request)
	if err != nil {
		u.response.Error(w, r, err, prefixLog, general.ConstUnableUnmarshal)
		return
	}

	var (
		dataList []entity.HalodocListPenyakit
	)

	dataList, err = u.halodocService.GetListPenyakit()
	if err != nil {
		u.response.Error(w, r, err, prefixLog, "gagal mendapatkan nama penyakit")
		return
	}

	// var listUrl []entity.AlodokterValidation
	fmt.Printf("total: %#v\n", len(dataList))

	// listUrl, err = u.alodokterService.CheckUrlIsExist(dataList)
	// if err != nil {
	// 	u.response.Error(w, r, err, prefixLog, "gagal check url")
	// 	return
	// }

	// var listNotExist []entity.AlodokterValidation

	// listNotExist = make([]entity.AlodokterValidation, 0)

	// for _, b := range listUrl {

	// 	if !b.IsExist {
	// 		fmt.Printf("url: %#v, %#v\n", b.IsExist, b.Url)
	// 		listNotExist = append(listNotExist, b)
	// 	}
	// }

	// fmt.Printf("not exist : %#v\n", len(listNotExist))

}
