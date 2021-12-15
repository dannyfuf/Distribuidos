package fulcrum

import (
	"log"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "bufio"

	// "google.golang.org/grpc"
	// "golang.org/x/net/context"

	"lab3/common"
)

type Server struct {
	Relojes  map[string][]int
	N_server int
}

func (s * Server) SendFile(stream FulcrumService_SendFileServer) error {
	
	// receive Planet name
	planet, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el archivo") {
		return err
	}

	if planet.Request == "" {
		return nil
	}

	// receive file
	file, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el archivo") {
		return err
	}
	
	// write file to data folder
	planet_map := common.Get_string_file_as_map(file.Request)
	
	os.Remove("data/planets/" + planet.Request)
	os.Create("data/planets/" + planet.Request)

	common.Write_map_to_file(planet_map, planet.Request)

	// send response to client
	err = stream.Send(&FulcrumResponse{Response: "OK"})
	if common.Check_error(err, "Error al enviar el texto al cliente") {
		return err
	}

	common.Check_line("data/planets/"+planet.Request)

	os.Remove("data/log.txt")
	os.Create("data/log.txt")

	return nil
}

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
	req, err := stream.Recv()
	req_split := strings.Split(req.Request, " ")

	
	if common.Check_file_exists("data/planets/" + req_split[0] + ".txt") {
		text := common.Get_file_as_string("data/planets/" + req_split[0] + ".txt")
		planet_map := common.Get_string_file_as_map(text)	
		
		if _, ok := planet_map[req_split[1]]; ok {
			if _, ok := s.Relojes[req_split[0]]; ok {
				clock := common.Array_as_string(s.Relojes[req_split[0]])
				err = stream.Send(&FulcrumResponse{
					Response: strconv.Itoa(planet_map[req_split[1]])+" "+clock,
				})
				common.Check_error(err, "Error enviar la cantidad de soldados al cliente")
			}
		} else {
			//send -1 to client
			err = stream.Send(&FulcrumResponse{
				Response: "-1 ",
			})
			common.Check_error(err, "Error al notificar que la cuidad no existe")
		}
	} else {
		// send -1 to client
		err = stream.Send(&FulcrumResponse{
			Response: "-1 ",
		})
		common.Check_error(err, "Error al notificar que el planeta no existe")
	}
	return nil
}

func (s * Server) RequestConnectionFulcrum(stream FulcrumService_RequestConnectionFulcrumServer) error {
	//recibir mensaje del informante y manejar solicitud

	req, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el mensaje del informante") {
		return err
	}
	
	split_req := strings.Split(req.Request, ",")

	fmt.Printf("%s\n",req.Request)

	var command string
	var planet string
	var city string
	var opt string
	readed, _ := fmt.Sscanf(split_req[0], "%s %s %s %s", &command, &planet, &city, &opt)
	
	file_exist := common.Check_file_exists("data/planets/" + planet+ ".txt") 

	if !file_exist {
		s.Relojes[planet] = make([]int, 3)
	}

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
			common.Append_line_to_file(req.Request, "data/log.txt")
			s.Relojes[planet][s.N_server] += 1
		}
		
	} else if command == "UpdateName" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			tmp := planet_map[city] // "cuidaa" : 1 -> int
			delete(planet_map, city)
			planet_map[opt] = tmp
			common.Append_line_to_file(req.Request, "data/log.txt")
			s.Relojes[planet][s.N_server] += 1 // Reloj = {"tierra": [2,2,3] }
		}

	} else if command == "UpdateNumber" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			planet_map[city], _ = strconv.Atoi(opt)
			common.Append_line_to_file(req.Request, "data/log.txt")
			s.Relojes[planet][s.N_server] += 1
		}
	} else if command == "DeleteCity" && file_exist {
		if _, bool1 := planet_map[city]; bool1 {
			os.Remove("data/planets/" + planet+".txt")
			_, err := os.Create("data/planets/" + planet+".txt")
			common.Check_error(err, "Error al crear el archivo")
			delete(planet_map, city)
			common.Append_line_to_file(req.Request, "data/log.txt")
			s.Relojes[planet][s.N_server] += 1
		}
	}

	// Write planet file
	common.Write_map_to_file(planet_map, planet+".txt")

	// print s.Relojes
	for planet_tmp, relojes := range s.Relojes {
		fmt.Printf("%s: %v\n", planet_tmp, relojes)
	}

	clock := common.Array_as_string(s.Relojes[planet])
	err = stream.Send(&FulcrumResponse{
		Response: clock,
	})
	common.Check_error(err, "Error sending response")

	return nil
}

func (s * Server) GetFileList(stream FulcrumService_GetFileListServer) error {
	req, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el nombre del archivo") {
		return err
	}

	if req.Type == 11 {
		// create a list of files
		file_list := common.Create_file_list("data/planets/")
		// send list to client
		err = stream.Send(&FulcrumResponse{
			Response: strings.Join(file_list, ","),
		})
		if common.Check_error(err, "Error al enviar el texto al cliente") {
			return err
		}
	}
	return nil
}

func (s * Server) GetClock(stream FulcrumService_GetClockServer) error {
	fmt.Printf("------------------------------------------\n")
	req, err := stream.Recv()
	fmt.Printf("\nSolicitando planeta: %s\n", req.Request)
	if common.Check_error(err, "Error al recibir el nombre del archivo") {
		return err
	}

	req_split := strings.Split(req.Request, ".")
	if _, ok := s.Relojes[req_split[0]]; ok {
		// split req.Request by "."
		clock := common.Array_as_string(s.Relojes[req_split[0]])

		fmt.Printf("Enviando Reloj: %s\n", clock)

		err = stream.Send(&FulcrumResponse{
			Response: clock,
		})
		if common.Check_error(err, "Error al enviar el texto al cliente") {
			return err
		}
	} else {
		err = stream.Send(&FulcrumResponse{
			Response: "",
		})
		if common.Check_error(err, "Error al enviar el texto al cliente") {
			return err
		}
	}

	return nil
}

func (s * Server) SendClock(stream FulcrumService_SendClockServer) error {
	
	req, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el nombre del archivo") {
		return err
	}

	req2, err := stream.Recv()
	if common.Check_error(err, "Error al recibir el nombre del archivo") {
		return err
	}
	fmt.Printf("------------------------------------------\nRecibiendo Reloj de %s: %s\n", strings.Split(req.Request, ".")[0], req2.Request)
	clock_array := common.String_as_array(req2.Request)

	req_split := strings.Split(req.Request, ".")
	s.Relojes[req_split[0]] = clock_array

	fmt.Println("-------------------- Nuevo reloj ----------------------\n")
	for planet, relojes := range s.Relojes {
		fmt.Printf("%s: %v\n", planet, relojes)
	}

	err = stream.Send(&FulcrumResponse{
		Response: "OK",
	})

	return nil
}
