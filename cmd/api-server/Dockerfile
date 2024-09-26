# Step 1 - Build executable
FROM public.ecr.aws/docker/library/golang:1.21.0-alpine3.18 AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./

ARG GITHUB_USERNAME
ARG GITHUB_PASSWORD

RUN apk add --no-cache git
RUN git config --global credential.username $GITHUB_USERNAME
RUN git config --global credential.helper '!f() { echo "password=${GITHUB_PASSWORD}"; }; f'
RUN go env -w GOPRIVATE=github.com/nable-fusion/*,github.com/logicnow/*
RUN go mod download

COPY . .
RUN go build -buildvcs=false -o /go/bin/api-server ./cmd/api-server/

# Step 2 - Run application
FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/api-server .
USER 65534

ENTRYPOINT [ "./api-server" ]
