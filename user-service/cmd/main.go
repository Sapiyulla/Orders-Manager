package main

import (
	"os"
	"user-service/config"
	"user-service/internal/application"

	"user-service/internal/infrastructure/http"
	"user-service/internal/infrastructure/repository"
	"user-service/pkg/database"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

var Configuration *config.Config

func init() {
	Conf, err := config.Load()
	if err != nil {
		logrus.Error(err)
		return
	}
	Configuration = Conf
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
}

func main() {
	postgres, err := database.NewPostgres(Configuration.Database, os.Getenv("DATABASE_PASSWORD"))
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	defer postgres.Close()

	repo := repository.NewPGUserRepository(postgres.Pool)
	usecase := application.NewAccountRepository(repo)
	handler := http.NewUserHandler(usecase)

	handler.SetHandler("POST", "/api/v1/register", handler.Register)
	handler.SetHandler("POST", "/api/v1/login", handler.Login)

	logrus.Fatal(handler.Run(Configuration.REST))
}
