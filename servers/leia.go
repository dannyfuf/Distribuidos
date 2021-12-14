package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	"src/servers/broker"
	"src/common"
)

func GetNumberRebelds(nombre_planeta string, nombre_ciudad string) string {

	fmt.Println("Solicitando informacion sobre rebeldes en " + nombre_planeta + ", en la ciudad " + nombre_ciudad)

	var mensaje string 

	mensaje = nombre_planeta+" "+nombre_ciudad

	var ipBroker string = common.Get_env_var("IP_SERVER_17")
	var portBroker string = common.Get_env_var("BROKER_PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ipBroker, portBroker), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := broker.NewBrokerServiceClient(conn)

	stream, err := c.RequestConnectionLeia(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	stream.Send(&broker.BrokerRequest{
		Type: 1, 
		Request: mensaje,
	})
	
	response, _ := stream.Recv()

	recep := "La cantidad de rebeldes es "+response.Response
	
	//Dependiendo de como se envie la informacion se procede a trabajarla mas o enviarla tal y como llegue
	return recep

}

func main(){
	respuesta := -1
	for respuesta != 0 {
	
		fmt.Printf("Que desea hacer Princesa Leia?\n\n0: Salir\n1: Ver cantidad de rebeldes en un planeta\n\n")
		fmt.Scanf("%d",&respuesta)

		if respuesta == 0 {
			return

		} else if respuesta == 1{

			fmt.Println("Contador de Rebeldes segun planeta")

			fmt.Println("Ingrese el nombre del planeta a comprobar:")
			var planeta string
			fmt.Scanln(&planeta)

			fmt.Println("Ingrese el nombre de la ciudad a comprobar:")
			var ciudad string
			fmt.Scanln(&ciudad)

			result := GetNumberRebelds(planeta, ciudad)
			fmt.Println(result)
			
		} else {
			respuesta = -1
			fmt.Printf("Seleccione una opcion valida\n")
		}
		
	}
	return
}
