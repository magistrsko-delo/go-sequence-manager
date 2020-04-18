package services

import (
	"encoding/json"
	DTO_input "go-sequence-manager/Models/DTO/input"
	"go-sequence-manager/grpc_client"
)

type PublishSequenceService struct {
	sequenceGrpcClient *grpc_client.SequenceServiceClient
	rabbitmq *RabbitMQ
}

func (publishSequenceService *PublishSequenceService) PublishSequence(publishInput *DTO_input.InputPublishSequence) (bool, error) {

	_, err := publishSequenceService.sequenceGrpcClient.GetSequenceMedia(publishInput.SequenceId)
	if err != nil {
		return false, err
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