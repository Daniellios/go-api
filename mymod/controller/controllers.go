package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Daniellios/mymodules/model"
	"github.com/gorilla/mux"
)

func HandleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>My API</h1>"))

}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получить все заказы")
	w.Header().Set("Content-Type", "application/json")

	if len(model.Orders) > 0 {
		json.NewEncoder(w).Encode(model.Orders)

	} else {
		json.NewEncoder(w).Encode("Список заказов пуст")
	}
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получить Заказ")
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for _, order := range model.Orders {
		if order.OrderId == param["id"] {
			json.NewEncoder(w).Encode(order)
			return
		}
	}

	json.NewEncoder(w).Encode("Заказа с таким Id не существует")
	return

}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Создать заказ")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Пустое тело запроса")
	}

	var order model.Order
	_ = json.NewDecoder(r.Body).Decode(&order)

	if order.IsEmpty() {
		json.NewEncoder(w).Encode("Укажите свое имя")
		return
	}

	rand.Seed(time.Now().UnixNano())

	order.OrderId = strconv.Itoa(rand.Intn(999))

	model.Orders = append(model.Orders, order)
	json.NewEncoder(w).Encode(order)
	return
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Обновить Заказ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, order := range model.Orders {
		if order.OrderId == params["id"] {
			model.Orders = append(model.Orders[:index], model.Orders[index+1:]...)
			var order model.Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.OrderId = params["id"]
			model.Orders = append(model.Orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}

}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Удалить заказ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, order := range model.Orders {
		if order.OrderId == params["id"] {
			model.Orders = append(model.Orders[:index], model.Orders[index+1:]...)
			json.NewEncoder(w).Encode("Заказ удален")
			break
		} else {
			json.NewEncoder(w).Encode("Такого заказа не найдено")
		}

	}

	json.NewEncoder(w).Encode(model.Orders)
	return
}
