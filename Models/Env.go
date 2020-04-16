package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

type Env struct {
	OriginAllowed string
	Env string
	Port string
	MediaMetadataGrpcServer string
	MediaMetadataGrpcPort string
	MediaChunkMetadataServer string
	MediaChunkMetadataPort string
	SequenceServiceServer string
	SequenceServicePort string
}

func InitEnv()  {
	envStruct = &Env{
		OriginAllowed:  			os.Getenv("ORIGIN_ALLOWED"),
		Env: 			  			os.Getenv("ENV"),
		Port: 						os.Getenv("PORT"),
		MediaMetadataGrpcServer:	os.Getenv("MEDIA_METADATA_GRPC_SERVER"),
		MediaMetadataGrpcPort:		os.Getenv("MEDIA_METADATA_GRPC_PORT"),
		MediaChunkMetadataServer:	os.Getenv("MEDIA_CHUNKS_METADATA_GRPC_SERVER"),
		MediaChunkMetadataPort:		os.Getenv("MEDIA_CHUNKS_METADATA_GRPC_PORT"),
		SequenceServiceServer:		os.Getenv("SEQUENCE_SERVICE_GRPC_SERVER"),
		SequenceServicePort:		os.Getenv("SEQUENCE_SERVICE_GRPC_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}