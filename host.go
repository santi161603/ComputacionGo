package main

import (
	"fmt"
	"os"
)

func host() {

	nombre, err := os.Hostname()
	if err != nil {
		fmt.Println("Error obteniendo el nombre del host:", err)
		return
	}
	fmt.Println("Nombre del host:", nombre)

}
