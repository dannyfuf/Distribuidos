package fulcrum

import (
	"log"
	"os"
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

	


	

	return nil
}

