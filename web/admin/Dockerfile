#FROM node:18.16.0-alpine3.16 AS build
FROM node:lts-alpine AS build

WORKDIR /opt/www
COPY . /opt/www/

ARG devinggo_NODE_ENV=production

ENV devinggo_NODE_ENV $devinggo_NODE_ENV

RUN echo "devinggo_NODE_ENV=$devinggo_NODE_ENV"

RUN npm install && \
    if [ "$devinggo_NODE_ENV" = "development" ]; then npm run build --mode development; fi && \
    if [ "$devinggo_NODE_ENV" = "production" ]; then npm run build --mode production; fi && \
    npm cache clean --force

RUN npm run generate:version

FROM nginx:alpine

COPY --from=build /opt/www/dist /usr/share/nginx/html
