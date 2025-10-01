package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Estructura para almacenar tanto la imagen en Base64 como el nombre del archivo
type ImageData struct {
	Name   string
	Base64 string
}

// ImagenABase64 lee imágenes de los directorios especificados y las convierte a base64.
// Selecciona 2 imágenes de cada directorio.
func ImagenABase64(directorios []string) []ImageData {
	var images []ImageData

	// Iteramos por cada directorio
	for _, directorio := range directorios {
		files, err := os.ReadDir(directorio)
		if err != nil {
			fmt.Println("Error leyendo el directorio:", err)
			return nil
		}

		var nombres []string // Slice para guardar los nombres

		// Guardar los nombres de las imágenes
		for _, file := range files {
			nombre := file.Name()
			// Buscar el último punto para quitar la extensión
			if idx := strings.LastIndex(nombre, "."); idx != -1 {
				nombre = nombre[:idx]
			}
			nombres = append(nombres, nombre) // Añadir al slice
		}

		if len(nombres) == 0 {
			fmt.Println("No se encontraron imágenes en el directorio:", directorio)
			continue
		}

		// Seleccionamos 2 imágenes aleatorias por directorio
		r := rand.New(rand.NewSource(time.Now().UnixNano())) // Generador local

		// Limitamos a 2 imágenes por directorio
		numImages := 2
		if len(nombres) < 2 {
			numImages = len(nombres) // Si hay menos de 2 imágenes, seleccionamos todas
		}

		// Seleccionar las imágenes aleatoriamente
		for i := 0; i < numImages; i++ {
			indice := r.Intn(len(nombres)) // Índice aleatorio
			nombreAleatorio := nombres[indice]
			base64img, err := buscarYConvertirImagen(directorio, nombreAleatorio)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				images = append(images, ImageData{
					Name:   nombreAleatorio + ".png", // Asumimos que la extensión es .png, ajusta si es necesario
					Base64: base64img,
				})
			}
		}
	}

	return images
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
				data, err := os.ReadFile(ruta)
				if err != nil {
					return "", err
				}
				return base64.StdEncoding.EncodeToString(data), nil
			}
		}
	}
	return "", fmt.Errorf("no se encontró la imagen original para: %s", nombreSinExtension)
}

func iniciarHttp(port, directorio1, directorio2 string) {
	// Preparar las imágenes en base64 y sus nombres, desde ambos directorios
	imagenes := ImagenABase64([]string{directorio1, directorio2})
	if imagenes == nil {
		fmt.Println("No se pudieron cargar imágenes, cerrando el servidor.")
		return
	}

	// Cargar el archivo HTML desde el disco
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error al cargar la plantilla HTML:", err)
		return
	}

	// Servir los archivos estáticos (CSS, imágenes, etc.)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Renderizar la plantilla HTML con el nombre del host y las imágenes
		host, err := os.Hostname()
		if err != nil {
			http.Error(w, "Error al obtener el nombre del host", http.StatusInternalServerError)
			return
		}

		data := struct {
			Host   string
			Images []ImageData
		}{
			Host:   host,
			Images: imagenes, // Pasar las imágenes con nombre y base64
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		}
	})

	fmt.Printf("Servidor escuchando en http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func main() {
	// Recibir el puerto y los directorios como argumentos
	port := flag.String("port", "8080", "Puerto donde correrá el servidor")
	directorio1 := flag.String("dir1", "./imagenes/directorio1", "Directorio de imágenes 1")
	directorio2 := flag.String("dir2", "./imagenes/directorio2", "Directorio de imágenes 2")
	flag.Parse()

	// Iniciar el servidor HTTP con el puerto y directorios indicados
	iniciarHttp(*port, *directorio1, *directorio2)
}
