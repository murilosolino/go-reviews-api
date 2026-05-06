package validations_test

import (
	"testing"

	"github.com/murilosolino/challenge-backend-7/internal/dto"
	"github.com/murilosolino/challenge-backend-7/internal/validations"
)

var reviewTest = []struct {
	in       string
	rev      dto.Review
	errorMsg string
}{
	{
		in: "Teste Review Vazia",
		rev: dto.Review{
			Id:         1,
			Review:     "",
			AuthorName: "ABC",
			Url_photo:  "jfds",
		},
		errorMsg: "a review não pode ser vazia",
	},
	{
		in: "Teste Nome do autor vazio",
		rev: dto.Review{
			Id:         1,
			Review:     "ABC",
			AuthorName: "",
			Url_photo:  "jfds",
		},
		errorMsg: "o nome do author não pode ser vazio nem exceder 255 caracteres",
	},
	{
		in: "Teste Nome do autor maior que o max permitido",
		rev: dto.Review{
			Id:         1,
			Review:     "ABC",
			AuthorName: up255char(),
			Url_photo:  "jfds",
		},
		errorMsg: "o nome do author não pode ser vazio nem exceder 255 caracteres",
	},
}

func TestReviewValidation(t *testing.T) {
	for _, tt := range reviewTest {
		t.Run(tt.in, func(t *testing.T) {
			err := validations.ValidateReview(tt.rev)
			if err.Error() != tt.errorMsg {
				t.Errorf("Esperado: %v | Recebido: %v", tt.errorMsg, err.Error())
			}
		})
	}
}
