#!/bin/sh

# 启动 Nuxt3 应用
cd /app/site && node ./.output/server/index.mjs &

mkdir -p /run/nginx

# 启动 Nginx
nginx &

# 启动 Go 应用
#cd /app && ./devinggo --gf.gmod=develop -config=./manifest/config/config.docker.yaml
cd /app && ./devinggo --gf.gmod=develop