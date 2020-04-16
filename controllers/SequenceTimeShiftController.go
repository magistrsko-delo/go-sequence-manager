package controllers

import (
	"github.com/gorilla/mux"
	"go-sequence-manager/services"
	"net/http"
	"strconv"
)

type SequenceTimeShiftController struct {
	SequenceTimeShiftService *services.SequenceTimeShiftService
}

func (SequenceTimeShiftController *SequenceTimeShiftController) GetSequenceChunks (w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	_, err := strconv.Atoi(params["sequenceId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
