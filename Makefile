.PHONY: default install build clean test fmt vet lint

default: build

build: check_go_version
	go build -o ./bin/clientgofix $(shell ./build/print-ldflags.sh) ./

install: check_go_version
	go install $(shell ./build/print-ldflags.sh) ./

clean:
	rm -fr bin

test:
	go test -v ./...

# Capture output and force failure when there is non-empty output
fmt:
	@echo gofmt -l ./pkg
	@OUTPUT=`gofmt -l ./pkg 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

vet:
	go vet ./pkg

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
# Capture output and force failure when there is non-empty output
lint:
	@echo golint ./...
	@OUTPUT=`golint ./... 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "golint errors:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

check_go_version:
	@OUTPUT=`go version`; \
	case "$$OUTPUT" in \
	*"go1.13."*);; \
	*"go1.14."*);; \
	*"devel"*);; \
	*) \
		echo "Expected: go version go1.13.*, go1.14.*, or devel"; \
		echo "Found:    $$OUTPUT"; \
		exit 1; \
	;; \
	esac
