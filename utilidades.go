package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Imagen struct {
	Nombre string
	Base64 string
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Desconocido"
	}
	return hostname
}

func obtenerImagenesAleatorias(directorio string, cantidad int) ([]Imagen, error) {
	var imagenes []Imagen
	var imagenesValidas []string

	rand.Seed(time.Now().UnixNano())

	err := filepath.Walk(directorio, func(ruta string, info os.FileInfo, err error) error {
		if !info.IsDir() && verificarImagen(ruta) {
			imagenesValidas = append(imagenesValidas, ruta)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; len(imagenes) < 3; i++ {

		rutaAbsoluta := imagenesValidas[rand.Intn(len(imagenesValidas))]
		nombre := filepath.Base(rutaAbsoluta)
		contenido, _ := ioutil.ReadFile(rutaAbsoluta)
		base64str := base64.StdEncoding.EncodeToString(contenido)

		imagen := Imagen{
			Nombre: nombre,
			Base64: base64str,
		}

		imagenes = append(imagenes, imagen)

	}

	return imagenes, nil
}

func verificarImagen(nombre string) bool {
	extensiones := []string{".jpg", ".jpeg", ".png"}
	extension := strings.ToLower(filepath.Ext(nombre))
	for i := 0; i < len(extensiones); i++ {
		if extension == extensiones[i] {
			return true
		}
	}
	return false
}
