package service_test

import (
	"github.com/nable-fusion/fusion-cloud-common/pkg/log"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (suite *ServiceTestSuite) SetupTest() {
	_ = log.NewLogger(log.NewConfig())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
