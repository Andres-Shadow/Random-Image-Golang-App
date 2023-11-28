package main

import (
	"html/template"
	"net/http"
)

// StartServer inicia el servidor web.
func startServer(port, imageDir string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := getHostname()
		imagenes, err := obtenerImagenesAleatorias(imageDir, 3)
		if err != nil {
			http.Error(w, "No se encontraron imágenes", http.StatusInternalServerError)
			return
		}
		renderizarPagina(w, hostname, imagenes)
	})

	http.ListenAndServe(":"+port, mux)
}

// RenderizarPagina renderiza la página HTML.
func renderizarPagina(w http.ResponseWriter, hostname string, imagenes []Imagen) {
	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, "Error con el HTML", http.StatusInternalServerError)
		return
	}


	data := struct {
		Hostname string
		Imagenes []Imagen
	}{
		Hostname: hostname,
		Imagenes: imagenes,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error con el HTML", http.StatusInternalServerError)
	}
}
