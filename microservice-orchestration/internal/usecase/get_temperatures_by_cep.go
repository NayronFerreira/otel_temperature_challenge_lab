package usecase

import (
	"context"

	"github.com/NayronFerreira/microservice-orchestration/internal/model/entity"
	"github.com/NayronFerreira/microservice-orchestration/internal/service"
)

type GetTemperaturesUseCase struct {
	viaCepService     service.ViaCepServiceInterface
	weatherApiService service.WeatherApiServiceInterface
}

func NewGetTemperatureUseCase(
	viaCepService service.ViaCepServiceInterface,
	weatherApiService service.WeatherApiServiceInterface,
) *GetTemperaturesUseCase {
	return &GetTemperaturesUseCase{
		viaCepService:     viaCepService,
		weatherApiService: weatherApiService,
	}
}

func (u *GetTemperaturesUseCase) Execute(ctx context.Context, cep string) (retVal entity.Temperatures, err error) {

	cepData, err := u.viaCepService.GetCEPData(ctx, cep)
	if err != nil {
		return retVal, err
	}

	weatherData, err := u.weatherApiService.GetWeatherData(ctx, cepData.Localidade)
	if err != nil {
		return retVal, err
	}

	tempF := weatherData.Current.TempC*1.8 + 32
	tempK := weatherData.Current.TempC + 273

	retVal = entity.Temperatures{
		City:  cepData.Localidade,
		TempC: weatherData.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}

	return retVal, nil
}
