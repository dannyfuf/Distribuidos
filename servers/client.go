
package main

import (
	"log"
	"os"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"

	"servers/broker"
)

func get_env_var(key string) string {

	// load .env file
	err := godotenv.Load(".env")	
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func main() {
	var ip string = get_env_var("IP_BROKER")
	var port string = get_env_var("PORT_BROKER")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := broker.NewBrokerServiceClient(conn)

	stream, err := c.RequestConnection(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	stream.Send(&broker.BrokerRequest{
		Type: 1,
	})
	for {
		// get response
		response, _ := stream.Recv()
		//print response
		fmt.Println(response)
	}
	

}