package handler

import (
	"net/http"

	"github.com/NayronFerreira/microservice-ratelimiter/model/dto"
	"github.com/NayronFerreira/microservice-ratelimiter/utils"
)

func DummyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, dto.ResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "Request allowed",
			Success:    true,
		})
	})
}
