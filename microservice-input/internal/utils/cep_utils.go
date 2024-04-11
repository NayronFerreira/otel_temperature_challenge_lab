package utils

import "regexp"

func IsValidCEP(cep string) bool {
	regex := regexp.MustCompile(`^\d{8}$`)

	if len(cep) != 8 {
		return false
	}

	if !regex.MatchString(cep) {
		return false
	}

	return true
}
