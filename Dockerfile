FROM node:19.8-alpine3.16 as ui

WORKDIR /home

RUN #sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN wget -qO- https://get.pnpm.io/install.sh | ENV="$HOME/.shrc" SHELL="$(which sh)" sh -

COPY ui /home/ui

RUN source /root/.shrc && cd /home/ui && pnpm i && pnpm build

FROM golang:1.19-alpine3.17 as builder

WORKDIR /home

COPY . /home/

COPY --from=0  /home/ui/dist  /home/agent/venus/api/ui

RUN ls -l && go env -w GO111MODULE=on  \
    && go env -w GOPROXY="https://goproxy.cn,direct" \
    && go mod tidy \
    && go mod vendor \
    && go build -o venus agent/main.go \
    && go build -o vtlcli vtl/main.go

FROM alpine:3.17 as package

WORKDIR /home

COPY --from=1  /home/venus  /usr/bin
COPY --from=1  /home/vtlcli  /usr/bin/vtl

RUN ls -l && chmod +x /usr/bin/venus
RUN ls -l && chmod +x /usr/bin/vtl

CMD ["venus"]
