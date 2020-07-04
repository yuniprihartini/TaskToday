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

type progressHandler struct {
	progressUseCase usecase.ProgressUseCase
}

func ProgressControllers(r *mux.Router, service usecase.ProgressUseCase) {
	progressHandler := progressHandler{service}
	r.HandleFunc("/allprogress", progressHandler.GetAllProgress).Methods(http.MethodGet)
	r.HandleFunc("/progress/{id}", progressHandler.ProgressById).Methods(http.MethodGet)
	r.HandleFunc("/progress", progressHandler.CreateProgress).Methods(http.MethodPost)
	r.HandleFunc("/progress/{id}", progressHandler.DeleteProgress).Methods(http.MethodDelete)
	r.HandleFunc("/progress/{id}", progressHandler.UpdateProgress).Methods(http.MethodPut)
}

func (kt progressHandler) GetAllProgress(w http.ResponseWriter, r *http.Request) {
	progresss, err := kt.progressUseCase.GetAllDataProgress()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfprogresss, _ := json.Marshal(progresss)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfprogresss)
}

func (kt progressHandler) ProgressById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param["id"])
	progress, err := kt.progressUseCase.GetDataById(param["id"])
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfprogresss, _ := json.Marshal(progress)
	if err != nil {
		w.Write([]byte("Oops, somethin when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfprogresss)
}

func (p progressHandler) CreateProgress(w http.ResponseWriter, r *http.Request) {
	var progress models.Progress
	err := json.NewDecoder(r.Body).Decode(&progress)

	var isTrue = utils.ValidationProgress(&progress)
	tools.PrintlnErr(err)
	if isTrue == false {
		w.Write([]byte("Customer Id or Service Id must be filled"))
	} else {
		err = p.progressUseCase.InsertDataProgress(progress)
		tools.PrintlnErr(err)
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}

}

func (p progressHandler) DeleteProgress(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := p.progressUseCase.DeleteDataProgressById(param["id"])
	tools.PrintlnErr(err)
	log.Println("Delete successful")
	w.Write([]byte("Delete successful"))
}

func (p progressHandler) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	var progress models.Progress
	param := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&progress)
	var isTrue = utils.ValidationProgress(&progress)
	if isTrue == false {
		w.Write([]byte("Customer Id or Service Id must be filled"))
	} else {
		err := p.progressUseCase.UpdateDataProgress(param["id"], progress)
		tools.PrintlnErr(err)
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
