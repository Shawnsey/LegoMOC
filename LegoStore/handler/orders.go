package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Order struct {
	db sql.DB
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
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
