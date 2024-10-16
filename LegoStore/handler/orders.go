package handler

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shawnsey/LegoMOC/LegoStore/database/daos"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
)

type OrderHandler struct {
    OrderDao daos.OrderDao
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	err := o.OrderDao.Insert(r.Context(), order)
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
func (o *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add List orders")

}
func (o *OrderHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Get order")

}
func (o *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	parsedId, err := o.parseInt32(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	updatedOrder, err := o.OrderDao.Update(r.Context(), parsedId)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	
	res, err := json.Marshal(updatedOrder)
	if err != nil {
		fmt.Println("failed to marshall", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusAccepted)

}
func (o *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	parsedId, err := o.parseInt32(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}


	err = o.OrderDao.Delete(r.Context(),parsedId)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (o *OrderHandler) parseInt32(id string) (int32, error) {

    num, err := strconv.ParseInt(id, 10, 32) // Base 10, 32-bit size
    if err != nil {
        return 0, err // Return the parsing error if it fails
    }
	
    // Convert the int64 to int32 after ensuring it's in the valid range
    if num < int64(^int32(0)) && num > int64(int32(uint32(0))) { // Check the int32 range
        return 0, fmt.Errorf("value out of range for int32: %s", id)
    }

    return int32(num), nil
}
