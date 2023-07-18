package structs

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type LoanApplicationDetails struct {
	AmountRequired   string `json:"AmountRequired"`
	Term             string `json:"Term"`
	Title            string `json:"Title"`
	FirstName        string `json:"FirstName" `
	LastName         string `json:"LastName"`
	DateOfBirth      string `json:"DateOfBirth"`
	Mobile           string `json:"Mobile"`
	Email            string `json:"Email"`
	Repayment        string `json:"Repayment"`
	EstablishmentFee string `json:"EstablishmentFee"`
	TotalInterest    string `json:"TotalInterest"`
}

type LoanApplicationRedirection struct {
	URL string `json:"url"`
}
