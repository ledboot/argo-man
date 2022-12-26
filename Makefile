PROJECT_NAME := $(shell grep 'module ' go.mod | awk '{print $$2}' | sed 's/github.com\/ledbot\///g')
NAME := $(notdir $(PROJECT_NAME))
NEW_NAME :=$(shell echo $(NAME) | tr '_' '-')
NOW := $(shell date +'%Y%m%d%H%M%S')
TAG := $(shell git describe --always --tags --abbrev=0 | tr -d "[\r\n]")
COMMIT := $(shell git rev-parse --short HEAD| tr -d "[ \r\n\']")
VERSION_PKG := github.com/ledboot/argo-man/config
LD_FLAGS_BASE := -X $(VERSION_PKG).serviceName=$(NAME) -X $(VERSION_PKG).version=$(TAG)-$(COMMIT) -X $(VERSION_PKG).buildTime=$(shell date +%Y%m%d%H%M%S)
LD_FLAGS := -s -w $(LD_FLAGS_BASE)

IMPORTANT_GO_ENV_VARS := "GOPATH|GO111MODULE|GOARCH|GOCACHE|GOMODCACHE|GONOPROXY|GONOSUMDB|GOPRIVATE|GOPROXY|GOSUMDB|GOMOD|CGO"


.PHONY: build/binary
build/binary: export CGO_ENABLED=1
build/binary: export GOARCH=amd64
build/binary: export GOOS=linux
build/binary:
	@echo "\n###### building $(NAME)"
	@go env | grep -E $(IMPORTANT_GO_ENV_VARS)
	go build -trimpath -ldflags="$(LD_FLAGS)" -o $(NAME)

fmt:
	command -v gofumpt || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install mvdan.cc/gofumpt@latest && cd $(WORK))
	gofumpt -w -d .

lint:
	command -v golangci-lint || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0 && cd $(WORK))
	golangci-lint run  -v

ci/lint: export GOPATH=/go
ci/lint: export GO111MODULE=on
ci/lint: export GOPROXY=https://goproxy.cn
ci/lint: export GOPRIVATE=code.skyhorn.net
ci/lint: export GOOS=linux
ci/lint: export CGO_ENABLED=1
ci/lint: lint
