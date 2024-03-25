package auth

import (
	"context"
	"fmt"

	"github.com/Sysleec/chat-server/internal/connect"
	access "github.com/Sysleec/chat-server/pkg/access_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func CheckToken(ctx context.Context) error {
	conn, err := connect.AuthServer()
	if err != nil {
		return fmt.Errorf("failed to get connection to auth server %v", err)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("metadata is not provided")
	}
	if err != nil {
		return err
	}

	outCtx := context.Background()
	outCtx = metadata.NewOutgoingContext(outCtx, md)

	client := access.NewAccessV1Client(conn)
	_, err = client.Check(outCtx, &access.CheckRequest{
		EndpointAddress: "",
	})

	if err != nil {
		return fmt.Errorf("token not valid %v", err)
	}

	return nil
}

func GetName(ctx context.Context) (string, error) {
	conn, err := connect.AuthServer()
	if err != nil {
		return "", fmt.Errorf("failed to get connection to auth server %v", err)
	}
	md, err := extractMetadata(ctx)
	if err != nil {
		return "", err
	}
	outCtx := context.Background()
	outCtx = metadata.NewOutgoingContext(outCtx, md)

	client := access.NewAccessV1Client(conn)
	res, err := client.GetName(outCtx, &emptypb.Empty{})
	if err != nil {
		return "", fmt.Errorf("failed to get name %v", err)
	}
	return res.GetName(), nil
}

func extractMetadata(ctx context.Context) (metadata.MD, error) {
	incomMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}
	return incomMd, nil
}
