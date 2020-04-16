package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-sequence-manager/Models/DTO"
	"go-sequence-manager/services"
	"net/http"
	"strconv"
)

type SequenceMediaController struct {
	SequenceMediaService *services.SequenceMediaService
}

func (sequenceMediaController *SequenceMediaController) GetSequenceMedia(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	sequenceId, err := strconv.Atoi(params["sequenceId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp, err := sequenceMediaController.SequenceMediaService.GetSequenceMedias(int32(sequenceId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&DTO.ResponseDTO{
		Status:  0,
		Message: "",
		Data:    rsp,
	})
}
