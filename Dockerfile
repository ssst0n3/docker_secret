FROM golang:1.15 AS builder
COPY . /build
WORKDIR /build
RUN GO111MODULE="on" GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@http://ftp.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
sed -i "s@http://security.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
sed -i "s@http://deb.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
apt update && \
apt install -y upx
RUN upx docker_secret

FROM scratch
LABEL maintainer="ssst0n3@gmail.com"
COPY --from=builder /build/docker_secret /docker_secret
ENTRYPOINT ["/docker_secret"]