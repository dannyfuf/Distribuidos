package broker

import (
	// "log"
	"time"

	// "google.golang.org/grpc"
	// "golang.org/x/net/context"

	"src/common"
)


type Server struct {
}

func (s * Server) RequestConnection(stream BrokerService_RequestConnectionServer) error {
	// receibe message
	_, err := stream.Recv()
	common.Check_error(err, "Error receiving message")

	for {
		// send response
		err = stream.Send(&BrokerResponse{
			Response: "Hello world",
		})
		common.Check_error(err, "Error sending response")
		time.Sleep(2 * time.Second)
	}	
}