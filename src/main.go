package main

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/AnthuanGarcia/practicaNoSql/db"
)

func ListarColeccion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {

		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Metodo Invalido")
		return

	}

	collection := r.URL.Query().Get("collection")
	response, err := db.ObtenerColeccion(collection)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
		return

	}

	w.WriteHeader(http.StatusOK)

	for _, res := range response {

		w.Write(res)

	}

}

func InsertarDocumento(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Metodo Invalido")
		return

	}

	var doc interface{}

	collection := r.URL.Query().Get("collection")
	err := json.NewDecoder(r.Body).Decode(&doc) // Se decodifica el cuerpo de la peticion

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(err)
		return

	}

	if collection == "" {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	err = db.InsertarDocumento(doc, collection)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(err)
		return

	}

	w.WriteHeader(http.StatusOK)

}

func main() {

	http.HandleFunc("/listar/", ListarColeccion)
	http.HandleFunc("/insertar/", InsertarDocumento)

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)

}
