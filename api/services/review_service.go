package services

import (
	"fmt"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/api/model"
)

type ReviewService struct {
	rm model.ReviewModel
}

func NewReviewService(rm model.ReviewModel) *ReviewService {
	return &ReviewService{rm: rm}
}

func (svc *ReviewService) GetAllReviews() ([]model.ReviewsRow, error) {
	return svc.rm.List()
}

func (svc *ReviewService) Get3RandomReviews() ([]model.ReviewsRow, error) {
	return svc.rm.FindRandomRegisters(3)
}

func (svc *ReviewService) CreateReview(rev model.ReviewsRow) error {
	return svc.rm.Save(rev)
}

func (svc *ReviewService) UpdateReview(id int, r map[string]interface{}) error {
	if len(r) == 0 {
		err := fmt.Errorf("Nenhum campo enviado para a atualização")
		slog.Error("[ReviewService][UpdateReview] " + err.Error())
		return err
	}
	return svc.rm.Update(id, r)
}

func (svc *ReviewService) ExceludeReview(id int) error {
	return svc.rm.Delete(id)
}

func (svc *ReviewService) SearchById(id int) (model.ReviewsRow, error) {
	return svc.rm.FindById(id)
}
