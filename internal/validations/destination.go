package validations

import (
	"fmt"

	"github.com/murilosolino/challenge-backend-7/internal/dto"
)

func ValidateDestination(dest dto.Destination) error {
	if len(dest.Name) == 0 || len(dest.Name) > 255 {
		return fmt.Errorf("o nome do destino não pode ser vazio nem superior a 255 caracteres")
	}

	if dest.Price <= 0 {
		return fmt.Errorf("o preço do destino não pode negativo ou igual a 0")
	}
	return nil
}
