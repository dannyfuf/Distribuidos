package main

import (
	"fmt"
)

func GetNumberRebelds(nombre_planeta string, nombre_ciudad string) string {

	fmt.Println("Solicitando informacion sobre rebeldes en " + nombre_planeta + ", en la ciudad " + nombre_ciudad)
	//funcion que envia hacia broker aqui

	//funcion que recibe desde broker aqui

	//Dependiendo de como se envie la informacion se procede a trabajarla mas o enviarla tal y como llegue
	return "yes"

}

func main() {
	fmt.Println("Contador de Rebeldes segun planeta")

	fmt.Println("Ingrese el nombre del planeta a comprobar:")
	var planeta string
	fmt.Scanln(&planeta)

	fmt.Println("Ingrese el nombre de la ciudad a comprobar:")
	var ciudad string
	fmt.Scanln(&ciudad)

	result := GetNumberRebelds(planeta, ciudad)
	fmt.Println(result)

}
