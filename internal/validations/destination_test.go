package validations_test

import (
	"testing"

	"github.com/murilosolino/challenge-backend-7/internal/dto"
	"github.com/murilosolino/challenge-backend-7/internal/validations"
)

func up255char() string {
	var str string
	for i := 0; i < 257; i++ {
		str += "a"
	}
	return str
}

var flagtests = []struct {
	in       string
	dest     dto.Destination
	errorMsg string
}{
	{
		in: "Valida preço não negativo",
		dest: dto.Destination{
			Name:  "Russia",
			Price: -45,
		},
		errorMsg: "o preço do destino não pode negativo ou igual a 0",
	},
	{
		in: "Valida preço não zero",
		dest: dto.Destination{
			Name:  "Russia",
			Price: 0,
		},
		errorMsg: "o preço do destino não pode negativo ou igual a 0",
	},
	{
		in: "Valida nome não vazio",
		dest: dto.Destination{
			Name:  "",
			Price: 45,
		},
		errorMsg: "o nome do destino não pode ser vazio nem superior a 255 caracteres",
	},
	{
		in: "Valida limite maximo de caracteres no nome",
		dest: dto.Destination{
			Name:  up255char(),
			Price: 45,
		},
		errorMsg: "o nome do destino não pode ser vazio nem superior a 255 caracteres",
	},
}

func TestValidateDestinations(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			err := validations.ValidateDestination(tt.dest)
			if err.Error() != tt.errorMsg {
				t.Fatalf("Recebida: %v. Esperado: %v", err.Error(), tt.errorMsg)
			}
		})
	}
}
