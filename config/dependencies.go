package dependencies

import (
	"github.com/murilosolino/challenge-backend-7/api/controllers"
	"github.com/murilosolino/challenge-backend-7/api/model"
	"github.com/murilosolino/challenge-backend-7/api/services"
	"github.com/murilosolino/challenge-backend-7/config/database"
)

var dependecies map[string]func() any = map[string]func() any{
	"ReviewController": func() any {
		db := database.GetConnection()
		model := model.NewReviewModel(db)
		svc := services.NewReviewService(*model)
		return controllers.NewReviewController(*svc)
	},
	"HealthCheckController": func() any {
		return controllers.NewHealthCheck()
	},
}

func LoadDependencies() map[string]func() any {
	return dependecies
}
