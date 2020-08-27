GO=GO111MODULE=on go
.PHONY: test gocomply_metaschema clean build vendor

build:
	go build ./...
	go build -o oscalkit -v cli/main.go

test:
	@echo "Running Oscalkit test Utility"
	@echo "Running remaining tests"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v $(shell go list ./... | grep -v "/vendor/")

gocomply_metaschema:
ifeq ("$(wildcard $(GOPATH)/bin/gocomply_metaschema)","")
	go get -u -v github.com/gocomply/metaschema/cli/gocomply_metaschema
endif

generate: OSCAL gocomply_metaschema
	gocomply_metaschema generate ./OSCAL/src/metaschema github.com/gocomply/oscalkit types/oscal

OSCAL:
	git clone --depth 1 https://github.com/usnistgov/OSCAL

clean:
	rm -rf ./OSCAL
	rm -rf ./oscalkit

vendor:
	$(GO) mod tidy
	$(GO) mod vendor
	$(GO) mod verify
