package main

import (
	"housy/database"
	"housy/pkg/mysql"
	"housy/routes"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		fmt.Println(err)
	}
}