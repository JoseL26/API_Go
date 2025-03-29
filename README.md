# Fiber App - Go + Docker + AWS Serverless

Este proyecto es una API construida con [Go](https://go.dev/) y [Fiber](https://gofiber.io/), contenedorizada con [Docker](https://www.docker.com/) y desplegada en AWS utilizando [Serverless Framework](https://www.serverless.com/).

## 📌 Requisitos

Asegúrate de tener instalados los siguientes componentes en tu sistema:

- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/get-started) (Versión 28.0.1, build 068a01e)
- [Serverless Framework](https://www.serverless.com/framework/docs/getting-started) (Core: 3.40.0, Plugin: 7.2.3, SDK: 4.5.1)
- [AWS CLI](https://aws.amazon.com/cli/)

## 🚀 Instalación y Ejecución en Local

### 1. Clonar el repositorio

```sh
git clone https://github.com/tuusuario/fiber-app.git
cd fiber-app
```

### 2. Construir la imagen de Docker

```sh
$ docker build --no-cache -t fiber-app .
```

### 3. Ejecutar el contenedor

```sh
$ docker run -p 3000:3000 fiber-app
```

# 📦 Despliegue en AWS con Serverless Framework

### 1. Configurar credenciales de AWS

```sh
$ aws configure
```

### 2. Instalar dependencias del proyecto

```sh
$ go mod tidy
```

### 3. Desplegar con Serverless

```sh
$ serverless deploy
```

4. Obtener la URL de la API desplegada

```sh
$ serverless info
```
Esto te mostrará la URL de la API en AWS API Gateway.

# 🛑 Eliminar la API de AWS

Si deseas eliminar el servicio de AWS, usa:

```sh
$ serverless remove
```

# 📖 Endpoints

# 🛠 Variables de Entorno
Asegúrate de configurar las siguientes variables de entorno antes de ejecutar la API:

```sh
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=TU_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY=TU_SECRET_KEY
```
| Método | Ruta             | Descripción  |
| :---:  | :---:            | :---:           |
| POST   | /login           | Autenticación de usuarios   |
| POST   | /factorizar-qr   | Procesa códigos QR   |
