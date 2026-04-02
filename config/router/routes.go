package router

import (
	"log/slog"
	"net/http"

	"github.com/murilosolino/challenge-backend-7/api/controllers"
)

func InitServer(
	controllerReviews *controllers.ReviewsController,
	conntrollerHc *controllers.HealthCheckController,
) {

	r := http.NewServeMux()
	r.HandleFunc("GET /", conntrollerHc.HealthCheck)
	r.HandleFunc("GET /reviews", controllerReviews.ListAllReviews)
	r.HandleFunc("GET /reviews/{id}", controllerReviews.FindReviewById)
	r.HandleFunc("POST /post-review", controllerReviews.CreateNewReview)
	r.HandleFunc("PATCH /update-review/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("PUT /update-review/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("DELETE /remove-review/{id}", controllerReviews.ExceludeReview)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Error("Erro ao iniciar o servidor web", "erro", err)
		panic(err)
	}
}
