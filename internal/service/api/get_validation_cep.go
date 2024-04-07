package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (a API) isValidCEP(cep string) error {

	jsonBody, err := json.Marshal(map[string]string{"CEP": cep})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://localhost"+a.Config.WebValidationCEPServerPort, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnprocessableEntity {
		return errors.New("invalid zipcode")
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(string(res.Status))
	}

	return nil
}
