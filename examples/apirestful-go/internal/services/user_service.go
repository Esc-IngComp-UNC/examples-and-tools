package user

import (
	"apirestful-go/internal/dtos"
	"apirestful-go/internal/mappers"
	"apirestful-go/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetAll() ([]dtos.UserGet, error)
	GetByID(id string) (dtos.UserGet, error)
	Create(userRequest dtos.UserUpsert) (dtos.UserGet, error)
	Update(id string, userUpdate dtos.UserUpsert) (dtos.UserGet, error)
	Delete(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAll() ([]dtos.UserGet, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return mappers.ToUserGetList(users), nil
}

func (s *userService) GetByID(id string) (dtos.UserGet, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.UserGet{}, err
	}

	user, err := s.repo.GetByID(objectID)
	if err != nil {
		return dtos.UserGet{}, err
	}

	return mappers.ToUserGet(user), nil
}

func (s *userService) Create(userRequest dtos.UserUpsert) (dtos.UserGet, error) {
	user := mappers.ToUserModel(userRequest)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return dtos.UserGet{}, err
	}

	return mappers.ToUserGet(createdUser), nil
}

func (s *userService) Update(id string, userUpdate dtos.UserUpsert) (dtos.UserGet, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.UserGet{}, err
	}

	user, err := s.repo.GetByID(objectID)
	if err != nil {
		return dtos.UserGet{}, err
	}

	user = mappers.ToUserModel(userUpdate)

	user.UpdatedAt = time.Now()

	updatedUser, err := s.repo.Update(objectID, user)
	if err != nil {
		return dtos.UserGet{}, err
	}

	return mappers.ToUserGet(updatedUser), nil
}

func (s *userService) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(objectID)
}
