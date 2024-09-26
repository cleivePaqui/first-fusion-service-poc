package main

import "github.com/nable-fusion/fusion-service-generator/pkg/service-generator"

func main() {
	serviceGenerator := service_generator.NewApiGenerator()
	serviceGenerator.GenerateServiceFromJsonPath("api-template.json", ".")
}
