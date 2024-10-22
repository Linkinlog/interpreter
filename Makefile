MAIN := main.go
BINARY := mag
PWD := $(shell pwd)

all: $(BINARY)

$(BINARY): $(MAIN)
	go build -o $@ $^

clean:
	rm -f $(BINARY)

lint:
	command golangci-lint run

fmt:
	command gofmt -d -w $(PWD)

test:
	command go test ./... -v

neat: lint fmt

check: neat test

.PHONY: all clean lint fmt test neat
