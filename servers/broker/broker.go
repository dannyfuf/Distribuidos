package broker

import (
	"log"
	"time"
	"fmt"
	"math/rand"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	"src/common"
	"src/servers/fulcrum"
)
func checkVector(x int, y int, z int) int{
	s1 := rand.NewSource(time.Now().UnixNano())	
	r1 := rand.New(s1)
	var election int
	if x == y && x == z {
		election= r1.Intn(3)
		return election

	} else if x == y && x > z{
		election = r1.Intn(2)

	} else if x == z && x > y{
		election := r1.Intn(2)
		if election == 0{
			return 0
		} else {
			return 2
		}

	} else if y == z && x < y {
		election = r1.Intn(2) + 1

	} else if x > y && x > z {
		return 0

	} else if y > x && y > z {
		return 1

	} else { //z > x && z > y
		return 2
	}
	return 0
}
func ConnectFulcrum (mensaje string) string{
	var ipFulcrum string = common.Get_env_var("IP_SERVER_20")
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ipFulcrum, portFulcrum), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := fulcrum.NewFulcrumServiceClient(conn)

	stream, err := c.ConnectionBrokerFulcrum(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	stream.Send(&fulcrum.FulcrumRequest{
		Type: 1, 
		Request: mensaje,
	})
	
	response, _ := stream.Recv()
	return response.Response
}

type Server struct {
}

func (s * Server) RequestConnectionInf(stream BrokerService_RequestConnectionInfServer) error {
	req, err := stream.Recv()
	log.Println("Request: ", req.Request)
	common.Check_error(err, "Error receiving message")

	s1 := rand.NewSource(time.Now().UnixNano())	
	r1 := rand.New(s1)
	election := r1.Intn(3)
	
	var ip string	
	//election = checkVector (x, y, z)
	if election == 0{
		ip = common.Get_env_var("IP_SERVER_18")

	} else if election == 1{
		ip = common.Get_env_var("IP_SERVER_19")

	} else if election == 2{
		ip = common.Get_env_var("IP_SERVER_20")
	}
	// send response
	err = stream.Send(&BrokerResponse{
		Response: ip,
	})
	common.Check_error(err, "Error sending response")
	return nil
}

func (s * Server) RequestConnectionLeia(stream BrokerService_RequestConnectionLeiaServer) error {

	// receibe message
	req, err := stream.Recv()
	peticion := req.Request
	common.Check_error(err, "Error receiving message")
	log.Printf("Request: %s\n", peticion)
	var answer string
	answer = ConnectFulcrum (peticion)

	err = stream.Send(&BrokerResponse{
		Response: answer,
	})
	common.Check_error(err, "Error sending response")
	return nil
}