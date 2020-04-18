package router

import (
	"github.com/gorilla/mux"
	"go-sequence-manager/controllers"
	"go-sequence-manager/services"
)

type SequenceManagerRouter struct {
	Router *mux.Router
}



func (sequenceManagerRouter *SequenceManagerRouter) RegisterHandlers()  {
	sequenceMediaController := controllers.SequenceMediaController{SequenceMediaService: services.InitSequenceMediaService()}
	seqeunceTimeshiftController := controllers.SequenceTimeShiftController{SequenceTimeShiftService:services.InitSequenceTimeShiftService()}

	(*sequenceManagerRouter).Router.HandleFunc("/sequence/{sequenceId}/media", sequenceMediaController.GetSequenceMedia).Methods("GET")
	(*sequenceManagerRouter).Router.HandleFunc("/sequence/{sequenceId}/media/{mediaId}", sequenceMediaController.AddMediaToSequence).Methods("POST")
	(*sequenceManagerRouter).Router.HandleFunc("/sequence/{sequenceId}/media/{mediaId}", sequenceMediaController.DeleteMediaFromSequence).Methods("DELETE")

	(*sequenceManagerRouter).Router.HandleFunc("/sequence/{sequenceId}/chunks", seqeunceTimeshiftController.GetSequenceChunks).Methods("GET")
}