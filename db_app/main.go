package main

import (
	"db_app/postgres"
	"fmt"
)

type DataStore interface {
	GetEmployee(id int) string
}

type App struct {
	ds DataStore
}

func main() {
	app := App{}
	app.ds = &postgres.PsqlStore{}
	//app.ds = &postgres.PsqlStore{}
	app.ds = &postgres.PsqlStore{}
	fmt.Printf(app.ds.GetEmployee(22))
}
