package models

import "samplelab-go/src/enum"

type User struct {
	ID       int64     `gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"-"`
	Role     enum.Role `json:"role"`
}

func (User) TableName() string {
	return "users"
}
