FROM node:20-alpine AS builder

ENV APP_HOME ./
WORKDIR "$APP_HOME"

COPY . ./
RUN npm install --registry=https://registry.npmmirror.com
RUN npm run build:docker

FROM node:20-alpine

COPY --from=builder ./.output /app/.output

WORKDIR /app

EXPOSE 3000

ENV ENV=docker
CMD ["node", "./.output/server/index.mjs"]
