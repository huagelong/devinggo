###############################################################################
#                                build
###############################################################################
FROM golang:1.23.4-alpine AS go-builder
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
# 安装 Make 及其他依赖
RUN apk add --no-cache make git wget nodejs npm \
    && npm install -g corepack
WORKDIR /app
COPY . ./
RUN mv ./manifest/config/config.docker.yaml ./manifest/config/config.yaml
RUN mv ./hack/config.docker.yaml ./hack/config.yaml
RUN rm -rf ./admin-ui/.env.production
RUN mv ./admin-ui/.env.docker ./admin-ui/.env.production
# make build 会自动安装 gf CLI 并构建前端+后端（GOPROXY 已设国内代理）
RUN make build
RUN chmod +x ./bin/v2.0.0/linux_amd64/devinggo
RUN cd ./bin/v2.0.0/linux_amd64/ && ./devinggo unpack
RUN ls -la ./bin/v2.0.0/linux_amd64

###############################################################################
#                                INSTALLATION
###############################################################################
FROM alpine:3.20
LABEL maintainer="hpuwang@gmail.com"

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置在容器内执行时当前的目录
ENV WORKDIR=/app
WORKDIR $WORKDIR

# 添加Go应用可执行文件，并设置执行权限
COPY --from=go-builder /app/bin/v2.0.0/linux_amd64/ ./
# 设置权限
RUN chmod +x $WORKDIR/devinggo

# 暴露应用端口
EXPOSE 8070
###############################################################################
#                                   START
###############################################################################

CMD ["/app/devinggo", "--gf.gmod=develop"]
