PLATFORM_RESOURCE_PATH = "./resource/public/system"
UI_PATH = "./web/system"
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
	@go mod tidy && gf run main.go

# Build binary using configuration from hack/config.yaml.
.PHONY: build
build: cli.install ui.build
	@rm -rf $(PLATFORM_RESOURCE_PATH)
	@mkdir $(PLATFORM_RESOURCE_PATH)
	@cd $(UI_PATH) && \cp -rf ./dist/*  ../../$(PLATFORM_RESOURCE_PATH)
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@rm -rf internal/packed/packed.go
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

.PHONY: config-init
config-init:
	@echo "开始备份配置文件并替换敏感配置信息..."
	# 备份原始配置文件
	@cp hack/config.example.yaml hack/config.yaml
	@cp manifest/config/config.example.yaml manifest/config/config.yaml
	@echo "success"
