.PHONY: test test-unit test-integration test-coverage

test: test-unit test-integration

test-unit:
		go test -v ./internal/... ./pkg/...

test-integration:
		go test -v ./test/integration/...

test-coverage:
		go test -coverprofile=coverage.out ./...

build:
	  go build -o bin/runtime cmd/runtime/main.go

run:
	  go run cmd/runtime/main.go

clean:
	  rm -rf bin/
