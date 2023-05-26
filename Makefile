.PHONY: build
build:
	go build -v ./cmd/orders-service

.PHONY: test
test:
	go test -v -race -timeout 30s .\...

.PHONY: swagger
swagger:
	swag init -g internal/controller/http/v1/router.go --output docs/order-service

.DEFAULT_GOAL := build