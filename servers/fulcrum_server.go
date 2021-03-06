package main

import (
	"fmt"
	"log"
	"os"
	"net"
	"strings"
	// "io/ioutil"
	"time"
	
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
		os.Create("data/log.txt")
		os.Mkdir("data/planets", 0777)
	} else {
		// if exist, check if planets folder exist inside of data folder, if not then create it
		_, err := os.Stat("data/planets")
		if os.IsNotExist(err) {
			log.Println("Creating planets folder")
			os.Mkdir("data/planets", 0777)
		}

		// check if log.txt exist inside of data folder, if not then create it
		_, err = os.Stat("data/log.txt")
		if os.IsNotExist(err) {
			log.Println("Creating log.txt")
			os.Create("data/log.txt")
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

	fmt.Println("Solving inconsistency lider")
	// add lider delta
	for city, amount := range lider { // {'melipilla': 1}
		_, bool1 := neighbor1[city]
		_, bool2 := neighbor2[city]
		if bool1 && bool2 { // if city exist in both neighbors
			values := []int{amount, neighbor1[city], neighbor2[city]}
			merge_map[city] = common.Get_max(values)
		} else if !bool1 && bool2 { // ciudad modificada en vecino 1
			_, bool1 := neighbor1_log[city]
			
			if bool1 && neighbor1_log[city] != "" {
				merge_map[neighbor1_log[city]] = neighbor1[neighbor1_log[city]]
			}

		} else if bool1 && !bool2 { // ciudad modificada en vecino 2
			_, bool2 := neighbor2_log[city]
			
			if bool2 && neighbor2_log[city] != "" {
				merge_map[neighbor2_log[city]] = neighbor2[neighbor2_log[city]]
			}
		} else {
			merge_map[city] = amount
		}
	}

	fmt.Println("Solving inconsistency neighbor1")
	// add neighbor1 delta
	for city, amount := range neighbor1 { // {'melipilla': 1}
		_, v2 := neighbor2[city]
		_, l_log := lider_log[city]
		_, merge := merge_map[city]

		if !merge && !v2{
			merge_map[city] = amount
		} else if !merge && v2 {
			if l_log && lider_log[city] != "" {
				merge_map[lider_log[city]] = amount
			}
		}
	}

	fmt.Println("Solving inconsistency neighbor2")
	// add neighbor2 delta
	for city, amount := range neighbor2 {
		if _, bool1 := merge_map[city]; !bool1 {
			merge_map[city] = amount
		}
	}

	return merge_map
}

func send_merged_file(file string, planet_name string, ip string) {
	// conect to neighbor as client
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, common.Get_env_var("FULCRUM_PORT")), grpc.WithInsecure())
	common.Check_error(err, "Error al crear la conexion con el vecino")
	defer conn.Close()

	// create client
	client := fulcrum.NewFulcrumServiceClient(conn)
	stream, err := client.SendFile(context.Background())

	// send planet name to neighbor
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 10, Request: planet_name})
	common.Check_error(err, "Error al enviar el planeta al servidor vecino. ip: "+ip)

	// send file name to neighbor
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 10, Request: file})
	common.Check_error(err, "Error al enviar el archivo al servidor vecino. ip: "+ip)

	// receive response from neighbor
	response, err := stream.Recv()
	common.Check_error(err, "Error al recibir la respuesta del servidor vecino. ip: "+ip)

	if response.Response != "OK" {
		log.Fatal("Error al enviar el archivo al servidor vecino. ip: "+ip)
	}
}

func get_planet_list(ip string) []string {
	// conect to neighbor as client
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, common.Get_env_var("FULCRUM_PORT")), grpc.WithInsecure())
	common.Check_error(err, "Error al crear la conexion con el vecino")
	defer conn.Close()

	// create fulcrum client
	client := fulcrum.NewFulcrumServiceClient(conn)
	stream, err := client.GetFileList(context.Background())
	common.Check_error(err, "Error al conectar con el servidor vecino. ip: "+ip)

	// ask for planet list
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 11, Request: ""})
	common.Check_error(err, "Error al enviar la peticion de planetas al servidor vecino. ip: "+ip)

	// receive response from neighbor
	response, err := stream.Recv()
	common.Check_error(err, "Error al recibir la respuesta del servidor vecino. ip: "+ip)

	return strings.Split(response.Response, ",")
}

func create_planet_list(lider []string, neighbor1 []string, neighbor2 []string) []string {
	// create a list with elements from lider, neighbor1 and neighbor2
	planet_list := make([]string, 0)
	
	for _, element := range lider {
		planet_list = append(planet_list, element)	
	}

	for _, element := range neighbor1 {
		if !common.Contains(element, planet_list) {
			planet_list = append(planet_list, element)
		}
	}

	for _, element := range neighbor2 {
		if !common.Contains(element, planet_list) {
			planet_list = append(planet_list, element)
		}
	}
	return planet_list
}

func merge_clocks(lider_clock map[string][]int, neighbor1_clock map[string][]int, neighbor2_clock map[string][]int) map[string][]int {
	fmt.Println("-------------- MERGING CLOCKS --------------\n")
	merge_map := make(map[string][]int)

	for planet, clock := range lider_clock {
		merge_map[planet] = []int{neighbor1_clock[planet][0], neighbor2_clock[planet][1], clock[2]}
	}
	return merge_map
}

