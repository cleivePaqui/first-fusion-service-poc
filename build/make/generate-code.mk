# OPEN API CODEGEN
.PHONY: oapi-gen
oapi-gen:
	oapi-codegen -generate server -package ${OAPI_PACKAGE} ${OAPI_SPEC_PATH} > ${OAPI_DIRECTORY}/server.gen.go
	oapi-codegen -generate types -package ${OAPI_PACKAGE} ${OAPI_SPEC_PATH} > ${OAPI_DIRECTORY}/models.gen.go
	oapi-codegen -generate client -package ${OAPI_PACKAGE} ${OAPI_SPEC_PATH} > ${OAPI_DIRECTORY}/client.gen.go

# GENERATE MOCKS
.PHONY: generate-mocks
generate-mocks:
	echo "https://github.com/vektra/mockery is required for this"
	rm -fr mocks
	mockery --all --with-expecter --keeptree
