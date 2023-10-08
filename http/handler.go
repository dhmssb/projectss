package http

import (
	"projectss/loan"
	"projectss/utils"

	"github.com/gofiber/fiber/v2"
)

type LoanRequest struct {
	Amount float64 `json:"amount"`
	Tenor  int     `json:"tenor"`
}

func LoanSimulationHandler(c *fiber.Ctx) error {
	// Parse loan amount and tenor from the request
	var loanReq LoanRequest
	if err := c.BodyParser(&loanReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	loanService := loan.Service{}

	// Calculate the loan simulations
	result := loanService.CalculateLoanSimulations(loanReq.Amount, loanReq.Tenor)

	// Respond with the loan simulation results
	return c.JSON(fiber.Map{
		"loan_simulation": result,
	})
}

func AhmadMonthlyPaymentsHandler(c *fiber.Ctx) error {
	// Get the repayment amounts from the previous result
	loanSimulationResult := utils.CalculateLoanSimulations(6000000, 3) // ini untuk soal 1
	feeStampDuty := loanSimulationResult.TotalFeeStampDuty
	interestPerMonth := loanSimulationResult.InterestPerMonth
	principalPerMonth := loanSimulationResult.PrincipalPerMonth

	payments := []fiber.Map{
		{"bulan": "february", "payment": principalPerMonth + feeStampDuty},
		{"bulan": "march", "payment": principalPerMonth + interestPerMonth + feeStampDuty},
		{"bulan": "april", "payment": principalPerMonth + interestPerMonth + feeStampDuty},
		{"bulan": "may", "payment": principalPerMonth + interestPerMonth + feeStampDuty},
	}

	// Calculate total payments for each month
	var total float64
	for _, payment := range payments {
		total += payment["payment"].(float64)
	}

	// Include total payment and fee stamp duty
	payments = append(payments, fiber.Map{
		"bulan":   "total",
		"payment": total + feeStampDuty,
	})

	return c.JSON(fiber.Map{
		"pembayaran_tagihan_normal": payments,
	})
}
