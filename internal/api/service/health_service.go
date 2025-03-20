package service

import (
	"github.com/Matheus-Lara/orare/internal/api/dto"
	"github.com/Matheus-Lara/orare/internal/i18n"
)

type HealthService struct{}

func (*HealthService) GetHealth() dto.HealthResponseDTO {
	return *dto.NewHealthResponseDTO(i18n.Message("HealthCheck.Response.Success"))
}

func NewHealthService() *HealthService {
	return &HealthService{}
}
