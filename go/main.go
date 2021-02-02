package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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
	Cost   int       `json:"cost"`
	Cpc    float32   `json:"cpc"`
	Cpm    float32   `json:"cpm"`
}

func (a *App) initRoutes() {
	a.Router.Path("/save_stats/").HandlerFunc(a.saveStatistics).Methods("POST")
}

func (a *App) initialize(user, password, dbname, addr string) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "192.168.1.40", 5432, user, password, dbname)

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
	addr := os.Getenv("ADDR")
	a := App{}

	a.initialize(user, password, dbname, addr)
	a.initRoutes()
	//a.checkSellerOfferExistence(1,1)
	log.Println("Welcome, dear inhabitant of Avito world!")
	a.Run()
}
