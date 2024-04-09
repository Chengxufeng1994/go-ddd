package dto

type CustomerCreationRequest struct {
	Email     string `json:"email"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CustomerCreationResponse struct{}
