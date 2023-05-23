.PHONY: clean deps simplify run dev test coverage build update-deps patch-deps

env:
		go env -w GOPRIVATE=github.com/figarocms

clean:
		rm -rf target; \
		rm -f coverage.*

deps: env clean
		go get -d -v ./...

simplify:
		gofmt -s -l -w .

lint:
		golangci-lint run --exclude-use-default=true --deadline=120s --skip-dirs dev

run: deps
		go run -tags=jsoniter *.go

run-api: deps
		go run -tags=jsoniter *.go api

dev:
		docker-compose -f ./scripts/dev/docker-compose.yaml up

test: deps
		go test -count=1 -v ./...

build: test
		CGO_ENABLED=0 GOOS=linux \
		go build \
		-a -installsuffix cgo \
		-tags=jsoniter -o target/app .
