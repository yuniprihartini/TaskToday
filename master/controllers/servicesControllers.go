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

type ServicesHandler struct {
	ServicesUseCase usecase.ServicesUseCase
}

func ServicesControllers(r *mux.Router, service usecase.ServicesUseCase) {
	ServicesHandler := ServicesHandler{service}
	r.HandleFunc("/allservices", ServicesHandler.GetAllServices).Methods(http.MethodGet)
	r.HandleFunc("/services/{id}", ServicesHandler.ServicesById).Methods(http.MethodGet)
	r.HandleFunc("/services", ServicesHandler.CreateServices).Methods(http.MethodPost)
	r.HandleFunc("/services/{id}", ServicesHandler.DeleteServices).Methods(http.MethodDelete)
	r.HandleFunc("/services/{id}", ServicesHandler.UpdateServices).Methods(http.MethodPut)
}

func (kt ServicesHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {
	Servicess, err := kt.ServicesUseCase.GetAllDataService()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfServicess, _ := json.Marshal(Servicess)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfServicess)
}

func (kt ServicesHandler) ServicesById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param["id"])
	Services, err := kt.ServicesUseCase.GetDataById(param["id"])
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfServicess, _ := json.Marshal(Services)
	if err != nil {
		w.Write([]byte("Oops, somethin when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfServicess)
}

func (p ServicesHandler) CreateServices(w http.ResponseWriter, r *http.Request) {
	var Services models.Services
	err := json.NewDecoder(r.Body).Decode(&Services)

	var isTrue = utils.ValidationServices(&Services)
	tools.PrintlnErr(err)
	if isTrue == false {
		w.Write([]byte("Vehicle Type or Service Type must be filled"))
	} else {
		err = p.ServicesUseCase.InsertDataService(Services)
		tools.PrintlnErr(err)
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}

}

func (p ServicesHandler) DeleteServices(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := p.ServicesUseCase.DeleteDataServiceById(param["id"])
	tools.PrintlnErr(err)
	log.Println("Delete successful")
	w.Write([]byte("Delete successful"))
}

func (p ServicesHandler) UpdateServices(w http.ResponseWriter, r *http.Request) {
	var Services models.Services
	param := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&Services)
	var isTrue = utils.ValidationServices(&Services)
	if isTrue == false {
		w.Write([]byte("Vehicle Type or Service Type must be filled"))
	} else {
		err := p.ServicesUseCase.UpdateDataService(param["id"], Services)
		tools.PrintlnErr(err)
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
