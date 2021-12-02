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

// check if exist the data folder, if not then create it
func check_data_folder() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		log.Println("Creating data folder")
		os.Mkdir("data", 0777)
	}
}

// check if file exist in data folder, if not create it
func check_file(file_name string) {
	_, err := os.Stat("data/" + file_name)
	if os.IsNotExist(err) {
		log.Println("Creating Planet file: " + file_name)
		file, err := os.Create("data/" + file_name)
		check_error(err, "Error al crear el archivo")
		file.Close()
	}
}

// write a register to a planet file in data folder
func write_register(planet string, line string) {
	file, err := os.OpenFile("data/" + planet, os.O_APPEND|os.O_WRONLY, 0600)
	check_error(err, "Error al abrir el archivo")
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	check_error(err, "Error al escribir en el archivo")
}

// delete file in data folder
func delete_file(file_name string) {
	err := os.Remove("data/" + file_name)
	check_error(err, "Error al eliminar el archivo")
}

func main() {
	var ip string = get_env_var("IP_SERVER_1")
	var port string = get_env_var("PORT_BROKER")

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

	check_error(err, "Error al iniciar el servidor de registro de jugadores")

}