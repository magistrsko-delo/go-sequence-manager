package DTO

import (
	pbMediaMediaMetadata "go-sequence-manager/proto/media_metadata"
)

type SequenceMediaDTO struct {
	Sequence *SequenceData `json:"sequence"`
	Medias [] *pbMediaMediaMetadata.MediaMetadataResponse
}

type SequenceData struct {
	SequenceId int32 `json:"sequenceId"`
	Name string `json:"name"`
	ProjectId int32 `json:"projectId"`
	Status int32 `json:"status"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
} 
