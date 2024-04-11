package utils

import (
	"regexp"

	"github.com/NayronFerreira/microservice-orchestration/internal/exceptions"
)

func FormatCEP(cep string) (string, error) {
	cepRegEx := `^\d{5}-\d{3}$`

	if regexp.MustCompile(cepRegEx).MatchString(cep) {
		return cep, nil
	}

	cepWithoutDash := regexp.MustCompile(`^\d{8}$`)
	if cepWithoutDash.MatchString(cep) {
		return cep[:5] + "-" + cep[5:], nil
	}

	return "", exceptions.ErrInvalidCEP
}
