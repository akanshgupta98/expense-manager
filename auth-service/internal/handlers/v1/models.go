package v1

type RegistrationPayload struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,gte=8"`
	FirstName string `json:"first_name" binding:"required" `
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
}

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RegistrationResponse struct {
	APIResponse APIResponse `json:"response"`
}

type User RegistrationPayload
type FetchUsersResponse struct {
	Users []User
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	APIResponse APIResponse `json:"response"`
}
