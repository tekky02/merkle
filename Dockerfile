FROM golang:latest as builder

ADD merkle /root/

RUN go env -w "GO111MODULE=auto" && go env -w "GOPROXY=https://goproxy.cn,direct" \
  cd /root/merkle && GOOS=linux GOARCH=amd64 go build main/main.go


FROM alpine:latest

COPY --from=builder /root/merkle /root/

