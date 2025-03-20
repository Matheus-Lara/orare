package dto

type HealthResponseDTO struct {
	Message string `json:"message"`
}

func NewHealthResponseDTO(message string) *HealthResponseDTO {
	return &HealthResponseDTO{Message: message}
}
