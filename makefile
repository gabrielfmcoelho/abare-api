

dev:
	go run main.go

prod:
	go build -o bin/main main.go
	./bin/main