.PHONY: generate

generate:
	@mkdir -p gen
	@protoc \
		   --go_out=gen --go_opt=paths=source_relative \
		   --go-grpc_out=gen --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		   --validate_out="lang=go:gen" --validate_opt=paths=source_relative \
		   -I proto \
		   proto/pokemon.proto