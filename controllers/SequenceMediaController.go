package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-sequence-manager/Models/DTO"
	DTO_input "go-sequence-manager/Models/DTO/input"
	"go-sequence-manager/services"
	"net/http"
	"strconv"
)

type SequenceMediaController struct {
	SequenceMediaService *services.SequenceMediaService
	PublishSequenceService *services.PublishSequenceService
}

func (sequenceMediaController *SequenceMediaController) AddMediaToSequence (w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	sequenceId, err := strconv.Atoi(params["sequenceId"])
	mediaId, err := strconv.Atoi(params["mediaId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp, err := sequenceMediaController.SequenceMediaService.AddMediaToSequence(int32(sequenceId), int32(mediaId))

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

func (sequenceMediaController *SequenceMediaController) DeleteMediaFromSequence (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sequenceId, err := strconv.Atoi(params["sequenceId"])
	mediaId, err := strconv.Atoi(params["mediaId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp, err := sequenceMediaController.SequenceMediaService.DeleteMediaFromSequence(int32(sequenceId), int32(mediaId))

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
		Status:  200,
		Message: "",
		Data:    rsp,
	})
}

// sequence publish request..
func (sequenceMediaController *SequenceMediaController) PublishSequence(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	sequenceId, err := strconv.Atoi(params["sequenceId"])

	inputPublish := &DTO_input.InputPublishSequence{}
	err = json.NewDecoder(r.Body).Decode(inputPublish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if int32(sequenceId) != inputPublish.SequenceId {
		http.Error(w, "sequence id-s does not match", http.StatusBadRequest)
		return
	}

	rsp, err := sequenceMediaController.PublishSequenceService.PublishSequence(inputPublish)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&DTO.ResponseDTO{
		Status:  200,
		Message: "Sequence added to publish process",
		Data:    rsp,
	})
}
