package validations

import (
	"fmt"

	"github.com/murilosolino/challenge-backend-7/internal/dto"
)

func ValidateReview(rev dto.Review) error {

	if len(rev.Review) == 0 {
		return fmt.Errorf("a review não pode ser vazia")
	}

	if len(rev.AuthorName) == 0 || len(rev.AuthorName) > 255 {
		return fmt.Errorf("o nome do author não pode ser vazio nem exceder 255 caracteres")
	}

	return nil
}
