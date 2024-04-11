package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NayronFerreira/microservice-orchestration/internal/exceptions"
	"github.com/NayronFerreira/microservice-orchestration/internal/model/response"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

type ViaCepServiceInterface interface {
	GetCEPData(ctx context.Context, cep string) (*response.ViaCEPResponse, error)
}

type ViaCepService struct {
	client *http.Client
}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{client: &http.Client{}}
}

func (s *ViaCepService) GetCEPData(ctx context.Context, cep string) (retVal *response.ViaCEPResponse, err error) {
	tracer := otel.Tracer(viper.GetString("SERVICE_NAME"))
	ctx, span := tracer.Start(ctx, "ViaCEPService.GetCEPData")
	defer span.End()

	url := viper.GetString("VIACEP_HOST_API") + cep + "/json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return retVal, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return retVal, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return retVal, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&retVal); err != nil {
		return retVal, err
	}

	if retVal.Erro == "true" {
		return retVal, exceptions.ErrCannotFindZipcode
	}

	return retVal, nil
}
