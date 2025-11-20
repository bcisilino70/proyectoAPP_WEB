# Usamos una imagen base de Go oficial
FROM golang:1.24

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el código fuente
COPY . .

# Construimos la aplicación
RUN go build -o main .

# Exponemos el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]