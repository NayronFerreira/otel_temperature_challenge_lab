package handlers

import (
	"net/http"

	"github.com/NayronFerreira/microservice-orchestration/internal/exceptions"
	"github.com/NayronFerreira/microservice-orchestration/internal/model/dto"
	"github.com/NayronFerreira/microservice-orchestration/internal/service"
	"github.com/NayronFerreira/microservice-orchestration/internal/usecase"
	"github.com/NayronFerreira/microservice-orchestration/internal/utils"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Handler struct {
	viaCepService     service.ViaCepServiceInterface
	weatherApiService service.WeatherApiServiceInterface
}

func NewHandler(
	viaCepService service.ViaCepServiceInterface,
	weatherApiService service.WeatherApiServiceInterface,
) *Handler {
	return &Handler{
		viaCepService:     viaCepService,
		weatherApiService: weatherApiService,
	}
}

func (h *Handler) GetTemperatures(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)
	tracer := otel.Tracer(viper.GetString("SERVICE_NAME"))

	ctx, span := tracer.Start(ctx, "get-temperatures-handler")
	defer span.End()

	cepParam := r.URL.Query().Get("cep")
	cep, err := utils.FormatCEP(cepParam)
	if err != nil {
		utils.JsonResponse(w, dto.ResponseDTO{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	getTemperaturesUseCase := usecase.NewGetTemperatureUseCase(h.viaCepService, h.weatherApiService)
	data, err := getTemperaturesUseCase.Execute(ctx, cep)
	if err != nil {
		if err == exceptions.ErrCannotFindZipcode {
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
