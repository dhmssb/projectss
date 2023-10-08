package loan

import "projectss/utils"

type Service struct{}

func (s *Service) CalculateLoanSimulations(amount float64, tenor int) *utils.LoanCalculationResult {
	return utils.CalculateLoanSimulations(amount, tenor)
}