func send_clock_to_neighbor(ip string, planet string, clock []int) bool{
	// connect to neighbor as client
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, common.Get_env_var("FULCRUM_PORT")), grpc.WithInsecure())
	common.Check_error(err, "Error al crear la conexion con el vecino")
	defer conn.Close()

	// create fulcrum client
	client := fulcrum.NewFulcrumServiceClient(conn)
	stream, err := client.SendClock(context.Background())
	common.Check_error(err, "Error al conectar con el servidor vecino")

	// send planet list to neighbor
	fmt.Printf("Enviando planeta %s a %s\n", planet, ip)
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 13, Request: planet})
	common.Check_error(err, "Error al enviar la peticion de planetas al servidor vecino")

	// send clock to neighbor
	fmt.Printf("\nEnviando Clock: %v a: %s\n", clock, ip)
	err = stream.Send(&fulcrum.FulcrumRequest{Type: 13, Request: common.Array_as_string(clock)})
	common.Check_error(err, "Error al enviar la peticion de planetas al servidor vecino")

	// receive response from neighbor
	res, err := stream.Recv()
	common.Check_error(err, "Error al recibir la respuesta del servidor vecino")
	
	if res.Response != "OK" {
		return false
	} else {
		return true
	}
	return false
}

func get_clocks(planet_list []string, ip string) map[string][]int {
	clock_map := make(map[string][]int)

	for _, planet := range planet_list {

		// connect to neighbor as client
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, common.Get_env_var("FULCRUM_PORT")), grpc.WithInsecure())
		common.Check_error(err, "Error al crear la conexion con el vecino")
		defer conn.Close()

		// create fulcrum client
		client := fulcrum.NewFulcrumServiceClient(conn)
		stream, err := client.GetClock(context.Background())
		common.Check_error(err, "Error al conectar con el servidor vecino")

		err = stream.Send(&fulcrum.FulcrumRequest{Type: 13, Request: planet})
		common.Check_error(err, "Error al enviar la peticion de planetas al servidor vecino en get_clocks")

		// receive response from neighbor
		response, err := stream.Recv()
		common.Check_error(err, "Error al recibir la respuesta del servidor vecino")

		planet_name := strings.Split(planet, ".")[0]
		if response.Response != "" {
			clock_map[planet_name] = common.String_as_array(response.Response)
		} else {
			clock_map[planet_name] = []int{0,0,0}
		}

	}

	fmt.Printf("\nSe han recibido los relojes de %s\n", ip)
	for planet, clock := range clock_map {
		fmt.Printf("%s: %v\n", planet, clock)
	}
	return clock_map
}

func merge(port string, n int) {
	// log "Merging files"
	fmt.Println("-------------------- Merging files --------------------\n")

	neighbors := common.Get_neighbors_fulcrum(n)

	// get planet list from neighbors
	planets_neighbor1 := get_planet_list(common.Get_env_var("IP_SERVER_18"))
	planets_neighbor2 := get_planet_list(common.Get_env_var("IP_SERVER_19"))
	planets_lider := common.Create_file_list("data/planets")
	planet_list := create_planet_list(planets_lider, planets_neighbor1, planets_neighbor2)
	fmt.Printf("Planetas encontrados:\n %v\n", planet_list)

	fmt.Printf("Pidiendo relojes a los vecinos\n")
	clocks1 := get_clocks(planet_list, common.Get_env_var("IP_SERVER_18"))
	clocks2 := get_clocks(planet_list, common.Get_env_var("IP_SERVER_19"))
	lider := get_clocks(planet_list, common.Get_env_var("IP_SERVER_20"))

	merged_clocks := merge_clocks(lider, clocks1, clocks2)

	for i := 0; i < len(planet_list); i++ {
		name := planet_list[i]

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

		// write merged file
		common.Write_map_to_file(merge_map, name)
		common.Check_line("data/planets/"+name)

		// send to neighbor
		fmt.Printf("\nEnviando relojes\n")
		// split name by "."
		file_name := strings.Split(name, ".")[0]
		send_clock_to_neighbor(common.Get_env_var("IP_SERVER_18"), name, merged_clocks[file_name])
		send_clock_to_neighbor(common.Get_env_var("IP_SERVER_19"), name, merged_clocks[file_name])
		send_clock_to_neighbor(common.Get_env_var("IP_SERVER_20"), name, merged_clocks[file_name])

		merged := common.Get_file_as_string("data/planets/"+name)
		send_merged_file(merged, name, common.Get_env_var("IP_SERVER_18"))
		send_merged_file(merged, name, common.Get_env_var("IP_SERVER_19"))
	}
}

func main() {
	var port string = common.Get_env_var("FULCRUM_PORT")

	//read n from keyboard
	var n int
	fmt.Print("Ingresa el n??mero de la m??quina: ")
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
		go func () {
			for {
				time.Sleep(time.Minute * 2)
				log.Println("Merging files")
				merge(port, n)
			}
		}()
	}

	// listen for requests to fulcrum
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	fmt.Println("Escuchando en el puerto: " + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	relojes := make(map[string][]int)

	server := &fulcrum.Server{
		Relojes: relojes,
		N_server: n-18,
	}

	grpcServer := grpc.NewServer()
	fulcrum.RegisterFulcrumServiceServer(grpcServer, server)
	err = grpcServer.Serve(lis) // bind server

	common.Check_error(err, "Error al iniciar el servidor fulcrum")
}