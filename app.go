package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/denisenkom/go-mssqldb"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host, user, password, dbname string) {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", user, password, host, dbname)

	var err error
	a.DB, err = sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {

}