package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"gonum.org/v1/gonum/mat"
	"API-Go/middleware"
	"API-Go/utils"
)

type MatrizSolicitud struct {
	Matriz [][]float64 `json:"matriz"`
}

type MatrizRespuesta struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

func calcularQR(Matriz [][]float64) ([][]float64, [][]float64, error) {
	filas := len(Matriz)
	columnas := len(Matriz[0])

	a := mat.NewDense(filas, columnas, nil)
	for i := 0; i < filas; i++ {
		a.SetRow(i, Matriz[i])
	}
	
	var qr mat.QR
	qr.Factorize(a)

	var q, r mat.Dense
	qr.QTo(&q)
	qr.RTo(&r)

	Q := make([][]float64, filas)
	R := make([][]float64, columnas)

	for i := 0; i < filas; i++ {
		Q[i] = q.RawRowView(i)
	}

	for i := 0; i < columnas; i++ {
		R[i] = r.RawRowView(i)
	}

	return Q, R, nil
}

func main() {
	utils.LoadEnv()

	app := fiber.New()

	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.CorsMiddleware())

	app.Post("/login", middleware.Login)

	app.Post("/factorizar-qr", middleware.JwtMiddleware(), func(c *fiber.Ctx) error {
		var req MatrizSolicitud
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		
		Q, R, err := calcularQR(req.Matriz)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error processing matriz"})
		}

		return c.JSON(MatrizRespuesta{
			Q: Q, R: R,
		})
	})

	port := utils.GetPort()
	log.Println("Servidor corriendo en http://localhost:", port)
	log.Fatal(app.Listen(port))
}
