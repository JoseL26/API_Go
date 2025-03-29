# Etapa de construcción
FROM golang:1.24.1 AS builder
WORKDIR /app

# Copiar archivos de módulos
COPY go.mod go.sum ./

# Descargar dependencias antes de copiar el código
RUN go mod download

# Copiar el código fuente
COPY src/ .

# Verificar módulos y construir la aplicación de forma compatible con Alpine Linux
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Etapa final (más liviana)
FROM alpine:latest
WORKDIR /app

# Copiar binario desde la etapa anterior y asegurarse de que tenga permisos de ejecución
COPY --from=builder /app/app /app/app
RUN chmod +x /app/app

# Exponer el puerto
EXPOSE 3000

# Ejecutar la aplicación
CMD ["./app"]
