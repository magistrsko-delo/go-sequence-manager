package services

import (
	"go-sequence-manager/Models/DTO"
	"go-sequence-manager/grpc_client"
	pbMediaMediaMetadata "go-sequence-manager/proto/media_metadata"
)

type SequenceMediaService struct {
	sequenceServiceGrpcClient *grpc_client.SequenceServiceClient
	mediaMetadataGrpcClient *grpc_client.MediaMetadataClient
}

func (sequenceMediaService *SequenceMediaService) AddMediaToSequence(sequenceId int32, mediaId int32) (*DTO.SequenceMediaDTO, error) {

	statusRsp, err := sequenceMediaService.sequenceServiceGrpcClient.AddMediaToSequence(sequenceId, mediaId)

	if err != nil || statusRsp.GetStatus() != 200 {
		return nil, err
	}

	sequenceMedias, err := sequenceMediaService.sequenceServiceGrpcClient.GetSequenceMedia(sequenceId)
	if err != nil {
		return nil, err
	}

	mediaMetadataRsp, err := sequenceMediaService.mediaMetadataGrpcClient.GetMediaMetadata(sequenceMedias.GetMediaIds()[0]) // first media thumbnail

	if err != nil {
		return nil, err
	}

	_, err = sequenceMediaService.sequenceServiceGrpcClient.UpdateSequenceMetadata(
		sequenceMedias.GetSequence().GetSequenceId(),
		sequenceMedias.GetSequence().GetName(),
		sequenceMedias.GetSequence().GetStatus(),
		sequenceMedias.GetSequence().GetProjectId(),
		mediaMetadataRsp.GetThumbnail())

	if err != nil {
		return nil, err
	}

	sequenceMediaDTO, err := sequenceMediaService.GetSequenceMedias(sequenceId)

	if err != nil {
		return nil, err
	}

	return sequenceMediaDTO, nil
}

func (sequenceMediaService *SequenceMediaService) DeleteMediaFromSequence(sequenceId int32, mediaId int32) (*DTO.SequenceMediaDTO, error) {

	statusRsp, err := sequenceMediaService.sequenceServiceGrpcClient.DeleteMediaFromSequence(sequenceId, mediaId)

	if err != nil || statusRsp.GetStatus() != 200 {
		return nil, err
	}

	sequenceMediaDTO, err := sequenceMediaService.GetSequenceMedias(sequenceId)

	if err != nil {
		return nil, err
	}

	if len(sequenceMediaDTO.Medias) == 0 {

		_, _ = sequenceMediaService.sequenceServiceGrpcClient.UpdateSequenceMetadata(
			sequenceMediaDTO.Sequence.SequenceId,
			sequenceMediaDTO.Sequence.Name,
			sequenceMediaDTO.Sequence.Status,
			sequenceMediaDTO.Sequence.ProjectId,
			"")

		sequenceMediaDTO, err = sequenceMediaService.GetSequenceMedias(sequenceId)

		if err != nil {
			return nil, err
		}

	}

	return sequenceMediaDTO, nil
}


func (sequenceMediaService *SequenceMediaService) GetSequenceMedias(sequenceId int32) (*DTO.SequenceMediaDTO, error)  {

	sequenceMediaData, err := sequenceMediaService.sequenceServiceGrpcClient.GetSequenceMedia(sequenceId)

	if err != nil {
		return nil, err
	}

	sequenceMediaDTO := &DTO.SequenceMediaDTO{
		Sequence: &DTO.SequenceData{
			SequenceId: sequenceMediaData.GetSequence().GetSequenceId(),
			Name:       sequenceMediaData.GetSequence().GetName(),
			ProjectId:  sequenceMediaData.GetSequence().GetProjectId(),
			Status:     sequenceMediaData.GetSequence().GetStatus(),
			Thumbnail: 	sequenceMediaData.GetSequence().GetThumbnail(),
			CreatedAt:  sequenceMediaData.GetSequence().GetCreatedAt(),
			UpdatedAt:  sequenceMediaData.GetSequence().GetUpdatedAt(),
		},
		Medias:   [] *pbMediaMediaMetadata.MediaMetadataResponse{},
	}

	for i := 0; i < len(sequenceMediaData.GetMediaIds()); i++ {
		mediaMetadata, err := sequenceMediaService.mediaMetadataGrpcClient.GetMediaMetadata(sequenceMediaData.GetMediaIds()[i])

		if err != nil {
			continue
		}
		sequenceMediaDTO.Medias = append(sequenceMediaDTO.Medias, mediaMetadata)
	}


	return sequenceMediaDTO, nil
}

func InitSequenceMediaService() *SequenceMediaService  {
	return &SequenceMediaService{
		sequenceServiceGrpcClient: grpc_client.InitSequenceServiceMetadata(),
		mediaMetadataGrpcClient:grpc_client.InitMediaMetadataClient(),
	}
}