PROJECT=github.com/wesovilabs/baboon
GO  = GOFLAGS=-mod=vendor go
GOBUILD  = CGO_ENABLED=0 $(GO) build
CMD_PATH=$(PROJECT)/cmd/baboon
REPORT_TESTUNIT=baboon-testUnit

all: clean deps lint test testInt build-all

.PHONY: clean
clean: ; @ ## Remove temporal files
	rm -fR build test_report

.PHONY: fmt
fmt: ; @ ## Code formatter
	for pkg in $(shell $(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/ ); do \
		gofmt -l -w  -e $$pkg/*.go; \
	done

.PHONY: lint
lint: ; @ ## Code analysis
	$(GO) run -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose

.PHONY: test
test: ; @ ## Run unit tests
	mkdir -p test_report
	$(GO) test -p=1  $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{.ImportPath}}{{ end }}' ./... | grep -v test ) -v -timeout 10s -short -cover| tee $(PWD)/test_report/$(REPORT_TESTUNIT).output; \
	status=$$?; \
	$(GO) run -mod=vendor github.com/tebeka/go2xunit -fail -input $(PWD)/test_report/$(REPORT_TESTUNIT).output -output $(PWD)/test_report/$(REPORT_TESTUNIT).xml; \
	exit $$status;

.PHONY: testInt
testInt: ; @ ## Run integration tests
	mkdir -p test_report; \
	cd test; \
	TEST_REPORT_DIR=$(PWD)/test_report DATASET_PATH=$(PWD)/testdata/paths go test -v -mod=vendor; \
	status=$$?; \
	cd .. ; \
	exit $$status;

.PHONY: deps
deps: ; @ ## Download project dependencies
	go mod tidy;
	go mod vendor;

.PHONY: build
build:
	$(GOBUILD) -o build/$(PROJECT)  $(CMD_PATH)

.PHONY: build-all
build-all:
	GOARCH=amd64 GOOS=linux  $(GOBUILD) -o build/$(PROJECT).linux  $(CMD_PATH)
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o build/$(PROJECT).darwin $(CMD_PATH)
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o build/$(PROJECT).exe $(CMD_PATH)

.PHONY: run
run: ; @ ## Run this project
ifeq ($(config),)
	APP_CONFIG=$(PWD)/scripts/config/local/app.yml $(GO) run cmd/api/main.go;
else
	APP_CONFIG=$(config) $(GO) run cmd/api/main.go;
endif
