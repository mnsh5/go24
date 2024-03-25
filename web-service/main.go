package main

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Creamos un mapa con la respuesta que queremos enviar
	respuesta := map[string]string{"message": "Ping Pong"}

	// Convertimos el mapa a JSON
	respuestaJSON, err := json.Marshal(respuesta)
	if err != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
		return
	}

	// Escribimos la respuesta JSON en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuestaJSON)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/api", hello)
	http.ListenAndServe(":3000", nil)
}
