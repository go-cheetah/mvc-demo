package health

import (
	"demo/internal/config"
	"demo/internal/pkg/database"
	"demo/internal/pkg/response"
	healthModel "demo/internal/model/health"
)

type Health struct{
	Config *config.Config
	DB     database.DB
	HealthModelClient *healthModel.Client
}

func (h *Health) Status() response.Response {
	return response.Success("health")
}
