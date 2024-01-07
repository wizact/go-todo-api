# Dockerfile used for the build image
FROM golang:1.19-alpine

LABEL org.opencontainers.image.source=https://github.com/wizact/go-todo-api
LABEL org.opencontainers.image.description="Todo API Builder Image"
LABEL org.opencontainers.image.licenses=MIT

# required for CGO builds
RUN apk add build-base