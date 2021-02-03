package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestSaveStatistics(t *testing.T) {
	client := &http.Client{}

	url := "http://0.0.0.0:7890/save_stats/"
	params := map[string]string{
		"date":   "2017-11-30",
		"views":  "120",
		"clicks": "240",
		"cost":   "34r 20k",
	}

	err, res := makePostMultipartRequest(client, url, params)
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestRetrieveStatistics(t *testing.T) {
	from := "2017-11-30"
	to := "2077-12-25"

	url := "http://0.0.0.0:7890/retrieve_stats/?to=" + to + "&from=" + from
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestDeleteStatistics(t *testing.T) {
	url := "http://0.0.0.0:7890/delete_stats/"

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Error(err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func makePostMultipartRequest(client *http.Client, url string, params map[string]string) (err error, res *http.Response) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for key, value := range params {
		err := w.WriteField(key, value)
		if err != nil {
			log.Println(err)
		}
	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err = client.Do(req)
	if err != nil {
		return err, nil
	}

	return nil, res
}
