package structs

type Request struct {
	AmountRequired string `json:"AmountRequired"`
	Term           string `json:"Term"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	DateOfBirth    string `json:"DateOfBirth"`
	Mobile         string `json:"Mobile"`
	Email          string `json:"Email"`
}
