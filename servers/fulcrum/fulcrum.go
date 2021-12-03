package fulcrum

import (
	"log"
	"src/common"
)


type Server struct {
	clocks		map[string][]int
}

// check if file exist in data folder, if not create it
func check_file(file_name string) {
	_, err := os.Stat("data/" + file_name)
	if os.IsNotExist(err) {
		log.Println("Creating Planet file: " + file_name)
		file, err := os.Create("data/" + file_name)
		check_error(err, "Error al crear el archivo")
		file.Close()
	}
}

// write a register to a planet file in data folder
func write_register(planet string, line string) {
	file, err := os.OpenFile("data/" + planet, os.O_APPEND|os.O_WRONLY, 0600)
	check_error(err, "Error al abrir el archivo")
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	check_error(err, "Error al escribir en el archivo")
}

// delete file in data folder
func delete_file(file_name string) {
	err := os.Remove("data/" + file_name)
	check_error(err, "Error al eliminar el archivo")
}