FROM golang:1.17.6-alpine3.15 as build
LABEL maintainer="117503445"
ENV GO111MODULE=on
WORKDIR /root/project
COPY go.mod go.sum ./
RUN go mod download
RUN wget -O gf https://github.com/gogf/gf-cli/releases/download/v2.0.0-rc/gf_linux_amd64 && chmod +x gf && ./gf install
COPY . .
RUN gf swagger --pack -y
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o server_bin
FROM alpine:3.15 as prod
EXPOSE 8080
# https://www.cnblogs.com/flipped/p/15808681.html
ENV TZ Asia/Shanghai
RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone
WORKDIR /root
COPY --from=build /root/project/server_bin /root/server_bin
ENTRYPOINT /root/server_bin
