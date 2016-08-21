TARGET ?= ./...
COVERAGEFILE ?= $(TMPDIR)/coverage.out

bootstrap:
	go get -t -v $(TARGET)

generate:
	go generate

test: generate
	go tool vet .
	go test -cover $(TARGET)

benchmark: generate
	go test -bench=. $(TARGET)

coverage: generate
	go test -coverprofile=$(COVERAGEFILE)
	go tool cover -func=$(COVERAGEFILE)

shared: generate
	go install -buildmode=shared -linkshared $(TARGET)

clean:
	go clean -ix $(TARGET)
