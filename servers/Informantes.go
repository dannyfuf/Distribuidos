package main
import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"src/servers/broker"
	"src/servers/fulcrum"
	"src/common"

)

type infData struct {
	servers []int
	ip string
}

var Relojes = make(map[string]infData)

func Save(planet string, city string, ip string, vector string){
	Relojes[planet+" "+city] = infData{servers: common.String_as_array(vector), ip: ip}
}

func ConnectBroker(mensaje string) string{
	var ipBroker string = common.Get_env_var("IP_SERVER_17")
	var portBroker string = common.Get_env_var("BROKER_PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ipBroker, portBroker), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := broker.NewBrokerServiceClient(conn)

	stream, err := c.RequestConnectionInf(context.Background())
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

func ConnectFulcrum(mensaje string, ip string) string{
	ipFulcrum := ip
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ipFulcrum, portFulcrum), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := fulcrum.NewFulcrumServiceClient(conn)

	stream, err := c.RequestConnectionFulcrum(context.Background())
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

func AddCity(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {
	var mensaje string

	if nuevo_valor >= 0 {
		mensaje = "AddCity "+nombre_planeta+" "+nombre_ciudad+" "+strconv.Itoa(nuevo_valor)+","+Relojes[nombre_planeta+" "+nombre_ciudad].ip

	} else {
		mensaje = "AddCity "+nombre_planeta+" "+nombre_ciudad+","+Relojes[nombre_planeta+" "+nombre_ciudad].ip

	}
	fmt.Printf("%s\n", mensaje)

	ipFulcrum := ConnectBroker(mensaje)
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	fmt.Printf("Se va a conectar al servidor con ip: %s:%s\n\n",ipFulcrum, portFulcrum)
	
	response := ConnectFulcrum(mensaje, ipFulcrum)
	Save(nombre_planeta, nombre_ciudad, ipFulcrum, response)

	fmt.Printf("%s %s:\n%v\n",nombre_planeta,nombre_ciudad,Relojes[nombre_planeta+" "+nombre_ciudad].ip)
	fmt.Printf("%v\n",Relojes[nombre_planeta+" "+nombre_ciudad].servers)

}

func UpdateName(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	
	mensaje := "UpdateName "+nombre_planeta+" "+nombre_ciudad+" "+nuevo_valor+","+Relojes[nombre_planeta+" "+nombre_ciudad].ip

	fmt.Printf("%s\n", mensaje)

	ipFulcrum := ConnectBroker(mensaje)
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	fmt.Printf("Se va a conectar al servidor con ip: %s:%s\n\n",ipFulcrum, portFulcrum)

	response := ConnectFulcrum(mensaje, ipFulcrum)
	Save(nombre_planeta, nuevo_valor, ipFulcrum, response)
	delete(Relojes, nombre_planeta+" "+nombre_ciudad)
	fmt.Printf("%s %s:\n%v\n",nombre_planeta,nuevo_valor,Relojes[nombre_planeta+" "+nuevo_valor].ip)
	fmt.Printf("%v\n",Relojes[nombre_planeta+" "+nuevo_valor].servers)

}

func UpdateNumber(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {

	mensaje := "UpdateNumber "+nombre_planeta+" "+nombre_ciudad+" "+strconv.Itoa(nuevo_valor)+","+Relojes[nombre_planeta+" "+nombre_ciudad].ip

	fmt.Printf("%s\n", mensaje)

	ipFulcrum := ConnectBroker(mensaje)
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	fmt.Printf("Se va a conectar al servidor con ip: %s:%s\n\n",ipFulcrum, portFulcrum)

	response := ConnectFulcrum(mensaje, ipFulcrum)
	Save(nombre_planeta, nombre_ciudad, ipFulcrum, response)

	fmt.Printf("%s %s:\n%v\n",nombre_planeta,nombre_ciudad,Relojes[nombre_planeta+" "+nombre_ciudad].ip)
	fmt.Printf("%v\n",Relojes[nombre_planeta+" "+nombre_ciudad].servers)

}

func DeleteCity(nombre_planeta string, nombre_ciudad string) {

	mensaje := "DeleteCity "+nombre_planeta+" "+nombre_ciudad+","+Relojes[nombre_planeta+" "+nombre_ciudad].ip

	fmt.Printf("%s\n", mensaje)

	ipFulcrum := ConnectBroker(mensaje)
	var portFulcrum string = common.Get_env_var("FULCRUM_PORT")
	fmt.Printf("Se va a conectar al servidor con ip: %s:%s\n\n",ipFulcrum, portFulcrum)

	//response := ConnectFulcrum(mensaje, ipFulcrum)
	ConnectFulcrum(mensaje, ipFulcrum)
	//Save(nombre_planeta, nombre_ciudad, ipFulcrum, response)
	delete(Relojes, nombre_planeta+" "+nombre_ciudad)
	fmt.Printf("La siguiente ip y reloj deberian estar vacios por eleminarlos,\n")
	fmt.Printf("%s %s:\n%v\n",nombre_planeta,nombre_ciudad,Relojes[nombre_planeta+" "+nombre_ciudad].ip)
	fmt.Printf("%v\n",Relojes[nombre_planeta+" "+nombre_ciudad].servers)

}

func menu() error{

	respuesta := -1
	fmt.Printf("Bienvenido al nuevo sistema rebelde.\n")
	for respuesta != 0 {

		fmt.Printf("Que desea hacer?\n\n0: Salir\n1: AddCity\n2: UpdateName\n3: UpdateNumber\n4: DeleteCity\n\n")
		fmt.Scanf("%d",&respuesta)

		if respuesta == 0 {

			return nil
			
		} else if respuesta == 1{
			var planeta string
			var ciudad string
			cantidad := -1
			fmt.Printf("Se procedera a agregar una ciudad...\n\n")
			fmt.Printf("Escriba el nombre del planeta: \n\n")
			fmt.Scanf("%s",&planeta)
			fmt.Printf("Escriba el nombre de la ciudad: \n\n")
			fmt.Scanf("%s",&ciudad)
			fmt.Printf("Ingrese numero de rebeldes en el planeta (puede dejar vacio): \n\n")
			fmt.Scanf("%d",&cantidad)
			AddCity(planeta, ciudad, cantidad)
			
		} else if respuesta == 2 {
			var planeta string
			var ciudad string
			var nuevo_valor string
			fmt.Printf("Escriba el nombre del planeta: \n\n")
			fmt.Scanf("%s",&planeta)
			fmt.Printf("Escriba el nombre de la ciudad: \n\n")
			fmt.Scanf("%s",&ciudad)
			fmt.Printf("Ingrese el nuevo nombre de la ciudad: \n\n")
			fmt.Scanf("%s",&nuevo_valor)
			UpdateName(planeta, ciudad, nuevo_valor)
			
		} else if respuesta == 3 {
			var planeta string
			var ciudad string
			var cantidad int
			fmt.Printf("Escriba el nombre del planeta: \n\n")
			fmt.Scanf("%s",&planeta)
			fmt.Printf("Escriba el nombre de la ciudad: \n\n")
			fmt.Scanf("%s",&ciudad)
			fmt.Printf("Ingrese la nueva cantidad de rebeldes en el planeta: \n\n")
			fmt.Scanf("%d",&cantidad)
			UpdateNumber(planeta, ciudad, cantidad)
			
		} else if respuesta == 4 {
			var planeta string
			var ciudad string
			fmt.Printf("Escriba el nombre del planeta: \n\n")
			fmt.Scanf("%s",&planeta)
			fmt.Printf("Escriba el nombre de la ciudad: \n\n")
			fmt.Scanf("%s",&ciudad)
			DeleteCity(planeta, ciudad)
			
		} else {
			respuesta = -1
			fmt.Printf("Seleccione una opcion valida\n")
		}
	}
	return nil
}


func main() {
	
	menu()
	//fmt.Printf("%v\n",Relojes["Planet City"].servers)
	//fmt.Printf("%s\n",Relojes["Planet City"].ip)

	fmt.Printf("Terminando programa, rebelde\n")
}



