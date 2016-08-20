TARGET ?= ./...
deps := golang.org/x/tools/cmd/goimports

bootstrap:
	go get -t -v $(TARGET)
	go get -u -v $(deps)

generate:
	go generate

test: generate
	go tool vet .
	go test $(TARGET)

benchmark: generate
	go test -bench=. $(TARGET)

shared: generate
	go install -buildmode=shared -linkshared $(TARGET)

clean:
	go clean -ix $(TARGET)
