package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"strings"
	"io/ioutil"
	// "time"
	
	"google.golang.org/grpc"
	"golang.org/x/net/context"

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

func get_string_log_as_map(text string) map[string]string {
	// split text by |
	cities := strings.Split(text, "|")

	var planet string
	var city string
	var command string
	var opt string
	var cities_map = make(map[string]string)
	for i := 0; i < len(cities); i++ {

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

func get_neighbor_files(file_name string, ip string, port string) string {
	// conect to neighbor as client
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	common.Check_error(err, "Error al crear la conexion con el vecino")
	defer conn.Close()

	// create client
	client := fulcrum.NewFulcrumServiceClient(conn)
	stream, err := client.GetFile(context.Background())
	common.Check_error(err, "Error al conectar con el servidor vecino. ip: "+ip)

	// send file name to neighbor
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 9, Request: file_name})
	common.Check_error(err, "Error al enviar el archivo al servidor vecino. ip: "+ip)
	// receive response from neighbor
	response, err := stream.Recv()
	common.Check_error(err, "Error al recibir la respuesta del servidor vecino. ip: "+ip)

	return response.Response
}

func solve_inconsistency(lider map[string]int, lider_log map[string]string, neighbor1 map[string]int, neighbor1_log map[string]string, neighbor2 map[string]int, neighbor2_log map[string]string) map[string]int {
	// merge maps
	merge_map := make(map[string]int)

	// add lider delta
	for city, amount := range lider { // {'melipilla': 1}
		_, bool1 := neighbor1[city]
		_, bool2 := neighbor2[city]
		if bool1 && bool2 { // if city exist in both neighbors
			values := []int{amount, neighbor1[city], neighbor2[city]}
			// get max value
			merge_map[city] = common.Get_max(values)
		} else {
			_, bool1 := neighbor1_log[city] 
			_, bool2 := neighbor2_log[city]
			if bool1 && bool2 {
				if neighbor1_log[city] == "" || neighbor2_log[city] == "" {
					continue
				} else if neighbor1_log[city] != "" {
					merge_map[neighbor1_log[city]] = neighbor1[neighbor1_log[city]]
				} else if neighbor2_log[city] != ""  {
					merge_map[neighbor2_log[city]] = neighbor2[neighbor2_log[city]]
				} else {
					merge_map[city] = amount
				}

			}
		}
	}

	// add neighbor1 delta
	for city, amount := range neighbor1 { // {'melipilla': 1}
		if _, bool2 := merge_map[city]; !bool2 {
			if _, bool1 := neighbor2[city]; bool1 {
				values := []int{neighbor1[city], neighbor2[city]}
				merge_map[city] = common.Get_max(values)
			} else {
				merge_map[city] = amount
			}
		}
	}

	// add neighbor2 delta
	for city, amount := range neighbor2 {
		if _, bool1 := merge_map[city]; !bool1 {
			merge_map[city] = amount
		}
	}

	return merge_map
}

func merge(port string, n int) {
	// log "Merging files"
	log.Println("Merging files")

	neighbors := common.Get_neighbors_fulcrum(n)

	// get all filenames in data/planets folder
	files, err := ioutil.ReadDir("data/planets")
	common.Check_error(err, "Error al leer los archivos de la carpeta data/planets")
	for i := 0; i < len(files); i++ {
		name := files[i].Name()

		// getting files from neighbors
		// get neighbor1
		text_neighbor2 := get_neighbor_files(name, neighbors[1], port)
		neighbor2 := common.Get_string_file_as_map(text_neighbor2)
		
		text_log_neighbor2 := get_neighbor_files("log.txt", neighbors[1], port)
		log_neighbor2 := get_string_log_as_map(text_log_neighbor2)
		
		// get neighbor2
		text_neighbor1 := get_neighbor_files(name, neighbors[0], port)
		neighbor1 := common.Get_string_file_as_map(text_neighbor1)
		
		text_log_neighbor1 := get_neighbor_files("log.txt", neighbors[0], port)
		log_neighbor1 := get_string_log_as_map(text_log_neighbor1)

		// get lider file as string
		lider := common.Get_string_file_as_map(common.Get_file_as_string("data/planets/"+name))
		log_lider := get_string_log_as_map(common.Get_file_as_string("data/log.txt"))

		// solve inconsistency
		merge_map := solve_inconsistency(lider, log_lider, neighbor1, log_neighbor1, neighbor2, log_neighbor2)

		// write merged map to file
		common.Write_map_to_file(merge_map, name)
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
	
	//go routine to merge files in data folder every two minutes
	if n == 20{
		//print "Este es el nodo lider"
		fmt.Println("Este es el nodo lider")
		///////////////////////////////////////////////////////
		// UNCOMMENT THIS TO MERGE FILES EVERY TWO MINUTES
		/////////////////////////////////////////////////////////
		// go func () {
		// 	for {
		// 		merge(port, n)
		// 		time.Sleep(time.Minute * 2)
		// 	}
		// 	}()
	}

	// listen for requests to fulcrum
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	fmt.Println("Escuchando en el puerto: " + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	server := &fulcrum.Server{}

	grpcServer := grpc.NewServer()
	fulcrum.RegisterFulcrumServiceServer(grpcServer, server)
	err = grpcServer.Serve(lis) // bind server

	common.Check_error(err, "Error al iniciar el servidor fulcrum")
}