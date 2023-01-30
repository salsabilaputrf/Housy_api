package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"
	"housy/pkg/middleware"

	"github.com/gorilla/mux"
)

func HouseRoutes(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	r.HandleFunc("/houses", middleware.Auth(h.FindHouses)).Methods("GET")
	r.HandleFunc("/house/{id_house}", h.GetHouse).Methods("GET")
	r.HandleFunc("/house", middleware.Auth(h.CreateHouse)).Methods("POST")
	
}