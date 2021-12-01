package broker

import (
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"
)

func get_env_var(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func check_error(e error, msg string) bool {
	if e != nil {
		log.Printf("%s", msg)
		log.Printf("Error: %v", e)
		return true
	}
	return false
}

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