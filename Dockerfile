# Etapa de construcción
FROM golang:1.24.1 AS builder
WORKDIR /app

# Copiar archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

# Copiar el código fuente
COPY . .

# Construir el binario (con nombre "bootstrap" para AWS Lambda)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bootstrap .

# Etapa final (imagen liviana)
FROM alpine:latest
WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/bootstrap /app/bootstrap
RUN chmod +x /app/bootstrap

# Exponer puerto 3000 solo si se ejecuta en Docker
EXPOSE 3000

# Si se ejecuta en Lambda, el binario ya es "bootstrap"
CMD ["./bootstrap"]
