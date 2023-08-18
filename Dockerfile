# build
FROM golang:1.18 as build

WORKDIR /build

ADD . .

RUN GOPROXY=https://goproxy.cn,direct go mod download && go build -o app

# release
FROM alpine as release

WORKDIR app

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /build/app ./go-sample

EXPOSE 8080

ENTRYPOINT ["./app/go-sample"]