package model

import (
	_ "github.com/go-playground/validator"
	"github.com/google/uuid"
)


type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
}
