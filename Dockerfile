###############################################################################
#                                build
###############################################################################
FROM golang:1.23-alpine AS builder
# ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
# 安装 Node.js、Yarn、Make 及其他依赖
RUN apk add --no-cache nodejs npm yarn make git wget
WORKDIR /app
COPY . ./
RUN mv ./manifest/config/config.docker.yaml ./manifest/config/config.yaml
RUN mv ./hack/config.docker.yaml ./hack/config.yaml
RUN rm -rf ./web/admin/.env.production
RUN mv ./web/admin/.env.docker ./web/admin/.env.production
RUN make cli
RUN make build
RUN chmod +x ./bin/v1.0.0/linux_amd64/devinggo
RUN cd ./bin/v1.0.0/linux_amd64/ && ./devinggo unpack
RUN ls -la ./bin/v1.0.0/linux_amd64
###############################################################################
#                                INSTALLATION
###############################################################################
FROM loads/alpine:3.8
LABEL maintainer="hpuwang@gmail.com"
# 设置在容器内执行时当前的目录
ENV WORKDIR /app
WORKDIR $WORKDIR
# 添加应用可执行文件，并设置执行权限
COPY --from=builder /app/bin/v1.0.0/linux_amd64/ ./
RUN ls -la ./
RUN chmod +x $WORKDIR/devinggo
###############################################################################
#                                   START
###############################################################################

CMD ./devinggo