FROM golang:alpine as builder

LABEL MAINTAINER="whm2416@qq.com"
ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./gvaServer ./
RUN chmod +x ./gvaServer
COPY ./config.docker.yaml ./
COPY  ./resource ./resource/

EXPOSE 8888
ENTRYPOINT ./gvaServer -c config.online.yaml
