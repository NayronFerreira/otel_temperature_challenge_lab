package response

type DataResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type GetTemperatureServiceResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    DataResponse `json:"data,omitempty"`
}
