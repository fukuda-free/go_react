# FROM golang:alpine3.15
# FROM golang:1.19.0
# FROM public.ecr.aws/docker/library/golang:1.19.0
FROM public.ecr.aws/docker/library/golang:1.18.5

# 権限関係
# ARG UID
# RUN adduser -D app -u ${UID:-1000}
# USER app


# ディレクトリ
ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src

# インストール系
RUN apt-get update
# vscodeの拡張機能のためにインストール
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest

# beegoで必要なプラグイン
# RUN go get "github.com/go-sql-driver/mysql"
# RUN go install github.com/go-sql-driver/mysql@latest

# RUN go get "github.com/beego/bee"
# RUN go get "github.com/astaxie/beego"
# # RUN go install github.com/beego/bee/v2@latest