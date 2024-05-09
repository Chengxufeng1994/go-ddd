FROM golang:1.22.2 AS build-stage
ARG upx_version=4.2.4

RUN apt-get update && apt-get install -y --no-install-recommends xz-utils

RUN curl -Ls https://github.com/upx/upx/releases/download/v${upx_version}/upx-${upx_version}-amd64_linux.tar.xz -o - | tar xvJf - -C /tmp && \
  cp /tmp/upx-${upx_version}-amd64_linux/upx /usr/local/bin/ && \
  chmod +x /usr/local/bin/upx

RUN apt-get remove -y xz-utils && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -ldflags="-s -w" -o main ./cmd/main.go
RUN upx -9 main

FROM alpine:latest

WORKDIR /usr/local/app

COPY --from=build-stage /usr/local/app/main .

EXPOSE 3030

CMD [ "/usr/local/app/main" ]
