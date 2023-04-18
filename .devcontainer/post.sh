#!/usr/bin/env bash
sudo apt-get update && sudo apt-get install -y python3-pip unzip
echo '{
  "insecure-registries": [
    "172.2.2.11:5000"
  ]
}' | sudo tee /etc/docker/daemon.json
curl -LO --proxy socks5://172.2.0.230:7891 https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip
unzip protoc-21.5-linux-x86_64.zip -d /go
rm -rf protoc-21.5-linux-x86_64.zip
http_proxy=socks5://172.2.0.230:7891
https_proxy=socks5://172.2.0.230:7891
go env -w CGO_ENABLED="0"
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go install git.aimap.io/LBM/bedrock/cmd/protoc-gen-gokit-endpoint@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/google/wire/cmd/wire@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest