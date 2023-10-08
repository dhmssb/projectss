package main

import (
	"projectss/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define the API endpoint for loan simulation
	app.Post("/loan_simulation", http.LoanSimulationHandler)

	// Define the API endpoint for Ahmad's monthly payments
	app.Get("/ahmad_monthly_payments", http.AhmadMonthlyPaymentsHandler)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
