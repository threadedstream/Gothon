package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

const (
	yyyyMMDDRegex = "([12]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01]))"
)

type Response struct {
	w               http.ResponseWriter
	statusCode      int
	responseMessage interface{}
}

func (r *Response) jsonResponse() {
	response, _ := json.Marshal(r.responseMessage)

	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(r.statusCode)
	r.w.Write(response)
}

func (a *App) saveStatistics(w http.ResponseWriter, r *http.Request) {
	date := r.PostFormValue("date")
	views := r.PostFormValue("views")
	clicks := r.PostFormValue("clicks")
	cost := r.PostFormValue("cost")
	var err error
	var matches bool
	if date == "" {
		response := Response{
			w:               w,
			statusCode:      http.StatusBadRequest,
			responseMessage: map[string]interface{}{"status": false, "result": "Date is the mandatory field"},
		}
		response.jsonResponse()
		return
	}

	matches, err = regexp.MatchString(yyyyMMDDRegex, date)
	if err != nil {
		panic(err)
	}
	if !matches {
		response := Response{
			w:               w,
			statusCode:      http.StatusBadRequest,
			responseMessage: map[string]interface{}{"status": false, "result": "The format of date should be 'YYYY-mm-dd'"},
		}
		response.jsonResponse()
		return
	}

	err = a.saveStatisticsToDatabase(date, views, clicks, cost)
	if err != nil {
		response := Response{
			w:               w,
			statusCode:      http.StatusBadRequest,
			responseMessage: map[string]interface{}{"status": false, "result": "Failed to save statistics"},
		}
		response.jsonResponse()
		return
	}

	response := Response{
		w:               w,
		statusCode:      http.StatusOK,
		responseMessage: map[string]interface{}{"status": true, "result": "OK"},
	}
	response.jsonResponse()
}

func (a *App) retrieveStatistics(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	to := r.FormValue("to")
	var err error
	var matchesFrom bool
	var matchesTo bool

	if from == "" || to == "" {
		response := Response{
			w:               w,
			statusCode:      http.StatusBadRequest,
			responseMessage: map[string]interface{}{"status": false, "result": "Please, fill out all necessary fields"},
		}
		response.jsonResponse()
		return
	}
	matchesFrom, err = regexp.MatchString(yyyyMMDDRegex, from)
	if err != nil {
		panic(err)
	}
	matchesTo, err = regexp.MatchString(yyyyMMDDRegex, to)
	if err != nil {
		panic(err)
	}
	if !(matchesFrom && matchesTo) {
		response := Response{
			w:               w,
			statusCode:      http.StatusBadRequest,
			responseMessage: map[string]interface{}{"status": false, "result": "The format of date should be 'YYYY-mm-dd'"},
		}
		response.jsonResponse()
		return
	}
}
