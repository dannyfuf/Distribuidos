package main

import (
	"fmt"
	"log"
	"os"
	"net"
	
	"google.golang.org/grpc"

	"src/servers/broker"
	"src/common"
)


func main() {
	var ip string = common.Get_env_var("IP_SERVER_17")
	var port string = common.Get_env_var("BROKER_PORT")

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

	common.Check_error(err, "Error al iniciar el servidor de registro de jugadores")

}