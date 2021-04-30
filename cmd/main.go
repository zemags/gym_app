/*
Main application for this project
*/

package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/zemags/gym_app"
	"github.com/zemags/gym_app/pkg/handler"
	"github.com/zemags/gym_app/pkg/repository"
	"github.com/zemags/gym_app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error loading configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gym_app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
