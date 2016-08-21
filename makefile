TARGET ?= ./...
COVERAGEFILE ?= $(TMPDIR)/coverage.out
deps := github.com/fzipp/gocyclo github.com/gordonklaus/ineffassign github.com/client9/misspell/cmd/misspell

bootstrap:
	go get -t $(TARGET)
	go get -u $(deps)

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

pretty:
	@gocyclo -over 10 . 1>&2
	@ineffassign . 1>&2
	@misspell . 1>&2

clean:
	go clean -ix $(TARGET)
