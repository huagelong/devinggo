include ./hack/hack-cli.mk
include ./hack/hack-cus.mk

# Update GoFrame and its CLI to latest stable version.
.PHONY: up
up: cli.install
	@gf up -a


# Parse api and generate controller/sdk.
.PHONY: ctrl
ctrl: cli.install
	@gf gen ctrl

# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao

# Parse current project go files and generate enums go file.
.PHONY: enums
enums: cli.install
	@gf gen enums

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@gf gen service
	@gf gen service -s=modules/system/logic -d=modules/system/service
	@gf gen service -s=modules/api/logic -d=modules/api/service

# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git describe --dirty --always --tags --abbrev=8 --match 'v*' | sed 's/-/./2' | sed 's/-/./2'))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@gf docker ${_PUSH} -tn $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push:
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_ENV};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	if [ $(DEPLOY_NAME) != "" ]; then \
		kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}"; \
	fi;


# Parsing protobuf files and generating go files.
.PHONY: pb
pb: cli.install
	@gf gen pb

# Generate protobuf files for database tables.
.PHONY: pbentity
pbentity: cli.install
	@gf gen pbentity

# Run golangci-lint to check code quality.
.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --config .golangci.yml

# Format code with goimports and gofmt.
.PHONY: fmt
fmt:
	@go run golang.org/x/tools/cmd/goimports@latest -w .
	@gofmt -s -w .


# ====================================================================
# Code Generator Commands
# ====================================================================

# Show help for all generator commands.
.PHONY: gen-help
gen-help:
	@echo "=== devinggo 代码生成器命令 ==="
	@echo ""
	@echo "模块管理："
	@echo "  make gen-module name=<模块名>           创建新模块"
	@echo "  make clone-module name=<新模块名> source=<源模块名>  克隆模块"
	@echo "  make export-module name=<模块名>        导出模块为zip"
	@echo "  make import-module file=<zip文件>       导入模块"
	@echo "  make list-modules                       列出所有模块"
	@echo "  make validate-module name=<模块名>      验证模块结构"
	@echo ""
	@echo "Worker生成："
	@echo "  make gen-worker module=<模块名> worker=<worker名>  创建worker"
	@echo ""
	@echo "CRUD生成："
	@echo "  make gen-crud table=<表名> [module=<模块名>]  生成CRUD代码"
	@echo ""

# Create a new module.
.PHONY: gen-module
gen-module:
	@if [ -z "$(name)" ]; then \
		echo "错误: 请指定模块名称，例如: make gen-module name=mymodule"; \
		exit 1; \
	fi
	@go run ./hack/generator/main.go module:create -name $(name)
	@echo "模块 $(name) 创建成功！"
	@echo "请运行以下命令更新代码："
	@echo "  make service"
	@echo "  make dao"

# Clone an existing module.
.PHONY: clone-module
clone-module:
	@if [ -z "$(name)" ]; then \
		echo "错误: 请指定新模块名称，例如: make clone-module name=newmodule source=system"; \
		exit 1; \
	fi
	@if [ -z "$(source)" ]; then \
		echo "错误: 请指定源模块名称，例如: make clone-module name=newmodule source=system"; \
		exit 1; \
	fi
	@go run ./hack/generator/main.go module:clone -source $(source) -target $(name)
	@echo "模块克隆成功！"
	@echo "请运行以下命令更新代码："
	@echo "  make service"

# Export a module to zip file.
.PHONY: export-module
export-module:
	@if [ -z "$(name)" ]; then \
		echo "错误: 请指定模块名称，例如: make export-module name=system"; \
		exit 1; \
	fi
	@go run ./hack/generator/main.go module:export -name $(name)

# Import a module from zip file.
.PHONY: import-module
import-module:
	@if [ -z "$(file)" ]; then \
		echo "错误: 请指定zip文件路径，例如: make import-module file=./system.zip"; \
		exit 1; \
	fi
	@go run ./hack/generator/main.go module:import -file $(file)
	@echo "模块导入成功！"
	@echo "请运行以下命令更新代码："
	@echo "  make service"

# List all modules.
.PHONY: list-modules
list-modules:
	@go run ./hack/generator/main.go module:list

# Validate module structure.
.PHONY: validate-module
validate-module:
	@if [ -z "$(name)" ]; then \
		echo "错误: 请指定模块名称，例如: make validate-module name=system"; \
		exit 1; \
	fi
	@go run ./hack/generator/main.go module:validate -name $(name)

# Create a new worker.
.PHONY: gen-worker
gen-worker:
	@if [ -z "$(module)" ]; then \
		echo "错误: 请指定模块名称，例如: make gen-worker module=system worker=MyWorker"; \
		exit 1; \
	fi
	@if [ -z "$(worker)" ]; then \
		echo "错误: 请指定worker名称，例如: make gen-worker module=system worker=MyWorker"; \
		exit 1; \
	fi
	@if [ -z "$(type)" ]; then \
		go run ./hack/generator/main.go worker:create -module $(module) -name $(worker); \
	else \
		go run ./hack/generator/main.go worker:create -module $(module) -name $(worker) -type $(type); \
	fi
	@echo "Worker创建成功！"

# Generate CRUD code for database table.
.PHONY: gen-crud
gen-crud:
	@if [ -z "$(table)" ]; then \
		echo "错误: 请指定表名称，例如: make gen-crud table=system_user"; \
		exit 1; \
	fi
	@if [ -z "$(module)" ]; then \
		if [ -z "$(name)" ]; then \
			go run ./hack/generator/main.go crud:generate -t=$(table) -n=$(table); \
		else \
			go run ./hack/generator/main.go crud:generate -t=$(table) -n=$(name); \
		fi; \
	else \
		if [ -z "$(name)" ]; then \
			go run ./hack/generator/main.go crud:generate -m=$(module) -t=$(table) -n=$(table); \
		else \
			go run ./hack/generator/main.go crud:generate -m=$(module) -t=$(table) -n=$(name); \
		fi; \
	fi
	@echo "CRUD代码生成成功！"
	@echo "请运行以下命令更新代码："
	@echo "  make service"
	@echo "  make ctrl"