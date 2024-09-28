package mappers

import (
	"apirestful-go/internal/dtos"
	"apirestful-go/internal/models"
)

func ToUserGet(user models.User) dtos.UserGet {
	return dtos.UserGet{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		DNI:      user.DNI,
		Phone:    user.Phone,
		Address:  user.Country + ", " + user.State + ", " + user.City + " (" + user.PostalCode + "), " + user.Address,
	}
}

func ToUserModel(userUpsert dtos.UserUpsert) models.User {
	return models.User{
		Username:   userUpsert.Username,
		Name:       userUpsert.Name,
		Email:      userUpsert.Email,
		Age:        userUpsert.Age,
		DNI:        userUpsert.DNI,
		Phone:      userUpsert.Phone,
		Country:    userUpsert.Country,
		State:      userUpsert.State,
		City:       userUpsert.City,
		Address:    userUpsert.Address,
		PostalCode: userUpsert.PostalCode,
		Password:   userUpsert.Password,
	}
}

func ToUserGetList(users []models.User) []dtos.UserGet {
	var userDtos []dtos.UserGet
	for _, user := range users {
		userDtos = append(userDtos, ToUserGet(user))
	}
	return userDtos
}
