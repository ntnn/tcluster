bootstrap_deps:=golang.org/x/tools/cmd/goimports

bootstrap:
	go get -u -t -v .
	go get -u -v $(bootstrap_deps)

generate:
	go generate

test: generate
	go tool vet .
	go test ./...

benchmark: generate
	go test -bench=. ./...
