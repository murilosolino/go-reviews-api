package dependencies

import (
	"github.com/murilosolino/challenge-backend-7/internal/config/database"
	"github.com/murilosolino/challenge-backend-7/internal/controllers"
	"github.com/murilosolino/challenge-backend-7/internal/model"
	"github.com/murilosolino/challenge-backend-7/internal/services"
)

var dependencies map[string]func() any = map[string]func() any{
	"ReviewController": func() any {
		db := database.GetConnection()
		baseModel := model.NewBaseModel(db)
		model := model.NewReviewModel(*baseModel)
		svc := services.NewReviewSvc(*model)
		return controllers.NewReviewController(svc)
	},
	"HealthCheckController": func() any {
		return controllers.NewHealthCheck()
	},
	"DestinationsController": func() any {
		db := database.GetConnection()
		baseModel := model.NewBaseModel(db)
		model := model.NewDestinationModel(*baseModel)
		svc := services.NewDestinationSvc(*model)
		return controllers.NewDestinationController(svc)
	},
}

func LoadDependencies() map[string]func() any {
	return dependencies
}
