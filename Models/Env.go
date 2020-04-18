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
	SequenceServiceServer string
	SequenceServicePort string
	RabbitUser string
	RabbitPassword string
	RabbitQueue string
	RabbitHost string
	RabbitPort string
}

func InitEnv()  {
	envStruct = &Env{
		OriginAllowed:  			os.Getenv("ORIGIN_ALLOWED"),
		Env: 			  			os.Getenv("ENV"),
		Port: 						os.Getenv("PORT"),
		MediaMetadataGrpcServer:	os.Getenv("MEDIA_METADATA_GRPC_SERVER"),
		MediaMetadataGrpcPort:		os.Getenv("MEDIA_METADATA_GRPC_PORT"),
		SequenceServiceServer:		os.Getenv("SEQUENCE_SERVICE_GRPC_SERVER"),
		SequenceServicePort:		os.Getenv("SEQUENCE_SERVICE_GRPC_PORT"),
		RabbitUser:       			os.Getenv("RABBIT_USER"),
		RabbitPassword:   			os.Getenv("RABBIT_PASSWORD"),
		RabbitQueue:      			os.Getenv("RABBIT_QUEUE"),
		RabbitHost:       			os.Getenv("RABBIT_HOST"),
		RabbitPort: 				os.Getenv("RABBIT_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}