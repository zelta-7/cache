all: install generate

install:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-types:
	oapi-codegen -generate=types -config server/types.cfg.yaml api-spec.yaml

generate-server:
	oapi-codegen -generate=gin,spec -config server/server.cfg.yaml api-spec.yaml

