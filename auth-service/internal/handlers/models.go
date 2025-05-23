package handlers

type RegistrationPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"pasword" binding:"required,gte=8"`
}

type RegistrationResponse struct {
	Msg string `json:"msg"`
}
