package main

import (
	"encoding/json"
	"log"
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

func responseShortcut(w http.ResponseWriter, statusCode int, responseMessage interface{}) {
	response := Response{
		w:               w,
		statusCode:      statusCode,
		responseMessage: responseMessage,
	}
	response.jsonResponse()
}

func (a *App) saveStatistics(w http.ResponseWriter, r *http.Request) {
	date := r.PostFormValue("date")
	views := r.PostFormValue("views")
	clicks := r.PostFormValue("clicks")
	cost := r.PostFormValue("cost")
	var err error
	var matches bool
	if date == "" {
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"status": false, "result": "Date is the mandatory field"})
		return
	}
	matches, err = regexp.MatchString(yyyyMMDDRegex, date)
	if err != nil {
		panic(err)
	}
	if !matches {
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"status": false, "result": "The format of date should be 'YYYY-mm-dd'"})
		return
	}

	err = a.saveStatisticsToDatabase(date, views, clicks, cost)
	if err != nil {
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"status": false, "result": "Failed to save statistics"})
		return
	}

	responseShortcut(w, http.StatusOK, map[string]interface{}{"status": true, "result": "OK"})
}

func (a *App) retrieveStatistics(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	to := r.FormValue("to")
	var err error
	var matchesFrom bool
	var matchesTo bool
	var statsList []Statistics

	if from == "" || to == "" {
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"status": false, "result": "Please, fill out all necessary fields"})
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
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"status": false, "result": "The format of date should be 'YYYY-mm-dd'"})
		return
	}

	data, err := a.retrieveStatisticsFromDatabase(from, to)
	if err != nil {
		responseShortcut(w, http.StatusBadRequest, map[string]interface{}{"status": false, "result": "Failed to load statistics"})
		return
	}
	for i := range data {
		var stats Statistics
		b, err := json.Marshal(data[i])
		if err != nil {
			log.Fatal(err)
		}
		err_ := json.Unmarshal(b, &stats)
		if err_ != nil {
			log.Fatal(err_)
		}
		statsList = append(statsList, stats)
	}

	responseShortcut(w, http.StatusOK, map[string]interface{}{"status": true, "result": statsList})
}

func (a *App) deleteAllStatistics(w http.ResponseWriter, r *http.Request) {
	err := a.deleteAllStatisticsFromDatabase()
	if err != nil {
		responseShortcut(w, http.StatusBadRequest,
			map[string]interface{}{"success": false, "result": "Failed to delete statistics from database"})
	}
}
