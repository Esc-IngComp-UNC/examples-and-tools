package server

import (
	"apirestful-go/config"
	"apirestful-go/internal/handlers"
	"apirestful-go/internal/repository"
	services "apirestful-go/internal/services"
	"apirestful-go/pkg/helpers"
	"context"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerChi struct {
	serverAddress string
	mongoURI      string
	databaseName  string
}

func NewServerChi(cfg *config.Config) *ServerChi {
	return &ServerChi{
		serverAddress: cfg.ServerAddr,
		mongoURI:      cfg.DBConnString,
		databaseName:  cfg.DatabaseName,
	}
}

func (a *ServerChi) Run() (err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(a.mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(a.databaseName)

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	apiRouter.Mount("/api/v1", rt)

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userHandler.RegisterRoutes(rt)

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithJSON(w, http.StatusOK, "pong")
	})

	err = http.ListenAndServe(a.serverAddress, apiRouter)

	return
}
