package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherApiServiceInterface interface {
	GetWeatherData(ctx context.Context, location string) (*WeatherAPIResponse, error)
}

type WeatherApiService struct {
	client *http.Client
}

func NewWeatherApiService() *WeatherApiService {
	return &WeatherApiService{
		client: &http.Client{},
	}
}

func (s *WeatherApiService) GetWeatherData(ctx context.Context, location string) (retVal *WeatherAPIResponse, err error) {
	tracer := otel.Tracer(viper.GetString("SERVICE_NAME"))
	ctx, span := tracer.Start(ctx, "weather-api.get-weather-data")
	defer span.End()

	WEATHER_API_KEY := viper.GetString("WEATHER_API_KEY")
	WEATHER_HOST_API := viper.GetString("WEATHER_HOST_API")
	urlString := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", WEATHER_HOST_API, WEATHER_API_KEY, url.QueryEscape(location))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlString, nil)
	if err != nil {
		return retVal, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return retVal, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return retVal, errors.New("cannot find weather data")
	}

	if err = json.NewDecoder(res.Body).Decode(&retVal); err != nil {
		return nil, err
	}

	return retVal, nil
}
