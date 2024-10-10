test:
	go run cmd/main.go

install:
	go mod tidy
	go install go.uber.org/mock/mockgen@latest