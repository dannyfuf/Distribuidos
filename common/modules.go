package common

import (
	"log"
	"os"
	"bufio"
	
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