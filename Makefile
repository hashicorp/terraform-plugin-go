default: test

lint:
	golangci-lint run ./...

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	cd tools; go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Protocol Buffers compilation is done outside of 'go generate' handling since
# the 'protoc' tool is not installable via 'go install'.
protoc:
	@cd tfprotov5/internal/tfplugin5 && \
		protoc \
			--proto_path=. \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			tfplugin5.proto
	@cd tfprotov6/internal/tfplugin6 && \
		protoc \
			--proto_path=. \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			tfplugin6.proto

test:
	go test ./...

# Generate copywrite headers
generate:
	cd tools; go generate ./...

.PHONY: default lint protoc test tools
