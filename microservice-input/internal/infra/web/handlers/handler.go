package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/microservice-input/internal/exceptions"
	"github.com/NayronFerreira/microservice-input/internal/model/dto"
	"github.com/NayronFerreira/microservice-input/internal/service"
	"github.com/NayronFerreira/microservice-input/internal/usecase"
	"github.com/NayronFerreira/microservice-input/internal/utils"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Handler struct {
	weatherApiService service.GetTemperatureServiceInterface
}

type InputDTO struct {
	Cep string `json:"cep"`
}

func NewHandler(weatherApiService service.GetTemperatureServiceInterface) *Handler {
	return &Handler{
		weatherApiService: weatherApiService,
	}
}

func (h *Handler) GetTemperatures(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)
	tracer := otel.Tracer(viper.GetString("SERVICE_NAME"))

	ctx, span := tracer.Start(ctx, "GetTemperaturesHandler")
	defer span.End()

	var input InputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JsonResponse(w, dto.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	if !utils.IsValidCEP(input.Cep) {
		utils.JsonResponse(w, dto.ResponseDTO{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    exceptions.ErrInvalidCEP.Error(),
			Success:    false,
		})
		return
	}

	getTemperaturesUseCase := usecase.NewGetTemperatureUseCase(h.weatherApiService)
	data, err := getTemperaturesUseCase.Execute(ctx, input.Cep)
	if err != nil {
		if err.Error() == exceptions.ErrInvalidCEP.Error() {
			utils.JsonResponse(w, dto.ResponseDTO{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    err.Error(),
				Success:    false,
			})
			return
		}

		if err.Error() == exceptions.ErrCannotFindZipcode.Error() {
			utils.JsonResponse(w, dto.ResponseDTO{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Success:    false,
			})
			return
		}

		utils.JsonResponse(w, dto.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	utils.JsonResponse(w, dto.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Success:    true,
		Data:       data,
	})
}
