tests:
	go test ./... -v


build:
	env GOSUMDB=off GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

clean:
	rm -rf ./bin

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml