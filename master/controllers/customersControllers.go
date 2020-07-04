package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ujian/master/models"
	"ujian/master/usecase"
	"ujian/master/utils"
	"ujian/tools"

	"github.com/gorilla/mux"
)

type CustomersHandler struct {
	CustomersUseCase usecase.CustomersUseCase
}

func CustomersControllers(r *mux.Router, service usecase.CustomersUseCase) {
	CustomersHandler := CustomersHandler{service}
	r.HandleFunc("/allcustomers", CustomersHandler.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{id}", CustomersHandler.CustomersById).Methods(http.MethodGet)
	r.HandleFunc("/customers", CustomersHandler.CreateCustomers).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id}", CustomersHandler.DeleteCustomers).Methods(http.MethodDelete)
	r.HandleFunc("/customers/{id}", CustomersHandler.UpdateCustomers).Methods(http.MethodPut)
}

func (kt CustomersHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	Customerss, err := kt.CustomersUseCase.GetAllDataCustomers()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfCustomerss, _ := json.Marshal(Customerss)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfCustomerss)
}

func (kt CustomersHandler) CustomersById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param["id"])
	Customers, err := kt.CustomersUseCase.GetDataById(param["id"])
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfCustomerss, _ := json.Marshal(Customers)
	if err != nil {
		w.Write([]byte("Oops, somethin when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfCustomerss)
}

func (p CustomersHandler) CreateCustomers(w http.ResponseWriter, r *http.Request) {
	var Customers models.Customers
	err := json.NewDecoder(r.Body).Decode(&Customers)

	var isTrue = utils.ValidationCustomers(&Customers)
	tools.PrintlnErr(err)
	if isTrue == false {
		w.Write([]byte("Customer Name, Booking Date, Number Plat must be filled"))
	} else {
		err = p.CustomersUseCase.InsertDataCustomers(Customers)
		tools.PrintlnErr(err)
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}

}

func (p CustomersHandler) DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := p.CustomersUseCase.DeleteDataCustomersById(param["id"])
	tools.PrintlnErr(err)
	log.Println("Delete successful")
	w.Write([]byte("Delete successful"))
}

func (p CustomersHandler) UpdateCustomers(w http.ResponseWriter, r *http.Request) {
	var Customers models.Customers
	param := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&Customers)
	var isTrue = utils.ValidationCustomers(&Customers)
	if isTrue == false {
		w.Write([]byte("Customer Name, Booking Date, Number Plat must be filled"))
	} else {
		err := p.CustomersUseCase.UpdateDataCustomers(param["id"], Customers)
		tools.PrintlnErr(err)
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
