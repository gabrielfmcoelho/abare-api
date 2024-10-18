.PHONY: default run build test docs clean

APP_NAME=Abare

default: docs run

run:
	@go run cmd/main.go

build:
	@go build -o $(APP_NAME) main.go

tests:
	@go test ./ ...

docs:
	@swag init -g cmd/main.go

clean:
	@rm -f $(APP_NAME)
	@rm -rf docs