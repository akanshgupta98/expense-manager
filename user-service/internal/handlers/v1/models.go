package v1

type CreateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    int    `json:"user_id" binding:"required"`
	Country   string `json:"country"`
	Email     string `json:"email" binding:"required,email"`
}

type CreateProfileResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
