package grpc

import (
	"log"

	userpb "github.com/Vostok1611/project-protos/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := userpb.NewUserServiceClient(conn)
	log.Printf("Connected to Users service at %s", addr)
	return client, conn, nil
}
