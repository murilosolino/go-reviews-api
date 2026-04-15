package services

import (
	"fmt"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/api/model"
)

type ReviewSvc struct {
	rm model.ReviewModel
}

func NewReviewSvc(rm model.ReviewModel) *ReviewSvc {
	return &ReviewSvc{rm: rm}
}

func (svc *ReviewSvc) GetAllReviews() ([]model.ReviewsRow, error) {
	return svc.rm.List()
}

func (svc *ReviewSvc) Get3RandomReviews() ([]model.ReviewsRow, error) {
	return svc.rm.FindRandomRegisters(3)
}

func (svc *ReviewSvc) CreateReview(rev model.ReviewsRow) error {
	return svc.rm.Save(rev)
}

func (svc *ReviewSvc) UpdateReview(id int, r map[string]interface{}) error {
	if len(r) == 0 {
		err := fmt.Errorf("Nenhum campo enviado para a atualização")
		slog.Error("[ReviewSvc][UpdateReview] " + err.Error())
		return err
	}
	return svc.rm.Update(id, r)
}

func (svc *ReviewSvc) ExceludeReview(id int) error {
	return svc.rm.Delete(id)
}

func (svc *ReviewSvc) SearchById(id int) (model.ReviewsRow, error) {
	return svc.rm.FindById(id)
}
