.PHONY: test gocomply_metaschema clean

test:
	@echo "Running Oscalkit test Utility"
	@echo "Running remaining tests"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v $(shell go list ./... | grep -v "/vendor/\|/test_util/src")

gocomply_metaschema:
ifeq ("$(wildcard $(GOPATH)/bin/gocomply_metaschema)","")
	go get -u -v github.com/gocomply/metaschema/cli/gocomply_metaschema
endif

generate: OSCAL gocomply_metaschema
	gocomply_metaschema generate ./OSCAL/src/metaschema github.com/docker/oscalkit types/oscal

clean:
	rm -rf ./OSCAL
