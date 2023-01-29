package main

import (
	"log"
	"net/http"

	с "github.com/Daniellios/mymodules/controller"
	m "github.com/Daniellios/mymodules/model"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	intialData := []m.Order{
		{OrderId: "2", OrderName: "TV", OrderPrice: 400, Client: &m.Client{Fullname: "Danill Blinnikov", Phone: "+79999999999", Password: "124r255"}},
		{OrderId: "5", OrderName: "SmartPhone", OrderPrice: 500, Client: &m.Client{Fullname: "Ivan Ivanov", Phone: "+7888888888", Password: "124r6666"}},
		{OrderId: "10", OrderName: "Drone", OrderPrice: 899, Client: &m.Client{Fullname: "Petr Semenov", Phone: "+7777777777", Password: "124r66zfw6"}},
	}

	m.Orders = append(m.Orders, intialData...)

	r.HandleFunc("/", с.HandleHomeRoute).Methods("GET")
	r.HandleFunc("/orders", с.GetOrders).Methods("GET")
	r.HandleFunc("/order/{id}", с.GetOrder).Methods("GET")
	r.HandleFunc("/order", с.CreateOrder).Methods("POST")
	r.HandleFunc("/order/{id}", с.UpdateOrder).Methods("PUT")
	r.HandleFunc("/order/{id}", с.DeleteOrder).Methods("DELETE")

	// listeting to the port
	log.Fatal(http.ListenAndServe(":4000", r))

}
