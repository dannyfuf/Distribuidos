package main

import (
	"fmt"
	// "io/ioutil"
	"os"
	// "bufio"
	// "strings"
)

func Check_file_exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	fmt.Printf("%d\n", Check_file_exists("servers/data/planets/tierra.txt"))
	fmt.Printf("%d\n", Check_file_exists("servers/data/planets/AAAAAA.txt"))
}