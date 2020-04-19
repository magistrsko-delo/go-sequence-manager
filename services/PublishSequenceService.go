package services

import (
	"encoding/json"
	"errors"
	DTO_input "go-sequence-manager/Models/DTO/input"
	"go-sequence-manager/grpc_client"
)

type PublishSequenceService struct {
	sequenceGrpcClient *grpc_client.SequenceServiceClient
	rabbitmq *RabbitMQ
}

func (publishSequenceService *PublishSequenceService) PublishSequence(publishInput *DTO_input.InputPublishSequence) (bool, error) {

	sequenceData, err := publishSequenceService.sequenceGrpcClient.GetSequenceMedia(publishInput.SequenceId)
	if err != nil {
		return false, err
	}

	if sequenceData.GetSequence().GetStatus() >= 3 {
		return false, errors.New("sequence status 3 or greater.. already published")
	}

	publishSequenceMessageForQueue, err := json.Marshal(publishInput)

	if err != nil {
		return false, err
	}

	err = publishSequenceService.rabbitmq.MessageToQueue(publishSequenceMessageForQueue)
	if err != nil {
		return false, err
	}

	return true, nil
}

func InitPublishSequenceService() *PublishSequenceService  {

	return &PublishSequenceService{
		sequenceGrpcClient: grpc_client.InitSequenceServiceMetadata(),
		rabbitmq:           InitRabbitMQ(),
	}
}