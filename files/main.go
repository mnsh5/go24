package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var filePath = "/home/a779433/repos/me/go24/go24/files/hello.txt"

func Files() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Crea un escáner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Lee y muestra cada línea del archivo
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Verifica si hubo errores durante la lectura del archivo
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Abrir el archivo en modo de escritura/creación (O_APPEND para agregar al final)
	file2, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Contenido adicional a agregar al archivo
	additionalContent := "\nMore content to add to hello.txt"

	// Escribir el contenido adicional al final del archivo
	_, err = file2.WriteString(additionalContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Content added successfully to hello.txt")
}
