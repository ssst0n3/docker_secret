FROM golang:1.16-alpine AS builder
#ENV GOPROXY="https://proxy.golang.org"
ENV GOPROXY="https://goproxy.io,direct"
COPY . /build
WORKDIR /build
RUN GO111MODULE="on" GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@https://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx docker_secret

FROM scratch
LABEL maintainer="ssst0n3@gmail.com"
COPY --from=builder /build/docker_secret /docker_secret
ENTRYPOINT ["/docker_secret"]