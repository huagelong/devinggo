###############################################################################
#                                build
###############################################################################
FROM golang:1.21-alpine AS builder
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GO111MODULE=auto \
    GOOS=linux

# 安装 Node.js、Yarn、Make 及其他依赖
RUN apk add --no-cache nodejs npm yarn make
WORKDIR /app
COPY . ./
RUN mv ./manifest/config/config.docker.yaml ./manifest/config/config.yaml
RUN mv ./hack/config.example.yaml ./hack/config.yaml
RUN rm -rf ./web/system/.env.production
RUN mv ./web/system/.env.docker ./web/system/.env.production
RUN make build
###############################################################################
#                                INSTALLATION
###############################################################################
FROM loads/alpine:3.8
LABEL maintainer="hpuwang@gmail.com"
# 设置在容器内执行时当前的目录
ENV WORKDIR /opt/devinggo
# 添加应用可执行文件，并设置执行权限
COPY --from=builder /app/bin/v1.0.0/linux_amd64/devinggo   $WORKDIR/devinggo
RUN chmod +x $WORKDIR/devinggo
###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./devinggo