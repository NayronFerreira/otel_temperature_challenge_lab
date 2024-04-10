package usecase

import (
	"context"

	"github.com/NayronFerreira/microservice-input/internal/model/entity"
	"github.com/NayronFerreira/microservice-input/internal/service"
)

type GetTemperaturesUseCase struct {
	weatherApiService service.GetTemperatureServiceInterface
}

func NewGetTemperatureUseCase(weatherApiService service.GetTemperatureServiceInterface) *GetTemperaturesUseCase {
	return &GetTemperaturesUseCase{
		weatherApiService: weatherApiService,
	}
}

func (u *GetTemperaturesUseCase) Execute(ctx context.Context, cep string) (retVal entity.Temperatures, err error) {

	weatherResData, err := u.weatherApiService.GetTemperatureService(ctx, cep)
	if err != nil {
		return retVal, err
	}

	retVal = entity.Temperatures{
		City:  weatherResData.Data.City,
		TempC: weatherResData.Data.TempC,
		TempF: weatherResData.Data.TempF,
		TempK: weatherResData.Data.TempK,
	}

	return retVal, nil

}
