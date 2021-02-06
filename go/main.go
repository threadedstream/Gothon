package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	Router *mux.Router
	Conn   *sql.DB
	Server *http.Server
}

type Statistics struct {
	Date   time.Time `json:"date"`
	Views  int       `json:"views"`
	Clicks int       `json:"clicks"`
	Cost   string    `json:"cost"`
	Cpc    string    `json:"cpc"`
	Cpm    string    `json:"cpm"`
}

func (a *App) initRoutes() {
	a.Router.Path("/save_stats/").HandlerFunc(a.saveStatistics).Methods("POST")
	a.Router.Path("/retrieve_stats/").HandlerFunc(a.retrieveStatistics).Methods("GET")
	a.Router.Path("/delete_stats/").HandlerFunc(a.deleteAllStatistics).Methods("DELETE")
}

func (a *App) initialize(user, password, dbname, addr, host string, port int) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	a.Conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		return
	}
	a.Router = mux.NewRouter()

	a.Server = &http.Server{
		Addr:         addr,
		Handler:      a.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//Initialize router and routes
	a.initRoutes()
}

func (a *App) Run() {
	log.Printf("Starting server on %s\n", a.Server.Addr)
	log.Fatal(a.Server.ListenAndServe())
}

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}
	addr := os.Getenv("ADDR")

	a := App{}

	a.initialize(user, password, dbname, addr, host, port)
	a.initRoutes()
	//a.checkSellerOfferExistence(1,1)
	log.Println("Welcome, dear inhabitant of Avito world!")
	a.Run()
}
