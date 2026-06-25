###############################################################################
#                                build
###############################################################################
FROM golang:1.23.4-alpine AS go-builder
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
# 安装 Make 及其他依赖
RUN apk add --no-cache make git wget nodejs yarn
WORKDIR /app
COPY . ./
RUN mv ./manifest/config/config.docker.yaml ./manifest/config/config.yaml
RUN mv ./hack/config.docker.yaml ./hack/config.yaml
RUN rm -rf ./admin-ui/.env.production
RUN mv ./admin-ui/.env.docker ./admin-ui/.env.production
# 预装 gf CLI（走 goproxy 国内代理，避免 wget 直连 GitHub 不稳定）
RUN go install github.com/gogf/gf/cmd/gf@latest
RUN make build
RUN chmod +x ./bin/v1.0.0/linux_amd64/devinggo
RUN cd ./bin/v1.0.0/linux_amd64/ && ./devinggo unpack
RUN ls -la ./bin/v1.0.0/linux_amd64

###############################################################################
#                                INSTALLATION
###############################################################################
FROM alpine:3.20
LABEL maintainer="hpuwang@gmail.com"

# 安装Nginx及运行时依赖
RUN apk add --no-cache nginx ca-certificates tzdata

# 设置在容器内执行时当前的目录
ENV WORKDIR=/app
WORKDIR $WORKDIR

# 添加Go应用可执行文件，并设置执行权限
COPY --from=go-builder /app/bin/v1.0.0/linux_amd64/ ./
COPY --from=go-builder /app/docs/docker/start.sh ./start.sh
# 复制Nginx配置文件
COPY --from=go-builder /app/docs/docker/nginx.conf /etc/nginx/http.d/default.conf
# 设置权限
RUN chmod +x $WORKDIR/devinggo

# 创建启动脚本
RUN chmod +x /app/start.sh

# 暴露Nginx端口
EXPOSE 80
###############################################################################
#                                   START
###############################################################################

CMD ["/app/start.sh"]