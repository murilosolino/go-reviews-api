package services

import (
	"fmt"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/internal/dto"
	"github.com/murilosolino/challenge-backend-7/internal/model"
)

type ReviewSvc struct {
	rm model.ReviewModel
}

func NewReviewSvc(rm model.ReviewModel) *ReviewSvc {
	return &ReviewSvc{rm: rm}
}

func (svc *ReviewSvc) GetAllReviews() ([]dto.Review, error) {
	return svc.rm.List()
}

func (svc *ReviewSvc) Get3RandomReviews() ([]dto.Review, error) {
	return svc.rm.FindRandomRegisters(3)
}

func (svc *ReviewSvc) CreateReview(rev map[string]interface{}) error {
	return svc.rm.Bm.Save(rev)
}

func (svc *ReviewSvc) UpdateReview(id int, r map[string]interface{}) error {
	if len(r) == 0 {
		err := fmt.Errorf("nenhum campo enviado para a atualização")
		slog.Error("[ReviewSvc][UpdateReview] " + err.Error())
		return err
	}
	return svc.rm.Bm.Update(id, r)
}

func (svc *ReviewSvc) ExceludeReview(id int) error {
	return svc.rm.Bm.Exclude(id)
}

func (svc *ReviewSvc) SearchById(id int) (dto.Review, error) {
	return svc.rm.FindById(id)
}
