package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	model "github.com/NayronFerreira/otel_temperature_challenge_lab/model/response"
)

func (a API) GetLocationByCEP(cep string) (locality, UF string, err error) {
	timeoutSeconds, err := strconv.Atoi(a.Config.TimeoutSeconds)
	if err != nil {
		return locality, UF, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s/json/", a.Config.ViaCepHost, cep), nil)
	if err != nil {
		return locality, UF, err
	}

	// Cria um http.Client com a verificação de certificado desativada
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return locality, UF, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return locality, UF, errors.New("invalid zipcode")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locality, UF, err
	}

	var viaCepRes model.ViaCepRes
	if err = json.Unmarshal(body, &viaCepRes); err != nil {
		return locality, UF, err
	}

	if resp.StatusCode == http.StatusOK && viaCepRes.Localidade == "" {
		return locality, UF, errors.New("can not find zipcode")
	}

	return viaCepRes.Localidade, viaCepRes.Uf, nil
}
