package main

import (
	"log/slog"

	dependencies "github.com/murilosolino/challenge-backend-7/internal/config"
	"github.com/murilosolino/challenge-backend-7/internal/config/database"
	"github.com/murilosolino/challenge-backend-7/internal/config/router"
	"github.com/murilosolino/challenge-backend-7/internal/controllers"
)

func main() {
	database.CreateConnection()
	deps := dependencies.LoadDependencies()

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
