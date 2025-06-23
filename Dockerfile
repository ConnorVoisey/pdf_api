# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pdf_api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...



FROM docker.io/chromedp/headless-shell:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /pdf_api /pdf_api
RUN mkdir /logs

EXPOSE 3000

ENTRYPOINT ["/pdf_api"]
