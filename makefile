test:
	go test -coverprofile=.coverage/coverage.out ./...

start:
	go run main.go

coverage:
	go tool cover -html=.coverage/coverage.out

build:
	go build