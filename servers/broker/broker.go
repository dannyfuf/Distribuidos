package broker

import (
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"
	"src/common"
)


type Server struct {
}

func (s * Server) RequestConnectionInf(stream BrokerService_RequestConnectionInfServer) error {

	// receibe message
	_, err := stream.Recv()
	check_error(err, "Error receiving message")

	r1 := rand.New(s1)
	rand := r1.Intn(3)
	
	var port string = common.Get_env_var("FULCRUM_PORT")
	var ip string	
	
	if rand == 0{
		ip = common.Get_env_var("IP_SERVER_18")

	} else if rand == 1{
		ip = common.Get_env_var("IP_SERVER_19")

	} else if rand == 2{
		ip = common.Get_env_var("IP_SERVER_20")
	}

	// send response
	err = stream.Send(&BrokerResponse{
		Response: ip,
	})
	check_error(err, "Error sending response")
	return
		
}

func (s * Server) RequestConnectionLeia(stream BrokerService_RequestConnectionLeiaServer) error {

	// for {
	// 	// receibe message
	// 	_, err := stream.Recv()
	// 	check_error(err, "Error receiving message")

	// 	r1 := rand.New(s1)
	// 		rand := r1.Intn(3)
			
	// 		var port string = common.Get_env_var("FULCRUM_PORT")
	// 		var ip string	
			
	// 		if rand == 0{
	// 			//ip = common.Get_env_var("IP_SERVER_1")
	// 			ip = "172.17.0.3"
	// 		} else if rand == 1{

	// 			//ip = common.Get_env_var("IP_SERVER_2")
	// 			ip = "172.17.0.4"
	// 		} else if rand == 2{

	// 			//ip = common.Get_env_var("IP_SERVER_3")
	// 			ip = "172.17.0.5"
	// 		}

	// 		// send response
	// 		err = stream.Send(&BrokerResponse{
	// 			Response: "Hello world",
	// 		})
	// 		check_error(err, "Error sending response")
	// 		time.Sleep(2 * time.Second)
	// }	
}