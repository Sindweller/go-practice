FROM golang:1.19 AS development
WORKDIR $GOPATH/src
COPY . .
RUN GOPROXY="https://goproxy.cn,direct" CGO_ENABLED=0 GOOS=linux go build -o httpserver

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/httpserver .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add curl bash tree tzdata git \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime \
    && echo -ne "Alpine Linux 3.4 image. (`uname -rsv`)\n" >> /root/.built
EXPOSE 8840
ENTRYPOINT ["./httpserver"]