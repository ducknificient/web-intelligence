package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ducknificient/web-intelligence/go/logger"
)

type Response interface {
	Error(w http.ResponseWriter, r *http.Request, err error, prefixLog string, errMsg string)
	Panic(w http.ResponseWriter, r *http.Request)
	Default(w http.ResponseWriter, httpStatus int, status bool, msg string)
	DefaultText(w http.ResponseWriter, httpStatus int, status bool, msg string)
	Forbidden(w http.ResponseWriter, r *http.Request, err error, prefixLog, msg string)
	CustomResponse(w http.ResponseWriter, r *http.Request, prefixLog string, msg string, any any)
}

type JsonRs struct {
	Sc  int    `json:"sc"`
	St  bool   `json:"st"`
	Msg string `json:"msg"`
}

type JsonRsCustom struct {
	Sc   int    `json:"sc"`
	St   bool   `json:"st"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type DefaultResponse struct {
	Logger logger.Logger
}

func (res *DefaultResponse) Error(w http.ResponseWriter, r *http.Request, err error, prefixLog string, errMsg string) {
	if err != nil {
		res.Logger.Error(prefixLog + `. ` + errMsg + `. Error : ` + err.Error())
	} else {
		res.Logger.Error(prefixLog + errMsg + ". Error var empty.")
	}

	// enableCors(r, &w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	rs, err := json.Marshal(JsonRs{http.StatusInternalServerError, false, "There is problem with your request. Please contact our representative."})
	// rs, err := json.Marshal(JsonRs{http.StatusInternalServerError, false, prefixLog + `. Error : ` + err.Error() + `. ` + errMsg})
	if err != nil {
		res.Logger.Error(prefixLog + errMsg + ". Detail : " + err.Error())
	}
	w.Write(rs)

	return
}

func (res *DefaultResponse) Panic(w http.ResponseWriter, r *http.Request) {
	if rc := recover(); rc != nil {
		debugStack := string(debug.Stack())
		res.Logger.Error(fmt.Sprintf("Panic - *Panic is occured*"))
		err := errors.New("Error inside Panic")

		ginres := DefaultResponse{res.Logger}
		ginres.Error(w, r, err, "Panic - ", fmt.Sprintf("Panic : %v", debugStack))
	}
}

func (res *DefaultResponse) Default(w http.ResponseWriter, httpStatus int, status bool, msg string) {
	rs, err := json.Marshal(JsonRs{httpStatus, status, msg})
	if err != nil {
		res.Logger.Error("Default - " + err.Error())
		// panic("Default: - " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(rs)

	return
}

func (res *DefaultResponse) DefaultText(w http.ResponseWriter, httpStatus int, status bool, msg string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(httpStatus)
	w.Write([]byte(msg))

	return
}

func (res *DefaultResponse) Forbidden(w http.ResponseWriter, r *http.Request, err error, prefixLog string, msg string) {
	if err != nil {
		res.Logger.Error(prefixLog + ". Detail : " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	rs, err := json.Marshal(JsonRs{http.StatusForbidden, false, msg})
	if err != nil {
		res.Logger.Error(prefixLog + ". Detail : " + err.Error())
	}
	w.Write(rs)
	return
}

func (res *DefaultResponse) CustomResponse(w http.ResponseWriter, r *http.Request, prefixLog string, msg string, any any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rs, err := json.Marshal(JsonRsCustom{http.StatusOK, true, msg, any})
	if err != nil {
		res.Logger.Error(prefixLog + err.Error() + ". Detail : " + err.Error())
	}
	w.Write(rs)
	return
}
