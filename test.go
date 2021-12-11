package main

import (
	"fmt"
	// "os"
	// "bufio"
	// "strings"
)

func main() {
	// read a file line by line
	// for each line, print the line

	// open the file
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // create a scanner
	// scanner := bufio.NewScanner(file)

	// text := ""
	// // loop through the file
	// for scanner.Scan() {
	// 	text += scanner.Text() + "|"
	// }
	// fmt.Print(text)

	// // split text into lines using | as separator
	// lines := strings.Split(text, "|")
	// // print lines
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	var word string
	var num int
	str := "hola 1"
	a,b := fmt.Sscanf(str, "%s %d", &word, &num)

	fmt.Println(word, num, a, b)

	str = "hola"
	c, d := fmt.Sscanf(str, "%s %d", &word, &num)

	fmt.Println(word, num, c, d)
}