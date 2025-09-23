package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Ingresa la ruta del directorio: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	directorio := scanner.Text()

	files, err := os.ReadDir(directorio)
	if err != nil {
		fmt.Println("Error leyendo el directorio:", err)
		return
	}

	for _, file := range files {
		nombre := file.Name()
		// Buscar el último punto
		if idx := strings.LastIndex(nombre, "."); idx != -1 {
			nombre = nombre[:idx]
		}
		fmt.Println("-", nombre)
	}
}
