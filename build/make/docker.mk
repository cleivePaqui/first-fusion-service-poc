.PHONY: docker-build-api-server
docker-build-api-server:
	@docker build -f cmd/api-server/Dockerfile -t api-server:$(VERSION) --build-arg GITHUB_USERNAME=$(GITHUB_USERNAME) --build-arg GITHUB_PASSWORD=$(GITHUB_PASSWORD) --no-cache .

.PHONY: docker-run-api-server
docker-run-api-server:
	docker run -p 1323:1323 -e APP_CERBERUS_CLIENT_SECRET=$(CLIENT_SECRET) -e APP_CERBERUS_CLIENT_ID=$(CLIENT_ID)  -it api-server:$(VERSION)
