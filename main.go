package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Seleccione una opción:")
	fmt.Println("1. convertir una imagen a base64 de manera aleatoria atraves de un directorio ingresado")
	fmt.Println("2. imprimir el nombre del host del sistema")
	fmt.Println("3. iniciar un servidor HTTP que responda con 'Hola Mundo!'")
	fmt.Print("Ingrese el número de la opción: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	opcion := scanner.Text()

	switch opcion {
	case "1":
		ImagenABase64()
	case "2":
		host()
	case "3":
		iniciarHttp()
	default:
		fmt.Println("Opción no válida")
	}
}
