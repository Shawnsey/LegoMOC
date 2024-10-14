package handler

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/shawnsey/LegoMOC/LegoStore/database/order"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
)

type Order struct {
	DB *order.Postgresdb
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerId int32   `json:"customer_id"`
		LineItems  *string `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	order := model.Order{
		OrderID:    rand.Int32(),
		CustomerID: body.CustomerId,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err := o.DB.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("failed to insert", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshall", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List orders")

}
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order")

}
func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order")

}
func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete order")

}
