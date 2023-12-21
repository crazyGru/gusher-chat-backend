FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY / ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gusher-backend ./cmd/backend

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /gusher-backend /gusher-backend
COPY --from=build-stage /app/config.yml /config.yml

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/gusher-backend", "config.yml"]