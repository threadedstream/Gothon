package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSaveStatisticsToDatabase(t *testing.T) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	a := App{}

	a.initialize(user, password, dbname, "")
	date := "2077-12-30"
	views := "2077"
	clicks := "2077"
	cost := "2077"
	err := a.saveStatisticsToDatabase(date, views, clicks, cost)
	assert.Equal(t, err, nil)
}

func TestRetrieveStatisticsFromDatabase(t *testing.T) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	a := App{}

	a.initialize(user, password, dbname, "")

	from := "2077-12-25"
	to := "2078-01-05"

	rows, err := a.retrieveStatisticsFromDatabase(from, to)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, rows, nil)
}
