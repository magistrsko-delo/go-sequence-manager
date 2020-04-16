package grpc_client

import (
	"fmt"
	"go-sequence-manager/Models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pbMediaMediaMetadata "go-sequence-manager/proto/media_metadata"
	"log"
)

type MediaMetadataClient struct {
	Conn *grpc.ClientConn
	client pbMediaMediaMetadata.MediaMetadataClient
}

func (mediaMetadataClient *MediaMetadataClient) GetMediaMetadata (mediaId int32) (*pbMediaMediaMetadata.MediaMetadataResponse, error)  {

	response, err := mediaMetadataClient.client.GetMediaMetadata(context.Background(), &pbMediaMediaMetadata.GetMediaMetadataRequest{
		MediaId:              mediaId,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}


func InitMediaMetadataClient() *MediaMetadataClient  {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING mediaMetadata client")
	conn, err := grpc.Dial(env.MediaMetadataGrpcServer + ":" + env.MediaMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION mediaMetadata client")

	client := pbMediaMediaMetadata.NewMediaMetadataClient(conn)
	return &MediaMetadataClient{
		Conn:    conn,
		client:  client,
	}
}