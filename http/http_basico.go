package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/holamundo", holaMundoHandler)
	log.Printf("Servidor http iniciando por el puerto %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola Mundo\n")
}
