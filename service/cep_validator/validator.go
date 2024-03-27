package service

func IsInvalidCEP(CEP string) bool {
	return len(CEP) != 8
}
