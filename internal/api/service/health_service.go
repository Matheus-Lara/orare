package service

type HealthService struct{}

func (*HealthService) GetHealthMessage() string {
	return "Welcome to Orare API!"
}

func NewHealthService() *HealthService {
	return &HealthService{}
}
