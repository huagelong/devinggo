ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "devinggo"
DOCKER_NAME = "devinggo"

# Windows 下使用 Git 的 sh.exe，让 POSIX 风格 recipe（set -e / if [ -d ] / rm -rf 等）直接生效
ifeq ($(OS),Windows_NT)
	SHELL := C:/Program Files/Git/bin/sh.exe
endif

include ./hack/hack.mk