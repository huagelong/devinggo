###############################################################################
#                                build
###############################################################################
FROM golang:1.23-alpine AS go-builder
# ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
# 安装 Make 及其他依赖
RUN apk add --no-cache make git wget nodejs npm yarn
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

# 构建site前端 (Nuxt3)
FROM node:20-alpine AS site-builder
WORKDIR /app
COPY ./web/site ./
COPY ./web/site/.env.docker ./.env.production
RUN yarn install && yarn build:docker

###############################################################################
#                                INSTALLATION
###############################################################################
FROM node:20-alpine
LABEL maintainer="hpuwang@gmail.com"

# 安装Nginx
RUN apk add --no-cache nginx

# 设置在容器内执行时当前的目录
ENV WORKDIR /app
WORKDIR $WORKDIR

# 添加Go应用可执行文件，并设置执行权限
COPY --from=go-builder /app/bin/v1.0.0/linux_amd64/ ./
COPY --from=go-builder /app/docs/docker/start.sh ./start.sh
# 复制Nginx配置文件
COPY --from=go-builder /app/docs/docker/nginx.conf /etc/nginx/http.d/default.conf
# 复制Nuxt3应用的构建结果
COPY --from=site-builder /app/.output /app/site/.output
# 设置权限
RUN chmod +x $WORKDIR/devinggo

# 创建启动脚本
RUN chmod +x /app/start.sh
###############################################################################
#                                   START
###############################################################################

CMD ["/app/start.sh"]