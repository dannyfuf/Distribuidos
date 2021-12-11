package main
import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"

	"servers/broker"
)

func get_env_var(key string) string {

	// load .env file
	err := godotenv.Load(".env")	
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func ConnectBroker() string{
	var ipBroker string = get_env_var("IP_SERVER_17")
	var portBroker string = get_env_var("BROKER_PORT")
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
	})
	
	response, _ := stream.Recv()
	return response
}

func AddCity(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {

	ipFulcrum = ConnectBroker()
	var portFulcrum string = get_env_var("FULCRUM_PORT")

	fmt.Printf("Se va a conectar al servidor con ip: %s:%s",ipFulcrum, portFulcrum)

	fmt.Printf("Estoy en AddCity\nEstoy agregando la ciudad %s al planeta %s ", nombre_ciudad, nombre_planeta)
	if nuevo_valor >= 0 {
		fmt.Printf("con %d rebeldes\n\n", nuevo_valor)
	} else {
		fmt.Printf("donde aun no hay rebeldes.\n\n")

	}

	



}

func UpdateName(nombre_planeta string, nombre_ciudad string, nuevo_valor string) {
	
	ipFulcrum = ConnectBroker()
	var portFulcrum string = get_env_var("FULCRUM_PORT")

	fmt.Printf("Se va a conectar al servidor con ip: %s:%s",ipFulcrum, portFulcrum)
	fmt.Printf("En el planeta %s, se esta actualizando el nombre de la ciudad %s a %s.\n\n", nombre_planeta, nombre_ciudad, nuevo_valor)
}

func UpdateNumber(nombre_planeta string, nombre_ciudad string, nuevo_valor int) {

	ipFulcrum = ConnectBroker()
	var portFulcrum string = get_env_var("FULCRUM_PORT")

	fmt.Printf("Se va a conectar al servidor con ip: %s:%s",ipFulcrum, portFulcrum)
	fmt.Printf("En el planeta %s, se esta actualizando la cantidad de rebeldes de la ciudad %s a %d.\n\n", nombre_planeta, nombre_ciudad, nuevo_valor)

}

func DeleteCity(nombre_planeta string, nombre_ciudad string) {

	ipFulcrum = ConnectBroker()
	var portFulcrum string = get_env_var("FULCRUM_PORT")

	fmt.Printf("Se va a conectar al servidor con ip: %s:%s",ipFulcrum, portFulcrum)
	fmt.Printf("La ciudad %s del planeta %s ha sido destruida...\n\n", nombre_ciudad, nombre_planeta)
}
func menu() {
	respuesta := -1
	fmt.Printf("Bienvenido al nuevo sistema rebelde.\n")
	for respuesta != 0 {

		fmt.Printf("Que desea hacer?\n\n0: Salir\n1: AddCity\n2: UpdateName\n3: UpdateNumber\n4: DeleteCity\n\n")
		fmt.Scanf("%d",&respuesta)

		if respuesta == 0 {
			
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
	return
}

func main() {

	menu()
	fmt.Printf("Terminando programa, rebelde\n")
}