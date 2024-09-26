package service

import (
	"github.com/nable-fusion/fusion-cloud-common/pkg/log"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	IntegrationPointsHealthCheck() error
}

type Service struct {
	logger *zap.Logger
}

func NewService() ServiceInterface {
	i := &Service{
		logger: log.Logger(),
	}
	i.logger.Info("service layer created")
	return i
}

// IntegrationPointsHealthCheck provides the service-tier details for the API (HEAD /).
func (i *Service) IntegrationPointsHealthCheck() error {
	//todo healthchecks added here
	return nil
}
