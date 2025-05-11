PLATFORM_RESOURCE_PATH = "./resource/public/admin"
UI_PATH = "./web/admin"
SITE_PATH = "./web/site"

VERSION = $(shell git describe --tags --always --match='v*')
SED = sed
ifneq ($(shell go env GOOS),windows)
	ifeq ($(shell uname), Darwin)
		SED = gsed
	endif
endif
##   run: Run devinggo for development environment
.PHONY: run
run: dao service
	@echo "******** gf run ********"
	@go mod tidy && go run main.go

# Build binary using configuration from hack/config.yaml.
.PHONY: build
build: cli.install ui.build
	@if [ -d $(PLATFORM_RESOURCE_PATH) ]; then rm -rf $(PLATFORM_RESOURCE_PATH); fi
	@mkdir $(PLATFORM_RESOURCE_PATH)
	@if [ -d $(UI_PATH)/dist ]; then cd $(UI_PATH) && \cp -rf ./dist/* ../../$(PLATFORM_RESOURCE_PATH); fi
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@if [ -f internal/packed/packed.go ]; then rm -rf internal/packed/packed.go; fi
	@go mod tidy
	@gf build -ew

.PHONY: install
install: 
	@echo "******** install ********"
	@go run main.go install

#node package install
.PHONY: ui.install
ui.install: cli.install
	@set -e;\
	cd $(UI_PATH);\
	yarn install;

#ui build
.PHONY: ui.build
ui.build: ui.install
	@set -e;\
	cd $(UI_PATH);\
	yarn build;

.PHONY: siteui.install
siteui.install: cli.install
	@set -e;\
	cd $(SITE_PATH);\
	yarn install;

#ui build
.PHONY: siteui.build
siteui.build: siteui.install
	@set -e;\
	cd $(SITE_PATH);\
	yarn build;
