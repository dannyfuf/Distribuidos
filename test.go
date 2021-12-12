package main

import (
	"fmt"
	"io/ioutil"
	// "os"
	// "bufio"
	// "strings"
)

func main() {
	files, _ := ioutil.ReadDir("servers/data/planets")
	//print file names
	for i:= 0; i < len(files); i++ {
		fmt.Println(files[i].Name())
	}
}