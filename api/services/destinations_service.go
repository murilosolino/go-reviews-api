package services

import (
	"context"
	"log/slog"
	"os"

	"github.com/murilosolino/challenge-backend-7/api/model"
	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type DestinationSvc struct {
	dm           model.DestinationModel
	openIAClient openai.Client
}

func NewDestinationSvc(dm model.DestinationModel) *DestinationSvc {
	opts := []option.RequestOption{}
	if apiKey := os.Getenv("OPENAI_API_KEY"); apiKey != "" {
		opts = append(opts, option.WithAPIKey(apiKey))
	}
	client := openai.NewClient(opts...)

	return &DestinationSvc{dm: dm, openIAClient: client}
}

func (svc *DestinationSvc) CreateDestination(m map[string]interface{}) error {
	return svc.dm.Bm.Save(m)
}

func (svc *DestinationSvc) FindByName(name string) (model.DestinationRow, error) {
	return svc.dm.FindByName(name)
}

func (svc *DestinationSvc) ListDestinations() ([]model.DestinationRow, error) {
	return svc.dm.ListAllDestinations()
}

func (svc *DestinationSvc) DeleteDestinationById(id int) error {
	return svc.dm.Bm.Exclude(id)
}

func (svc *DestinationSvc) UpdateDestination(id int, m map[string]interface{}) error {
	return svc.dm.Bm.Update(id, m)
}

func (svc *DestinationSvc) GenerateIADescriptiveText(destination string) string {
	prompt := "Escreva um texto descritivo de até 150 caracteres sobre o destino: " + destination + ". A resposta deve conter apenas o texto gerado"

	resp, err := svc.openIAClient.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: openai.ChatModelGPT5Nano,
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
	})
	if err != nil {
		slog.Error("[OPENAI] Falha ao gerar descriptive_text", "erro", err)
		return ""
	}

	return resp.OutputText()
}
