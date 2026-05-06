package router

import (
	"log/slog"
	"net/http"

	"github.com/murilosolino/challenge-backend-7/internal/controllers"
	"github.com/murilosolino/challenge-backend-7/internal/middleware"
)

func InitServer(
	controllerReviews *controllers.ReviewsController,
	conntrollerHc *controllers.HealthCheckController,
	controllerDest *controllers.DestinationController,
) {
	r := http.NewServeMux()

	//HealthCheck
	r.HandleFunc("GET /", conntrollerHc.HealthCheck)
	//Reviews
	r.HandleFunc("GET /reviews", controllerReviews.ListAllReviews)
	r.HandleFunc("GET /reviews/{id}", controllerReviews.FindReviewById)
	r.HandleFunc("GET /reviews-home", controllerReviews.GetRandomReviews)
	r.HandleFunc("POST /reviews", controllerReviews.CreateNewReview)
	r.HandleFunc("PATCH /reviews/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("PUT /reviews/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("DELETE /reviews/{id}", controllerReviews.ExceludeReview)
	//Destinations
	r.HandleFunc("POST /destinations", controllerDest.CreateNewDestination)
	r.HandleFunc("GET /destinations", controllerDest.ListAllDestinations)
	r.HandleFunc("DELETE /destinations/{id}", controllerDest.DeleteDestination)
	r.HandleFunc("PUT /destinations/{id}", controllerDest.UpdateDestination)

	h := middleware.CORSMiddleware(r)

	err := http.ListenAndServe(":8080", h)
	if err != nil {
		slog.Error("Erro ao iniciar o servidor web", "erro", err)
		panic(err)
	}
}
