package models

import (
	"github.com/google/uuid"
)

type User struct {
	//id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Firstname string `gorm:"type:charactar varying(64);not null"`
	Email     string `gorm:"type:charactar varying(64);not null"`
	Lastname  string `gorm:"type:charactar varying(64);not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	id        uuid.UUID `json:"id,omitempty"`
	firstname string    `json:"firstname,omitempty"`
	email     string    `json:"email,omitempty"`
	lastname  string    `json:"lastname,omitempty"`
}

type UserResponseX struct {
	// id        uuid.UUID `json:"id,omitempty"`
	firstname string
	email     string
	lastname  string
}

type Leads struct {
	// id        uuid.UUID `json:"id,omitempty"`
	firstname string
	email     string
	lastname  string
}
