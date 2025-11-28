package main

import (
	"fmt"
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Mi Web en Go</title>
		</head>
		<body>
			<h1>Soy alumno de la UOC</h1>
			<img src="/static/antidisturb.jpeg" alt="Imagen de prueba" width="300">
		</body>
	</html>
	`
	fmt.Fprint(rw, htmlContent)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Index)

	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Visita: http://localhost:3000")

	http.ListenAndServe(":3000", nil)
}
