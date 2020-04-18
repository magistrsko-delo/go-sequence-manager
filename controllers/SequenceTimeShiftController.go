package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-sequence-manager/Models/DTO"
	"go-sequence-manager/services"
	"net/http"
	"strconv"
)

type SequenceTimeShiftController struct {
	SequenceTimeShiftService *services.SequenceTimeShiftService
}

func (sequenceTimeShiftController *SequenceTimeShiftController) GetSequenceChunks (w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	sequenceId, err := strconv.Atoi(params["sequenceId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp, err := sequenceTimeShiftController.SequenceTimeShiftService.GetSequenceTimeShiftData(int32(sequenceId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&DTO.ResponseDTO{
		Status:  200,
		Message: "",
		Data:    rsp,
	})

}
