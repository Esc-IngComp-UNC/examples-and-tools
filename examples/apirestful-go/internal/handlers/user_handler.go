package handlers

import (
	"apirestful-go/internal/dtos"
	service "apirestful-go/internal/services"
	"apirestful-go/pkg/errors"
	"apirestful-go/pkg/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(rt *chi.Mux)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternal("Internal server error", err))
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, users)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest dtos.UserUpsert
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid request payload", err))
		return
	}

	validationErrors := userRequest.Validate()
	if validationErrors != nil {
		helpers.RespondWithError(w, errors.NewBadRequest(validationErrors.Error(), validationErrors))
		return
	}

	user, err := h.service.Create(userRequest)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternal("Failed to create user", err))
		return
	}
	helpers.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetByID(id)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternal("Failed to get user", err))
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, user)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var userUpdate dtos.UserUpsert
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		helpers.RespondWithError(w, errors.NewBadRequest("Invalid request payload", err))
		return
	}

	validationErrors := userUpdate.Validate()
	if validationErrors != nil {
		helpers.RespondWithError(w, errors.NewBadRequest(validationErrors.Error(), validationErrors))
		return
	}

	user, err := h.service.Update(id, userUpdate)
	if err != nil {
		helpers.RespondWithError(w, errors.NewInternal("Failed to update user", err))
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, user)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.Delete(id); err != nil {
		helpers.RespondWithError(w, errors.NewInternal("Failed to delete user", err))
		return
	}
	helpers.RespondWithoutData(w, http.StatusNoContent)
}

func (h *userHandler) RegisterRoutes(rt *chi.Mux) {
	rt.Get("/users", h.GetAll)
	rt.Post("/users", h.Create)
	rt.Get("/users/{id}", h.GetByID)
	rt.Put("/users/{id}", h.Update)
	rt.Delete("/users/{id}", h.Delete)
}
