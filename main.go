package main

import (
	"MusicMesh/api_geteway-MusicMesh/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.NewClient(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	compositionConn, err := grpc.NewClient(":5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	collaborationConn, err := grpc.NewClient(":5053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	discoveryConn, err := grpc.NewClient(":5054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	gin := api.API(userConn, compositionConn, collaborationConn, discoveryConn)
	panic(gin.Run(":8080"))
}
