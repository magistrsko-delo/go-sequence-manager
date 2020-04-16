package DTO

type SequenceTimeShiftDTO struct {
	Sequence *SequenceData `json:"sequence"`
	Data [] *ChunksResolutionsData `json:"data"`
}

type ChunksResolutionsData struct {
	Resolution string `json:"resolution"`
	Chunks [] *ChunkInfoResolution `json:"chunks"`
}

type ChunkInfoResolution struct {
	ChunkId int32 `json:"chunkId"`              
	Position int32 `json:"position"`            
	AwsBucketName string `json:"awsBucketName"`        
	AwsStorageName string `json:"awsStorageName"`       
	Length float64 `json:"length"`               
	CreatedAt int64 `json:"createdAt"`            
	ChunksUrl string `json:"chunksUrl"`
}