.PHONY: fmt test lint run-playground run-geomcli run-stringcli

fmt:
	go fmt ./...

test:
	go test ./...

lint:
	golangci-lint run

run-playground:
	go run ./cmd/playground

run-geomcli:
	go run ./cmd/geomcli

run-stringcli:
	go run ./cmd/stringcli

run-picli:
	go run ./cmd/picli --timeout 200ms --verbose

run-semaphore:
	go run ./cmd/semaphorecli