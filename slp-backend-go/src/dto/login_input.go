package dto

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
