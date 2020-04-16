package services

import (
	"go-sequence-manager/Models/DTO"
	"go-sequence-manager/grpc_client"
)

type SequenceTimeShiftService struct {
	sequenceServiceGrpcClient *grpc_client.SequenceServiceClient
	mediaChunksGrpcClient *grpc_client.MediaChunksClient
}

func (sequenceTimeShiftService *SequenceTimeShiftService) GetSequenceTimeShiftData(sequenceId int32) (*DTO.SequenceTimeShiftDTO, error)  {
	sequenceMediaDataResponse, err := sequenceTimeShiftService.sequenceServiceGrpcClient.GetSequenceMedia(sequenceId)

	if err != nil {
		return nil, err
	}

	sequenceTimeShiftDTO := &DTO.SequenceTimeShiftDTO{
		Sequence: &DTO.SequenceData{
			SequenceId: sequenceMediaDataResponse.GetSequence().GetSequenceId(),
			Name:       sequenceMediaDataResponse.GetSequence().GetName(),
			ProjectId:  sequenceMediaDataResponse.GetSequence().GetProjectId(),
			Status:     sequenceMediaDataResponse.GetSequence().GetStatus(),
			CreatedAt:  sequenceMediaDataResponse.GetSequence().GetCreatedAt(),
			UpdatedAt:  sequenceMediaDataResponse.GetSequence().GetUpdatedAt(),
		},
		Data:     [] *DTO.ChunksResolutionsData {},
	}

	resolutionArray, err := sequenceTimeShiftService.mediaChunksGrpcClient.GetAvailableResolutions()

	if err != nil {
		return nil, err
	}

	for r := 0; r < len(resolutionArray.GetResolutions()); r++ {
		mediaIds := sequenceMediaDataResponse.GetMediaIds()

		sequenceTimeShiftDTO.Data = append(sequenceTimeShiftDTO.Data, &DTO.ChunksResolutionsData{
			Resolution: resolutionArray.GetResolutions()[r],
			Chunks:     [] *DTO.ChunkInfoResolution{},
		})

		for i := 0; i < len(mediaIds); i++ {
			chunkInfoResolutionRsp, err := sequenceTimeShiftService.mediaChunksGrpcClient.GetMediaChunksInfoResolution(mediaIds[i], resolutionArray.GetResolutions()[r])

			if err != nil {
				continue
			}
			for j := 0; j < len(chunkInfoResolutionRsp.GetData()); j++ {
				sequenceTimeShiftDTO.Data[r].Chunks = append(sequenceTimeShiftDTO.Data[r].Chunks, &DTO.ChunkInfoResolution{
					ChunkId:        chunkInfoResolutionRsp.GetData()[j].GetChunk().GetChunkId(),
					Position:       chunkInfoResolutionRsp.GetData()[j].GetPosition(),
					AwsBucketName:  chunkInfoResolutionRsp.GetData()[j].GetChunk().GetAwsBucketName(),
					AwsStorageName: chunkInfoResolutionRsp.GetData()[j].GetChunk().GetAwsStorageName(),
					Length:         chunkInfoResolutionRsp.GetData()[j].GetChunk().GetLength(),
					CreatedAt:      chunkInfoResolutionRsp.GetData()[j].GetChunk().GetCreatedAt(),
					ChunksUrl:      "testUrl",
				})
			}

		}

	}

	return sequenceTimeShiftDTO, nil
}

func InitSequenceTimeShiftService() *SequenceTimeShiftService  {
	return &SequenceTimeShiftService{
		sequenceServiceGrpcClient: grpc_client.InitSequenceServiceMetadata(),
		mediaChunksGrpcClient:     grpc_client.InitMediaChunksClient(),
	}
}