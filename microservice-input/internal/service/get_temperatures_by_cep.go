package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/NayronFerreira/microservice-input/internal/model/response"
)

type GetTemperatureServiceInterface interface {
	GetTemperatureService(ctx context.Context, cep string) (response.GetTemperatureServiceResponse, error)
}

type GetTemperatureService struct {
	client *http.Client
}

func NewGetTemperatureService() *GetTemperatureService {
	return &GetTemperatureService{
		client: &http.Client{},
	}
}

func (s *GetTemperatureService) GetTemperatureService(ctx context.Context, cep string) (retval response.GetTemperatureServiceResponse, err error) {

	URL := viper.GetString("WEATHER_SERVICE_URL") + "?cep=" + cep

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return retval, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	res, err := s.client.Do(req)
	if err != nil {
		return retval, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&retval); err != nil {
		return retval, err
	}

	if !retval.Success {
		return retval, errors.New(retval.Message)
	}

	return retval, nil
}
