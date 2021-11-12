package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sachinsmc/hubuc-task/model"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type Response struct {
	Message string `json:"message"`
	User model.User `json:"user"`
}


func Add(u *model.User) Response {
	u.ID = uuid.New()
	validate = validator.New()
	err := validate.Struct(u)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			fmt.Println(err.StructField())
			return Response{
				Message: err.StructField() + " is required or invalid",
				User: *u,
			}
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(hash)
	if err != nil {
		fmt.Println("Failed to Create password hash : ", err)
	}

	return Response{
		Message: "User created successfully",
		User: *u,
	}
}
