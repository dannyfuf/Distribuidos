package main

import (
	"fmt"
	"log"
	"os"
	"net"
	
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

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

func check_error(e error, msg string) bool {
	if e != nil {
		log.Printf("%s", msg)
		log.Printf("Error: %v", e)
		return true
	}
	return false
}


func main() {
	var ip string = get_env_var("IP_BROKER")
	var port string = get_env_var("PORT_BROKER")

	//print ip and port in the same line
	
	fmt.Println("---------------------------------------------------------")
	fmt.Println("------------- Iniciando Broker Mos Eisley ---------------")
	fmt.Println("---------------------------------------------------------\n")
	fmt.Printf("%s:%s\n", ip, port)
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	server := &broker.Server{}

	grpcServer := grpc.NewServer()
	broker.RegisterBrokerServiceServer(grpcServer, server)
	err = grpcServer.Serve(lis) // bind server

	check_error(err, "Error al iniciar el servidor de registro de jugadores")

}