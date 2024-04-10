package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	model "github.com/NayronFerreira/otel_temperature_challenge_lab/internal/model/response"
)

func (a API) GetCelciusByLocality(city, state string) (celcius float64, err error) {
	timeoutSeconds, err := strconv.Atoi(a.Config.TimeoutSeconds)
	if err != nil {
		return celcius, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	location := url.QueryEscape(fmt.Sprintf("%s,%s", city, state))
	weatherUrl := fmt.Sprintf("%s/current.json?key=%s&q=%s", a.Config.WeatherHost, a.Config.WeatherKey, location)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, weatherUrl, nil)
	if err != nil {
		return celcius, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return celcius, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return celcius, err
	}
	var weatherRes model.WeatherRes
	if err = json.Unmarshal(body, &weatherRes); err != nil {
		return celcius, err
	}

	return weatherRes.Current.TempC, nil
}
