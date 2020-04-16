package grpc_client

import (
	"fmt"
	"go-sequence-manager/Models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pbMediaChunks "go-sequence-manager/proto/media_chunks_metadata"
	"log"
)

type MediaChunksClient struct {
	Conn *grpc.ClientConn
	client pbMediaChunks.MediaMetadataClient
}

func (mediaChunksClient *MediaChunksClient) GetMediaChunksResolution(mediaId int32, resolution string) (*pbMediaChunks.MediaChunkInfoResponseRepeated, error)  {

	response, err := mediaChunksClient.client.GetMediaChunksResolution(context.Background(), &pbMediaChunks.MediaChunkResolutionRequest{
		Resolution:           resolution,
		MediaId:              mediaId,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func InitMediaChunksClient() *MediaChunksClient {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING media chunks client")
	conn, err := grpc.Dial(env.MediaChunkMetadataServer + ":" + env.MediaChunkMetadataPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION media chunks client")

	client := pbMediaChunks.NewMediaMetadataClient(conn)
	return &MediaChunksClient{
		Conn:    conn,
		client:  client,
	}
}
