default: fmt test

fmt:
	go fmt ./...

test:
	go test ./...

install:
	go install github.com/a-skua/fzbz/cmd/fzbz
