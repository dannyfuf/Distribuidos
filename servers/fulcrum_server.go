package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"strings"
	
	"google.golang.org/grpc"
	// "golang.org/x/net/context"

	"src/servers/fulcrum"
	"src/common"
)

// check if exist the data folder, if not then create it
func check_data_folder() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		log.Println("Creating data folder")
		os.Mkdir("data", 0777)
		os.Mkdir("data/planets", 0777)
	} else {
		// if exist, check if planets folder exist inside of data folder, if not then create it
		_, err := os.Stat("data/planets")
		if os.IsNotExist(err) {
			log.Println("Creating planets folder")
			os.Mkdir("data/planets", 0777)
		}
	}
}

func get_string_file_as_map(text string) map[string]int {
	// split text by |
	cities := strings.Split(text, "|")

	var planet string
	var city string
	var amount int
	var cities_map = make(map[string]int)
	for i = 0; i < len(cities); i++ {
		fmt.Sscanf(cities[i], "%s %s %d", &planet, &city, &amount)
		cities_map[city] = amount
	}
	return cities_map
}
func get_string_log_as_map(text string) map[string]string {
	// split text by |
	cities := strings.Split(text, "|")

	var planet string
	var city string
	var command int
	var opt string
	var cities_map = make(map[string]string)
	for i = 0; i < len(cities); i++ {

		fmt.Sscanf(cities[i], "%s %s %s %s", &command, &planet, &city, &opt)
		
		if command == "UpdateName" {
			if _, flag := cities_map[city]; flag {
				if cities_map[city] != "" {
					cities_map[city] = opt
				}
			} else {
				cities_map[city] = opt
			}
		} else if command == "DeleteCity" {
			cities_map[city] = ""
		} else {
			cities_map[city] = city
		}
	}
	return cities_map
}

func get_neighbor_files(file_name string, ip string, port string) map[string]int {
	// conect to neighbor as client
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	common.Check_error(err, "Error al crear la conexion con el vecino")
	defer conn.Close()

	// create client
	client := fulcrum.NewFulcrumServiceClient(conn)
	stream, err := client.Get_neighbor_files(context.Background())
	common.Check_error(err, "Error al conectar con el servidor vecino. ip: "+ip)

	// send file name to neighbor
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 9, Request: file_name})
	common.Check_error(err, "Error al enviar el archivo al servidor vecino. ip: "+ip)
	// receive response from neighbor
	response, err := stream.Recv()
	common.Check_error(err, "Error al recibir la respuesta del servidor vecino. ip: "+ip)

	if file_name == "log.txt" {
		return get_string_log_as_map(response.Response)
	} else {
		return get_string_file_as_map(response.Response)
	}

}

func solve_inconsistency(lider map[string]int, lider_log map[string]string, neighbor1 map[string]int, neighbor1_log map[string]string, neighbor2 map[string]int, neighbor2_log map[string]string) map[string]int {
	// merge maps
	merge_map = make(map[string]int)
	for city, amount := range lider {
		if _, bool1 := neighbor1[city]; _, bool2 := neighbor2[city]; bool1; bool2 { // if city exist in both neighbors
			values := []int{amount, neighbor1[city], neighbor2[city]}
			// get max value
			merge_map[city] = common.Get_max(values)
		} else {
			// TODO: verify the log to check if city was deleted or his name was changed
			if _, bool1  := neighbor1_log[city]; _, bool2  := neighbor2_log[city]; bool1; bool2 {
				// check if any is == to " sa"
				if neighbor1_log[city] == "" || neighbor2_log[city] == "" {
					merge_map[city] = 0
				} else {
					merge_map[city] = amount
				}

			}
		}
	}
}


func write_map_to_file(merge_map map[string]int, planet_name string) {
	f, err := os.OpenFile("data/planets/"+planet_name+".txt", os.O_RDWR, 0755)
	common.Check_error(err, "Error al abrir el archivo")
	defer f.Close()

	// write map to file
	for city, amount := range merge_map {
		f.WriteString(fmt.Sprintf("%s %s %d\n", planet_name, city, amount))
	}
}

func merge(port string, n int) {
	neighbors = common.Get_neighbors_fulcrum(n)

	// get all filenames in data/planets folder
	files, err := ioutil.ReadDir("data/planets")
	common.Check_error(err, "Error al leer los archivos de la carpeta data/planets")
	for i = 0; i < len(files); i++ {
		name := files[i]

		// getting files from neighbors
		neighbor2 := get_neighbor_files(name, neighbors[1], port)
		log_neighbor2 := get_neighbor_files("log.txt", neighbors[1], port)

		neighbor1 := get_neighbor_files(name, neighbors[0], port)
		log_neighbor1 := get_neighbor_files("log.txt", neighbors[0], port)
		
		// get lider file as string
		lider := get_string_file_as_map(common.Get_file_as_string("data/planets/"+name))
		log_lider := get_string_file_as_map(common.Get_file_as_string("data/log.txt")

		// solve inconsistency
		merge_map = solve_inconsistency(lider, neighbor1, neighbor2)

		planet_name = strings.Split(name, ".")[0]
		// write merged map to file
		write_map_to_file(merge_map, planet_name)
	}
}

func main() {
	var port string = common.Get_env_var("FULCRUM_PORT")

	//read n from keyboard
	var n int
	fmt.Print("Ingresa el número de la máquina: ")
	fmt.Scanf("%d", &n)

	fmt.Println("-----------------------------------------------")
	fmt.Println("------------- Iniciando Fulcrum ---------------")
	fmt.Println("-----------------------------------------------\n")

	// check if exist the data folder, if not then create it
	check_data_folder()

	// listen for requests to fulcrum
	go func () {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		server := &fulcrum.Server{}
 
		grpcServer := grpc.NewServer()
		fulcrum.RegisterFulcrumServiceServer(grpcServer, server)
		err = grpcServer.Serve(lis) // bind server

		common.Check_error(err, "Error al iniciar el servidor fulcrum")
	}()

	//go routine to merge files in data folder every two minutes
	if n == 20{
		go func () {
			for {
				merge(port, n)
				time.Sleep(time.Minute * 2)
			}
		}
	}
}