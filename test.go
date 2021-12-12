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

func Append_line_to_file(line string, path string) {
	// append a line to a file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(line + "\n"); err != nil {
		fmt.Println(err)
	}

}

func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	line := "This is a test"
	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}
}