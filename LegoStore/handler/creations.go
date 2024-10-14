package handler

import (
	"github.com/shawnsey/LegoMOC/LegoStore/database/daos"
	"fmt"
	"net/http"
)

type CreationsHandler struct {
    CreationsDao daos.CreationsDao
}

func NewUserHandler(CreationsDao daos.CreationsDao) *CreationsHandler {
    return &CreationsHandler{CreationsDao: CreationsDao}
}

func (c *CreationsHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add creation")
}

func (c *CreationsHandler) Search(w http.ResponseWriter, r *http.Request) {
	queryParameters := r.URL.Query()
	parameterList := make(map[string][]string)

	for key,values := range queryParameters {
		parameterList[key] = values
	}
	for key, values := range parameterList {
		fmt.Printf("Parameter key: %s, value: %s", key, values)
	}

}

func (o *CreationsHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List orders")

}
func (o *CreationsHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order")

}
func (o *CreationsHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order")

}
func (o *CreationsHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete order")

}
