package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
		fmt.Println("Failed to create uuid",err)
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

	writeJSONResponse(w,order, 201)
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

	writeJSONResponse(w, orders, 200)
}

func (o *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	parsedUUID := getValidUuid(r)

	order,err := o.OrderDao.GetById(r.Context(), parsedUUID)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, order, 200)
}

func (o *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {

	var body struct {
		OrderId uuid.UUID
		ShippedAt  *string `json:"shipped_at"`
		CompletedAt *string `json:"completed_at"`
	}
	requestBody := daos.OrderUpdateBody{}

	parsedUUID := getValidUuid(r)

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: Request body malformation"))
		return
	}
	if (body.ShippedAt == nil && body.CompletedAt == nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: No data to update in request"))
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
	
	writeJSONResponse(w, updatedOrder, 202)
}

func (o *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

	parsedUUID := getValidUuid(r)


	err := o.OrderDao.Delete(r.Context(), parsedUUID)
	if err != nil {
		fmt.Println("Failed to delete order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
