# stage 1: build src code to binary
# 获取对应版本号 https://hub.docker.com/_/golang
FROM golang:1.18-alpine3.11 as builder
ENV GOPROXY="https://goproxy.io"
ENV GOPATH=$HOME/go
ENV GOBIN=$HOME/go/bin
ENV PATH=$PATH:$GOPATH/bin
COPY .. /app/
# 下载指定的包，go.mod已经记录，可以直接使用
RUN cd /app && go build -o app cmd/app/main.go

# stage 2: use alpine as base image
FROM alpine:3.10

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  && \
    apk update && \
    apk --no-cache add tzdata ca-certificates && \
    cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
# 使用--from的参数做到拷贝使用
COPY --from=builder /app/app /app
COPY --from=builder /app/config/config.yaml /etc/demo_config.yaml
CMD ["sh","-c","/app -c /etc/demo_config.yaml"]