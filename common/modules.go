package common

import (
	"log"
	"os"
	"bufio"
	"strings"
	"fmt"
	"io/ioutil"
	"strconv"
	
	"github.com/joho/godotenv"

)

func Get_env_var(key string) string {
	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Check_error(e error, msg string) bool {
	if e != nil {
		log.Printf("%s", msg)
		log.Printf("Error: %v", e)
		return true
	}
	return false
}

func Get_neighbors_fulcrum(n int) []string {
	if n == 18 {
		return []string{Get_env_var("IP_SERVER_19"), Get_env_var("IP_SERVER_20")}
	} else if n == 19 {
		return []string{Get_env_var("IP_SERVER_18"), Get_env_var("IP_SERVER_20")}
	} else if n == 20 {
		return []string{Get_env_var("IP_SERVER_18"), Get_env_var("IP_SERVER_19")}
	} else {
		return nil
	}
}

func Get_file_as_string(path string) string {
	file, err := os.Open(path)
	Check_error(err, "Error al abrir el archivo")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	text := ""
	for scanner.Scan() {
		text += scanner.Text() + "|"
	}
	return text
}

func Get_max(array []int) int {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}

func Get_string_file_as_map(text string) map[string]int {
	// split text by |
	cities := strings.Split(text, "|")

	var planet string
	var city string
	var amount int
	var cities_map = make(map[string]int)
	for i := 0; i < len(cities); i++ {
		fmt.Sscanf(cities[i], "%s %s %d", &planet, &city, &amount)
		cities_map[city] = amount
	}

	return cities_map
}

func Check_file_exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Write_map_to_file(merge_map map[string]int, planet_name string) {
	log.Println("Writing map to file")	

	if Check_file_exists("data/planets/"+planet_name) {
		os.Remove("data/planets/"+planet_name)
	}

	fmt.Println("CREANDO EL ARCHIVO: " + "data/planets/" + planet_name)
	os.Create("data/planets/"+planet_name)

	if planet_name != "" {
		f, err := os.OpenFile("data/planets/"+planet_name, os.O_RDWR, 0755)
		Check_error(err, "Error al abrir el archivo al escribir el map")
		defer f.Close()

		//split filename by .
		name := strings.Split(planet_name, ".")[0]

		// write map to file
		for city, amount := range merge_map {
			f.WriteString(fmt.Sprintf("%s %s %d\n", name, city, amount))
		}
	}
}

func Append_line_to_file(line string, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	Check_error(err, "Error al abrir el archivo")
	defer f.Close()

	_, err = f.WriteString(line+"\n")
	Check_error(err, "Error al escribir en el archivo")
}

func Create_file_list(path string) []string {
	// create a list of files
	files, err := ioutil.ReadDir(path)
	Check_error(err, "Error al leer el directorio")
	var file_list []string
	for _, file := range files {
		file_list = append(file_list, file.Name())
	}
	return file_list
}

func Contains(elem string, list []string) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}
	return false
}

func Array_as_string(array []int) string {
	str := fmt.Sprintf("%d,%d,%d", array[0], array[1], array[2])
	return str
}

func String_as_array(str string) []int {
	array := strings.Split(str, ",")
	var int_array []int
	for _, v := range array {
		n, _ := strconv.Atoi(v)
		int_array = append(int_array, n)
	}
	return int_array
}

func Check_line(path string){
	// read file line by line
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	Check_error(err, "Error al abrir el archivo")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var planet string
	var city string
	var amount int
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := fmt.Sscanf(line, "%s %s %d", &planet, &city, &amount)
		if n == 3 && city != "0" {
			lines = append(lines, line)
		}
	}

	os.Remove(path)
	os.Create(path)

	// write lines to file
	n_file, err := os.OpenFile(path, os.O_RDWR, 0755)
	Check_error(err, "Error al abrir el archivo")
	defer file.Close()
	fmt.Println("---------------------------")
	for _, line := range lines {
		fmt.Println(line)
		n_file.WriteString(line+"\n")
	}
}