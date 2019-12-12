package cmd

import (
	"context"
	"github.com/mkyc/go-grpc-pb-test1/pkg/protocol/grpc"
	"github.com/mkyc/go-grpc-pb-test1/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()
	v1API := v1.NewCommunicationTestingServiceServer()
	return grpc.RunServer(ctx, v1API, "8080")
}
