package main

import (
	"fmt"
	// "io/ioutil"
	// "os"
	// "bufio"
	// "strings"
	// "strconv"
	// "src/common"
)

type infData struct {
	Servers []int
	Ip string
}

var Relojes = make(map[string]infData)

func main() {
	// Relojes["1"] = infData{Servers: []int{1, 2, 3}, Ip: "1.1.1.1"}
	// Relojes["2"] = infData{Servers: []int{1, 2, 3}, Ip: "1.2.3.4"}

	// // print Relojes
	// for k, v := range Relojes {
	// 	fmt.Println(k, v)
	// }
	// fmt.Printf("----------------------------\n")
	// // update ip of Relojes["1"]
	// t := Relojes["1"]
	// t.Ip = "1.2.3.4"
	// Relojes["1"] = t
	// for k, v := range Relojes {
	// 	fmt.Println(k, v)
	// }
	
	//Relojes[planet+" "+city] = infData{servers: t2, ip: ip}

	a := make(map[string][]int)
	a["1"] = []int{1, 2, 3}
	a["2"] = []int{1, 2, 3}

	a["1"] = []int{3, 2, 1}
	
	//print map
	for k, v := range a {
		fmt.Println(k, v)
	}

	return
}