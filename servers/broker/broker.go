package broker

import (
	"log"
	"time"
	"fmt"
	"strings"
	"math/rand"
	
	"google.golang.org/grpc"
	"golang.org/x/net/context"

	"src/common"
	"src/servers/fulcrum"
)
func check_ip(ip string) string{
	s1 := rand.NewSource(time.Now().UnixNano())	
	r1 := rand.New(s1)
	election := r1.Intn(3)
	ips := []string{common.Get_env_var("IP_SERVER_18"), common.Get_env_var("IP_SERVER_19"), common.Get_env_var("IP_SERVER_20")}
	if ip == "" {
		return ips[election]
	}
	return ip
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
	
	split_req := strings.Split(req.Request, ",")	
	
	election := check_ip (split_req[1])
	
	err = stream.Send(&BrokerResponse{
		Response: election,
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