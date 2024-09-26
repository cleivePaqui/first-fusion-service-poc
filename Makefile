APP_NAME := api-template
APP_NAMESPACE := fusion-api-template

BRANCH := $(if $(BRANCH_NAME),$(BRANCH_NAME),$(shell git rev-parse --symbolic-full-name --abbrev-ref HEAD))
COMMIT_ID := $(shell git rev-parse --short HEAD)
VERSION ?= $(strip $(if $(filter-out $(BRANCH),main),$(shell semtag getfinal)-$(COMMIT_ID)-SNAPSHOT,$(shell semtag getfinal)))

OAPI_PACKAGE := v1
OAPI_DIRECTORY := api/${OAPI_PACKAGE}
OAPI_SPEC_PATH := ${OAPI_DIRECTORY}/openapi3_0.yaml

DOCKER_REPO_SUFFIX := $(if $(filter-out $(BRANCH),main),-snapshot,)
DOCKER_REPO := 616350704275.dkr.ecr.eu-west-1.amazonaws.com/fusion
API_SERVER_DOCKER_REPO := ${DOCKER_REPO}/${APP_NAME}-server${DOCKER_REPO_SUFFIX}

.PHONY: update-from-template
update-from-template:
	go mod tidy
	go get github.com/nable-fusion/fusion-service-generator@v1
	go run tools/applytemplate/main.go
	go get github.com/nable-fusion/fusion-service-generator@v1

.PHONY: version
version:
	@echo "${VERSION}"

.PHONY: release
release:
	semtag final -s minor

.PHONY: go-env
go-env:
	@go env -w GOPRIVATE=github.com/nable-fusion/*,github.com/logicnow/*
	@go mod download

# OAPI / MOCKS
include build/make/generate-code.mk

# DOCKER
include build/make/docker.mk

# TERRAFORM
include build/make/terraform.mk

# TEST
include build/make/integration-tests.mk
include build/make/unit-tests.mk

# PRE-COMMIT
include build/make/pre-commit.mk

# LINTING
include build/make/lint.mk

# SECRETS

.PHONY: generate-cerberus-sealed-secret-dev
generate-cerberus-sealed-secret-dev:
	aws eks update-kubeconfig --name=slr-cdo-dev-01-euwe1 --region eu-west-1 --profile cdo_dev-developer
	kubectl create secret generic cerberus --namespace ${APP_NAMESPACE} --dry-run=client --from-literal clientid=$(CLIENT_ID) --from-literal clientsecret=$(CLIENT_SECRET) -o yaml | \
	kubeseal --controller-namespace=sealedsecrets --controller-name=forge-sealed-secrets --format=yaml > deployments/k8s/overlays/dev/eu-west-1/resources/sealedsecret-cerberus.yaml
	kubectl config use-context docker-desktop

.PHONY: generate-cerberus-sealed-secret-stg
generate-cerberus-sealed-secret-stg:
	aws eks update-kubeconfig --name=slr-cdo-stg-01-euwe1 --region eu-west-1 --profile cdo_stg-developer
	kubectl create secret generic cerberus --namespace ${APP_NAMESPACE} --dry-run=client --from-literal clientid=$(CLIENT_ID) --from-literal clientsecret=$(CLIENT_SECRET) -o yaml | \
	kubeseal --controller-namespace=sealedsecrets --controller-name=forge-sealed-secrets --format=yaml > deployments/k8s/overlays/stg/eu-west-1/resources/sealedsecret-cerberus.yaml
	kubectl config use-context docker-desktop

.PHONY: generate-local-cerberus-secrets
generate-local-cerberus-secrets:
	kubectl create secret generic cerberus --namespace ${APP_NAMESPACE} --dry-run=client --from-literal clientid=$(CLIENT_ID) --from-literal clientsecret=$(CLIENT_SECRET) -o yaml > \
	deployments/k8s/overlays/local/resources/cerberus-secrets.yaml


.PHONY: generate-localstack-secrets
generate-localstack-secrets:
	kubectl create secret generic localstack --namespace ${APP_NAMESPACE} --dry-run=client --from-literal apikey=$(LOCALSTACK_API_KEY) -o yaml > \
	deployments/k8s/overlays/local/resources/localstack-api-key.yaml
