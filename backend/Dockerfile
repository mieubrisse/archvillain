FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install oapi-codegen for code generation
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate Go code from OpenAPI spec
RUN oapi-codegen -generate types -package main openapi.yaml > generated_types.go
RUN oapi-codegen -generate gorilla -package main openapi.yaml > generated_server.go

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates docker-cli
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]