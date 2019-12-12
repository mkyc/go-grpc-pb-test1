package v1

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	v1 "github.com/mkyc/go-grpc-pb-test1/pkg/api/v1"
	"log"
)

type communicationTestingServiceServer struct {
}

func NewCommunicationTestingServiceServer() v1.CommunicationTestingServiceServer {
	return &communicationTestingServiceServer{}
}

func (s *communicationTestingServiceServer) Handle(ctx context.Context, req *v1.TaskRequest) (*v1.TaskResponse, error) {
	log.Println("I just got hit!")
	resp := v1.TaskResponse{
		SenderCounter:      req.GetSenderCounter(),
		SenderName:         req.GetSenderName(),
		SentTimestamp:      req.GetSentTimestamp(),
		ReceiverCounter:    2,
		ReceiverName:       "server1",
		ReceivedTimestamp:  ptypes.TimestampNow(),
		RespondedTimestamp: ptypes.TimestampNow(),
		TaskType:           req.GetTaskType(),
	}
	return &resp, nil
}
