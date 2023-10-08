package utils

// LoanCalculationResult represents the result of loan calculations
type LoanCalculationResult struct {
	TotalPayment        float64
	PrincipalPerMonth   float64
	FeePerMonth         float64
	InterestPerMonth    float64
	TotalInterest       float64
	TotalFeeStampDuty   float64
	TotalDemandPerMonth float64
}

// CalculateLoanSimulations calculates the loan simulations
func CalculateLoanSimulations(amount float64, tenor int) *LoanCalculationResult {
	feePercentage := 0.05
	monthlyInterest := 0.0199
	stampDuty := 10000.0

	principalPerMonth := amount / float64(tenor)
	feePerMonth := amount * feePercentage
	interestPerMonth := amount * monthlyInterest
	totalInterest := interestPerMonth * float64(tenor)
	totalFeeStampDuty := feePerMonth*float64(tenor) + stampDuty
	totalPayment := amount + totalInterest + totalFeeStampDuty
	totalDemandPerMonth := totalPayment / float64(tenor)

	return &LoanCalculationResult{
		TotalPayment:        totalPayment,
		PrincipalPerMonth:   principalPerMonth,
		FeePerMonth:         feePerMonth,
		InterestPerMonth:    interestPerMonth,
		TotalInterest:       totalInterest,
		TotalFeeStampDuty:   totalFeeStampDuty,
		TotalDemandPerMonth: totalDemandPerMonth,
	}
}
