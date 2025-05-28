package dto

import "samplelab-go/src/enum"

type User struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password,omitempty" binding:"required"`
	Role     enum.Role `json:"role"`
}

type RegisterInput struct {
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" binding:"required,email"`
	Role  enum.Role `json:"role"`
}
