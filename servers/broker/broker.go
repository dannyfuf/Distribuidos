package broker

import (
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"
	"src/common"
)


type Server struct {
}

func (s * Server) RequestConnection(stream BrokerService_RequestConnectionServer) error {
	// receibe message
	_, err := stream.Recv()
	check_error(err, "Error receiving message")

	for {
		// send response
		err = stream.Send(&BrokerResponse{
			Response: "Hello world",
		})
		check_error(err, "Error sending response")
		time.Sleep(2 * time.Second)
	}	
}