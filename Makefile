default: test

lint:
	golangci-lint run ./...

# Protocol Buffers compilation is done outside of 'go generate' handling since
# the 'protoc' tool is not installable via 'go install'.
protobuf:
	go run ./tools/protobuf-compile .

test:
	go test ./...

# Generate copywrite headers
generate:
	cd tools; go generate ./...

.PHONY: default lint protoc test tools
