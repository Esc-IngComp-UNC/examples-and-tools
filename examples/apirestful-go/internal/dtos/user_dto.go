package dtos

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserGet struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	DNI      string `json:"dni"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserUpsert struct {
	Username   string `json:"username" validate:"required,min=3,max=100,regexp=^[a-zA-Z0-9_]+$"`
	Name       string `json:"name" validate:"required,min=3,max=100"`
	Email      string `json:"email" validate:"required,email"`
	Age        int    `json:"age" validate:"required,min=18,max=100"`
	DNI        string `json:"dni" validate:"required,min=8,max=8"`
	Phone      string `json:"phone" validate:"required,min=10,max=10"`
	Country    string `json:"country" validate:"required,min=3,max=100"`
	State      string `json:"state" validate:"required,min=3,max=100"`
	City       string `json:"city" validate:"required,min=3,max=100"`
	Address    string `json:"address" validate:"required,min=3,max=100"`
	PostalCode string `json:"postal_code" validate:"required,min=4,max=4"`
	Password   string `json:"password" validate:"required,min=8,max=100"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(fl.Param())
		return re.MatchString(fl.Field().String())
	})
}

func (u *UserUpsert) Validate() error {
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "regexp":
			message = fmt.Sprintf("El campo %s debe contener solo letras, números y guiones bajos", err.Field())
		case "required":
			message = fmt.Sprintf("El campo %s es requerido", err.Field())
		case "min":
			message = fmt.Sprintf("El campo %s debe tener al menos %s caracteres", err.Field(), err.Param())
		case "max":
			message = fmt.Sprintf("El campo %s debe tener como máximo %s caracteres", err.Field(), err.Param())
		case "len":
			message = fmt.Sprintf("El campo %s debe tener exactamente %s caracteres", err.Field(), err.Param())
		}
		errorMessages = append(errorMessages, message)
	}

	return errors.New(strings.Join(errorMessages, ", "))
}
