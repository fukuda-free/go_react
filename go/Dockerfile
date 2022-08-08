FROM golang:alpine3.15
RUN apk update && apk add alpine-sdk

ARG UID
RUN adduser -D app -u ${UID:-1000}
USER app


ENV ROOT=/go/src
# ENV CGO_ENABLED 0
WORKDIR ${ROOT}



# COPY src/go.mod .
# COPY src/go.sum .
# RUN go mod download



# # vscodeの拡張機能のためにインストール
# RUN go install golang.org/x/tools/gopls@latest
# RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go get github.com/gorilla/sessions # gorillaセッションの利用
RUN go get github.com/go-sql-driver/mysql # mysql
RUN go get -u github.com/jinzhu/gorm # DBのテーブルとプログラミング元を関連付けて扱えるようにする


RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/joho/godotenv
