package main

import (
	"fmt"
	"net/http"
	"os" // Necesario para leer variables de entorno
)

// Variable global para almacenar el nombre del Pod
var podName string

func Index(rw http.ResponseWriter, r *http.Request) {
	// 1. Añadimos el nombre del Pod a la respuesta HTML
	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Mi Web en Go</title>
		</head>
		<body>
			<h1>Soy alumno de la UOC</h1>
			<p>Respondiendo desde el Pod: <strong>%s</strong></p> 
			<img src="/static/antidisturb.jpeg" alt="Imagen de prueba" width="300">
		</body>
	</html>
	`, podName) // Inyectamos el valor de podName en el HTML

	fmt.Fprint(rw, htmlContent)
}

func main() {
	// Leer el nombre del pod al iniciar la aplicación
	// Kubernetes establece la variable HOSTNAME como el nombre del Pod
	podName = os.Getenv("HOSTNAME")
	if podName == "" {
		podName = "Nombre de Pod no encontrado (Entorno Local)"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Index)

	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Visita: http://localhost:3000")

	// Usamos el puerto 3000, que es el puerto TargetPort/ContainerPort corregido.
	http.ListenAndServe(":3000", nil)
}
