FROM golang:1.25 AS builder

# Directorio de trabajo dentro del contenedor temporal
WORKDIR /app

# Copiamos solo el archivo de módulo para las dependencias (si lo tienes)
COPY go.mod ./
# Copiamos el código fuente principal y la carpeta estática
COPY app.go ./
COPY static ./static 

# Compilamos la aplicación apuntando al archivo app.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main app.go

# --- ETAPA 2: Producción (Runner) ---
# Usamos una imagen "Alpine" vacía y ligera para la imagen final 
FROM alpine:latest

WORKDIR /app

# Copiamos SOLO el binario compilado de la etapa anterior 
COPY --from=builder /app/main .

# Copiamos la carpeta 'static' para que se vea tu imagen
COPY --from=builder /app/static ./static

# Documentamos que la app usa el puerto 3000 
EXPOSE 3000

# Comando para arrancar tu servidor web 
CMD ["./main"]