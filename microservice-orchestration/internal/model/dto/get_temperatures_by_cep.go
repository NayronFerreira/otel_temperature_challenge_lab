package dto

type ResponseDTO struct {
	StatusCode int
	Success    bool
	Message    string
	Data       interface{}
}
