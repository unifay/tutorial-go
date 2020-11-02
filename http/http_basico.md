# Servidor Http basico en Go
Creamos el archivo http_basico.go con el siguiente contenido:
```
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


```
El comando para correrlo es 

```
$ go run ./http_basico.go
```
Despues del siguiente mensaje por terminal:

*Servidor http iniciando por el puerto 8080*

Vamos a la siguiente ruta:
```
http://localhost:8080/holamundo
```
Y deberiamos ver el mensaje : *Hola Mundo*