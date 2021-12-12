package broker

import (
	// "log"
	"time"
	"math/rand"

	// "google.golang.org/grpc"
	// "golang.org/x/net/context"

	"src/common"
	"src/servers/fulcrum"
)

func ConnectFulcrum (mensaje string) string{
	var ipFulcrum string = common.Get_env_var("IP_SERVER_20")
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ipBroker, portBroker), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := fulcrum.NewBrokerServiceClient(conn)

	stream, err := c.ConnectionBrokerFulcrum(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	stream.Send(&broker.BrokerRequest{
		Type: 1, 
		Request: mensaje,
	})
	
	response, _ := stream.Recv()
	return response.Response
}

type Server struct {
}

func (s * Server) RequestConnectionInf(stream BrokerService_RequestConnectionInfServer) error {

	// receibe message
	_, err := stream.Recv()
	common.Check_error(err, "Error receiving message")

	s1 := rand.NewSource(time.Now().UnixNano())	
	r1 := rand.New(s1)
	rand := r1.Intn(3)
	
	// var port string = common.Get_env_var("FULCRUM_PORT")
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
	common.Check_error(err, "Error sending response")
	return nil
}

func (s * Server) RequestConnectionLeia(stream BrokerService_RequestConnectionLeiaServer) error {

	// receibe message
	msg, err := stream.Recv()
	mensaje := msg.Request
	common.Check_error(err, "Error receiving message")
	
	// send response
	err = stream.Send(&BrokerResponse{
		Response: ConnectFulcrum(mensaje),
	})
	common.Check_error(err, "Error sending response")
	return nil
}