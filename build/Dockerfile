# Dockerfile used for the build image
ARG GO_VERSION=1.27.1
FROM golang:${GO_VERSION}-alpine3.24

LABEL org.opencontainers.image.source=https://github.com/wizact/go-todo-api
LABEL org.opencontainers.image.description="Todo API Builder Image"
LABEL org.opencontainers.image.licenses=MIT

# required for CGO builds
RUN apk add build-base
