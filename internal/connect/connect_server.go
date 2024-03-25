package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthServer() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50061", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
