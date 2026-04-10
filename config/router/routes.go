package router

import (
	"log/slog"
	"net/http"

	"github.com/murilosolino/challenge-backend-7/api/controllers"
	"github.com/murilosolino/challenge-backend-7/api/middleware"
)

func InitServer(
	controllerReviews *controllers.ReviewsController,
	conntrollerHc *controllers.HealthCheckController,
) {

	r := http.NewServeMux()
	r.HandleFunc("GET /", conntrollerHc.HealthCheck)
	r.HandleFunc("GET /reviews", controllerReviews.ListAllReviews)
	r.HandleFunc("GET /reviews/{id}", controllerReviews.FindReviewById)
	r.HandleFunc("GET /reviews-home", controllerReviews.GetRandomReviews)
	r.HandleFunc("POST /reviews", controllerReviews.CreateNewReview)
	r.HandleFunc("PATCH /reviews/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("PUT /reviews/{id}", controllerReviews.UpdateReview)
	r.HandleFunc("DELETE /review/{id}", controllerReviews.ExceludeReview)

	var handler http.Handler = r
	handler = middleware.CORSMiddleware(r)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		slog.Error("Erro ao iniciar o servidor web", "erro", err)
		panic(err)
	}
}
