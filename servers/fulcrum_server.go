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

// check if exist the data folder, if not then create it
func check_data_folder() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		log.Println("Creating data folder")
		os.Mkdir("data", 0777)
	}
}

func main() {
	var ip string = common.Get_env_var("IP_SERVER_1")
	var port string = common.Get_env_var("FULCRUM_PORT")

	//print ip and port in the same line
	
	fmt.Println("-----------------------------------------------")
	fmt.Println("------------- Iniciando Fulcrum ---------------")
	fmt.Println("-----------------------------------------------\n")
	fmt.Printf("%s:%s\n", ip, port)
	
	// check if exist the data folder, if not then create it
	check_data_folder()

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