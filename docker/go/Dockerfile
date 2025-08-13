# 本番環境中間ステージ
FROM golang:1.25-alpine as builder

RUN apk add --update --no-cache ca-certificates git && \
    apk add vim && \
    apk add --no-cache gcc && \
    apk add --no-cache musl-dev

WORKDIR /go/src/app/
COPY ./ /go/src/app

RUN go mod tidy

RUN GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# 本番環境
FROM alpine as prod
WORKDIR /go/src/app/
COPY --from=builder /go/src/app/main /go/src/app/main
CMD [ "/go/src/app/main" ]
EXPOSE 8888

# 開発環境用
FROM golang:1.25 as local

WORKDIR /go/src/app/
COPY ./ /go/src/app

RUN go get github.com/pilu/fresh
RUN go get github.com/jinzhu/gorm
RUN go get github.com/rubenv/sql-migrate/...
RUN go get github.com/rakyll/gotest

RUN go mod tidy

EXPOSE 8888