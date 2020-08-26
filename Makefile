GOOS := darwin
GOARCH := amd64
VERSION := 0.2.0
BUILD := $(shell git rev-parse --short HEAD)-dev
DATE := $(shell date "+%Y-%m-%d")

NAMESPACE := docker
REPO := oscalkit
BINARY=oscalkit_$(GOOS)_$(GOARCH)

.DEFAULT_GOAL := $(BINARY)
.PHONY: test build-docker push $(BINARY) clean

test:
	@echo "Running Oscalkit test Utility"
	@echo "Running remaining tests"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v $(shell go list ./... | grep -v "/vendor/\|/test_util/src")

build-docker:
	docker image build --build-arg VERSION=$(VERSION) --build-arg BUILD=$(BUILD) --build-arg DATE=$(DATE) -t $(NAMESPACE)/$(REPO):$(VERSION)-$(BUILD) .

push: build-docker
	docker image push $(NAMESPACE)/$(REPO):$(BUILD)

build:
	docker image build -f Dockerfile.build \
		--build-arg GOOS=$(GOOS) \
		--build-arg GOARCH=$(GOARCH) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD=$(BUILD) \
		--build-arg DATE=$(DATE) \
		--build-arg BINARY=$(BINARY) \
		-t $(NAMESPACE)/$(REPO):$(VERSION)-$(BUILD)-builder .

$(BINARY): build
	$(eval ID := $(shell docker create $(NAMESPACE)/$(REPO):$(VERSION)-$(BUILD)-builder))
	@docker cp $(ID):/$(BINARY) .
	@docker rm $(ID) >/dev/null

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
