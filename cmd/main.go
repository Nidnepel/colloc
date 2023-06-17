package main

import (
	"SoftDesignColloc/internal/config"
	"SoftDesignColloc/internal/database"
	"SoftDesignColloc/internal/handler"
	"SoftDesignColloc/internal/repository"
	"SoftDesignColloc/internal/service"
	"SoftDesignColloc/server"
	"fmt"
	"log"
)

const (
	serverPort = "7000"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.New(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDbHost, cfg.PostgresDb),
	)
	if err != nil {
		log.Fatalf("невозможно подключиться к базе: %v", err)
	}

	repos := repository.NewRepository(database.NewPGX(db))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(serverPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
