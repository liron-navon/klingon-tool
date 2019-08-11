test:
	go test -coverprofile=_coverage/coverage.out ./...

start:
	go run main.go

coverage:
	go tool cover -html=_coverage/coverage.out
