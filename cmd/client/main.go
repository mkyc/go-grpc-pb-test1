package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	v1 "github.com/mkyc/go-grpc-pb-test1/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewCommunicationTestingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req1 := v1.TaskRequest{
		SenderCounter: 1,
		SenderName:    "client1",
		SentTimestamp: ptypes.TimestampNow(),
		TaskType:      v1.TaskType_PING,
	}

	res1, err := c.Handle(ctx, &req1)

	if err != nil {
		log.Fatalf("Handle failed: %v", err)
	}
	log.Printf("Handle result: <%+v>\n\n", res1)
}
