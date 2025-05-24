package handlers

type RegistrationPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"pasword" binding:"required,gte=8"`
}

type APIResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RegistrationResponse struct {
	APIResponse APIResponse
}
