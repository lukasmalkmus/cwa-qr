# TOOLCHAIN
GO				:= CGO_ENABLED=0 GOBIN=$(CURDIR)/bin go
GO_BIN_IN_PATH  := CGO_ENABLED=0 go
GOFMT			:= $(GO)fmt

# ENVIRONMENT
VERBOSE	:=
GOPATH	:= $(GOPATH)

# APPLICATION INFORMATION
BUILD_DATE	:= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
REVISION	:= $(shell git rev-parse --short HEAD)
RELEASE		:= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)-dev
USER		:= $(shell whoami)

# GO TOOLS
GOLANGCI_LINT	:= bin/golangci-lint
GORELEASER		:= bin/goreleaser
GOTESTSUM		:= bin/gotestsum
PROTOC_GEN_GO	:= bin/protoc-gen-go
STRINGER		:= bin/stringer

# MISC
COVERPROFILE	:= coverage.out
DIST_DIR		:= dist

# GOTAGS
GOTAGS := osusergo netgo static_build

# GOLDFLAGS
GOLDFLAGS 	:= -s -w -extldflags "-fno-PIC -static -Wl -z now -z relro"
GOLDFLAGS 	+= -X github.com/axiomhq/pkg/version.release=$(RELEASE)
GOLDFLAGS 	+= -X github.com/axiomhq/pkg/version.revision=$(REVISION)
GOLDFLAGS 	+= -X github.com/axiomhq/pkg/version.buildDate=$(BUILD_DATE)
GOLDFLAGS 	+= -X github.com/axiomhq/pkg/version.buildUser=$(USER)

# FLAGS
GOFLAGS 			:= -buildmode=pie -installsuffix=cgo -trimpath  -tags='$(GOTAGS)' -ldflags='$(GOLDFLAGS)'
GO_TEST_FLAGS		:= -race -coverprofile=$(COVERPROFILE)
GORELEASER_FLAGS	:= --snapshot --rm-dist

# DEPENDENCIES
GOMODDEPS := go.mod go.sum

# Enable verbose test output if explicitly set.
GOTESTSUM_FLAGS =
ifdef VERBOSE
	GOTESTSUM_FLAGS += --format=standard-verbose
endif

# FUNCTIONS
# func go-list-pkg-sources(package)
go-list-pkg-sources = $(GO) list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $(1)
# func go-pkg-sourcefiles(package)
go-pkg-sourcefiles = $(shell $(call go-list-pkg-sources,$(strip $1)))

.PHONY: all
all: dep generate fmt lint test build ## Run dep, generate, fmt, lint, test and build

.PHONY: build
build: $(GORELEASER) dep.stamp $(call go-pkg-sourcefiles, ./...) ## Build the binary
	@echo ">> building binary"
	@$(GORELEASER) build $(GORELEASER_FLAGS)

.PHONY: clean
clean: ## Remove build and test artifacts
	@echo ">> cleaning up artifacts"
	@rm -rf $(COVERPROFILE) $(DIST_DIR)

.PHONY: cover
cover: $(COVERPROFILE) ## Calculate the code coverage score
	@echo ">> calculating code coverage"
	@$(GO) tool cover -func=$(COVERPROFILE) | tail -n1

.PHONY: dep-clean
dep-clean: ## Remove obsolete dependencies
	@echo ">> cleaning dependencies"
	@$(GO) mod tidy

.PHONY: dep-upgrade
dep-upgrade: ## Upgrade all direct dependencies to their latest version
	@echo ">> upgrading dependencies"
	@$(GO) get -d $(shell $(GO) list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	@make dep

.PHONY: dep
dep: dep-clean dep.stamp ## Install and verify dependencies and remove obsolete ones

dep.stamp: $(GOMODDEPS)
	@echo ">> installing dependencies"
	@$(GO) mod download
	@$(GO) mod verify
	@touch $@

.PHONY: fmt
fmt: ## Format and simplify the source code using `gofmt`
	@echo ">> formatting code"
	@! $(GOFMT) -s -w $(shell find . -path -prune -o -name '*.go' -print) | grep '^'

.PHONY: generate
generate: event_string.go internal/pb/qr_code.pb.go ## Generate code using `go generate` or other tools

.PHONY: install
install: $(GOPATH)/bin/cwa-qr ## Install the binary into the $GOPATH/bin directory

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint the source code
	@echo ">> linting code"
	@$(GOLANGCI_LINT) run

.PHONY: test
test: $(GOTESTSUM) ## Run all unit tests. Run with VERBOSE=1 to get verbose test output ('-v' flag)
	@echo ">> running tests"
	@$(GOTESTSUM) $(GOTESTSUM_FLAGS) -- $(GO_TEST_FLAGS) ./...

.PHONY: tools
tools: dep.stamp $(GOLANGCI_LINT) $(GORELEASER) $(GOTESTSUM) $(PROTOC_GEN_GO) $(STRINGER) ## Install all tools into the projects local $GOBIN directory

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# GO GENERATE TARGETS

%_string.go: %.go | $(STRINGER)
	@echo ">> generating $@ from $<"
	@$(GO) generate $<

internal/pb/%.pb.go: proto/%.proto | $(PROTOC_GEN_GO)
	@echo ">> generating $@ from $<"
	@protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --go_out=. $<

# MISC TARGETS

$(COVERPROFILE): test

# INSTALL TARGETS

$(GOPATH)/bin/cwa-qr: dep.stamp $(call go-pkg-sourcefiles, ./...)
	@echo ">> installing cwa-qr binary"
	@$(GO_BIN_IN_PATH) install $(GOFLAGS) ./cmd/cwa-qr

# GO TOOLS

$(GOLANGCI_LINT): dep.stamp $(call go-pkg-sourcefiles, github.com/golangci/golangci-lint/cmd/golangci-lint)
	@echo ">> installing golangci-lint"
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint

$(GORELEASER): dep.stamp $(call go-pkg-sourcefiles, github.com/goreleaser/goreleaser)
	@echo ">> installing goreleaser"
	@$(GO) install github.com/goreleaser/goreleaser

$(GOTESTSUM): dep.stamp $(call go-pkg-sourcefiles, gotest.tools/gotestsum)
	@echo ">> installing gotestsum"
	@$(GO) install gotest.tools/gotestsum

$(PROTOC_GEN_GO): dep.stamp $(call go-pkg-sourcefiles, google.golang.org/protobuf/cmd/protoc-gen-go)
	@echo ">> installing protoc-gen-go"
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go

$(STRINGER): dep.stamp $(call go-pkg-sourcefiles, golang.org/x/tools/cmd/stringer)
	@echo ">> installing stringer"
	@$(GO) install golang.org/x/tools/cmd/stringer
