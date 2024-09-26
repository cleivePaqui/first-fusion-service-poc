package main

import (
	"github.com/nable-fusion/fusion-cloud-common/pkg/log"
	"github.com/nable-fusion/fusion-cloud-common/pkg/server"
	v1 "github.com/nable-fusion/fusion-api-template/api/v1"
	"github.com/nable-fusion/fusion-api-template/internal/service"
)

func main() {
	_ = log.NewLogger(log.NewConfig())

	apiService := service.NewService()
	api := v1.NewApi(apiService)

	apiServer := server.NewServer("Fusion Api Template")
	v1.RegisterHandlers(apiServer.GetRouter(), &api)
	apiServer.Start()
}
