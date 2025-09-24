package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ImagenABase64 lee una imagen desde la ruta especificada y la convierte a una cadena base64.
func ImagenABase64() {

	fmt.Print("Ingresa la ruta del directorio: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	directorio := scanner.Text()

	files, err := os.ReadDir(directorio)
	if err != nil {
		fmt.Println("Error leyendo el directorio:", err)
		return
	}

	var nombres []string // Slice para guardar los nombres

	for _, file := range files {
		nombre := file.Name()
		// Buscar el último punto
		if idx := strings.LastIndex(nombre, "."); idx != -1 {
			nombre = nombre[:idx]
		}
		nombres = append(nombres, nombre) // Añadir al slice
	}

	fmt.Println("Tamaño del array:", len(nombres))

	if len(nombres) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) // Generador local
		indice := r.Intn(len(nombres))                       // Índice aleatorio
		nombreAleatorio := nombres[indice]
		fmt.Println("Nombre aleatorio:", nombreAleatorio)

		base64img, err := buscarYConvertirImagen(directorio, nombreAleatorio)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Imagen en base64:", base64img)
		}

	} else {
		fmt.Println("La lista está vacía, elija una ruta con archivos.")
	}
}

// buscarYConvertirImagen busca el archivo original (con extensión) y lo convierte a base64
func buscarYConvertirImagen(directorio, nombreSinExtension string) (string, error) {
	files, err := os.ReadDir(directorio)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		nombreArchivo := file.Name()
		if idx := strings.LastIndex(nombreArchivo, "."); idx != -1 {
			if nombreArchivo[:idx] == nombreSinExtension {
				ruta := filepath.Join(directorio, nombreArchivo)
				data, err := ioutil.ReadFile(ruta)
				if err != nil {
					return "", err
				}
				return base64.StdEncoding.EncodeToString(data), nil
			}
		}
	}
	return "", fmt.Errorf("no se encontró la imagen original para: %s", nombreSinExtension)
}
