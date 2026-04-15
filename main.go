package main

import (
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/api/controllers"
	dependencies "github.com/murilosolino/challenge-backend-7/config"
	"github.com/murilosolino/challenge-backend-7/config/database"
	"github.com/murilosolino/challenge-backend-7/config/router"
)

func main() {
	database.CreateConnection()
	var deps map[string]func() any = dependencies.LoadDependencies()

	controllerAny := deps["ReviewController"]()
	controllerHealthCheck := deps["HealthCheckController"]()
	controllerDestinations := deps["DestinationsController"]()

	controller := controllerAny.(*controllers.ReviewsController)
	controllerHC := controllerHealthCheck.(*controllers.HealthCheckController)
	controllerDest := controllerDestinations.(*controllers.DestinationController)

	slog.Info("iniciando o servidor")
	router.InitServer(controller, controllerHC, controllerDest)
	slog.Info("Descendo o servidor")
}
