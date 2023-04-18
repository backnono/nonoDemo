FROM 172.2.2.11:5000/lmi/golang:1.18 as builder

WORKDIR /build
COPY . .

RUN go env -w CGO_ENABLED="0"
ENV http_proxy=socks5://172.2.0.230:7891
ENV https_proxy=socks5://172.2.0.230:7891
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN export GIN_MODE=release
RUN cd cmd/app && go build

FROM 172.2.2.11:5000/lmi/ubuntu-base as publish
WORKDIR /app
COPY --from=builder /build/cmd/app .
ENV TZ=Asia/Shanghai
ENTRYPOINT ["./app"]
