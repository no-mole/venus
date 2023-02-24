FROM golang:1.19-alpine3.17 as builder

WORKDIR /home

COPY . /home/

RUN ls -l && go env -w GO111MODULE=on  \
    && go env -w GOPROXY="https://goproxy.cn,direct" \
    && go mod tidy \
    && go mod vendor \
    && go build -o venus agent/main.go \
    && go build -o vtlcli vtl/main.go

FROM alpine:3.17 as package

WORKDIR /home

COPY --from=0  /home/venus  /usr/bin
COPY --from=0  /home/vtlcli  /usr/bin/vtl

RUN ls -l && chmod +x /usr/bin/venus
RUN ls -l && chmod +x /usr/bin/vtl

CMD ["venus"]
