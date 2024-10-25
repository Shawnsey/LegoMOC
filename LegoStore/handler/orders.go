package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/shawnsey/LegoMOC/LegoStore/database/daos"
	"github.com/shawnsey/LegoMOC/LegoStore/sql/LegoMOC/public/model"
)

type OrderHandler struct {
    OrderDao daos.OrderDao
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerId *uuid.UUID   `json:"customer_id"`
		LineItems  *string `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order_id, err := uuid.NewUUID()
	if err != nil {
		fmt.Errorf("Failed to create uuid",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()

	order := model.Orders{
		OrderID:    order_id,
		CustomerID: body.CustomerId,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err = o.OrderDao.Insert(r.Context(), order)
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
	// todo: need to sort out how to get ownerId from a request to return list of 
	// orders that a customer has. Could implement keycloak jwt tokens and authorization
	// for now just have a static id to return orders from.
	var customerId = "98311391-88ca-48c7-ad1d-5ccb7fcb4e19"
	orders, err := o.OrderDao.List(r.Context(), uuid.MustParse(customerId)) 
	if err != nil {
		log.Printf("Error with Database, error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(orders)
	if err != nil {
		log.Printf("Error with marshaling orders, error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusAccepted)
}

func (o *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid uuid")
		w.Write([]byte("Bad Request: Invalid uuid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order,err := o.OrderDao.GetById(r.Context(), parsedUUID)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error with marshaling orders, error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusAccepted)

}

func (o *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {

	var body struct {
		OrderId uuid.UUID
		ShippedAt  *string `json:"shipped_at"`
		CompletedAt *string `json:"completed_at"`
	}
	requestBody := daos.OrderUpdateBody{}

	id := chi.URLParam(r, "id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid uuid")
		w.Write([]byte("Bad Request: Invalid uuid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode", err)
		w.Write([]byte("Bad Request: Request body malformation"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (body.ShippedAt == nil && body.CompletedAt == nil) {
		w.Write([]byte("Bad Request: No data to update in request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if body.ShippedAt != nil{
		shippedAt, err := time.Parse(time.RFC3339, *body.ShippedAt)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		fmt.Printf("shippedAT: %s", shippedAt)
		requestBody.ShippedAt = shippedAt
	}
	if body.CompletedAt != nil {
		completeAt, err := time.Parse(time.RFC3339, *body.CompletedAt)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		fmt.Printf("completedAt: %s", completeAt)
		requestBody.CompletedAt = completeAt
	}
	requestBody.OrderId = parsedUUID
	

	updatedOrder, err := o.OrderDao.Update(r.Context(), requestBody)
	if err != nil {
		fmt.Println("Failed to Update order")
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

	id := chi.URLParam(r, "id")
	fmt.Printf("uuid from path: %s", id)

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid uuid")
		w.Write([]byte("Bad Request: Invalid uuid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = o.OrderDao.Delete(r.Context(), parsedUUID)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
