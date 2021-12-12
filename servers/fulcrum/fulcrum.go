package fulcrum

import (
	"log"
	"fmt"
	"os"
	"strconv"
	// "bufio"

	// "google.golang.org/grpc"
	// "golang.org/x/net/context"

	"src/common"
)


type Server struct {
	// clocks		map[string][]int
}

// check if file exist in data folder, if not create it
func check_file(file_name string) {
	_, err := os.Stat("data/" + file_name)
	if os.IsNotExist(err) {
		log.Println("Creating Planet file: " + file_name)
		file, err := os.Create("data/" + file_name)
		common.Check_error(err, "Error al crear el archivo")
		file.Close()
	}
}

// write a register to a planet file in data folder
func write_register(planet string, line string) {
	file, err := os.OpenFile("data/" + planet, os.O_APPEND|os.O_WRONLY, 0600)
	common.Check_error(err, "Error al abrir el archivo")
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	common.Check_error(err, "Error al escribir en el archivo")
}

// delete file in data folder
func delete_file(file_name string) {
	err := os.Remove("data/" + file_name)
	common.Check_error(err, "Error al eliminar el archivo")
}

// func (s * Server) RequestConnection(stream FulcrumService_RequestConnectionServer) error {
// 	return nil
// }

func (s * Server) GetFile(stream FulcrumService_GetFileServer) error {
	// receive file name
	req, err := stream.Recv()
	log.Printf("%s", req.Request)
	if common.Check_error(err, "Error al recibir el nombre del archivo") {
		return err
	}
	var text string
	if req.Request == "log.txt" {
		text = common.Get_file_as_string("data/log.txt")
	} else {
		text = common.Get_file_as_string("data/planets/" + req.Request)
	}

	// send text to client
	err = stream.Send(&FulcrumResponse{Response: text})
	if common.Check_error(err, "Error al enviar el texto al cliente") {
		return err
	}
	return nil
}

func (s * Server) ConnectionBrokerFulcrum(stream FulcrumService_ConnectionBrokerFulcrumServer) error {
	//recibir mensajes de broker, hacer consulta, y devolverlo a broker
	return nil
}

func (s * Server) RequestConnectionFulcrum(stream FulcrumService_RequestConnectionFulcrumServer) error {
	//recibir mensaje del informante y manejar solicitud

	req, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el mensaje del informante") {
		return err
	}
	fmt.Printf("%s\n",req.Request)

	var command string
	var planet string
	var city string
	var opt string
	readed, _ := fmt.Sscanf(req.Request, "%s %s %s %s", &command, &planet, &city, &opt)
	
	file_exist := common.Check_file_exists("data/planets/" + planet+ ".txt") 

	var planet_map map[string]int
	if file_exist {
		txt_planet := common.Get_file_as_string("data/planets/" + planet+ ".txt")
		planet_map = common.Get_string_file_as_map(txt_planet)
	}
	
	//manejar cada tipo de mensaje
	if command == "AddCity" {
		if readed == 3 {
			opt = "0"
		}
		if !file_exist {
			planet_map = make(map[string]int)
			file, err := os.Create("data/planets/" + planet+".txt")
			common.Check_error(err, "Error al crear el archivo")
			defer file.Close()	
		}
		if _, ok := planet_map[city]; !ok {
			tmp, _ := strconv.Atoi(opt)
			planet_map[city] = tmp
		}
		
	} else if command == "UpdateName" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			tmp := planet_map[city]
			delete(planet_map, city)
			planet_map[opt] = tmp
		}

	} else if command == "UpdateNumber" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			planet_map[city], _ = strconv.Atoi(opt)
		}
	} else if command == "DeleteCity" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			os.Remove("data/planets/" + planet+".txt")
			_, err := os.Create("data/planets/" + planet+".txt")
			common.Check_error(err, "Error al crear el archivo")
			delete(planet_map, city)
		}

	}

	//pritn map
	for key, value := range planet_map {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Write planet file
	common.Write_map_to_file(planet_map, planet+".txt")

	//respuesta a infitrados

	err = stream.Send(&FulcrumResponse{
		Response: "he vuelto",
	})
	common.Check_error(err, "Error sending response")

	return nil
}

